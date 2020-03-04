// Code generated by protoc-gen-go. DO NOT EDIT.
// source: result.proto

package result

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

type Result_Status int32

const (
	Result_STATUS_SUCCESS Result_Status = 0
	Result_STATUS_FAILED  Result_Status = 1
)

var Result_Status_name = map[int32]string{
	0: "STATUS_SUCCESS",
	1: "STATUS_FAILED",
}

var Result_Status_value = map[string]int32{
	"STATUS_SUCCESS": 0,
	"STATUS_FAILED":  1,
}

func (x Result_Status) String() string {
	return proto.EnumName(Result_Status_name, int32(x))
}

func (Result_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4feee897733d2100, []int{0, 0}
}

type Result struct {
	Status               Result_Status `protobuf:"varint,1,opt,name=status,proto3,enum=result.Result_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_4feee897733d2100, []int{0}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetStatus() Result_Status {
	if m != nil {
		return m.Status
	}
	return Result_STATUS_SUCCESS
}

func init() {
	proto.RegisterEnum("result.Result_Status", Result_Status_name, Result_Status_value)
	proto.RegisterType((*Result)(nil), "result.Result")
}

func init() { proto.RegisterFile("result.proto", fileDescriptor_4feee897733d2100) }

var fileDescriptor_4feee897733d2100 = []byte{
	// 123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4a, 0x2d, 0x2e,
	0xcd, 0x29, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x32, 0xb8, 0xd8,
	0x82, 0xc0, 0x2c, 0x21, 0x5d, 0x2e, 0xb6, 0xe2, 0x92, 0xc4, 0x92, 0xd2, 0x62, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0x3e, 0x23, 0x51, 0x3d, 0xa8, 0x06, 0x88, 0xbc, 0x5e, 0x30, 0x58, 0x32, 0x08, 0xaa,
	0x48, 0x49, 0x9f, 0x8b, 0x0d, 0x22, 0x22, 0x24, 0xc4, 0xc5, 0x17, 0x1c, 0xe2, 0x18, 0x12, 0x1a,
	0x1c, 0x1f, 0x1c, 0xea, 0xec, 0xec, 0x1a, 0x1c, 0x2c, 0xc0, 0x20, 0x24, 0xc8, 0xc5, 0x0b, 0x15,
	0x73, 0x73, 0xf4, 0xf4, 0x71, 0x75, 0x11, 0x60, 0x4c, 0x62, 0x03, 0x5b, 0x6c, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x8b, 0xba, 0x07, 0xa4, 0x88, 0x00, 0x00, 0x00,
}