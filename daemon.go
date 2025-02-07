package main

import (
	"github.com/arnaudlcm/container-engine/common/log"
	"github.com/arnaudlcm/container-engine/internal/core"
)

func main() {
	log.Info("Starting container engine Deamon")

	engine := core.NewEngineDaemon()
	if _, err := engine.CreateContainer(); err != nil {
		log.Error("%w", err)
	}
}
