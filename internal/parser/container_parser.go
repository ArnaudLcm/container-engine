package parser

import (
	"bufio"
	"os"
	"strings"
)

type ContainerConfig struct {
	Layer   string
	Cmd     []string
	WorkDir string
}

func ParseContainerConfig(filename string) (ContainerConfig, error) {
	var result ContainerConfig

	file, err := os.Open(filename)
	if err != nil {
		return result, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction := strings.TrimSpace(scanner.Text())

		switch {
		case strings.HasPrefix(instruction, "LAYER"):
			result.Layer = strings.TrimSpace(strings.TrimPrefix(instruction, "LAYER"))
		case strings.HasPrefix(instruction, "CMD"):
			cmdContent := strings.TrimSpace(strings.TrimPrefix(instruction, "CMD"))
			cmdContent = strings.Trim(cmdContent, "[]\"")
			result.Cmd = strings.Split(cmdContent, " ")
		case strings.HasPrefix(instruction, "WORKDIR"):
			result.WorkDir = strings.TrimSpace(strings.TrimPrefix(instruction, "WORKDIR"))
		}
	}

	if err := scanner.Err(); err != nil {
		return result, err
	}

	return result, nil
}
