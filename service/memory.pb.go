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
	Value                []byte               `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Expiration           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=expiration" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Memory) Reset()         { *m = Memory{} }
func (m *Memory) String() string { return proto.CompactTextString(m) }
func (*Memory) ProtoMessage()    {}
func (*Memory) Descriptor() ([]byte, []int) {
	return fileDescriptor_memory_8f6b82fe3900ee84, []int{0}
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

func (m *Memory) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
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
	return fileDescriptor_memory_8f6b82fe3900ee84, []int{1}
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
	return fileDescriptor_memory_8f6b82fe3900ee84, []int{2}
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

func init() { proto.RegisterFile("memory.proto", fileDescriptor_memory_8f6b82fe3900ee84) }

var fileDescriptor_memory_8f6b82fe3900ee84 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x1b, 0x28, 0x41, 0xbd, 0x66, 0x00, 0x53, 0xa1, 0x28, 0x0c, 0x8d, 0x3c, 0xa0, 0xb0,
	0xb8, 0x52, 0x90, 0x18, 0x18, 0x58, 0x00, 0x31, 0xb1, 0x58, 0xbc, 0x80, 0x2b, 0x1d, 0x55, 0xa4,
	0x3a, 0x76, 0xdd, 0x0b, 0x82, 0xb7, 0x47, 0xb5, 0x1d, 0x54, 0xda, 0x6c, 0xbe, 0xf3, 0x77, 0xfe,
	0xfc, 0x1f, 0x64, 0x1a, 0xb5, 0x71, 0x3f, 0xc2, 0x3a, 0x43, 0xa6, 0x98, 0xaf, 0x8c, 0x59, 0xad,
	0x71, 0xe1, 0xab, 0x65, 0xf7, 0xb9, 0xa0, 0x46, 0xe3, 0x96, 0x94, 0xb6, 0x11, 0xb8, 0x39, 0x04,
	0x50, 0x5b, 0x8a, 0xd3, 0xbc, 0x85, 0xf4, 0xdd, 0xbf, 0xc6, 0x18, 0x8c, 0x5b, 0xa5, 0x31, 0x4f,
	0xca, 0xa4, 0x9a, 0x48, 0x7f, 0x66, 0x33, 0x38, 0xfb, 0x52, 0xeb, 0x0e, 0xf3, 0x93, 0x32, 0xa9,
	0x32, 0x19, 0x0a, 0xf6, 0x08, 0x80, 0xdf, 0xb6, 0x71, 0x8a, 0x1a, 0xd3, 0xe6, 0xa7, 0x65, 0x52,
	0x4d, 0xeb, 0x42, 0x04, 0x8b, 0xe8, 0x2d, 0xe2, 0xa3, 0xff, 0x86, 0xdc, 0xa3, 0xf9, 0x2d, 0x5c,
	0xbc, 0x21, 0x05, 0xa5, 0xc4, 0x4d, 0x87, 0x5b, 0x1a, 0x32, 0xf3, 0x07, 0xb8, 0x7a, 0x76, 0xa8,
	0x08, 0xff, 0xa3, 0x73, 0x48, 0x43, 0x78, 0x0f, 0x4f, 0xeb, 0x73, 0x11, 0xef, 0x63, 0xbb, 0xde,
	0xc0, 0xf8, 0x65, 0x97, 0xe6, 0x0e, 0x26, 0x7f, 0x1e, 0x76, 0x29, 0x0e, 0x9d, 0x45, 0x3f, 0xc8,
	0x47, 0xec, 0x09, 0xb2, 0x7d, 0x15, 0x9b, 0x89, 0x01, 0x73, 0x71, 0x7d, 0x14, 0xf0, 0x75, 0xb7,
	0x46, 0x3e, 0x5a, 0xa6, 0xbe, 0x73, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x71, 0x5c, 0xa3, 0x51,
	0x97, 0x01, 0x00, 0x00,
}
