package core

import (
	"context"

	pb "github.com/arnaudlcm/container-engine/service/proto"
	"github.com/google/uuid"
)

type ContainerStatus string

const (
	CONTAINER_RUNNING ContainerStatus = "running"
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

func (d *EngineDeamon) GetContainers(ctx context.Context, req *pb.ContainersRequest) (*pb.ContainersResponse, error) {

	containers := make([]*pb.ContainerInfos, 0, len(d.containers))

	for _, c := range d.containers {
		containers = append(containers, &pb.ContainerInfos{
			Id:     c.ID.String(),
			Status: convertStatusToProto(c.Status),
		})
	}

	return &pb.ContainersResponse{Containers: containers}, nil
}
