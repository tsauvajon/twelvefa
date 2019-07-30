// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calc.proto

package calc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MaxRequest struct {
	A                    int64    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int64    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MaxRequest) Reset()         { *m = MaxRequest{} }
func (m *MaxRequest) String() string { return proto.CompactTextString(m) }
func (*MaxRequest) ProtoMessage()    {}
func (*MaxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b9900dc883ea68, []int{0}
}

func (m *MaxRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaxRequest.Unmarshal(m, b)
}
func (m *MaxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaxRequest.Marshal(b, m, deterministic)
}
func (m *MaxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaxRequest.Merge(m, src)
}
func (m *MaxRequest) XXX_Size() int {
	return xxx_messageInfo_MaxRequest.Size(m)
}
func (m *MaxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MaxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MaxRequest proto.InternalMessageInfo

func (m *MaxRequest) GetA() int64 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *MaxRequest) GetB() int64 {
	if m != nil {
		return m.B
	}
	return 0
}

type MaxResponse struct {
	Max                  int64    `protobuf:"varint,1,opt,name=max,proto3" json:"max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MaxResponse) Reset()         { *m = MaxResponse{} }
func (m *MaxResponse) String() string { return proto.CompactTextString(m) }
func (*MaxResponse) ProtoMessage()    {}
func (*MaxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b9900dc883ea68, []int{1}
}

func (m *MaxResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaxResponse.Unmarshal(m, b)
}
func (m *MaxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaxResponse.Marshal(b, m, deterministic)
}
func (m *MaxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaxResponse.Merge(m, src)
}
func (m *MaxResponse) XXX_Size() int {
	return xxx_messageInfo_MaxResponse.Size(m)
}
func (m *MaxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MaxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MaxResponse proto.InternalMessageInfo

func (m *MaxResponse) GetMax() int64 {
	if m != nil {
		return m.Max
	}
	return 0
}

type AddRequest struct {
	A                    int64    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int64    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b9900dc883ea68, []int{2}
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

func (m *AddRequest) GetA() int64 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *AddRequest) GetB() int64 {
	if m != nil {
		return m.B
	}
	return 0
}

type AddResponse struct {
	Sum                  int64    `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b9900dc883ea68, []int{3}
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

func (m *AddResponse) GetSum() int64 {
	if m != nil {
		return m.Sum
	}
	return 0
}

type NthPrimesRequest struct {
	N                    []uint64 `protobuf:"varint,1,rep,packed,name=n,proto3" json:"n,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NthPrimesRequest) Reset()         { *m = NthPrimesRequest{} }
func (m *NthPrimesRequest) String() string { return proto.CompactTextString(m) }
func (*NthPrimesRequest) ProtoMessage()    {}
func (*NthPrimesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b9900dc883ea68, []int{4}
}

func (m *NthPrimesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NthPrimesRequest.Unmarshal(m, b)
}
func (m *NthPrimesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NthPrimesRequest.Marshal(b, m, deterministic)
}
func (m *NthPrimesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NthPrimesRequest.Merge(m, src)
}
func (m *NthPrimesRequest) XXX_Size() int {
	return xxx_messageInfo_NthPrimesRequest.Size(m)
}
func (m *NthPrimesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NthPrimesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NthPrimesRequest proto.InternalMessageInfo

func (m *NthPrimesRequest) GetN() []uint64 {
	if m != nil {
		return m.N
	}
	return nil
}

type NthPrimesResponse struct {
	NthPrimes            []uint64 `protobuf:"varint,1,rep,packed,name=nthPrimes,proto3" json:"nthPrimes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NthPrimesResponse) Reset()         { *m = NthPrimesResponse{} }
func (m *NthPrimesResponse) String() string { return proto.CompactTextString(m) }
func (*NthPrimesResponse) ProtoMessage()    {}
func (*NthPrimesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b9900dc883ea68, []int{5}
}

func (m *NthPrimesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NthPrimesResponse.Unmarshal(m, b)
}
func (m *NthPrimesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NthPrimesResponse.Marshal(b, m, deterministic)
}
func (m *NthPrimesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NthPrimesResponse.Merge(m, src)
}
func (m *NthPrimesResponse) XXX_Size() int {
	return xxx_messageInfo_NthPrimesResponse.Size(m)
}
func (m *NthPrimesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NthPrimesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NthPrimesResponse proto.InternalMessageInfo

func (m *NthPrimesResponse) GetNthPrimes() []uint64 {
	if m != nil {
		return m.NthPrimes
	}
	return nil
}

func init() {
	proto.RegisterType((*MaxRequest)(nil), "calc.MaxRequest")
	proto.RegisterType((*MaxResponse)(nil), "calc.MaxResponse")
	proto.RegisterType((*AddRequest)(nil), "calc.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "calc.AddResponse")
	proto.RegisterType((*NthPrimesRequest)(nil), "calc.NthPrimesRequest")
	proto.RegisterType((*NthPrimesResponse)(nil), "calc.NthPrimesResponse")
}

func init() { proto.RegisterFile("calc.proto", fileDescriptor_a2b9900dc883ea68) }

var fileDescriptor_a2b9900dc883ea68 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0x86, 0x30,
	0x18, 0xc7, 0xdf, 0x35, 0x09, 0x7c, 0xea, 0xa0, 0x3b, 0x94, 0x48, 0x90, 0xec, 0xe4, 0x21, 0x84,
	0xea, 0x1e, 0x48, 0x67, 0x23, 0xfc, 0x06, 0xd3, 0x0d, 0x0a, 0xdc, 0x34, 0xa7, 0xe0, 0xe7, 0xe9,
	0x93, 0xc6, 0xa6, 0x6b, 0xa3, 0x2e, 0xef, 0xed, 0x79, 0xfe, 0xfc, 0xf6, 0x63, 0xff, 0x0d, 0xa0,
	0x67, 0x43, 0x5f, 0x4d, 0xf3, 0xb8, 0x8c, 0x24, 0x32, 0x33, 0x2d, 0x01, 0x1a, 0xb6, 0xb5, 0xe2,
	0x6b, 0x15, 0x7a, 0x21, 0xd7, 0x80, 0x58, 0x86, 0x0a, 0x54, 0xe2, 0x16, 0x31, 0xb3, 0x75, 0xd9,
	0xc5, 0xbe, 0x75, 0xf4, 0x1e, 0xae, 0x2c, 0xa9, 0xa7, 0x51, 0x69, 0x41, 0x12, 0xc0, 0x92, 0x6d,
	0x07, 0x6c, 0x46, 0xa3, 0xaa, 0x39, 0x3f, 0x53, 0x65, 0x49, 0xaf, 0xd2, 0xab, 0x74, 0x2a, 0xbd,
	0x4a, 0x5a, 0x40, 0xf2, 0xb6, 0x7c, 0xbc, 0xcf, 0x9f, 0x52, 0xe8, 0x40, 0xa8, 0x32, 0x54, 0xe0,
	0x32, 0x6a, 0x91, 0xa2, 0x8f, 0x90, 0x06, 0xc4, 0x21, 0xba, 0x83, 0x58, 0xb9, 0xf0, 0x40, 0x7d,
	0xf0, 0xf4, 0x8d, 0x20, 0x7a, 0x65, 0x43, 0x4f, 0x1e, 0x00, 0x37, 0x6c, 0x23, 0x49, 0x65, 0x5f,
	0xc3, 0xd7, 0xcf, 0xd3, 0x20, 0xd9, 0x95, 0xf4, 0x64, 0xe8, 0x9a, 0x73, 0x47, 0xfb, 0x86, 0x8e,
	0x0e, 0x9a, 0xd0, 0x13, 0x79, 0x81, 0xf8, 0xf7, 0x5e, 0xe4, 0x66, 0x27, 0xfe, 0x56, 0xc9, 0x6f,
	0xff, 0xe5, 0xee, 0x7c, 0x77, 0x69, 0x3f, 0xe7, 0xf9, 0x27, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x29,
	0x85, 0x2b, 0xaa, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalcClient is the client API for Calc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalcClient interface {
	Max(ctx context.Context, in *MaxRequest, opts ...grpc.CallOption) (*MaxResponse, error)
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	NthPrimes(ctx context.Context, in *NthPrimesRequest, opts ...grpc.CallOption) (*NthPrimesResponse, error)
}

type calcClient struct {
	cc *grpc.ClientConn
}

func NewCalcClient(cc *grpc.ClientConn) CalcClient {
	return &calcClient{cc}
}

func (c *calcClient) Max(ctx context.Context, in *MaxRequest, opts ...grpc.CallOption) (*MaxResponse, error) {
	out := new(MaxResponse)
	err := c.cc.Invoke(ctx, "/calc.Calc/Max", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/calc.Calc/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcClient) NthPrimes(ctx context.Context, in *NthPrimesRequest, opts ...grpc.CallOption) (*NthPrimesResponse, error) {
	out := new(NthPrimesResponse)
	err := c.cc.Invoke(ctx, "/calc.Calc/NthPrimes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalcServer is the server API for Calc service.
type CalcServer interface {
	Max(context.Context, *MaxRequest) (*MaxResponse, error)
	Add(context.Context, *AddRequest) (*AddResponse, error)
	NthPrimes(context.Context, *NthPrimesRequest) (*NthPrimesResponse, error)
}

// UnimplementedCalcServer can be embedded to have forward compatible implementations.
type UnimplementedCalcServer struct {
}

func (*UnimplementedCalcServer) Max(ctx context.Context, req *MaxRequest) (*MaxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Max not implemented")
}
func (*UnimplementedCalcServer) Add(ctx context.Context, req *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedCalcServer) NthPrimes(ctx context.Context, req *NthPrimesRequest) (*NthPrimesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NthPrimes not implemented")
}

func RegisterCalcServer(s *grpc.Server, srv CalcServer) {
	s.RegisterService(&_Calc_serviceDesc, srv)
}

func _Calc_Max_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MaxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServer).Max(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.Calc/Max",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServer).Max(ctx, req.(*MaxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calc_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.Calc/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calc_NthPrimes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NthPrimesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServer).NthPrimes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.Calc/NthPrimes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServer).NthPrimes(ctx, req.(*NthPrimesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calc.Calc",
	HandlerType: (*CalcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Max",
			Handler:    _Calc_Max_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Calc_Add_Handler,
		},
		{
			MethodName: "NthPrimes",
			Handler:    _Calc_NthPrimes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calc.proto",
}