package core

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

type Process struct {
	Args              []string
	UID, GID          int
	Stdin             io.Reader
	Stdout            io.Writer
	CommunicationPipe *os.File // Used for communication between the container daemon and the process
}

func (p *Process) Start() error {
	cmd := exec.Command(strings.Join(p.Args, " "))
	cmd.Stdin = p.Stdin
	cmd.Stdout = p.Stdout
	cmd.Stderr = p.Stdout

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

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("error starting the exec.Command - %w", err)
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("error starting the exec.Command - %w", err)
	}

	return nil
}
