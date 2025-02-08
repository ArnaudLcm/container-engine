package core

import (
	"fmt"
	"net"
	"sync"

	"github.com/arnaudlcm/container-engine/common/log"
	pb "github.com/arnaudlcm/container-engine/service/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type EngineDaemon struct {
	mu         sync.Mutex
	containers map[uuid.UUID]Container
	pb.ContainerDaemonServiceServer
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
	pb.RegisterContainerDaemonServiceServer(s, g)
	log.Info("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}
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
