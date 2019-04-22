// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

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

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UserPageRequest struct {
	Page                 int32    `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	PageNo               int32    `protobuf:"varint,2,opt,name=PageNo,proto3" json:"PageNo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserPageRequest) Reset()         { *m = UserPageRequest{} }
func (m *UserPageRequest) String() string { return proto.CompactTextString(m) }
func (*UserPageRequest) ProtoMessage()    {}
func (*UserPageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserPageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserPageRequest.Unmarshal(m, b)
}
func (m *UserPageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserPageRequest.Marshal(b, m, deterministic)
}
func (m *UserPageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserPageRequest.Merge(m, src)
}
func (m *UserPageRequest) XXX_Size() int {
	return xxx_messageInfo_UserPageRequest.Size(m)
}
func (m *UserPageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserPageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserPageRequest proto.InternalMessageInfo

func (m *UserPageRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *UserPageRequest) GetPageNo() int32 {
	if m != nil {
		return m.PageNo
	}
	return 0
}

type UserPageResponse struct {
	Page                 int32    `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	PageNo               int32    `protobuf:"varint,2,opt,name=PageNo,proto3" json:"PageNo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserPageResponse) Reset()         { *m = UserPageResponse{} }
func (m *UserPageResponse) String() string { return proto.CompactTextString(m) }
func (*UserPageResponse) ProtoMessage()    {}
func (*UserPageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *UserPageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserPageResponse.Unmarshal(m, b)
}
func (m *UserPageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserPageResponse.Marshal(b, m, deterministic)
}
func (m *UserPageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserPageResponse.Merge(m, src)
}
func (m *UserPageResponse) XXX_Size() int {
	return xxx_messageInfo_UserPageResponse.Size(m)
}
func (m *UserPageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserPageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserPageResponse proto.InternalMessageInfo

func (m *UserPageResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *UserPageResponse) GetPageNo() int32 {
	if m != nil {
		return m.PageNo
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*UserPageRequest)(nil), "pb.UserPageRequest")
	proto.RegisterType((*UserPageResponse)(nil), "pb.UserPageResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe2, 0x62, 0x09, 0x2d,
	0x4e, 0x2d, 0x12, 0xe2, 0xe3, 0x62, 0xf2, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x62,
	0xf2, 0x4c, 0x11, 0x12, 0xe2, 0x62, 0xf1, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x52, 0x60, 0xd4, 0xe0,
	0x0c, 0x02, 0xb3, 0x95, 0x6c, 0xb9, 0xf8, 0x41, 0x6a, 0x03, 0x12, 0xd3, 0x53, 0x83, 0x52, 0x0b,
	0x4b, 0x53, 0x8b, 0x4b, 0x40, 0xca, 0x40, 0x5c, 0xb0, 0x46, 0xd6, 0x20, 0x30, 0x5b, 0x48, 0x8c,
	0x8b, 0x0d, 0x44, 0xfb, 0xe5, 0x83, 0x35, 0xb3, 0x06, 0x41, 0x79, 0x4a, 0x76, 0x5c, 0x02, 0x08,
	0xed, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0xa4, 0xe8, 0x37, 0x4a, 0xe0, 0xe2, 0x06, 0xe9, 0x0f,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x92, 0xe1, 0x62, 0xf1, 0xcc, 0x4b, 0xcb, 0x17, 0xe2,
	0xd0, 0x2b, 0x48, 0xd2, 0x03, 0x49, 0x48, 0xc1, 0x59, 0x4a, 0x0c, 0x42, 0xc6, 0x5c, 0x2c, 0x3e,
	0x99, 0xc5, 0x25, 0x42, 0xc2, 0x30, 0x31, 0x24, 0x57, 0x4b, 0x89, 0xa0, 0x0a, 0x42, 0xdc, 0xa2,
	0xc4, 0x90, 0xc4, 0x06, 0x0e, 0x17, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xe1, 0xbe,
	0x53, 0x25, 0x01, 0x00, 0x00,
}
