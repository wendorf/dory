// Code generated by protoc-gen-go. DO NOT EDIT.
// source: memory.proto

package memory

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Memory struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value                string               `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Expiration           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=expiration" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Memory) Reset()         { *m = Memory{} }
func (m *Memory) String() string { return proto.CompactTextString(m) }
func (*Memory) ProtoMessage()    {}
func (*Memory) Descriptor() ([]byte, []int) {
	return fileDescriptor_memory_46b5bbb990324aed, []int{0}
}
func (m *Memory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Memory.Unmarshal(m, b)
}
func (m *Memory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Memory.Marshal(b, m, deterministic)
}
func (dst *Memory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Memory.Merge(dst, src)
}
func (m *Memory) XXX_Size() int {
	return xxx_messageInfo_Memory.Size(m)
}
func (m *Memory) XXX_DiscardUnknown() {
	xxx_messageInfo_Memory.DiscardUnknown(m)
}

var xxx_messageInfo_Memory proto.InternalMessageInfo

func (m *Memory) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Memory) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Memory) GetExpiration() *timestamp.Timestamp {
	if m != nil {
		return m.Expiration
	}
	return nil
}

type GetMemoryRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMemoryRequest) Reset()         { *m = GetMemoryRequest{} }
func (m *GetMemoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetMemoryRequest) ProtoMessage()    {}
func (*GetMemoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_memory_46b5bbb990324aed, []int{1}
}
func (m *GetMemoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMemoryRequest.Unmarshal(m, b)
}
func (m *GetMemoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMemoryRequest.Marshal(b, m, deterministic)
}
func (dst *GetMemoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMemoryRequest.Merge(dst, src)
}
func (m *GetMemoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetMemoryRequest.Size(m)
}
func (m *GetMemoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMemoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMemoryRequest proto.InternalMessageInfo

func (m *GetMemoryRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateMemoryRequest struct {
	Memory               *Memory  `protobuf:"bytes,1,opt,name=memory" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMemoryRequest) Reset()         { *m = CreateMemoryRequest{} }
func (m *CreateMemoryRequest) String() string { return proto.CompactTextString(m) }
func (*CreateMemoryRequest) ProtoMessage()    {}
func (*CreateMemoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_memory_46b5bbb990324aed, []int{2}
}
func (m *CreateMemoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMemoryRequest.Unmarshal(m, b)
}
func (m *CreateMemoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMemoryRequest.Marshal(b, m, deterministic)
}
func (dst *CreateMemoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMemoryRequest.Merge(dst, src)
}
func (m *CreateMemoryRequest) XXX_Size() int {
	return xxx_messageInfo_CreateMemoryRequest.Size(m)
}
func (m *CreateMemoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMemoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMemoryRequest proto.InternalMessageInfo

func (m *CreateMemoryRequest) GetMemory() *Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func init() {
	proto.RegisterType((*Memory)(nil), "Memory")
	proto.RegisterType((*GetMemoryRequest)(nil), "GetMemoryRequest")
	proto.RegisterType((*CreateMemoryRequest)(nil), "CreateMemoryRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DoryClient is the client API for Dory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DoryClient interface {
	GetMemory(ctx context.Context, in *GetMemoryRequest, opts ...grpc.CallOption) (*Memory, error)
	CreateMemory(ctx context.Context, in *CreateMemoryRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type doryClient struct {
	cc *grpc.ClientConn
}

func NewDoryClient(cc *grpc.ClientConn) DoryClient {
	return &doryClient{cc}
}

func (c *doryClient) GetMemory(ctx context.Context, in *GetMemoryRequest, opts ...grpc.CallOption) (*Memory, error) {
	out := new(Memory)
	err := c.cc.Invoke(ctx, "/Dory/GetMemory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doryClient) CreateMemory(ctx context.Context, in *CreateMemoryRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Dory/CreateMemory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DoryServer is the server API for Dory service.
type DoryServer interface {
	GetMemory(context.Context, *GetMemoryRequest) (*Memory, error)
	CreateMemory(context.Context, *CreateMemoryRequest) (*empty.Empty, error)
}

func RegisterDoryServer(s *grpc.Server, srv DoryServer) {
	s.RegisterService(&_Dory_serviceDesc, srv)
}

func _Dory_GetMemory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoryServer).GetMemory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dory/GetMemory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoryServer).GetMemory(ctx, req.(*GetMemoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dory_CreateMemory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMemoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoryServer).CreateMemory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dory/CreateMemory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoryServer).CreateMemory(ctx, req.(*CreateMemoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dory_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Dory",
	HandlerType: (*DoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMemory",
			Handler:    _Dory_GetMemory_Handler,
		},
		{
			MethodName: "CreateMemory",
			Handler:    _Dory_CreateMemory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "memory.proto",
}

func init() { proto.RegisterFile("memory.proto", fileDescriptor_memory_46b5bbb990324aed) }

var fileDescriptor_memory_46b5bbb990324aed = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x4d, 0xcd, 0xcd,
	0x2f, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x92, 0x4f, 0xcf, 0xcf, 0x4f, 0xcf, 0x49,
	0xd5, 0x07, 0xf3, 0x92, 0x4a, 0xd3, 0xf4, 0x4b, 0x32, 0x73, 0x53, 0x8b, 0x4b, 0x12, 0x73, 0x0b,
	0xa0, 0x0a, 0xa4, 0xd1, 0x15, 0xa4, 0xe6, 0x16, 0x94, 0x40, 0x75, 0x2b, 0xe5, 0x71, 0xb1, 0xf9,
	0x82, 0x4d, 0x13, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x02, 0xb3, 0x85, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x98, 0xc0, 0x82,
	0x10, 0x8e, 0x90, 0x15, 0x17, 0x57, 0x6a, 0x45, 0x41, 0x66, 0x51, 0x62, 0x49, 0x66, 0x7e, 0x9e,
	0x04, 0xb3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x94, 0x1e, 0xc4, 0x16, 0x3d, 0x98, 0x2d, 0x7a, 0x21,
	0x30, 0x67, 0x04, 0x21, 0xa9, 0x56, 0x52, 0xe3, 0x12, 0x70, 0x4f, 0x2d, 0x81, 0x58, 0x19, 0x94,
	0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x82, 0xcd, 0x66, 0x25, 0x33, 0x2e, 0x61, 0xe7, 0xa2, 0xd4, 0xc4,
	0x92, 0x54, 0x54, 0xa5, 0xf2, 0x5c, 0x6c, 0x10, 0xcf, 0x83, 0x15, 0x73, 0x1b, 0xb1, 0xeb, 0x41,
	0xe5, 0xa1, 0xc2, 0x46, 0x85, 0x5c, 0x2c, 0x2e, 0x20, 0xdf, 0x68, 0x72, 0x71, 0xc2, 0xed, 0x11,
	0x12, 0xd4, 0x43, 0xb7, 0x53, 0x0a, 0xa6, 0x51, 0x89, 0x41, 0xc8, 0x8e, 0x8b, 0x07, 0xd9, 0x2a,
	0x21, 0x11, 0x3d, 0x2c, 0x36, 0x4b, 0x89, 0x61, 0x78, 0xd0, 0x15, 0x14, 0x8c, 0x4a, 0x0c, 0x49,
	0x6c, 0x60, 0x11, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x2e, 0x4a, 0x9e, 0x97, 0x01,
	0x00, 0x00,
}
