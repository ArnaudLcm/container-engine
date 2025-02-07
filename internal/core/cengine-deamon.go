package core

import (
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/arnaudlcm/container-engine/common/log"
	pb "github.com/arnaudlcm/container-engine/service/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type EngineDeamon struct {
	mu         sync.Mutex
	containers map[uuid.UUID]Container
	pb.DaemonServiceServer
}

const maxAttemptUUID int = 50

func NewEngineDeamon() *EngineDeamon {

	engineDeamon := EngineDeamon{
		containers: make(map[uuid.UUID]Container),
	}

	go runRPCServer(&engineDeamon)

	return &engineDeamon
}

func (g *EngineDeamon) LenContaienrs() int {
	return len(g.containers)

}
func runRPCServer(g *EngineDeamon) {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDaemonServiceServer(s, g)
	log.Info("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}
}

func (d *EngineDeamon) CreateContainer() (Container, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
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
	d.containers[uuid] = container

	log.Debug("Current length %d", len(d.containers))

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
