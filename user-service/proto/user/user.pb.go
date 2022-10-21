// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package laracom_user_service

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
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Status               string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	StripeId             string   `protobuf:"bytes,6,opt,name=stripe_id,json=stripeId,proto3" json:"stripe_id,omitempty"`
	CardBrand            string   `protobuf:"bytes,7,opt,name=card_brand,json=cardBrand,proto3" json:"card_brand,omitempty"`
	CardLastFour         string   `protobuf:"bytes,8,opt,name=card_last_four,json=cardLastFour,proto3" json:"card_last_four,omitempty"`
	TrialEndsAt          string   `protobuf:"bytes,9,opt,name=trial_ends_at,json=trialEndsAt,proto3" json:"trial_ends_at,omitempty"`
	DeletedAt            string   `protobuf:"bytes,10,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	RememberToken        string   `protobuf:"bytes,11,opt,name=remember_token,json=rememberToken,proto3" json:"remember_token,omitempty"`
	CreatedAt            string   `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
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

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *User) GetStripeId() string {
	if m != nil {
		return m.StripeId
	}
	return ""
}

func (m *User) GetCardBrand() string {
	if m != nil {
		return m.CardBrand
	}
	return ""
}

func (m *User) GetCardLastFour() string {
	if m != nil {
		return m.CardLastFour
	}
	return ""
}

func (m *User) GetTrialEndsAt() string {
	if m != nil {
		return m.TrialEndsAt
	}
	return ""
}

func (m *User) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

func (m *User) GetRememberToken() string {
	if m != nil {
		return m.RememberToken
	}
	return ""
}

func (m *User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *User) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "laracom.user.service.User")
	proto.RegisterType((*Request)(nil), "laracom.user.service.Request")
	proto.RegisterType((*Response)(nil), "laracom.user.service.Response")
	proto.RegisterType((*Error)(nil), "laracom.user.service.Error")
}

func init() {
	proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf)
}

var fileDescriptor_116e343673f7ffaf = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xbd, 0x6e, 0xd4, 0x40,
	0x10, 0xc7, 0xb9, 0x2f, 0xe7, 0x3c, 0xce, 0x5d, 0x31, 0x8a, 0xa2, 0xd5, 0x45, 0x41, 0x27, 0x0b,
	0x24, 0x2a, 0x0b, 0x25, 0x35, 0xc5, 0x25, 0x4a, 0xa2, 0x48, 0x54, 0x06, 0x1a, 0x1a, 0x6b, 0xcf,
	0x3b, 0x48, 0x16, 0xb6, 0xd7, 0xcc, 0xae, 0xe1, 0x79, 0x78, 0x38, 0x5e, 0x03, 0xa1, 0x1d, 0xfb,
	0x10, 0xc5, 0x41, 0x0a, 0x2a, 0x7b, 0x7e, 0xff, 0xff, 0x7c, 0x68, 0x67, 0x17, 0xa0, 0x77, 0xc4,
	0x59, 0xc7, 0xd6, 0x5b, 0x3c, 0xab, 0x35, 0xeb, 0xd2, 0x36, 0x99, 0x30, 0x47, 0xfc, 0xb5, 0x2a,
	0x29, 0xfd, 0x39, 0x85, 0xf9, 0x07, 0x47, 0x8c, 0x6b, 0x98, 0x56, 0x46, 0x4d, 0xb6, 0x93, 0x57,
	0x71, 0x3e, 0xad, 0x0c, 0x22, 0xcc, 0x5b, 0xdd, 0x90, 0x9a, 0x0a, 0x91, 0x7f, 0x3c, 0x83, 0x05,
	0x35, 0xba, 0xaa, 0xd5, 0x4c, 0xe0, 0x10, 0xe0, 0x06, 0x96, 0x9d, 0x76, 0xee, 0x9b, 0x65, 0xa3,
	0xe6, 0x22, 0xfc, 0x8e, 0xf1, 0x1c, 0x22, 0xe7, 0xb5, 0xef, 0x9d, 0x5a, 0x88, 0x32, 0x46, 0x78,
	0x01, 0xb1, 0xf3, 0x5c, 0x75, 0x54, 0x54, 0x46, 0x45, 0x43, 0xd2, 0x00, 0x1e, 0x0d, 0x5e, 0x02,
	0x94, 0x9a, 0x4d, 0xb1, 0x67, 0xdd, 0x1a, 0x75, 0x22, 0x6a, 0x1c, 0xc8, 0x4d, 0x00, 0xf8, 0x02,
	0xd6, 0x22, 0xd7, 0xda, 0xf9, 0xe2, 0x93, 0xed, 0x59, 0x2d, 0xc5, 0x72, 0x1a, 0xe8, 0x5b, 0xed,
	0xfc, 0xbd, 0xed, 0x19, 0x53, 0x58, 0x79, 0xae, 0x74, 0x5d, 0x50, 0x6b, 0x5c, 0xa1, 0xbd, 0x8a,
	0xc5, 0x94, 0x08, 0xbc, 0x6b, 0x8d, 0xdb, 0xf9, 0xd0, 0xc8, 0x50, 0x4d, 0x9e, 0x4c, 0x30, 0xc0,
	0xd0, 0x68, 0x24, 0x3b, 0x8f, 0x2f, 0x61, 0xcd, 0xd4, 0x50, 0xb3, 0x27, 0x2e, 0xbc, 0xfd, 0x4c,
	0xad, 0x4a, 0xc4, 0xb2, 0x3a, 0xd0, 0xf7, 0x01, 0xca, 0xb8, 0x4c, 0x7a, 0xac, 0x72, 0x3a, 0x8e,
	0x3b, 0x90, 0xa1, 0x49, 0xdf, 0x99, 0x83, 0xbc, 0x1a, 0xe4, 0x91, 0xec, 0x7c, 0x1a, 0xc3, 0x49,
	0x4e, 0x5f, 0x7a, 0x72, 0x3e, 0xfd, 0x3e, 0x81, 0x65, 0x4e, 0xae, 0xb3, 0xad, 0x23, 0xcc, 0x60,
	0x1e, 0x16, 0x25, 0x1b, 0x49, 0xae, 0x36, 0xd9, 0xb1, 0xed, 0x65, 0x61, 0x73, 0xb9, 0xf8, 0xf0,
	0x35, 0x2c, 0xc2, 0xd7, 0xa9, 0xe9, 0x76, 0xf6, 0x44, 0xc2, 0x60, 0xc4, 0x6b, 0x88, 0x88, 0xd9,
	0xb2, 0x53, 0x33, 0x49, 0xb9, 0x38, 0x9e, 0x72, 0x17, 0x3c, 0xf9, 0x68, 0x4d, 0xdf, 0xc0, 0x42,
	0x40, 0xb8, 0x1f, 0xa5, 0x35, 0x24, 0xf3, 0x2d, 0x72, 0xf9, 0xc7, 0x2d, 0x24, 0x86, 0x5c, 0xc9,
	0x55, 0xe7, 0x2b, 0xdb, 0x8e, 0x57, 0xe7, 0x4f, 0x74, 0xf5, 0x63, 0x02, 0x49, 0x98, 0xe1, 0xdd,
	0x50, 0x1c, 0xef, 0x21, 0xba, 0x95, 0x93, 0xc2, 0x7f, 0x0c, 0xbc, 0x79, 0x7e, 0x5c, 0x3b, 0x9c,
	0x55, 0xfa, 0x0c, 0x6f, 0x61, 0xf6, 0x40, 0xfe, 0x3f, 0x8b, 0x3c, 0x42, 0xf4, 0x40, 0x7e, 0x57,
	0xd7, 0x78, 0xf9, 0x37, 0xaf, 0x2c, 0xea, 0xe9, 0x52, 0x37, 0xe7, 0x1f, 0x8f, 0x3e, 0xb7, 0x7d,
	0x24, 0x6f, 0xf1, 0xfa, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x19, 0xf4, 0xe2, 0xab, 0x99, 0x03,
	0x00, 0x00,
}
