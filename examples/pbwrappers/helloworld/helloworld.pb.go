// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package helloworld

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type HelloRequest struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	StrVal               *wrappers.StringValue `protobuf:"bytes,2,opt,name=strVal,proto3" json:"strVal,omitempty"`
	FloatVal             *wrappers.FloatValue  `protobuf:"bytes,3,opt,name=floatVal,proto3" json:"floatVal,omitempty"`
	DoubleVal            *wrappers.DoubleValue `protobuf:"bytes,4,opt,name=doubleVal,proto3" json:"doubleVal,omitempty"`
	BoolVal              *wrappers.BoolValue   `protobuf:"bytes,5,opt,name=boolVal,proto3" json:"boolVal,omitempty"`
	BytesVal             *wrappers.BytesValue  `protobuf:"bytes,6,opt,name=bytesVal,proto3" json:"bytesVal,omitempty"`
	Int32Val             *wrappers.Int32Value  `protobuf:"bytes,7,opt,name=int32Val,proto3" json:"int32Val,omitempty"`
	Uint32Val            *wrappers.UInt32Value `protobuf:"bytes,8,opt,name=uint32Val,proto3" json:"uint32Val,omitempty"`
	Int64Val             *wrappers.Int64Value  `protobuf:"bytes,9,opt,name=int64Val,proto3" json:"int64Val,omitempty"`
	Uint64Val            *wrappers.UInt64Value `protobuf:"bytes,10,opt,name=uint64Val,proto3" json:"uint64Val,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HelloRequest) GetStrVal() *wrappers.StringValue {
	if m != nil {
		return m.StrVal
	}
	return nil
}

func (m *HelloRequest) GetFloatVal() *wrappers.FloatValue {
	if m != nil {
		return m.FloatVal
	}
	return nil
}

func (m *HelloRequest) GetDoubleVal() *wrappers.DoubleValue {
	if m != nil {
		return m.DoubleVal
	}
	return nil
}

func (m *HelloRequest) GetBoolVal() *wrappers.BoolValue {
	if m != nil {
		return m.BoolVal
	}
	return nil
}

func (m *HelloRequest) GetBytesVal() *wrappers.BytesValue {
	if m != nil {
		return m.BytesVal
	}
	return nil
}

func (m *HelloRequest) GetInt32Val() *wrappers.Int32Value {
	if m != nil {
		return m.Int32Val
	}
	return nil
}

func (m *HelloRequest) GetUint32Val() *wrappers.UInt32Value {
	if m != nil {
		return m.Uint32Val
	}
	return nil
}

func (m *HelloRequest) GetInt64Val() *wrappers.Int64Value {
	if m != nil {
		return m.Int64Val
	}
	return nil
}

func (m *HelloRequest) GetUint64Val() *wrappers.UInt64Value {
	if m != nil {
		return m.Uint64Val
	}
	return nil
}

type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor_17b8c58d586b62f2) }

var fileDescriptor_17b8c58d586b62f2 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcd, 0x8a, 0xd4, 0x40,
	0x14, 0x85, 0xe9, 0xd8, 0xf6, 0x4f, 0x8d, 0x0b, 0x29, 0xa4, 0x09, 0xb1, 0x91, 0xa1, 0x17, 0x32,
	0xab, 0x04, 0x7a, 0x9a, 0x11, 0x5c, 0x0e, 0xe2, 0xcf, 0x36, 0x83, 0xb3, 0xe8, 0x5d, 0x85, 0xb9,
	0xd3, 0x06, 0x6a, 0x52, 0xb1, 0x7e, 0x1c, 0xc2, 0x90, 0x8d, 0x3b, 0x77, 0x82, 0xef, 0xe0, 0xeb,
	0xb8, 0x70, 0xe9, 0xd6, 0x07, 0x91, 0xba, 0xa9, 0x4a, 0x45, 0xc6, 0xb8, 0x4b, 0xe5, 0x9c, 0xef,
	0x9e, 0x4a, 0xee, 0x21, 0x8f, 0x3f, 0x00, 0xe7, 0xe2, 0x56, 0x48, 0x7e, 0x95, 0xd6, 0x52, 0x68,
	0x41, 0x49, 0x78, 0x93, 0xac, 0x0f, 0x42, 0x1c, 0x38, 0x64, 0xac, 0x2e, 0x33, 0x56, 0x55, 0x42,
	0x33, 0x5d, 0x8a, 0x4a, 0x75, 0xce, 0xe4, 0x99, 0x53, 0xf1, 0x54, 0x98, 0xeb, 0xec, 0x56, 0xb2,
	0xba, 0x06, 0xe9, 0xf4, 0xcd, 0xf7, 0x29, 0x79, 0xf4, 0xd6, 0x0e, 0xcb, 0xe1, 0xa3, 0x01, 0xa5,
	0x29, 0x25, 0xd3, 0x8a, 0xdd, 0x40, 0x3c, 0x39, 0x9e, 0x9c, 0x2c, 0x73, 0x7c, 0xa6, 0x3b, 0x32,
	0x53, 0x5a, 0x5e, 0x32, 0x1e, 0x47, 0xc7, 0x93, 0x93, 0xa3, 0xed, 0x3a, 0xed, 0xa6, 0xa6, 0x7e,
	0x6a, 0x7a, 0xa1, 0x65, 0x59, 0x1d, 0x2e, 0x19, 0x37, 0x90, 0x3b, 0x2f, 0x7d, 0x41, 0x16, 0xd7,
	0x5c, 0x30, 0x6d, 0xb9, 0x07, 0xc8, 0x3d, 0xbd, 0xc7, 0xbd, 0x76, 0x06, 0x03, 0x79, 0x6f, 0xa6,
	0x2f, 0xc9, 0xf2, 0x4a, 0x98, 0x82, 0x83, 0x25, 0xa7, 0x23, 0x89, 0xaf, 0xbc, 0xc3, 0x40, 0x1e,
	0xec, 0x74, 0x47, 0xe6, 0x85, 0x10, 0xdc, 0x92, 0x0f, 0x91, 0x4c, 0xee, 0x91, 0xe7, 0x9d, 0x6e,
	0x20, 0xf7, 0x56, 0x7b, 0xd5, 0xa2, 0xd1, 0xa0, 0x2c, 0x36, 0x1b, 0xb9, 0xea, 0xb9, 0x33, 0xd8,
	0xab, 0x7a, 0xb3, 0x05, 0xcb, 0x4a, 0x9f, 0x6e, 0x2d, 0x38, 0x1f, 0x01, 0xdf, 0x39, 0x83, 0x05,
	0xbd, 0xd9, 0x7e, 0xa3, 0xe9, 0xc9, 0xc5, 0xc8, 0x37, 0xbe, 0x1f, 0xa0, 0xc1, 0xee, 0x42, 0xcf,
	0x76, 0x16, 0x5d, 0x8e, 0x87, 0xa2, 0xc1, 0x85, 0xe2, 0xb3, 0x0f, 0xed, 0x48, 0xf2, 0x9f, 0x50,
	0x8f, 0x06, 0xfb, 0xe6, 0x39, 0x21, 0xae, 0x27, 0x35, 0x6f, 0x68, 0x4c, 0xe6, 0x37, 0xa0, 0x14,
	0x3b, 0xf8, 0xa2, 0xf8, 0xe3, 0xf6, 0x57, 0x44, 0xe6, 0x6f, 0x24, 0x80, 0x06, 0x49, 0x7f, 0x44,
	0x64, 0x71, 0xc1, 0x1a, 0xe4, 0x68, 0x9c, 0x0e, 0x6a, 0x3c, 0xac, 0x5c, 0xb2, 0xfa, 0x87, 0x52,
	0xf3, 0x66, 0xf3, 0x35, 0xfa, 0xfc, 0xf3, 0xf7, 0xb7, 0xe8, 0x4b, 0x44, 0x8f, 0x32, 0xc5, 0x9a,
	0xec, 0xce, 0x96, 0xb1, 0xdd, 0xaf, 0xe8, 0x13, 0x3c, 0x2a, 0x2d, 0x3f, 0x31, 0x9e, 0xdd, 0x75,
	0x75, 0x6b, 0xf7, 0x09, 0x8d, 0xf1, 0x3d, 0xf6, 0x08, 0x15, 0xdf, 0xa8, 0x76, 0xbf, 0xa6, 0x09,
	0x6a, 0x5d, 0x51, 0x50, 0xec, 0x3b, 0xd3, 0xee, 0x63, 0xba, 0x42, 0xd5, 0xd6, 0x01, 0x35, 0xd7,
	0x8b, 0x30, 0x13, 0x17, 0xde, 0x49, 0x6e, 0xf5, 0x41, 0xc3, 0xbd, 0xa0, 0xe6, 0x37, 0x14, 0xf2,
	0x4c, 0x10, 0x4d, 0x50, 0x07, 0xe4, 0xd9, 0xce, 0x93, 0xf8, 0x9b, 0xff, 0x26, 0x9d, 0xd8, 0x2f,
	0xa1, 0x2d, 0x66, 0xb8, 0xa6, 0xd3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x65, 0xa5, 0x95, 0xa8,
	0x13, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}
