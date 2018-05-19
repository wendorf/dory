package server

import (
	"fmt"
	"net"
	"os"
	"syscall"
)

func NewListener(socketPath string) (net.Listener, error) {
	_, err := os.Lstat(socketPath)
	if err != nil {
		fmt.Printf("could not find socket: %s\n", socketPath)
	} else {
		err = syscall.Unlink(socketPath)
		if err != nil {
			return nil, fmt.Errorf("could not unlink socket: %v", err)
		}
	}

	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		return nil, fmt.Errorf("could not listen on socket: %v", err)
	}

	fmt.Println("Began listening")

	return listener, nil
}