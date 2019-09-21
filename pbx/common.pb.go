// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pbx/common.proto

package pbx

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

type RequestHeader struct {
	Platform             string   `protobuf:"bytes,1,opt,name=platform,proto3" json:"platform,omitempty"`
	Domain               string   `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	ClientKey            string   `protobuf:"bytes,3,opt,name=client_key,json=clientKey,proto3" json:"client_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestHeader) Reset()         { *m = RequestHeader{} }
func (m *RequestHeader) String() string { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()    {}
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_773f3b3a3f55901c, []int{0}
}

func (m *RequestHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestHeader.Unmarshal(m, b)
}
func (m *RequestHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestHeader.Marshal(b, m, deterministic)
}
func (m *RequestHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestHeader.Merge(m, src)
}
func (m *RequestHeader) XXX_Size() int {
	return xxx_messageInfo_RequestHeader.Size(m)
}
func (m *RequestHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RequestHeader proto.InternalMessageInfo

func (m *RequestHeader) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *RequestHeader) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RequestHeader) GetClientKey() string {
	if m != nil {
		return m.ClientKey
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestHeader)(nil), "pbx.RequestHeader")
}

func init() { proto.RegisterFile("pbx/common.proto", fileDescriptor_773f3b3a3f55901c) }

var fileDescriptor_773f3b3a3f55901c = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x48, 0xaa, 0xd0,
	0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x48,
	0xaa, 0x50, 0x4a, 0xe2, 0xe2, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xf1, 0x48, 0x4d, 0x4c,
	0x49, 0x2d, 0x12, 0x92, 0xe2, 0xe2, 0x28, 0xc8, 0x49, 0x2c, 0x49, 0xcb, 0x2f, 0xca, 0x95, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x85, 0xc4, 0xb8, 0xd8, 0x52, 0xf2, 0x73, 0x13, 0x33,
	0xf3, 0x24, 0x98, 0xc0, 0x32, 0x50, 0x9e, 0x90, 0x2c, 0x17, 0x57, 0x72, 0x4e, 0x66, 0x6a, 0x5e,
	0x49, 0x7c, 0x76, 0x6a, 0xa5, 0x04, 0x33, 0x58, 0x8e, 0x13, 0x22, 0xe2, 0x9d, 0x5a, 0x99, 0xc4,
	0x06, 0xb6, 0xcf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x77, 0x86, 0x4a, 0x1d, 0x83, 0x00, 0x00,
	0x00,
}