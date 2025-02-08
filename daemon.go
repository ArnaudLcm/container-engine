package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/arnaudlcm/container-engine/common/log"
	"github.com/arnaudlcm/container-engine/internal/core"
)

func main() {
	log.Info("Starting container engine Deamon")

	daemon := core.NewEngineDaemon()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	sigReceived := <-sigs
	fmt.Printf("Received signal: %s\n", sigReceived)
	daemon.Cleanup()
}
