package client

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/wendorf/dory/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)


type DoryClient struct {
	client memory.DoryClient
}

func NewClient(socketPath string) (*DoryClient, error) {
	conn, err := grpc.Dial(fmt.Sprintf("passthrough:///unix://%s", socketPath), grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("could not dial socket: %v", err)
	}

	return &DoryClient{
		client: memory.NewDoryClient(conn),
	}, nil
}

func (d *DoryClient) CreateMemory(key, value string, duration time.Duration) error {
	expiration, err := ptypes.TimestampProto(time.Now().Add(duration))
	if err != nil {
		return fmt.Errorf("could not create timestamp: %v", err)
	}
	mem := &memory.Memory{Name: key, Value: value, Expiration: expiration}

	_, err = d.client.CreateMemory(context.Background(), &memory.CreateMemoryRequest{Memory: mem})
	if err != nil {
		return fmt.Errorf("failed create memory: %v\n", err)
	}

	return nil
}

func (d *DoryClient) GetMemory(key string) (string, error) {
	mem, err := d.client.GetMemory(context.Background(), &memory.GetMemoryRequest{Name: key})
	if err != nil {
		return "", fmt.Errorf("failed get memory: %v\n", err)
	}

	return mem.Value, nil
}