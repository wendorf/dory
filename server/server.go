package server

import (
	"fmt"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/wendorf/dory/fish"
	"github.com/wendorf/dory/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DoryServer struct {
	socketPath string
	memories   *fish.Brain
}

func NewDoryServer(socketPath string) *DoryServer {
	server := &DoryServer{
		socketPath: socketPath,
		memories:   fish.NewBrain(),
	}

	return server
}

func (d *DoryServer) Run() (error) {
	listener, err := NewListener(d.socketPath)
	if err != nil {
		return fmt.Errorf("could not listen to socket: %v", err)
	}

	grpcServer := grpc.NewServer()
	memory.RegisterDoryServer(grpcServer, d)
	return grpcServer.Serve(listener)
}

func (d *DoryServer) GetMemory(ctx context.Context, request *memory.GetMemoryRequest) (*memory.Memory, error) {
	val, err := d.memories.Get(request.Name)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "memory does not exist: %s", request.Name)
	}

	return &memory.Memory{Value: string(val)}, nil
}

func (d *DoryServer) CreateMemory(ctx context.Context, request *memory.CreateMemoryRequest) (*empty.Empty, error) {
	if request.Memory.Expiration == nil {
		return nil, status.Error(codes.InvalidArgument, "memory must expire")
	}

	expiration, err := ptypes.Timestamp(request.Memory.Expiration)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid expiration: %s", request.Memory.Expiration)
	}
	err = d.memories.Set(request.Memory.Name, []byte(request.Memory.Value), expiration)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid memory value: %s", request.Memory.Value)
	}
	return &empty.Empty{}, nil
}
