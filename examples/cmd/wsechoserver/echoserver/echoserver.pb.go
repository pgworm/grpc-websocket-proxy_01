// Code generated by protoc-gen-go.
// source: echoserver.proto
// DO NOT EDIT!

/*
Package echoserver is a generated protocol buffer package.

It is generated from these files:
	echoserver.proto

It has these top-level messages:
	EchoRequest
	EchoResponse
	Heartbeat
	Empty
*/
package echoserver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

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

type Heartbeat_Status int32

const (
	Heartbeat_UNKNOWN Heartbeat_Status = 0
	Heartbeat_OK      Heartbeat_Status = 1
)

var Heartbeat_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "OK",
}
var Heartbeat_Status_value = map[string]int32{
	"UNKNOWN": 0,
	"OK":      1,
}

func (x Heartbeat_Status) String() string {
	return proto.EnumName(Heartbeat_Status_name, int32(x))
}
func (Heartbeat_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type EchoRequest struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *EchoRequest) Reset()                    { *m = EchoRequest{} }
func (m *EchoRequest) String() string            { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()               {}
func (*EchoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type EchoResponse struct {
	Message   string                      `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Timestamp *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *EchoResponse) Reset()                    { *m = EchoResponse{} }
func (m *EchoResponse) String() string            { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()               {}
func (*EchoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EchoResponse) GetTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type Heartbeat struct {
	Status Heartbeat_Status `protobuf:"varint,1,opt,name=status,enum=echoserver.Heartbeat_Status" json:"status,omitempty"`
}

func (m *Heartbeat) Reset()                    { *m = Heartbeat{} }
func (m *Heartbeat) String() string            { return proto.CompactTextString(m) }
func (*Heartbeat) ProtoMessage()               {}
func (*Heartbeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*EchoRequest)(nil), "echoserver.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "echoserver.EchoResponse")
	proto.RegisterType((*Heartbeat)(nil), "echoserver.Heartbeat")
	proto.RegisterType((*Empty)(nil), "echoserver.Empty")
	proto.RegisterEnum("echoserver.Heartbeat_Status", Heartbeat_Status_name, Heartbeat_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for EchoService service

type EchoServiceClient interface {
	Echo(ctx context.Context, opts ...grpc.CallOption) (EchoService_EchoClient, error)
	Stream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (EchoService_StreamClient, error)
	Heartbeats(ctx context.Context, opts ...grpc.CallOption) (EchoService_HeartbeatsClient, error)
}

type echoServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoServiceClient(cc *grpc.ClientConn) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Echo(ctx context.Context, opts ...grpc.CallOption) (EchoService_EchoClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EchoService_serviceDesc.Streams[0], c.cc, "/echoserver.EchoService/Echo", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServiceEchoClient{stream}
	return x, nil
}

type EchoService_EchoClient interface {
	Send(*EchoRequest) error
	Recv() (*EchoResponse, error)
	grpc.ClientStream
}

type echoServiceEchoClient struct {
	grpc.ClientStream
}

func (x *echoServiceEchoClient) Send(m *EchoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoServiceEchoClient) Recv() (*EchoResponse, error) {
	m := new(EchoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *echoServiceClient) Stream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (EchoService_StreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EchoService_serviceDesc.Streams[1], c.cc, "/echoserver.EchoService/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServiceStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EchoService_StreamClient interface {
	Recv() (*EchoResponse, error)
	grpc.ClientStream
}

type echoServiceStreamClient struct {
	grpc.ClientStream
}

func (x *echoServiceStreamClient) Recv() (*EchoResponse, error) {
	m := new(EchoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *echoServiceClient) Heartbeats(ctx context.Context, opts ...grpc.CallOption) (EchoService_HeartbeatsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EchoService_serviceDesc.Streams[2], c.cc, "/echoserver.EchoService/Heartbeats", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServiceHeartbeatsClient{stream}
	return x, nil
}

type EchoService_HeartbeatsClient interface {
	Send(*Empty) error
	Recv() (*Heartbeat, error)
	grpc.ClientStream
}

type echoServiceHeartbeatsClient struct {
	grpc.ClientStream
}

func (x *echoServiceHeartbeatsClient) Send(m *Empty) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoServiceHeartbeatsClient) Recv() (*Heartbeat, error) {
	m := new(Heartbeat)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for EchoService service

type EchoServiceServer interface {
	Echo(EchoService_EchoServer) error
	Stream(*Empty, EchoService_StreamServer) error
	Heartbeats(EchoService_HeartbeatsServer) error
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_Echo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServiceServer).Echo(&echoServiceEchoServer{stream})
}

type EchoService_EchoServer interface {
	Send(*EchoResponse) error
	Recv() (*EchoRequest, error)
	grpc.ServerStream
}

type echoServiceEchoServer struct {
	grpc.ServerStream
}

func (x *echoServiceEchoServer) Send(m *EchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoServiceEchoServer) Recv() (*EchoRequest, error) {
	m := new(EchoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EchoService_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EchoServiceServer).Stream(m, &echoServiceStreamServer{stream})
}

type EchoService_StreamServer interface {
	Send(*EchoResponse) error
	grpc.ServerStream
}

type echoServiceStreamServer struct {
	grpc.ServerStream
}

func (x *echoServiceStreamServer) Send(m *EchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _EchoService_Heartbeats_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServiceServer).Heartbeats(&echoServiceHeartbeatsServer{stream})
}

type EchoService_HeartbeatsServer interface {
	Send(*Heartbeat) error
	Recv() (*Empty, error)
	grpc.ServerStream
}

type echoServiceHeartbeatsServer struct {
	grpc.ServerStream
}

func (x *echoServiceHeartbeatsServer) Send(m *Heartbeat) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoServiceHeartbeatsServer) Recv() (*Empty, error) {
	m := new(Empty)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echoserver.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Echo",
			Handler:       _EchoService_Echo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Stream",
			Handler:       _EchoService_Stream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Heartbeats",
			Handler:       _EchoService_Heartbeats_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("echoserver.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xdd, 0x62, 0x13, 0x3a, 0xb1, 0x12, 0x57, 0xc4, 0x10, 0x2a, 0x96, 0xbd, 0x58, 0x3c,
	0x24, 0xa5, 0x7a, 0x10, 0xef, 0x15, 0xa1, 0xd0, 0x40, 0xaa, 0x78, 0x75, 0x53, 0xc6, 0x36, 0x60,
	0xb2, 0x31, 0xbb, 0x2d, 0x78, 0xf5, 0x15, 0x7c, 0x34, 0x5f, 0xc1, 0x77, 0xf0, 0x2a, 0xc9, 0xb6,
	0x69, 0xc4, 0xea, 0x71, 0x76, 0xbe, 0xf9, 0xe7, 0x9f, 0x7f, 0xc1, 0xc6, 0xe9, 0x5c, 0x48, 0xcc,
	0x97, 0x98, 0x7b, 0x59, 0x2e, 0x94, 0xa0, 0xb0, 0x79, 0x71, 0x3b, 0x33, 0x21, 0x66, 0xcf, 0xe8,
	0xf3, 0x2c, 0xf6, 0x79, 0x9a, 0x0a, 0xc5, 0x55, 0x2c, 0x52, 0xa9, 0x49, 0xf7, 0x74, 0xd5, 0x2d,
	0xab, 0x68, 0xf1, 0xe4, 0xab, 0x38, 0x41, 0xa9, 0x78, 0x92, 0x69, 0x80, 0x9d, 0x81, 0x35, 0x9c,
	0xce, 0x45, 0x88, 0x2f, 0x0b, 0x94, 0x8a, 0x3a, 0x60, 0x26, 0x28, 0x25, 0x9f, 0xa1, 0x43, 0xba,
	0xa4, 0xd7, 0x0a, 0xd7, 0x25, 0x8b, 0x60, 0x4f, 0x83, 0x32, 0x13, 0xa9, 0xc4, 0xbf, 0x49, 0x7a,
	0x05, 0xad, 0x6a, 0x8b, 0xd3, 0xe8, 0x92, 0x9e, 0x35, 0x70, 0x3d, 0xed, 0xc3, 0x5b, 0xfb, 0xf0,
	0xee, 0xd6, 0x44, 0xb8, 0x81, 0xd9, 0x23, 0xb4, 0x6e, 0x91, 0xe7, 0x2a, 0x42, 0xae, 0xe8, 0x25,
	0x18, 0x52, 0x71, 0xb5, 0x90, 0xa5, 0xfe, 0xfe, 0xa0, 0xe3, 0xd5, 0x72, 0xa8, 0x30, 0x6f, 0x52,
	0x32, 0xe1, 0x8a, 0x65, 0x27, 0x60, 0xe8, 0x17, 0x6a, 0x81, 0x79, 0x3f, 0x1e, 0x8d, 0x83, 0x87,
	0xb1, 0xbd, 0x43, 0x0d, 0x68, 0x04, 0x23, 0x9b, 0x30, 0x13, 0x9a, 0xc3, 0x24, 0x53, 0xaf, 0x83,
	0x2f, 0xa2, 0x0f, 0x9f, 0x60, 0xbe, 0x8c, 0xa7, 0x48, 0x03, 0xd8, 0x2d, 0x4a, 0x7a, 0x5c, 0xdf,
	0x52, 0x4b, 0xc6, 0x75, 0x7e, 0x37, 0x74, 0x12, 0xcc, 0x7e, 0xfb, 0xf8, 0x7c, 0x6f, 0x00, 0x6b,
	0xfa, 0x05, 0x71, 0x4d, 0xce, 0x7b, 0xa4, 0x4f, 0xe8, 0x4d, 0x61, 0x24, 0x47, 0x9e, 0xd0, 0x83,
	0x1f, 0x93, 0xc5, 0xf6, 0x7f, 0xc4, 0xda, 0xa5, 0x98, 0x49, 0xb5, 0x58, 0x9f, 0xd0, 0x00, 0xa0,
	0x3a, 0x56, 0x6e, 0xd3, 0x3a, 0xda, 0x9a, 0x0b, 0x3b, 0x2c, 0x85, 0xda, 0xcc, 0xf2, 0xe7, 0xd5,
	0x78, 0x61, 0x2c, 0x32, 0xca, 0x3f, 0xb8, 0xf8, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xea, 0x44,
	0xef, 0x57, 0x02, 0x00, 0x00,
}
