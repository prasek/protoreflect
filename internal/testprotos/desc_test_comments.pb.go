// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desc_test_comments.proto

package testprotos

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"
import _ "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Request_MarioCharacters int32

const (
	Request_MARIO     Request_MarioCharacters = 1
	Request_LUIGI     Request_MarioCharacters = 2
	Request_PEACH     Request_MarioCharacters = 3
	Request_BOWSER    Request_MarioCharacters = 4
	Request_WARIO     Request_MarioCharacters = 5
	Request_WALUIGI   Request_MarioCharacters = 6
	Request_SHY_GUY   Request_MarioCharacters = 7
	Request_HEY_HO    Request_MarioCharacters = 7
	Request_MAGIKOOPA Request_MarioCharacters = 8
	Request_KAMEK     Request_MarioCharacters = 8
)

var Request_MarioCharacters_name = map[int32]string{
	1: "MARIO",
	2: "LUIGI",
	3: "PEACH",
	4: "BOWSER",
	5: "WARIO",
	6: "WALUIGI",
	7: "SHY_GUY",
	// Duplicate value: 7: "HEY_HO",
	8: "MAGIKOOPA",
	// Duplicate value: 8: "KAMEK",
}
var Request_MarioCharacters_value = map[string]int32{
	"MARIO":     1,
	"LUIGI":     2,
	"PEACH":     3,
	"BOWSER":    4,
	"WARIO":     5,
	"WALUIGI":   6,
	"SHY_GUY":   7,
	"HEY_HO":    7,
	"MAGIKOOPA": 8,
	"KAMEK":     8,
}

func (x Request_MarioCharacters) Enum() *Request_MarioCharacters {
	p := new(Request_MarioCharacters)
	*p = x
	return p
}
func (x Request_MarioCharacters) String() string {
	return proto.EnumName(Request_MarioCharacters_name, int32(x))
}
func (x *Request_MarioCharacters) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Request_MarioCharacters_value, data, "Request_MarioCharacters")
	if err != nil {
		return err
	}
	*x = Request_MarioCharacters(value)
	return nil
}
func (Request_MarioCharacters) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorDescTestComments, []int{0, 0}
}

// We need a request for our RPC service below.
type Request struct {
	// A field comment
	Ids []int32 `protobuf:"varint,1,rep,packed,name=ids,json=|foo|" json:"ids,omitempty"`
	// label comment
	Name                         *string         `protobuf:"bytes,2,opt,name=name,def=fubar" json:"name,omitempty"`
	Extras                       *Request_Extras `protobuf:"group,3,opt,name=Extras,json=extras" json:"extras,omitempty"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptorDescTestComments, []int{0} }

var extRange_Request = []proto.ExtensionRange{
	{Start: 100, End: 200},
	{Start: 201, End: 250},
}

func (*Request) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Request
}

const Default_Request_Name string = "fubar"

func (m *Request) GetIds() []int32 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *Request) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return Default_Request_Name
}

func (m *Request) GetExtras() *Request_Extras {
	if m != nil {
		return m.Extras
	}
	return nil
}

// Group comment
type Request_Extras struct {
	Dbl *float64 `protobuf:"fixed64,1,opt,name=dbl" json:"dbl,omitempty"`
	Flt *float32 `protobuf:"fixed32,2,opt,name=flt" json:"flt,omitempty"`
	// Leading comment...
	Str              *string `protobuf:"bytes,3,opt,name=str" json:"str,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request_Extras) Reset()         { *m = Request_Extras{} }
func (m *Request_Extras) String() string { return proto.CompactTextString(m) }
func (*Request_Extras) ProtoMessage()    {}
func (*Request_Extras) Descriptor() ([]byte, []int) {
	return fileDescriptorDescTestComments, []int{0, 0}
}

func (m *Request_Extras) GetDbl() float64 {
	if m != nil && m.Dbl != nil {
		return *m.Dbl
	}
	return 0
}

func (m *Request_Extras) GetFlt() float32 {
	if m != nil && m.Flt != nil {
		return *m.Flt
	}
	return 0
}

func (m *Request_Extras) GetStr() string {
	if m != nil && m.Str != nil {
		return *m.Str
	}
	return ""
}

var E_Guid1 = &proto.ExtensionDesc{
	ExtendedType:  (*Request)(nil),
	ExtensionType: (*uint64)(nil),
	Field:         123,
	Name:          "foo.bar.guid1",
	Tag:           "varint,123,opt,name=guid1",
	Filename:      "desc_test_comments.proto",
}

var E_Guid2 = &proto.ExtensionDesc{
	ExtendedType:  (*Request)(nil),
	ExtensionType: (*uint64)(nil),
	Field:         124,
	Name:          "foo.bar.guid2",
	Tag:           "varint,124,opt,name=guid2",
	Filename:      "desc_test_comments.proto",
}

func init() {
	proto.RegisterType((*Request)(nil), "foo.bar.Request")
	proto.RegisterType((*Request_Extras)(nil), "foo.bar.Request.Extras")
	proto.RegisterEnum("foo.bar.Request_MarioCharacters", Request_MarioCharacters_name, Request_MarioCharacters_value)
	proto.RegisterExtension(E_Guid1)
	proto.RegisterExtension(E_Guid2)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RpcService service

type RpcServiceClient interface {
	// Method comment
	StreamingRpc(ctx context.Context, opts ...grpc.CallOption) (RpcService_StreamingRpcClient, error)
	UnaryRpc(ctx context.Context, in *Request, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type rpcServiceClient struct {
	cc *grpc.ClientConn
}

func NewRpcServiceClient(cc *grpc.ClientConn) RpcServiceClient {
	return &rpcServiceClient{cc}
}

func (c *rpcServiceClient) StreamingRpc(ctx context.Context, opts ...grpc.CallOption) (RpcService_StreamingRpcClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RpcService_serviceDesc.Streams[0], c.cc, "/foo.bar.RpcService/StreamingRpc", opts...)
	if err != nil {
		return nil, err
	}
	x := &rpcServiceStreamingRpcClient{stream}
	return x, nil
}

type RpcService_StreamingRpcClient interface {
	Send(*Request) error
	CloseAndRecv() (*Request, error)
	grpc.ClientStream
}

type rpcServiceStreamingRpcClient struct {
	grpc.ClientStream
}

func (x *rpcServiceStreamingRpcClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rpcServiceStreamingRpcClient) CloseAndRecv() (*Request, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Request)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rpcServiceClient) UnaryRpc(ctx context.Context, in *Request, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/foo.bar.RpcService/UnaryRpc", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RpcService service

type RpcServiceServer interface {
	// Method comment
	StreamingRpc(RpcService_StreamingRpcServer) error
	UnaryRpc(context.Context, *Request) (*google_protobuf.Empty, error)
}

func RegisterRpcServiceServer(s *grpc.Server, srv RpcServiceServer) {
	s.RegisterService(&_RpcService_serviceDesc, srv)
}

func _RpcService_StreamingRpc_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RpcServiceServer).StreamingRpc(&rpcServiceStreamingRpcServer{stream})
}

type RpcService_StreamingRpcServer interface {
	SendAndClose(*Request) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type rpcServiceStreamingRpcServer struct {
	grpc.ServerStream
}

func (x *rpcServiceStreamingRpcServer) SendAndClose(m *Request) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rpcServiceStreamingRpcServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RpcService_UnaryRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServiceServer).UnaryRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foo.bar.RpcService/UnaryRpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServiceServer).UnaryRpc(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _RpcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "foo.bar.RpcService",
	HandlerType: (*RpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryRpc",
			Handler:    _RpcService_UnaryRpc_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingRpc",
			Handler:       _RpcService_StreamingRpc_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "desc_test_comments.proto",
}

func init() { proto.RegisterFile("desc_test_comments.proto", fileDescriptorDescTestComments) }

var fileDescriptorDescTestComments = []byte{
	// 650 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x4d, 0x6b, 0xdb, 0x4a,
	0x14, 0xf5, 0x48, 0x23, 0x59, 0x9e, 0xbc, 0xc7, 0x1b, 0x06, 0x92, 0xe8, 0x29, 0x10, 0xfc, 0xc2,
	0xa3, 0x35, 0x59, 0xc8, 0xc4, 0xe9, 0x2a, 0xab, 0x3a, 0xc1, 0x8d, 0x1d, 0xd7, 0x38, 0x8c, 0x09,
	0x21, 0x5d, 0xd4, 0x48, 0xf2, 0xd8, 0x51, 0xb1, 0x3d, 0xea, 0x68, 0xdc, 0xe6, 0x6b, 0xd3, 0x5d,
	0x7e, 0x41, 0xb7, 0xdd, 0x14, 0x1a, 0x02, 0x85, 0x34, 0xbb, 0x6e, 0x4a, 0xbb, 0x6b, 0x96, 0xfd,
	0x01, 0xfd, 0x15, 0x5d, 0xcc, 0xa2, 0x9b, 0x22, 0xc9, 0x21, 0x90, 0x44, 0xab, 0xcb, 0x39, 0x47,
	0xe7, 0xcc, 0xe5, 0xde, 0x8b, 0xec, 0x1e, 0x8b, 0x83, 0xae, 0x64, 0xb1, 0xec, 0x06, 0x7c, 0x34,
	0x62, 0x63, 0x19, 0xbb, 0x91, 0xe0, 0x92, 0x93, 0x7c, 0x9f, 0x73, 0xd7, 0xf7, 0x84, 0xb3, 0x30,
	0xe0, 0x7c, 0x30, 0x64, 0xe5, 0x14, 0xf6, 0x27, 0xfd, 0x32, 0x1b, 0x45, 0xf2, 0x30, 0x53, 0x39,
	0xc5, 0xdb, 0x64, 0xe2, 0x27, 0xc2, 0x48, 0x72, 0x31, 0x55, 0xcc, 0xdf, 0x24, 0xf0, 0x48, 0x86,
	0x7c, 0x3c, 0x0d, 0x58, 0xfa, 0xa1, 0xa3, 0x3c, 0x65, 0x2f, 0x27, 0x2c, 0x96, 0xe4, 0x7f, 0xa4,
	0x87, 0xbd, 0xd8, 0x06, 0x45, 0xbd, 0x64, 0xac, 0xe3, 0x73, 0x05, 0x75, 0xcf, 0x0f, 0x2e, 0x15,
	0xd4, 0x0f, 0x0e, 0x8f, 0x30, 0xa0, 0xc6, 0x49, 0x9f, 0xf3, 0x13, 0xf2, 0x2f, 0x82, 0x63, 0x6f,
	0xc4, 0x6c, 0xad, 0x08, 0x4a, 0x85, 0x35, 0xa3, 0x3f, 0xf1, 0x3d, 0x41, 0x53, 0x88, 0x94, 0x91,
	0xc9, 0x0e, 0xa4, 0xf0, 0x62, 0x5b, 0x2f, 0x82, 0x12, 0xaa, 0xcc, 0xbb, 0xd3, 0xe7, 0xbb, 0xd3,
	0x08, 0xb7, 0x96, 0xd2, 0x74, 0x2a, 0x73, 0x9e, 0x20, 0x33, 0x43, 0x08, 0x46, 0x7a, 0xcf, 0x1f,
	0xda, 0xa0, 0x08, 0x4a, 0x80, 0x26, 0x65, 0x82, 0xf4, 0x87, 0x32, 0x8d, 0xd1, 0x68, 0x52, 0x26,
	0x48, 0x2c, 0x45, 0xea, 0x5d, 0xa0, 0x49, 0xb9, 0x66, 0x9e, 0x29, 0x98, 0xc3, 0xb9, 0xa5, 0x6f,
	0x00, 0xfd, 0xd3, 0xf2, 0x44, 0xc8, 0x37, 0xf6, 0x3d, 0xe1, 0x05, 0x92, 0x89, 0x98, 0xcc, 0x22,
	0xa3, 0x55, 0xa5, 0x8d, 0x36, 0x06, 0x0e, 0x3a, 0x53, 0xf0, 0xad, 0x76, 0xa1, 0xe0, 0x2f, 0x48,
	0x16, 0x90, 0xf1, 0x74, 0xa7, 0xb1, 0xd9, 0xc0, 0x9a, 0x83, 0xbf, 0x2a, 0xf8, 0x1d, 0x5c, 0x29,
	0xd8, 0xcb, 0x65, 0x1f, 0x29, 0x20, 0x63, 0xbb, 0x56, 0xdd, 0xa8, 0x63, 0x9d, 0x20, 0x64, 0xae,
	0xb7, 0x77, 0x3b, 0x35, 0x8a, 0x61, 0x02, 0xef, 0xa6, 0x56, 0x06, 0x99, 0x41, 0xf9, 0xdd, 0x6a,
	0x66, 0x60, 0x12, 0x1b, 0xe5, 0x3b, 0xf5, 0xbd, 0xee, 0xe6, 0xce, 0x1e, 0xce, 0x3b, 0x33, 0x9f,
	0x15, 0x9c, 0x3c, 0x9c, 0x1a, 0x21, 0x64, 0xd6, 0x6b, 0x7b, 0xdd, 0x7a, 0x1b, 0xe7, 0xc9, 0xdf,
	0xa8, 0xd0, 0xaa, 0x6e, 0x36, 0x9a, 0xed, 0xf6, 0x76, 0x15, 0x5b, 0x89, 0x59, 0xb3, 0xda, 0xaa,
	0x35, 0xb1, 0xe5, 0xcc, 0x9c, 0x29, 0x78, 0x7c, 0xa1, 0xe0, 0x1b, 0x03, 0x83, 0x65, 0xc3, 0xea,
	0xe1, 0x2b, 0xb0, 0xbc, 0x68, 0x5d, 0x01, 0xfc, 0x1b, 0x38, 0x73, 0xe7, 0x0a, 0x9a, 0x71, 0x34,
	0xf4, 0xe4, 0x7f, 0x97, 0x0a, 0x5a, 0x39, 0xa0, 0xe9, 0xd0, 0x30, 0xf3, 0x69, 0xcb, 0xc0, 0x06,
	0x5b, 0xd0, 0x42, 0x78, 0x76, 0x0b, 0x5a, 0x8b, 0x78, 0x95, 0xea, 0x7d, 0xce, 0xa9, 0x9e, 0xcc,
	0x41, 0xf7, 0xbd, 0xa3, 0xca, 0x07, 0x80, 0x10, 0x8d, 0x82, 0x0e, 0x13, 0xaf, 0xc2, 0x80, 0x91,
	0x47, 0xe8, 0xaf, 0x8e, 0x14, 0xcc, 0x1b, 0x85, 0xe3, 0x01, 0x8d, 0x02, 0x82, 0x6f, 0x4f, 0xc5,
	0xb9, 0x83, 0x94, 0x00, 0x69, 0x22, 0x6b, 0x67, 0xec, 0x89, 0xc3, 0xfb, 0xff, 0x98, 0x73, 0xb3,
	0x95, 0x73, 0xaf, 0x57, 0xce, 0xad, 0x25, 0xfb, 0xb8, 0x34, 0xfb, 0x51, 0xc1, 0xf7, 0x51, 0xad,
	0xfa, 0x49, 0xc1, 0xd7, 0x5f, 0xde, 0x39, 0xe5, 0x9f, 0xcf, 0x1f, 0x9f, 0x6a, 0x59, 0x33, 0x9a,
	0xd5, 0x3b, 0x57, 0xd0, 0x20, 0xba, 0xcf, 0xfd, 0x0b, 0x05, 0xc1, 0xa9, 0x96, 0x5b, 0x7b, 0x80,
	0x8c, 0xc1, 0x24, 0xec, 0xad, 0xdc, 0x4d, 0xb0, 0x8f, 0x8b, 0xa0, 0x04, 0x69, 0x46, 0x5f, 0xeb,
	0x2a, 0xf7, 0xe8, 0x4e, 0x6e, 0x74, 0x95, 0xf5, 0xd5, 0x67, 0x2b, 0x83, 0x50, 0xee, 0x4f, 0x7c,
	0x37, 0xe0, 0xa3, 0xf2, 0x8b, 0xfd, 0xc9, 0x28, 0xca, 0x8e, 0x42, 0xb0, 0xfe, 0x90, 0x05, 0xb2,
	0x1c, 0x8e, 0x25, 0x13, 0x63, 0x6f, 0x58, 0x4e, 0x4e, 0x21, 0x65, 0xe2, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x65, 0xae, 0x4e, 0xa8, 0x7e, 0x03, 0x00, 0x00,
}
