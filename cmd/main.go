package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/arnaudlcm/container-engine/internal/cmd"
	pb "github.com/arnaudlcm/container-engine/service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}
	defer conn.Close()

	client := pb.NewDaemonServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	response, err := client.GetContainers(ctx, &pb.ContainersRequest{})
	if err != nil {
		log.Fatalf("Error getting status: %v", err)
	}
	fmt.Println("CONTAINER \t STATUS")
	for _, c := range response.Containers {
		fmt.Println("%s \t %d", c.Id, c.Status)

	}
	defer cancel()

	rootCmd := cmd.GetRootCommand()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(255)
	}
}
