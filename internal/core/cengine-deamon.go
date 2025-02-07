package core

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/arnaudlcm/container-engine/service/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type EngineDeamon struct {
	containers map[uuid.UUID]Container
	pb.DaemonServiceServer
}

const maxAttemptUUID int = 50

func NewEngineDeamon() *EngineDeamon {

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDaemonServiceServer(s, &EngineDeamon{})
	log.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	return &EngineDeamon{
		containers: make(map[uuid.UUID]Container),
	}
}

func (d *EngineDeamon) CreateContainer() (Container, error) {
	container := Container{}

	uuid, err := d.getUniqueUUID()
	if err != nil {
		return container, err
	}

	container.ID = uuid
	d.containers[uuid] = container

	container.Manager, err = NewCGroupManager(container.ID)
	if err != nil {
		return container, fmt.Errorf("error during CGroupManager creation: %w", err)
	}

	process := Process{
		Args:              []string{"/bin/bash"},
		Stdin:             os.Stdin,
		Stdout:            os.Stdout,
		CommunicationPipe: nil,
		UID:               0,
		GID:               0,
	}

	container.Process = process
	container.Status = CONTAINER_RUNNING
	if err := process.Start(); err != nil {
		return container, err
	}

	return container, nil
}

func (d *EngineDeamon) getUniqueUUID() (uuid.UUID, error) {
	for i := 0; i < maxAttemptUUID; i++ {
		newUUID := uuid.New()
		if _, exists := d.containers[newUUID]; !exists {
			return newUUID, nil
		}
	}

	return uuid.UUID{}, fmt.Errorf("can't find a unique UUID")
}
