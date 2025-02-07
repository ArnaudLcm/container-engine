package core

import (
	"io"
	"os"
)

type Process struct {
	Args              []string
	Cmd               string
	UID, GID          int
	Stdin             io.Reader
	Stdout            io.Reader
	CommunicationPipe *os.File // Used for communication between the container daemon and the process
}
