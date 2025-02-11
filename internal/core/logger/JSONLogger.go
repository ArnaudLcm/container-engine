package logger

import (
	"bufio"
	"encoding/json"
	"os"
	"path"
	"sync"
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
	mu      sync.Mutex
}

func NewLogger(containerID string) (Logger, error) {
	logPath := path.Join(utils.LIB_FOLDER_LOGS_PATH, containerID) + ".json"
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	return &JSONLogger{logFile: file}, nil
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

	l.mu.Lock()
	defer l.mu.Unlock()
	_, err = l.logFile.Write(append(logJSON, '\n'))
	return err
}

func (l *JSONLogger) GetLastLogs(n int) ([]string, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	logPath := l.logFile.Name()
	file, err := os.Open(logPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	var logs []string
	var buffer []byte
	position := stat.Size()

	for position > 0 && len(logs) < n {
		position--
		file.Seek(position, 0)

		b := make([]byte, 1)
		_, err := file.Read(b)
		if err != nil {
			break
		}

		if b[0] == '\n' {
			line := reverseBytes(buffer)
			logs = append(logs, string(line))
			buffer = []byte{}
		} else {
			buffer = append(buffer, b[0])
		}
	}

	reverseSlice(logs)
	return logs, nil
}

func reverseBytes(data []byte) []byte {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return data
}

func reverseSlice(logs []string) {
	for i, j := 0, len(logs)-1; i < j; i, j = i+1, j-1 {
		logs[i], logs[j] = logs[j], logs[i]
	}
}
