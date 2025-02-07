package core

import (
	"fmt"
	"os"

	"github.com/google/uuid"
)

const cGroupv2Path string = "/sys/fs/cgroup"

type CGroupManager struct {
	cGroupPath string
}

func NewCGroupManager(containerUUID uuid.UUID) (CGroupManager, error) {

	manager := CGroupManager{}
	groupPath := fmt.Sprintf("%s/%s", cGroupv2Path, containerUUID)
	if err := os.MkdirAll(groupPath, 0644); err != nil {
		return manager, fmt.Errorf("failed to create cgroup for container %s: %w", containerUUID, err)
	}

	manager.cGroupPath = groupPath

	return manager, nil

}

func (c *CGroupManager) Add(pid int) error {
	procsPath := fmt.Sprintf("%s/%s", c.cGroupPath, "cpu.procs")

	file, err := os.OpenFile(procsPath, os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("%d", pid)); err != nil {
		return fmt.Errorf("failed to setup cgroup for the container: %w", err)
	}

	return nil

}

func (c *CGroupManager) Remove(pid int) error {
	return nil
}
