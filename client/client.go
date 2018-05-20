package client

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/wendorf/dory/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (d *DoryClient) CreateMemory(key string, value []byte, duration time.Duration) error {
	expiration, err := ptypes.TimestampProto(time.Now().Add(duration))
	if err != nil {
		return fmt.Errorf("invalid duration: %v", err)
	}
	mem := &memory.Memory{Name: key, Value: value, Expiration: expiration}

	_, err = d.client.CreateMemory(context.Background(), &memory.CreateMemoryRequest{Memory: mem})
	if err != nil {
		return fmt.Errorf("could not create memory: %v", err)
	}

	return nil
}

func (d *DoryClient) GetMemory(key string) ([]byte, error) {
	mem, err := d.client.GetMemory(context.Background(), &memory.GetMemoryRequest{Name: key})
	if err != nil {
		s, ok := status.FromError(err)
		if !ok {
			return nil, errors.New("could not get memory: server error")
		}

		switch s.Code() {
		case codes.FailedPrecondition:
			return nil, errors.New(s.Message())
		default:
			return nil, fmt.Errorf("could not get memory: %v", err)
		}
	}

	return mem.Value, nil
}
