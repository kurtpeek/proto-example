// Code generated by protoc-gen-go. DO NOT EDIT.
// source: book/book.proto

package book

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	author "github.com/kurtpeek/proto-example/gen/go/author"
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

type Book struct {
	Title                string         `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Author               *author.Author `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee9082fb44230b1b, []int{0}
}

func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Book) GetAuthor() *author.Author {
	if m != nil {
		return m.Author
	}
	return nil
}

func init() {
	proto.RegisterType((*Book)(nil), "book.Book")
}

func init() { proto.RegisterFile("book/book.proto", fileDescriptor_ee9082fb44230b1b) }

var fileDescriptor_ee9082fb44230b1b = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xca, 0xcf, 0xcf,
	0xd6, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0x20, 0xb6, 0x94, 0x70, 0x62,
	0x69, 0x49, 0x46, 0x7e, 0x91, 0x3e, 0x84, 0x82, 0x48, 0x29, 0xb9, 0x70, 0xb1, 0x38, 0xe5, 0xe7,
	0x67, 0x0b, 0x89, 0x70, 0xb1, 0x96, 0x64, 0x96, 0xe4, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x06, 0x41, 0x38, 0x42, 0x6a, 0x5c, 0x6c, 0x10, 0xd5, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46,
	0x7c, 0x7a, 0x50, 0xcd, 0x8e, 0x60, 0x2a, 0x08, 0x2a, 0xeb, 0xa4, 0x1f, 0xa5, 0x9b, 0x9e, 0x59,
	0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x5d, 0x5a, 0x54, 0x52, 0x90, 0x9a, 0x9a,
	0xad, 0x0f, 0xb6, 0x42, 0x37, 0xb5, 0x22, 0x31, 0xb7, 0x20, 0x27, 0x55, 0x3f, 0x3d, 0x35, 0x4f,
	0x3f, 0x3d, 0x1f, 0xec, 0xae, 0x24, 0x36, 0xb0, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x45,
	0x8f, 0x85, 0x29, 0xab, 0x00, 0x00, 0x00,
}
