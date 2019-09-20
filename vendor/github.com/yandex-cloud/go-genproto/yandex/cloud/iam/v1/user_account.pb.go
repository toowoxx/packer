// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/iam/v1/user_account.proto

package iam

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/validation"
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

// Currently represents only [Yandex.Passport account](/docs/iam/concepts/#passport).
type UserAccount struct {
	// ID of the user account.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to UserAccount:
	//	*UserAccount_YandexPassportUserAccount
	UserAccount          isUserAccount_UserAccount `protobuf_oneof:"user_account"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *UserAccount) Reset()         { *m = UserAccount{} }
func (m *UserAccount) String() string { return proto.CompactTextString(m) }
func (*UserAccount) ProtoMessage()    {}
func (*UserAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_101d694eed7724ca, []int{0}
}

func (m *UserAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAccount.Unmarshal(m, b)
}
func (m *UserAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAccount.Marshal(b, m, deterministic)
}
func (m *UserAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAccount.Merge(m, src)
}
func (m *UserAccount) XXX_Size() int {
	return xxx_messageInfo_UserAccount.Size(m)
}
func (m *UserAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAccount.DiscardUnknown(m)
}

var xxx_messageInfo_UserAccount proto.InternalMessageInfo

func (m *UserAccount) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type isUserAccount_UserAccount interface {
	isUserAccount_UserAccount()
}

type UserAccount_YandexPassportUserAccount struct {
	YandexPassportUserAccount *YandexPassportUserAccount `protobuf:"bytes,2,opt,name=yandex_passport_user_account,json=yandexPassportUserAccount,proto3,oneof"`
}

func (*UserAccount_YandexPassportUserAccount) isUserAccount_UserAccount() {}

func (m *UserAccount) GetUserAccount() isUserAccount_UserAccount {
	if m != nil {
		return m.UserAccount
	}
	return nil
}

func (m *UserAccount) GetYandexPassportUserAccount() *YandexPassportUserAccount {
	if x, ok := m.GetUserAccount().(*UserAccount_YandexPassportUserAccount); ok {
		return x.YandexPassportUserAccount
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UserAccount) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UserAccount_YandexPassportUserAccount)(nil),
	}
}

// A YandexPassportUserAccount resource. For more information, see [Yandex.Passport account](/docs/iam/concepts/#passport).
type YandexPassportUserAccount struct {
	// Login of the Yandex.Passport user account.
	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	// Default email of the Yandex.Passport user account.
	DefaultEmail         string   `protobuf:"bytes,2,opt,name=default_email,json=defaultEmail,proto3" json:"default_email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *YandexPassportUserAccount) Reset()         { *m = YandexPassportUserAccount{} }
func (m *YandexPassportUserAccount) String() string { return proto.CompactTextString(m) }
func (*YandexPassportUserAccount) ProtoMessage()    {}
func (*YandexPassportUserAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_101d694eed7724ca, []int{1}
}

func (m *YandexPassportUserAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_YandexPassportUserAccount.Unmarshal(m, b)
}
func (m *YandexPassportUserAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_YandexPassportUserAccount.Marshal(b, m, deterministic)
}
func (m *YandexPassportUserAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_YandexPassportUserAccount.Merge(m, src)
}
func (m *YandexPassportUserAccount) XXX_Size() int {
	return xxx_messageInfo_YandexPassportUserAccount.Size(m)
}
func (m *YandexPassportUserAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_YandexPassportUserAccount.DiscardUnknown(m)
}

var xxx_messageInfo_YandexPassportUserAccount proto.InternalMessageInfo

func (m *YandexPassportUserAccount) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *YandexPassportUserAccount) GetDefaultEmail() string {
	if m != nil {
		return m.DefaultEmail
	}
	return ""
}

func init() {
	proto.RegisterType((*UserAccount)(nil), "yandex.cloud.iam.v1.UserAccount")
	proto.RegisterType((*YandexPassportUserAccount)(nil), "yandex.cloud.iam.v1.YandexPassportUserAccount")
}

func init() {
	proto.RegisterFile("yandex/cloud/iam/v1/user_account.proto", fileDescriptor_101d694eed7724ca)
}

var fileDescriptor_101d694eed7724ca = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x50, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xb6, 0x05, 0x85, 0x65, 0x73, 0x87, 0xe8, 0x61, 0x13, 0x85, 0x31, 0x41, 0x76, 0x59, 0xc2,
	0xf4, 0x38, 0x3c, 0x38, 0x10, 0x3c, 0x4a, 0x41, 0x41, 0x2f, 0xe5, 0xad, 0x89, 0xf5, 0x41, 0x7e,
	0xd4, 0x36, 0x29, 0xee, 0xbf, 0xf1, 0x4f, 0x15, 0x93, 0x1e, 0x2a, 0x6c, 0xc7, 0xf7, 0xbd, 0xef,
	0x17, 0x1f, 0xb9, 0xd9, 0x81, 0x11, 0xf2, 0x9b, 0x17, 0xca, 0x7a, 0xc1, 0x11, 0x34, 0x6f, 0x57,
	0xdc, 0x37, 0xb2, 0xce, 0xa1, 0x28, 0xac, 0x37, 0x8e, 0x55, 0xb5, 0x75, 0x96, 0x9e, 0x45, 0x1e,
	0x0b, 0x3c, 0x86, 0xa0, 0x59, 0xbb, 0xba, 0xb8, 0xfa, 0x27, 0x6e, 0x41, 0xa1, 0x00, 0x87, 0xd6,
	0x44, 0xcd, 0xfc, 0x27, 0x21, 0xc3, 0x97, 0x46, 0xd6, 0x0f, 0xd1, 0x89, 0x8e, 0x49, 0x8a, 0x62,
	0x92, 0xcc, 0x92, 0xc5, 0x20, 0x4b, 0x51, 0xd0, 0x2f, 0x72, 0x19, 0x0d, 0xf2, 0x0a, 0x9a, 0xa6,
	0xb2, 0xb5, 0xcb, 0xfb, 0xc9, 0x93, 0x74, 0x96, 0x2c, 0x86, 0xb7, 0x8c, 0xed, 0x89, 0x66, 0x6f,
	0x01, 0x7b, 0xee, 0x74, 0xbd, 0x94, 0xa7, 0xa3, 0x6c, 0xba, 0x3b, 0xf4, 0xdc, 0x8c, 0xc9, 0xa8,
	0x1f, 0x31, 0x7f, 0x25, 0xd3, 0x83, 0x4e, 0xf4, 0x9c, 0x1c, 0x2b, 0x5b, 0xa2, 0xe9, 0x2a, 0xc7,
	0x83, 0x5e, 0x93, 0x53, 0x21, 0x3f, 0xc0, 0x2b, 0x97, 0x4b, 0x0d, 0xa8, 0x42, 0xcd, 0x41, 0x36,
	0xea, 0xc0, 0xc7, 0x3f, 0x6c, 0x73, 0xff, 0xbe, 0x2e, 0xd1, 0x7d, 0xfa, 0x2d, 0x2b, 0xac, 0xe6,
	0xb1, 0xcf, 0x32, 0xce, 0x54, 0xda, 0x65, 0x29, 0x4d, 0x58, 0x88, 0xef, 0x19, 0x7f, 0x8d, 0xa0,
	0xb7, 0x27, 0xe1, 0x7d, 0xf7, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xf9, 0xdc, 0xc7, 0x2e, 0x9e, 0x01,
	0x00, 0x00,
}
