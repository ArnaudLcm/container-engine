package core

import "github.com/google/uuid"

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
