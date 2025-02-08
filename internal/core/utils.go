package core

import (
	pb "github.com/arnaudlcm/container-engine/service/proto"
)

func ConvertContainerStatusToString(status pb.ContainerStatus) string {
	switch status {
	case pb.ContainerStatus_CONTAINER_RUNNING:
		return "Running"
	case pb.ContainerStatus_CONTAINER_HANGING:
		return "Hanging"
	case pb.ContainerStatus_CONTAINER_KILLED:
		return "Killed"
	default:
		return "Unknown"
	}
}
