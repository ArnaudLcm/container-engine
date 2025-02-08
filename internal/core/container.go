package core

import (
	"context"
	"fmt"
	"os"

	pb "github.com/arnaudlcm/container-engine/service/proto"
	"github.com/google/uuid"
)

type ContainerStatus string

const (
	CONTAINER_RUNNING ContainerStatus = "running"
	CONTAINER_HANGING ContainerStatus = "hanging"
	CONTAINER_KILLED  ContainerStatus = "killed"
)

type Container struct {
	ID         uuid.UUID
	RootFs     string // Path to the root fs
	Status     ContainerStatus
	Process    Process
	Manager    CGroupManager
	Namespaces map[NamespaceIdentifier]string // List of namespaces attached to the container with their paths
}

func (d *EngineDaemon) GetContainers(ctx context.Context, req *pb.GetContainersRequest) (*pb.GetContainersResponse, error) {

	containers := make([]*pb.ContainerInfos, 0, len(d.containers))

	for _, c := range d.containers {
		containers = append(containers, &pb.ContainerInfos{
			Id:     c.ID.String(),
			Status: convertStatusToProto(c.Status),
		})
	}

	return &pb.GetContainersResponse{Containers: containers}, nil
}

func (g *EngineDaemon) CreateContainer(ctx context.Context, req *pb.CreateContainerRequest) (*pb.CreateContainerResponse, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	container := Container{}

	uuid, err := g.getUniqueUUID()
	if err != nil {
		return &pb.CreateContainerResponse{Success: false}, err
	}

	container.ID = uuid
	g.containers[uuid] = container

	container.Manager, err = NewCGroupManager(container.ID)
	if err != nil {
		return &pb.CreateContainerResponse{Success: false}, fmt.Errorf("error during CGroupManager creation: %w", err)
	}

	process := Process{
		Args:              req.Config.Cmd,
		Stdin:             os.Stdin,
		Stdout:            os.Stdout,
		CommunicationPipe: nil,
		UID:               0,
		GID:               0,
	}

	container.Process = process
	container.Status = CONTAINER_KILLED
	g.containers[uuid] = container
	return &pb.CreateContainerResponse{Success: true}, nil
}
