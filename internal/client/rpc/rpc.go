package rpc

import (
	"context"
	"fmt"
	"time"

	pb "github.com/arnaudlcm/container-engine/service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type gRPCClient struct {
	Client pb.ContainerDaemonServiceClient
	Ctx    context.Context
	Cancel context.CancelFunc
}

func GetGRPCClient(ctx context.Context) (*gRPCClient, error) {
	value := ctx.Value("rpc")
	if value == nil {
		return nil, fmt.Errorf("gRPC client not found in context")
	}
	client, ok := value.(*gRPCClient)
	if !ok {
		return nil, fmt.Errorf("invalid gRPC client type")
	}
	return client, nil
}

func SetupGrpcClient() (context.Context, error) {
	// Establish gRPC connection
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC client: %v", err)
	}

	client := pb.NewContainerDaemonServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*180)

	grpcClient := &gRPCClient{
		Client: client,
		Ctx:    ctx,
		Cancel: cancel,
	}

	ctx = context.WithValue(ctx, "rpc", grpcClient)
	return ctx, nil
}
