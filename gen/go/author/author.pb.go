// Code generated by protoc-gen-go. DO NOT EDIT.
// source: author/author.proto

package author

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

type Author struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Author) Reset()         { *m = Author{} }
func (m *Author) String() string { return proto.CompactTextString(m) }
func (*Author) ProtoMessage()    {}
func (*Author) Descriptor() ([]byte, []int) {
	return fileDescriptor_41efdaee09960382, []int{0}
}

func (m *Author) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Author.Unmarshal(m, b)
}
func (m *Author) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Author.Marshal(b, m, deterministic)
}
func (m *Author) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Author.Merge(m, src)
}
func (m *Author) XXX_Size() int {
	return xxx_messageInfo_Author.Size(m)
}
func (m *Author) XXX_DiscardUnknown() {
	xxx_messageInfo_Author.DiscardUnknown(m)
}

var xxx_messageInfo_Author proto.InternalMessageInfo

func (m *Author) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Author)(nil), "author.Author")
}

func init() { proto.RegisterFile("author/author.proto", fileDescriptor_41efdaee09960382) }

var fileDescriptor_41efdaee09960382 = []byte{
	// 116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x2c, 0x2d, 0xc9,
	0xc8, 0x2f, 0xd2, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x92,
	0x0c, 0x17, 0x9b, 0x23, 0x98, 0x25, 0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8,
	0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x3b, 0x19, 0x46, 0xe9, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26,
	0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x67, 0x97, 0x16, 0x95, 0x14, 0xa4, 0xa6, 0x66, 0xeb, 0x83, 0x8d,
	0xd0, 0x4d, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0xd5, 0x4f, 0x4f, 0xcd, 0xd3, 0x4f, 0xcf, 0x87,
	0x1a, 0x9f, 0xc4, 0x06, 0x96, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xec, 0xb9, 0xfa, 0xeb,
	0x76, 0x00, 0x00, 0x00,
}
