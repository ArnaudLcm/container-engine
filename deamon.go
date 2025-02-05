package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/arnaudlcm/container-engine/common/log"
)

const (
	cgroupMemoryMountPoint = "/sys/fs/cgroup/memory/"
	cgroupName             = "offshore-engine"
	memoryLimit            = 100 * 1024 * 1024 // 100MB
)

func setupCgroup(pid string) error {
	cgroupPath := filepath.Join(cgroupMemoryMountPoint, cgroupName)

	// Create the cgroup directory
	if err := os.MkdirAll(cgroupPath, 0755); err != nil {
		return fmt.Errorf("failed to create cgroup: %w", err)
	}

	// Write memory limit
	if err := os.WriteFile(filepath.Join(cgroupPath, "memory.limit_in_bytes"),
		[]byte(strconv.Itoa(memoryLimit)), 0644); err != nil {
		return fmt.Errorf("failed to set memory limit: %w", err)
	}

	// Add the process to the cgroup
	if err := os.WriteFile(filepath.Join(cgroupPath, "cgroup.procs"), []byte(pid), 0644); err != nil {
		return fmt.Errorf("failed to add process to cgroup: %w", err)
	}

	return nil
}

func setupPivotRoot() {
	root := "/rootfs"
	if err := syscall.Mount(root, root, "", syscall.MS_BIND, ""); err != nil {
		fmt.Println("Mount error:", err)
	}
	runDir := filepath.Join(root, ".pivot_root")
	os.Mkdir(runDir, 0700)
	if err := syscall.PivotRoot(root, runDir); err != nil {
		fmt.Println("PivotRoot error:", err)
	}
	os.Chdir("/")
	os.Remove("/.pivot_root")
}

func setupNetworking() {
	exec.Command("ip", "netns", "add", "container_net").Run()
	exec.Command("ip", "link", "add", "veth0", "type", "veth", "peer", "name", "veth1").Run()
	exec.Command("ip", "link", "set", "veth1", "netns", "container_net").Run()
	exec.Command("ip", "addr", "add", "192.168.1.2/24", "dev", "veth0").Run()
	exec.Command("ip", "link", "set", "veth0", "up").Run()
}

func main() {
	log.Info("Starting container engine Deamon")
}
