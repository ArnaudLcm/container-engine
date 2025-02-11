package logger

import (
	"bufio"
	"os"
)

type Logger interface {
	GetStdOut() *os.File
	GetStdErr() *os.File
	ProcessOutput(scanner *bufio.Scanner, stream string)
	WriteLog(message string, stream string) error
}
