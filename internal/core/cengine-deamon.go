package core

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/google/uuid"
)

type EngineDeamon struct {
	containers map[uuid.UUID]Container
}

const maxAttemptUUID int = 50

func NewEngineDeamon() *EngineDeamon {
	return &EngineDeamon{
		containers: make(map[uuid.UUID]Container),
	}
}

func (d *EngineDeamon) CreateContainer() (Container, error) {
	container := Container{}

	uuid, err := d.getUniqueUUID()
	if err != nil {
		return container, err
	}

	container.ID = uuid
	d.containers[uuid] = container

	cmd := exec.Command("/bin/bash")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWUSER | syscall.CLONE_NEWNS,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getuid(),
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getgid(),
				Size:        1,
			},
		},
	}

	container.Manager, err = NewCGroupManager(container.ID)
	if err != nil {
		return container, fmt.Errorf("error during CGroupManager creation: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return container, fmt.Errorf("error starting the exec.Command - %w", err)
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
