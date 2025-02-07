package core

import (
	pb "github.com/arnaudlcm/container-engine/service/proto"
)

// Convert internal ContainerStatus to gRPC ContainerStatus
func convertStatusToProto(status ContainerStatus) pb.ContainerStatus {
	switch status {
	case CONTAINER_RUNNING:
		return pb.ContainerStatus_CONTAINER_RUNNING
	case CONTAINER_KILLED:
		return pb.ContainerStatus_CONTAINER_KILLED
	default:
		return pb.ContainerStatus_CONTAINER_UNKNOWN
	}
}
