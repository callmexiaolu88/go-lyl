// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mathFunction.proto

package mathFunction

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AddRequest struct {
	P1                   int32    `protobuf:"varint,1,opt,name=p1,proto3" json:"p1,omitempty"`
	P2                   int32    `protobuf:"varint,2,opt,name=p2,proto3" json:"p2,omitempty"`
	P3                   int32    `protobuf:"varint,3,opt,name=p3,proto3" json:"p3,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df79ca0ee6f52085, []int{0}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetP1() int32 {
	if m != nil {
		return m.P1
	}
	return 0
}

func (m *AddRequest) GetP2() int32 {
	if m != nil {
		return m.P2
	}
	return 0
}

func (m *AddRequest) GetP3() int32 {
	if m != nil {
		return m.P3
	}
	return 0
}

type AddResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_df79ca0ee6f52085, []int{1}
}

func (m *AddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResponse.Unmarshal(m, b)
}
func (m *AddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResponse.Marshal(b, m, deterministic)
}
func (m *AddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResponse.Merge(m, src)
}
func (m *AddResponse) XXX_Size() int {
	return xxx_messageInfo_AddResponse.Size(m)
}
func (m *AddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddResponse proto.InternalMessageInfo

func (m *AddResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type SubtractionRequest struct {
	P1                   int32    `protobuf:"varint,1,opt,name=p1,proto3" json:"p1,omitempty"`
	P2                   int32    `protobuf:"varint,2,opt,name=p2,proto3" json:"p2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubtractionRequest) Reset()         { *m = SubtractionRequest{} }
func (m *SubtractionRequest) String() string { return proto.CompactTextString(m) }
func (*SubtractionRequest) ProtoMessage()    {}
func (*SubtractionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df79ca0ee6f52085, []int{2}
}

func (m *SubtractionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubtractionRequest.Unmarshal(m, b)
}
func (m *SubtractionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubtractionRequest.Marshal(b, m, deterministic)
}
func (m *SubtractionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubtractionRequest.Merge(m, src)
}
func (m *SubtractionRequest) XXX_Size() int {
	return xxx_messageInfo_SubtractionRequest.Size(m)
}
func (m *SubtractionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubtractionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubtractionRequest proto.InternalMessageInfo

func (m *SubtractionRequest) GetP1() int32 {
	if m != nil {
		return m.P1
	}
	return 0
}

func (m *SubtractionRequest) GetP2() int32 {
	if m != nil {
		return m.P2
	}
	return 0
}

type SubtractionResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubtractionResponse) Reset()         { *m = SubtractionResponse{} }
func (m *SubtractionResponse) String() string { return proto.CompactTextString(m) }
func (*SubtractionResponse) ProtoMessage()    {}
func (*SubtractionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_df79ca0ee6f52085, []int{3}
}

func (m *SubtractionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubtractionResponse.Unmarshal(m, b)
}
func (m *SubtractionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubtractionResponse.Marshal(b, m, deterministic)
}
func (m *SubtractionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubtractionResponse.Merge(m, src)
}
func (m *SubtractionResponse) XXX_Size() int {
	return xxx_messageInfo_SubtractionResponse.Size(m)
}
func (m *SubtractionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubtractionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubtractionResponse proto.InternalMessageInfo

func (m *SubtractionResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*AddRequest)(nil), "AddRequest")
	proto.RegisterType((*AddResponse)(nil), "AddResponse")
	proto.RegisterType((*SubtractionRequest)(nil), "SubtractionRequest")
	proto.RegisterType((*SubtractionResponse)(nil), "SubtractionResponse")
}

func init() { proto.RegisterFile("mathFunction.proto", fileDescriptor_df79ca0ee6f52085) }

var fileDescriptor_df79ca0ee6f52085 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0x4d, 0x2c, 0xc9,
	0x70, 0x2b, 0xcd, 0x4b, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xb2,
	0xe1, 0xe2, 0x72, 0x4c, 0x49, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe3, 0x62,
	0x2a, 0x30, 0x94, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62, 0x2a, 0x30, 0x04, 0xf3, 0x8d, 0x24,
	0x98, 0xa0, 0x7c, 0x23, 0x30, 0xdf, 0x58, 0x82, 0x19, 0xca, 0x37, 0x56, 0x52, 0xe5, 0xe2, 0x06,
	0xeb, 0x2e, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12, 0xe3, 0x62, 0x2b, 0x4a, 0x2d, 0x2e, 0xcd,
	0x29, 0x81, 0x1a, 0x01, 0xe5, 0x29, 0x99, 0x70, 0x09, 0x05, 0x97, 0x26, 0x95, 0x14, 0x25, 0x82,
	0x6d, 0x26, 0xd2, 0x32, 0x25, 0x5d, 0x2e, 0x61, 0x14, 0x5d, 0xf8, 0x2d, 0x31, 0xca, 0xe2, 0xe2,
	0xf1, 0x45, 0xf2, 0x9f, 0x90, 0x02, 0x17, 0xb3, 0x63, 0x4a, 0x8a, 0x10, 0xb7, 0x1e, 0xc2, 0x7f,
	0x52, 0x3c, 0x7a, 0xc8, 0xce, 0xb5, 0xe0, 0xe2, 0x46, 0xb2, 0x40, 0x48, 0x58, 0x0f, 0xd3, 0x91,
	0x52, 0x22, 0x7a, 0x58, 0xdc, 0x90, 0xc4, 0x06, 0x0e, 0x3c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xa3, 0x8b, 0x97, 0x62, 0x52, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MathFunctionClient is the client API for MathFunction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MathFunctionClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	Subtraction(ctx context.Context, in *SubtractionRequest, opts ...grpc.CallOption) (*SubtractionResponse, error)
}

type mathFunctionClient struct {
	cc *grpc.ClientConn
}

func NewMathFunctionClient(cc *grpc.ClientConn) MathFunctionClient {
	return &mathFunctionClient{cc}
}

func (c *mathFunctionClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/MathFunction/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathFunctionClient) Subtraction(ctx context.Context, in *SubtractionRequest, opts ...grpc.CallOption) (*SubtractionResponse, error) {
	out := new(SubtractionResponse)
	err := c.cc.Invoke(ctx, "/MathFunction/Subtraction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MathFunctionServer is the server API for MathFunction service.
type MathFunctionServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
	Subtraction(context.Context, *SubtractionRequest) (*SubtractionResponse, error)
}

func RegisterMathFunctionServer(s *grpc.Server, srv MathFunctionServer) {
	s.RegisterService(&_MathFunction_serviceDesc, srv)
}

func _MathFunction_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathFunctionServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MathFunction/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathFunctionServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MathFunction_Subtraction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubtractionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathFunctionServer).Subtraction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MathFunction/Subtraction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathFunctionServer).Subtraction(ctx, req.(*SubtractionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MathFunction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "MathFunction",
	HandlerType: (*MathFunctionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _MathFunction_Add_Handler,
		},
		{
			MethodName: "Subtraction",
			Handler:    _MathFunction_Subtraction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mathFunction.proto",
}
