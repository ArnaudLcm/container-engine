package core

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/moby/sys/reexec"
)

type Process struct {
	Args              []string
	UID, GID          int
	Stdin             io.Reader
	Stdout            io.Writer
	CommunicationPipe *os.File // Used for communication between the container daemon and the process
	rootPath          string
	workingDirectory  string
	cmd               *exec.Cmd
}

func init() {
	reexec.Register("pivot_root", pivotRootReexec)

	if reexec.Init() {
		os.Exit(0)
	}
}

func pivotRootReexec() {
	rootPath := os.Getenv("NEW_ROOT")

	if rootPath == "" {
		fmt.Fprintf(os.Stderr, "missing required environment variables\n")
		os.Exit(1)
	}

	putold := filepath.Join(rootPath, "/.pivot_root")

	if err := os.MkdirAll(putold, 0700); err != nil {
		fmt.Fprintf(os.Stderr, "error creating pivot_root dir: %v\n", err)
		os.Exit(1)
	}

	if err := syscall.Mount(rootPath, rootPath, "", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		fmt.Fprintf(os.Stderr, "error binding root path: %v\n", err)
		os.Exit(1)
	}

	if err := syscall.PivotRoot(rootPath, putold); err != nil {
		fmt.Fprintf(os.Stderr, "error pivoting root: %v\n", err)
		os.Exit(1)
	}

	if err := os.Chdir("/"); err != nil {
		fmt.Fprintf(os.Stderr, "error changing directory: %v\n", err)
		os.Exit(1)
	}

	if err := syscall.Mount("proc", "/proc", "proc", 0, ""); err != nil {
		fmt.Println("failed to mount /proc: ", err)
		os.Exit(1)
	}

	putold = "/.pivot_root"
	if err := syscall.Unmount(putold, syscall.MNT_DETACH); err != nil {
		fmt.Fprintf(os.Stderr, "error unmounting pivot root: %v\n", err)
		os.Exit(1)
	}

	if err := os.RemoveAll(putold); err != nil {
		fmt.Fprintf(os.Stderr, "error removing pivot_root directory: %v\n", err)
		os.Exit(1)
	}

	runSubCommand()

}

func runSubCommand() error {
	execCmdStr := os.Getenv("EXEC_CMD")
	if execCmdStr == "" {
		fmt.Fprintf(os.Stderr, "error: EXEC_CMD not set\n")
		os.Exit(1)
	}
	execArgsStr := os.Getenv("EXEC_ARGS")
	execArgs := []string{}
	if execArgsStr != "" {
		execArgs = strings.Split(execArgsStr, " ")
	}
	execCmd := exec.Command(execCmdStr, execArgs...)

	execCmd.Stdin = os.Stdin
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stdout

	if err := execCmd.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "error starting exec command: %v\n", err)
		os.Exit(1)
	}

	if err := execCmd.Wait(); err != nil {
		fmt.Fprintf(os.Stderr, "error waiting for exec command: %v\n", err)
		os.Exit(1)
	}
	return nil
}

func (p *Process) Init() error {
	cmd := reexec.Command("pivot_root")
	cmd.Env = append(os.Environ(),
		"NEW_ROOT="+p.rootPath,
		"WORKING_DIR="+p.workingDirectory,
		"EXEC_CMD="+p.Args[0],
		"EXEC_ARGS="+strings.Join(p.Args[1:], " "),
	)
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

	p.cmd = cmd

	return nil
}
