package core

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/arnaudlcm/container-engine/internal/core/logger"
	"github.com/arnaudlcm/container-engine/internal/parser"
	pb "github.com/arnaudlcm/container-engine/service/proto"
	"github.com/google/uuid"
)

type Container struct {
	ID      uuid.UUID
	RootFs  string // Path to the root fs
	Status  pb.ContainerStatus
	Process Process
	Manager CGroupManager
	Logger  *logger.Logger
}

func (d *EngineDaemon) GetContainers(ctx context.Context, req *pb.GetContainersRequest) (*pb.GetContainersResponse, error) {

	containers := make([]*pb.ContainerInfos, 0, len(d.containers))

	for _, c := range d.containers {
		containers = append(containers, &pb.ContainerInfos{
			Id:     c.ID.String(),
			Status: c.Status,
		})
	}

	return &pb.GetContainersResponse{Containers: containers}, nil
}

func (g *EngineDaemon) CreateContainer(req *pb.CreateContainerRequest, stream pb.ContainerDaemonService_CreateContainerServer) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	container := Container{}

	// Send a message about the creation process starting
	if err := stream.Send(&pb.CreateContainerResponse{
		Success: false,
		Message: "Starting container creation",
	}); err != nil {
		return fmt.Errorf("error sending progress: %w", err)
	}

	uuid, err := g.getUniqueUUID()
	if err != nil {
		if err := stream.Send(&pb.CreateContainerResponse{
			Success: false,
			Message: "Failed to generate UUID",
		}); err != nil {
			return fmt.Errorf("error sending progress: %w", err)
		}
		return err
	}

	container.ID = uuid

	// Notify the client about CGroupManager creation
	container.Manager, err = NewCGroupManager(container.ID)
	if err != nil {
		if err := stream.Send(&pb.CreateContainerResponse{
			Success: false,
			Message: "Error during CGroupManager creation",
		}); err != nil {
			return fmt.Errorf("error sending progress: %w", err)
		}
		return fmt.Errorf("error during CGroupManager creation: %w", err)
	}

	if err := stream.Send(&pb.CreateContainerResponse{
		Success: false,
		Message: "Setting up the layer",
	}); err != nil {
		return fmt.Errorf("error sending progress: %w", err)
	}

	// First retrieve the manifest
	var reader io.ReadCloser
	// Check if layerUrl is a URL or a local file path
	if parsedURL, err := url.ParseRequestURI(req.Config.Layer); err == nil && (parsedURL.Scheme == "http" || parsedURL.Scheme == "https") {
		// HTTP(S) URL: Download the tarball
		resp, err := http.Get(req.Config.Layer)
		if err != nil {
			return fmt.Errorf("failed to download layer manifest: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("bad response status: %d", resp.StatusCode)
		}

		reader = resp.Body
	} else {
		// Local file path: Open the file
		file, err := os.Open(req.Config.Layer)
		if err != nil {
			return fmt.Errorf("failed to open local file: %w", err)
		}
		defer file.Close()

		reader = file
	}

	data, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	manifest, err := parser.ParseImageManifest(data)

	if err != nil {
		return err
	}

	path, err := g.fsManager.AddLayer(manifest, g.manifestKey, uuid.String())
	if err != nil {
		if err := stream.Send(&pb.CreateContainerResponse{
			Success: false,
			Message: "Error during layer setup",
		}); err != nil {
			return fmt.Errorf("error sending progress: %w", err)
		}
		return err
	}

	// Setup the logger
	logger, err := logger.NewLogger(container.ID.String())
	if err != nil {
		if err := stream.Send(&pb.CreateContainerResponse{
			Success: false,
			Message: "Error during logger setup",
		}); err != nil {
			return fmt.Errorf("error sending progress: %w", err)
		}
	}

	container.Logger = &logger

	process := Process{
		Args:              req.Config.Cmd,
		Stdin:             os.Stdin,
		Stdout:            os.Stdout,
		CommunicationPipe: nil,
		UID:               0,
		GID:               0,
		rootPath:          path,
		workingDirectory:  req.Config.Workdir,
	}

	container.Process = process
	container.Status = pb.ContainerStatus_CONTAINER_HANGING
	g.containers[uuid] = container

	if err := process.Init(); err != nil {
		return err
	}

	// Init the logger
	stdoutPipe, _ := process.cmd.StdoutPipe()
	stderrPipe, _ := process.cmd.StderrPipe()

	stdoutScanner := bufio.NewScanner(stdoutPipe)
	stderrScanner := bufio.NewScanner(stderrPipe)

	go logger.ProcessOutput(stdoutScanner, "stdout")
	go logger.ProcessOutput(stderrScanner, "stderr")

	//@TODO: Find a better way to manage pipes

	// Start the init process in the container and add it to the cgroup
	if err := process.cmd.Start(); err != nil {
		return err
	}

	if err := stream.Send(&pb.CreateContainerResponse{
		Success: true,
		Message: "Container process started",
	}); err != nil {
		return fmt.Errorf("error sending progress: %w", err)
	}

	container.Manager.Add(process.cmd.Process.Pid)

	// Send final success messages
	if err := stream.Send(&pb.CreateContainerResponse{
		Success: true,
		Message: "Container created successfully",
	}); err != nil {
		return fmt.Errorf("error sending progress: %w", err)
	}

	return nil
}
