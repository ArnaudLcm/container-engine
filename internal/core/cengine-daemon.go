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

type EngineDaemon struct {
	mu         sync.Mutex
	containers map[uuid.UUID]Container
	pb.DaemonServiceServer
}

const maxAttemptUUID int = 50

func NewEngineDaemon() *EngineDaemon {

	engineDaemon := EngineDaemon{
		containers: make(map[uuid.UUID]Container),
	}

	go runRPCServer(&engineDaemon)

	return &engineDaemon
}

func runRPCServer(g *EngineDaemon) {
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

func (g *EngineDaemon) CreateContainer() (Container, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	container := Container{}

	uuid, err := g.getUniqueUUID()
	if err != nil {
		return container, err
	}

	container.ID = uuid
	g.containers[uuid] = container

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
	g.containers[uuid] = container

	log.Debug("Current length %d", len(g.containers))

	if err := process.Start(); err != nil {
		return container, err
	}

	return container, nil
}

func (g *EngineDaemon) getUniqueUUID() (uuid.UUID, error) {
	for i := 0; i < maxAttemptUUID; i++ {
		newUUID := uuid.New()
		if _, exists := g.containers[newUUID]; !exists {
			return newUUID, nil
		}
	}

	return uuid.UUID{}, fmt.Errorf("can't find a unique UUID")
}
