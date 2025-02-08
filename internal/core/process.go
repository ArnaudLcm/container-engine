package core

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

type Process struct {
	Args              []string
	UID, GID          int
	Stdin             io.Reader
	Stdout            io.Writer
	CommunicationPipe *os.File // Used for communication between the container daemon and the process
	rootPath          string
	workingDirectory  string
}

func (p *Process) PivotRoot() error {

	putold := filepath.Join(p.rootPath, "/.pivot_root")

	err := os.MkdirAll(putold, 0700)
	if err != nil {
		return err
	}

	err = syscall.Mount(p.rootPath, p.rootPath, "", syscall.MS_BIND|syscall.MS_REC, "")
	if err != nil {
		return err
	}

	err = syscall.PivotRoot(p.rootPath, putold)
	if err != nil {
		return fmt.Errorf("error occured while pivoting root: %w", err)
	}

	err = os.Chdir(filepath.Join("/", p.workingDirectory))
	if err != nil {
		return fmt.Errorf("error occured during chdir to working dir: %w", err)
	}

	putold = "/.pivot_root"
	err = syscall.Unmount(putold, syscall.MNT_DETACH)
	if err != nil {
		return fmt.Errorf("error occured during unmounting of pivot root: %w", err)
	}

	err = os.RemoveAll(putold)
	if err != nil {
		return err
	}
	return nil
}

func (p *Process) Start() error {
	cmd := exec.Command(strings.Join(p.Args, " "))
	cmd.Stdin = p.Stdin
	cmd.Stdout = p.Stdout
	cmd.Stderr = p.Stdout

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS | syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNET | syscall.CLONE_NEWUSER,
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

	if err := p.PivotRoot(); err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("error starting the exec.Command - %w", err)
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("error starting the exec.Command - %w", err)
	}

	return nil
}
