// Code generated by protoc-gen-go. DO NOT EDIT.
// source: system/email/proto/message.proto

package system_micro_srv_email

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

type EmailRequest struct {
	To                   string   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Cc                   string   `protobuf:"bytes,2,opt,name=cc,proto3" json:"cc,omitempty"`
	ToName               string   `protobuf:"bytes,3,opt,name=toName,proto3" json:"toName,omitempty"`
	Subject              string   `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	Body                 string   `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	Bcc                  string   `protobuf:"bytes,6,opt,name=bcc,proto3" json:"bcc,omitempty"`
	IsHtml               bool     `protobuf:"varint,7,opt,name=isHtml,proto3" json:"isHtml,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailRequest) Reset()         { *m = EmailRequest{} }
func (m *EmailRequest) String() string { return proto.CompactTextString(m) }
func (*EmailRequest) ProtoMessage()    {}
func (*EmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7559485d539ec57a, []int{0}
}

func (m *EmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailRequest.Unmarshal(m, b)
}
func (m *EmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailRequest.Marshal(b, m, deterministic)
}
func (m *EmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailRequest.Merge(m, src)
}
func (m *EmailRequest) XXX_Size() int {
	return xxx_messageInfo_EmailRequest.Size(m)
}
func (m *EmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmailRequest proto.InternalMessageInfo

func (m *EmailRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *EmailRequest) GetCc() string {
	if m != nil {
		return m.Cc
	}
	return ""
}

func (m *EmailRequest) GetToName() string {
	if m != nil {
		return m.ToName
	}
	return ""
}

func (m *EmailRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *EmailRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *EmailRequest) GetBcc() string {
	if m != nil {
		return m.Bcc
	}
	return ""
}

func (m *EmailRequest) GetIsHtml() bool {
	if m != nil {
		return m.IsHtml
	}
	return false
}

type EmailResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailResponse) Reset()         { *m = EmailResponse{} }
func (m *EmailResponse) String() string { return proto.CompactTextString(m) }
func (*EmailResponse) ProtoMessage()    {}
func (*EmailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7559485d539ec57a, []int{1}
}

func (m *EmailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailResponse.Unmarshal(m, b)
}
func (m *EmailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailResponse.Marshal(b, m, deterministic)
}
func (m *EmailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailResponse.Merge(m, src)
}
func (m *EmailResponse) XXX_Size() int {
	return xxx_messageInfo_EmailResponse.Size(m)
}
func (m *EmailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmailResponse proto.InternalMessageInfo

func (m *EmailResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *EmailResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*EmailRequest)(nil), "system.micro.srv.email.EmailRequest")
	proto.RegisterType((*EmailResponse)(nil), "system.micro.srv.email.EmailResponse")
}

func init() { proto.RegisterFile("system/email/proto/message.proto", fileDescriptor_7559485d539ec57a) }

var fileDescriptor_7559485d539ec57a = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x4d, 0x1b, 0xd3, 0x38, 0x54, 0x91, 0x39, 0x94, 0xc5, 0x53, 0x09, 0x0a, 0x9e, 0x36,
	0xa0, 0xcf, 0x20, 0x78, 0xa9, 0x87, 0x14, 0x1f, 0x20, 0xd9, 0x0e, 0x12, 0xe9, 0x76, 0xeb, 0xce,
	0xb6, 0xd0, 0x97, 0xf1, 0x59, 0x4d, 0x26, 0x1b, 0x50, 0x10, 0x7a, 0xfb, 0xbf, 0x7f, 0x87, 0xd9,
	0x7f, 0x7e, 0x58, 0xf2, 0x89, 0x03, 0xd9, 0x92, 0x6c, 0xdd, 0x6e, 0xcb, 0xbd, 0x77, 0xc1, 0x95,
	0x96, 0x98, 0xeb, 0x0f, 0xd2, 0x42, 0xb8, 0x18, 0x26, 0xb4, 0x6d, 0x8d, 0x77, 0x9a, 0xfd, 0x51,
	0xcb, 0x6c, 0xf1, 0x9d, 0xc0, 0xfc, 0xa5, 0x57, 0x15, 0x7d, 0x1d, 0x88, 0x03, 0xde, 0xc0, 0x24,
	0x38, 0x95, 0x2c, 0x93, 0xc7, 0xab, 0xaa, 0x53, 0x3d, 0x1b, 0xa3, 0x26, 0x03, 0x1b, 0x83, 0x0b,
	0xc8, 0x82, 0x7b, 0xab, 0x2d, 0xa9, 0xa9, 0x78, 0x91, 0x50, 0xc1, 0x8c, 0x0f, 0xcd, 0x27, 0x99,
	0xa0, 0x52, 0x79, 0x18, 0x11, 0x11, 0xd2, 0xc6, 0x6d, 0x4e, 0xea, 0x52, 0x6c, 0xd1, 0x78, 0x0b,
	0xd3, 0xa6, 0x5b, 0x9b, 0x89, 0xd5, 0xcb, 0x7e, 0x6f, 0xcb, 0xaf, 0xc1, 0x6e, 0xd5, 0xac, 0x33,
	0xf3, 0x2a, 0x52, 0xb1, 0x82, 0xeb, 0x98, 0x8f, 0xf7, 0x6e, 0xc7, 0xf1, 0x23, 0x63, 0xba, 0xeb,
	0x24, 0x65, 0x5e, 0x8d, 0x88, 0x05, 0xcc, 0xc9, 0x7b, 0xe7, 0x57, 0xc3, 0xe5, 0x31, 0xf4, 0x1f,
	0xef, 0xa9, 0x86, 0x5c, 0xd6, 0xad, 0x8f, 0x06, 0xdf, 0x21, 0x5d, 0xd3, 0x6e, 0x83, 0xf7, 0xfa,
	0xff, 0x72, 0xf4, 0xef, 0x62, 0xee, 0x1e, 0xce, 0x4c, 0x0d, 0xf1, 0x8a, 0x8b, 0x26, 0x93, 0xc6,
	0x9f, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x46, 0x42, 0x55, 0x95, 0x01, 0x00, 0x00,
}