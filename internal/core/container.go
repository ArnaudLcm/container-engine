package core

import (
	"context"
	"fmt"
	"os"

	pb "github.com/arnaudlcm/container-engine/service/proto"
	"github.com/google/uuid"
)

type Container struct {
	ID      uuid.UUID
	RootFs  string // Path to the root fs
	Status  pb.ContainerStatus
	Process Process
	Manager CGroupManager
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
	g.containers[uuid] = container

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

	path, err := g.fsManager.AddLayer(req.Config.Env, uuid.String())
	if err != nil {
		if err := stream.Send(&pb.CreateContainerResponse{
			Success: false,
			Message: "Error during layer setup",
		}); err != nil {
			return fmt.Errorf("error sending progress: %w", err)
		}
		return err
	}

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

	if err := stream.Send(&pb.CreateContainerResponse{
		Success: true,
		Message: "Container process started",
	}); err != nil {
		return fmt.Errorf("error sending progress: %w", err)
	}

	go process.Start()

	// Send final success message
	if err := stream.Send(&pb.CreateContainerResponse{
		Success: true,
		Message: "Container created successfully",
	}); err != nil {
		return fmt.Errorf("error sending progress: %w", err)
	}

	return nil
}
