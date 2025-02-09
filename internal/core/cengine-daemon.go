package core

import (
	"crypto/ecdsa"
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
	pb.ContainerDaemonServiceServer
	fsManager   *FSManager
	manifestKey *ecdsa.PublicKey
}

const maxAttemptUUID int = 50
const LIB_FOLDER_PATH = "/var/lib/cengine"

func NewEngineDaemon(key *ecdsa.PublicKey) *EngineDaemon {

	engineDaemon := EngineDaemon{
		containers:  make(map[uuid.UUID]Container),
		fsManager:   NewFSManager(),
		manifestKey: key,
	}

	// Setup working directory
	os.MkdirAll(LIB_FOLDER_PATH, 0755)
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

func (g *EngineDaemon) Cleanup() {
	g.fsManager.CleanUp()
}
