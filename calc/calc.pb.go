// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calc.proto

package calc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
	A                    int64    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
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

func (m *AddResponse) GetA() int64 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *AddResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
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
	NthPrime             []uint64 `protobuf:"varint,1,rep,packed,name=nthPrime,proto3" json:"nthPrime,omitempty"`
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

func (m *NthPrimesResponse) GetNthPrime() []uint64 {
	if m != nil {
		return m.NthPrime
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
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x4e, 0x85, 0x30,
	0x14, 0x86, 0x6f, 0x2d, 0x31, 0xde, 0x73, 0x1d, 0x7a, 0x3b, 0x28, 0x61, 0x91, 0x74, 0x62, 0xd0,
	0x6b, 0xa2, 0xbb, 0x09, 0x71, 0xc6, 0x98, 0xbe, 0x41, 0xa1, 0x4d, 0x34, 0x81, 0x82, 0x2d, 0x26,
	0x3c, 0x8f, 0x4f, 0x6a, 0x5a, 0x0a, 0x54, 0x5d, 0xdc, 0xce, 0xf9, 0xf3, 0xf1, 0x85, 0xff, 0x14,
	0xa0, 0x11, 0x6d, 0x73, 0x1a, 0x4c, 0x3f, 0xf6, 0x34, 0x71, 0x33, 0x2b, 0x00, 0x2a, 0x31, 0x71,
	0xf5, 0xf1, 0xa9, 0xec, 0x48, 0x2f, 0x01, 0x89, 0x14, 0xe5, 0xa8, 0xc0, 0x1c, 0x09, 0xb7, 0xd5,
	0xe9, 0xd9, 0xbc, 0xd5, 0xec, 0x06, 0x0e, 0x9e, 0xb4, 0x43, 0xaf, 0xad, 0xa2, 0x04, 0x70, 0x27,
	0xa6, 0x00, 0xbb, 0xd1, 0xa9, 0x4a, 0x29, 0xff, 0xa3, 0xba, 0x83, 0x83, 0x27, 0x83, 0xea, 0x27,
	0x4a, 0x00, 0x2b, 0x63, 0x3c, 0xbc, 0xe7, 0x6e, 0x64, 0x39, 0x90, 0x97, 0xf1, 0xed, 0xd5, 0xbc,
	0x77, 0xca, 0x46, 0x7a, 0x9d, 0xa2, 0x1c, 0x17, 0x09, 0x47, 0x9a, 0xdd, 0xc3, 0x31, 0x22, 0x82,
	0x36, 0x83, 0x0b, 0x1d, 0xc2, 0x40, 0xae, 0xfb, 0xc3, 0x17, 0x82, 0xe4, 0x59, 0xb4, 0x0d, 0xbd,
	0x05, 0x5c, 0x89, 0x89, 0x92, 0x93, 0xbf, 0xcc, 0x76, 0x8a, 0xec, 0x18, 0x25, 0xb3, 0x90, 0xed,
	0x1c, 0x5d, 0x4a, 0xb9, 0xd0, 0x5b, 0xdb, 0x85, 0x8e, 0x5a, 0xb1, 0x1d, 0x7d, 0x82, 0xfd, 0xfa,
	0x57, 0xf4, 0x6a, 0x26, 0x7e, 0x17, 0xc9, 0xae, 0xff, 0xe4, 0xcb, 0xf7, 0xf5, 0xb9, 0x7f, 0xa8,
	0xc7, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x87, 0xf4, 0xee, 0xb6, 0x01, 0x00, 0x00,
}
