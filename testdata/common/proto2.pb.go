// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/proto2.proto

package common

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Service, which do smth
type Proto2Message struct {
	Scalar               *int32   `protobuf:"varint,1,opt,name=scalar" json:"scalar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Proto2Message) Reset()         { *m = Proto2Message{} }
func (m *Proto2Message) String() string { return proto.CompactTextString(m) }
func (*Proto2Message) ProtoMessage()    {}
func (*Proto2Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_749cd6b8b12ddea6, []int{0}
}

func (m *Proto2Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proto2Message.Unmarshal(m, b)
}
func (m *Proto2Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proto2Message.Marshal(b, m, deterministic)
}
func (m *Proto2Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proto2Message.Merge(m, src)
}
func (m *Proto2Message) XXX_Size() int {
	return xxx_messageInfo_Proto2Message.Size(m)
}
func (m *Proto2Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Proto2Message.DiscardUnknown(m)
}

var xxx_messageInfo_Proto2Message proto.InternalMessageInfo

func (m *Proto2Message) GetScalar() int32 {
	if m != nil && m.Scalar != nil {
		return *m.Scalar
	}
	return 0
}

func init() {
	proto.RegisterType((*Proto2Message)(nil), "common.Proto2Message")
}

func init() { proto.RegisterFile("common/proto2.proto", fileDescriptor_749cd6b8b12ddea6) }

var fileDescriptor_749cd6b8b12ddea6 = []byte{
	// 123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x37, 0xd2, 0x03, 0x53, 0x42, 0x6c, 0x10, 0x41,
	0x25, 0x75, 0x2e, 0xde, 0x00, 0xb0, 0xb8, 0x6f, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x90, 0x18,
	0x17, 0x5b, 0x71, 0x72, 0x62, 0x4e, 0x62, 0x91, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x94,
	0xe7, 0xa4, 0x1f, 0xa5, 0x9b, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0xef,
	0xea, 0x1e, 0xa2, 0x1b, 0x9a, 0x5d, 0x94, 0x98, 0x99, 0x97, 0xaa, 0x9f, 0x9e, 0x6f, 0x94, 0x5e,
	0x98, 0xa3, 0x5f, 0x92, 0x5a, 0x5c, 0x92, 0x92, 0x58, 0x92, 0xa8, 0x0f, 0x31, 0x19, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x06, 0x3b, 0x88, 0x20, 0x77, 0x00, 0x00, 0x00,
}