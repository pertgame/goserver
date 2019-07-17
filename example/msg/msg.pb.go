// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example/msg/msg.proto

package cmsg

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

type ReqHello struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqHello) Reset()         { *m = ReqHello{} }
func (m *ReqHello) String() string { return proto.CompactTextString(m) }
func (*ReqHello) ProtoMessage()    {}
func (*ReqHello) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6af7a2c9c40ad41, []int{0}
}

func (m *ReqHello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqHello.Unmarshal(m, b)
}
func (m *ReqHello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqHello.Marshal(b, m, deterministic)
}
func (m *ReqHello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqHello.Merge(m, src)
}
func (m *ReqHello) XXX_Size() int {
	return xxx_messageInfo_ReqHello.Size(m)
}
func (m *ReqHello) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqHello.DiscardUnknown(m)
}

var xxx_messageInfo_ReqHello proto.InternalMessageInfo

func (m *ReqHello) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RespHello struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespHello) Reset()         { *m = RespHello{} }
func (m *RespHello) String() string { return proto.CompactTextString(m) }
func (*RespHello) ProtoMessage()    {}
func (*RespHello) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6af7a2c9c40ad41, []int{1}
}

func (m *RespHello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespHello.Unmarshal(m, b)
}
func (m *RespHello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespHello.Marshal(b, m, deterministic)
}
func (m *RespHello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespHello.Merge(m, src)
}
func (m *RespHello) XXX_Size() int {
	return xxx_messageInfo_RespHello.Size(m)
}
func (m *RespHello) XXX_DiscardUnknown() {
	xxx_messageInfo_RespHello.DiscardUnknown(m)
}

var xxx_messageInfo_RespHello proto.InternalMessageInfo

func (m *RespHello) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ReqSession2Server struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSession2Server) Reset()         { *m = ReqSession2Server{} }
func (m *ReqSession2Server) String() string { return proto.CompactTextString(m) }
func (*ReqSession2Server) ProtoMessage()    {}
func (*ReqSession2Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6af7a2c9c40ad41, []int{2}
}

func (m *ReqSession2Server) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSession2Server.Unmarshal(m, b)
}
func (m *ReqSession2Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSession2Server.Marshal(b, m, deterministic)
}
func (m *ReqSession2Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSession2Server.Merge(m, src)
}
func (m *ReqSession2Server) XXX_Size() int {
	return xxx_messageInfo_ReqSession2Server.Size(m)
}
func (m *ReqSession2Server) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSession2Server.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSession2Server proto.InternalMessageInfo

func (m *ReqSession2Server) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RespSession2Server struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespSession2Server) Reset()         { *m = RespSession2Server{} }
func (m *RespSession2Server) String() string { return proto.CompactTextString(m) }
func (*RespSession2Server) ProtoMessage()    {}
func (*RespSession2Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6af7a2c9c40ad41, []int{3}
}

func (m *RespSession2Server) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespSession2Server.Unmarshal(m, b)
}
func (m *RespSession2Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespSession2Server.Marshal(b, m, deterministic)
}
func (m *RespSession2Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespSession2Server.Merge(m, src)
}
func (m *RespSession2Server) XXX_Size() int {
	return xxx_messageInfo_RespSession2Server.Size(m)
}
func (m *RespSession2Server) XXX_DiscardUnknown() {
	xxx_messageInfo_RespSession2Server.DiscardUnknown(m)
}

var xxx_messageInfo_RespSession2Server proto.InternalMessageInfo

func (m *RespSession2Server) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ReqServer2Server struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqServer2Server) Reset()         { *m = ReqServer2Server{} }
func (m *ReqServer2Server) String() string { return proto.CompactTextString(m) }
func (*ReqServer2Server) ProtoMessage()    {}
func (*ReqServer2Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6af7a2c9c40ad41, []int{4}
}

func (m *ReqServer2Server) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqServer2Server.Unmarshal(m, b)
}
func (m *ReqServer2Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqServer2Server.Marshal(b, m, deterministic)
}
func (m *ReqServer2Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqServer2Server.Merge(m, src)
}
func (m *ReqServer2Server) XXX_Size() int {
	return xxx_messageInfo_ReqServer2Server.Size(m)
}
func (m *ReqServer2Server) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqServer2Server.DiscardUnknown(m)
}

var xxx_messageInfo_ReqServer2Server proto.InternalMessageInfo

func (m *ReqServer2Server) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RespServer2Server struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespServer2Server) Reset()         { *m = RespServer2Server{} }
func (m *RespServer2Server) String() string { return proto.CompactTextString(m) }
func (*RespServer2Server) ProtoMessage()    {}
func (*RespServer2Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6af7a2c9c40ad41, []int{5}
}

func (m *RespServer2Server) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespServer2Server.Unmarshal(m, b)
}
func (m *RespServer2Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespServer2Server.Marshal(b, m, deterministic)
}
func (m *RespServer2Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespServer2Server.Merge(m, src)
}
func (m *RespServer2Server) XXX_Size() int {
	return xxx_messageInfo_RespServer2Server.Size(m)
}
func (m *RespServer2Server) XXX_DiscardUnknown() {
	xxx_messageInfo_RespServer2Server.DiscardUnknown(m)
}

var xxx_messageInfo_RespServer2Server proto.InternalMessageInfo

func (m *RespServer2Server) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ReqRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqRequest) Reset()         { *m = ReqRequest{} }
func (m *ReqRequest) String() string { return proto.CompactTextString(m) }
func (*ReqRequest) ProtoMessage()    {}
func (*ReqRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6af7a2c9c40ad41, []int{6}
}

func (m *ReqRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqRequest.Unmarshal(m, b)
}
func (m *ReqRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqRequest.Marshal(b, m, deterministic)
}
func (m *ReqRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqRequest.Merge(m, src)
}
func (m *ReqRequest) XXX_Size() int {
	return xxx_messageInfo_ReqRequest.Size(m)
}
func (m *ReqRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReqRequest proto.InternalMessageInfo

func (m *ReqRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RespRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespRequest) Reset()         { *m = RespRequest{} }
func (m *RespRequest) String() string { return proto.CompactTextString(m) }
func (*RespRequest) ProtoMessage()    {}
func (*RespRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6af7a2c9c40ad41, []int{7}
}

func (m *RespRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespRequest.Unmarshal(m, b)
}
func (m *RespRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespRequest.Marshal(b, m, deterministic)
}
func (m *RespRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespRequest.Merge(m, src)
}
func (m *RespRequest) XXX_Size() int {
	return xxx_messageInfo_RespRequest.Size(m)
}
func (m *RespRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RespRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RespRequest proto.InternalMessageInfo

func (m *RespRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqHello)(nil), "cmsg.ReqHello")
	proto.RegisterType((*RespHello)(nil), "cmsg.RespHello")
	proto.RegisterType((*ReqSession2Server)(nil), "cmsg.ReqSession2Server")
	proto.RegisterType((*RespSession2Server)(nil), "cmsg.RespSession2Server")
	proto.RegisterType((*ReqServer2Server)(nil), "cmsg.ReqServer2Server")
	proto.RegisterType((*RespServer2Server)(nil), "cmsg.RespServer2Server")
	proto.RegisterType((*ReqRequest)(nil), "cmsg.ReqRequest")
	proto.RegisterType((*RespRequest)(nil), "cmsg.RespRequest")
}

func init() { proto.RegisterFile("example/msg/msg.proto", fileDescriptor_e6af7a2c9c40ad41) }

var fileDescriptor_e6af7a2c9c40ad41 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0xcf, 0x2d, 0x4e, 0x07, 0x61, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21,
	0x96, 0xe4, 0xdc, 0xe2, 0x74, 0x25, 0x39, 0x2e, 0x8e, 0xa0, 0xd4, 0x42, 0x8f, 0xd4, 0x9c, 0x9c,
	0x7c, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20,
	0x30, 0x5b, 0x49, 0x9e, 0x8b, 0x33, 0x28, 0xb5, 0xb8, 0x00, 0xb7, 0x02, 0x75, 0x2e, 0xc1, 0xa0,
	0xd4, 0xc2, 0xe0, 0xd4, 0xe2, 0xe2, 0xcc, 0xfc, 0x3c, 0xa3, 0xe0, 0xd4, 0xa2, 0xb2, 0xd4, 0x22,
	0xac, 0x0a, 0x35, 0xb8, 0x84, 0x40, 0x26, 0x11, 0xa1, 0x52, 0x8d, 0x4b, 0x00, 0x6c, 0x24, 0x48,
	0x01, 0x3e, 0x75, 0x60, 0xab, 0x41, 0x26, 0x12, 0x52, 0xa8, 0xc0, 0xc5, 0x15, 0x94, 0x5a, 0x18,
	0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x82, 0x55, 0x85, 0x22, 0x17, 0x37, 0xc8, 0x28, 0x3c, 0x4a,
	0x92, 0xd8, 0xc0, 0xc1, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x65, 0xc9, 0x78, 0x28, 0x4f,
	0x01, 0x00, 0x00,
}