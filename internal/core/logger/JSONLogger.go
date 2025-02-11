package logger

import (
	"bufio"
	"encoding/json"
	"os"
	"path"
	"time"

	"github.com/arnaudlcm/container-engine/internal/core/utils"
)

type JSONLogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Stream    string    `json:"stream"`
	Message   string    `json:"message"`
}

type JSONLogger struct {
	logFile *os.File
}

func NewLogger(containerID string) (Logger, error) {
	logPath := path.Join(utils.LIB_FOLDER_LOGS_PATH, containerID) + ".json"
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	return &JSONLogger{logFile: file}, nil // Return pointer
}

func (l *JSONLogger) GetStdOut() *os.File {
	return l.logFile
}
func (l *JSONLogger) GetStdErr() *os.File {
	return l.logFile
}

func (l *JSONLogger) ProcessOutput(scanner *bufio.Scanner, stream string) {
	for scanner.Scan() {
		content := scanner.Text()
		_ = l.WriteLog(content, stream)
	}
}

func (l *JSONLogger) WriteLog(message string, stream string) error {
	entry := JSONLogEntry{
		Timestamp: time.Now(),
		Stream:    stream,
		Message:   message,
	}

	logJSON, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	_, err = l.logFile.Write(append(logJSON, '\n'))
	return err
}
