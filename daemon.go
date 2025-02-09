package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/arnaudlcm/container-engine/common/log"
	"github.com/arnaudlcm/container-engine/internal/core"
)

func main() {
	log.Info("Starting container engine Deamon")

	keyPath := flag.String("key", "", "Path to the ECC public key")

	flag.Parse()

	if *keyPath == "" {
		fmt.Println("Usage: go run daemon.go -key=PathToTheKey")
		return
	}

	key, err := core.LoadECCPublicKey(*keyPath)
	if err != nil {
		log.Fatal("%w", err)
	}

	daemon := core.NewEngineDaemon(key)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	sigReceived := <-sigs
	fmt.Printf("Received signal: %s\n", sigReceived)
	daemon.Cleanup()
}
