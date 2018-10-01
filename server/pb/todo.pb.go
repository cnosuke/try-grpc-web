// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

/*
Package todo is a generated protocol buffer package.

It is generated from these files:
	todo.proto

It has these top-level messages:
	GetLatestRequest
	GetLatestResponse
	Todo
*/
package todo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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

type GetLatestRequest struct {
}

func (m *GetLatestRequest) Reset()                    { *m = GetLatestRequest{} }
func (m *GetLatestRequest) String() string            { return proto.CompactTextString(m) }
func (*GetLatestRequest) ProtoMessage()               {}
func (*GetLatestRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetLatestResponse struct {
	Todo *Todo `protobuf:"bytes,1,opt,name=todo" json:"todo,omitempty"`
}

func (m *GetLatestResponse) Reset()                    { *m = GetLatestResponse{} }
func (m *GetLatestResponse) String() string            { return proto.CompactTextString(m) }
func (*GetLatestResponse) ProtoMessage()               {}
func (*GetLatestResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetLatestResponse) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type Todo struct {
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Text      string                     `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
}

func (m *Todo) Reset()                    { *m = Todo{} }
func (m *Todo) String() string            { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()               {}
func (*Todo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Todo) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Todo) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*GetLatestRequest)(nil), "todo.GetLatestRequest")
	proto.RegisterType((*GetLatestResponse)(nil), "todo.GetLatestResponse")
	proto.RegisterType((*Todo)(nil), "todo.Todo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TodoService service

type TodoServiceClient interface {
	GetLatest(ctx context.Context, in *GetLatestRequest, opts ...grpc.CallOption) (*GetLatestResponse, error)
}

type todoServiceClient struct {
	cc *grpc.ClientConn
}

func NewTodoServiceClient(cc *grpc.ClientConn) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) GetLatest(ctx context.Context, in *GetLatestRequest, opts ...grpc.CallOption) (*GetLatestResponse, error) {
	out := new(GetLatestResponse)
	err := grpc.Invoke(ctx, "/todo.TodoService/GetLatest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TodoService service

type TodoServiceServer interface {
	GetLatest(context.Context, *GetLatestRequest) (*GetLatestResponse, error)
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_GetLatest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).GetLatest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/GetLatest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).GetLatest(ctx, req.(*GetLatestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLatest",
			Handler:    _TodoService_GetLatest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0x31, 0x4f, 0x84, 0x40,
	0x10, 0x85, 0x83, 0x21, 0x26, 0x0c, 0x8d, 0x4e, 0xa1, 0x84, 0x42, 0x09, 0x15, 0xd5, 0x92, 0x40,
	0x63, 0x61, 0x6f, 0xa1, 0xd5, 0xca, 0x1f, 0x00, 0x19, 0x09, 0x89, 0x38, 0x1c, 0x3b, 0x5c, 0xee,
	0xe7, 0x5f, 0xd8, 0x3d, 0xb8, 0xcb, 0x5d, 0x37, 0xfb, 0xf2, 0xe5, 0xbd, 0x6f, 0x01, 0x84, 0x5b,
	0x56, 0xe3, 0xc4, 0xc2, 0xe8, 0x2f, 0x77, 0xfc, 0xda, 0x31, 0x77, 0x7f, 0x94, 0xdb, 0xac, 0x99,
	0x7f, 0x73, 0xe9, 0x07, 0x32, 0x52, 0x0f, 0xa3, 0xc3, 0x52, 0x84, 0x87, 0x0f, 0x92, 0xaf, 0x5a,
	0xc8, 0x88, 0xa6, 0xdd, 0x4c, 0x46, 0xd2, 0x12, 0x1e, 0x2f, 0x32, 0x33, 0xf2, 0xbf, 0x21, 0x7c,
	0x01, 0xdb, 0x18, 0x79, 0x89, 0x97, 0x85, 0x05, 0x28, 0x3b, 0x55, 0x71, 0xcb, 0xda, 0xe6, 0x69,
	0x05, 0xfe, 0xf2, 0xc2, 0x37, 0x08, 0xb6, 0x8d, 0x13, 0x1c, 0x2b, 0x67, 0xa1, 0x56, 0x0b, 0x55,
	0xad, 0x84, 0x3e, 0xc3, 0x88, 0xe0, 0x0b, 0x1d, 0x24, 0xba, 0x4b, 0xbc, 0x2c, 0xd0, 0xf6, 0x2e,
	0x3e, 0x21, 0x5c, 0x5a, 0xbf, 0x69, 0xda, 0xf7, 0x3f, 0x84, 0xef, 0x10, 0x6c, 0x66, 0xf8, 0xe4,
	0x1c, 0xae, 0xf5, 0xe3, 0xe7, 0x9b, 0xdc, 0x7d, 0xa1, 0xb9, 0xb7, 0xfb, 0xe5, 0x31, 0x00, 0x00,
	0xff, 0xff, 0xe4, 0xc3, 0x7a, 0x72, 0x27, 0x01, 0x00, 0x00,
}
