package http

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Config for HTTP proxy server.
type ServerConfig struct {
	Timeout              uint32            `protobuf:"varint,1,opt,name=timeout" json:"timeout,omitempty"` // Deprecated: Do not use.
	Accounts             map[string]string `protobuf:"bytes,2,rep,name=accounts" json:"accounts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	AllowTransparent     bool              `protobuf:"varint,3,opt,name=allow_transparent,json=allowTransparent" json:"allow_transparent,omitempty"`
	UserLevel            uint32            `protobuf:"varint,4,opt,name=user_level,json=userLevel" json:"user_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ServerConfig) Reset()         { *m = ServerConfig{} }
func (m *ServerConfig) String() string { return proto.CompactTextString(m) }
func (*ServerConfig) ProtoMessage()    {}
func (*ServerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_593308a66adc1606, []int{0}
}
func (m *ServerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerConfig.Unmarshal(m, b)
}
func (m *ServerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerConfig.Marshal(b, m, deterministic)
}
func (dst *ServerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerConfig.Merge(dst, src)
}
func (m *ServerConfig) XXX_Size() int {
	return xxx_messageInfo_ServerConfig.Size(m)
}
func (m *ServerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ServerConfig proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *ServerConfig) GetTimeout() uint32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *ServerConfig) GetAccounts() map[string]string {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *ServerConfig) GetAllowTransparent() bool {
	if m != nil {
		return m.AllowTransparent
	}
	return false
}

func (m *ServerConfig) GetUserLevel() uint32 {
	if m != nil {
		return m.UserLevel
	}
	return 0
}

// ClientConfig for HTTP proxy client.
type ClientConfig struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientConfig) Reset()         { *m = ClientConfig{} }
func (m *ClientConfig) String() string { return proto.CompactTextString(m) }
func (*ClientConfig) ProtoMessage()    {}
func (*ClientConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_593308a66adc1606, []int{1}
}
func (m *ClientConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientConfig.Unmarshal(m, b)
}
func (m *ClientConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientConfig.Marshal(b, m, deterministic)
}
func (dst *ClientConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientConfig.Merge(dst, src)
}
func (m *ClientConfig) XXX_Size() int {
	return xxx_messageInfo_ClientConfig.Size(m)
}
func (m *ClientConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ClientConfig proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ServerConfig)(nil), "v2ray.core.proxy.http.ServerConfig")
	proto.RegisterMapType((map[string]string)(nil), "v2ray.core.proxy.http.ServerConfig.AccountsEntry")
	proto.RegisterType((*ClientConfig)(nil), "v2ray.core.proxy.http.ClientConfig")
}

func init() {
	proto.RegisterFile("v2ray.com/core/proxy/http/config.proto", fileDescriptor_config_593308a66adc1606)
}

var fileDescriptor_config_593308a66adc1606 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4a, 0x33, 0x31,
	0x14, 0xc5, 0x99, 0x69, 0xbf, 0xcf, 0xf6, 0xda, 0x4a, 0x0d, 0x16, 0x46, 0x51, 0x28, 0x5d, 0x48,
	0x41, 0xc8, 0x60, 0xdd, 0x88, 0x5d, 0xd9, 0x22, 0xb8, 0x50, 0x28, 0x51, 0x5c, 0xb8, 0x29, 0x31,
	0x5c, 0xb5, 0x98, 0x26, 0x43, 0xe6, 0xce, 0xe8, 0xec, 0x7d, 0x1a, 0x9f, 0x52, 0x92, 0x5a, 0xff,
	0x40, 0x57, 0x49, 0x7e, 0xe7, 0xe4, 0xe4, 0x9e, 0xc0, 0x61, 0x39, 0x74, 0xb2, 0xe2, 0xca, 0x2e,
	0x52, 0x65, 0x1d, 0xa6, 0x99, 0xb3, 0x6f, 0x55, 0xfa, 0x4c, 0x94, 0xa5, 0xca, 0x9a, 0xc7, 0xf9,
	0x13, 0xcf, 0x9c, 0x25, 0xcb, 0xba, 0x2b, 0x9f, 0x43, 0x1e, 0x3c, 0xdc, 0x7b, 0xfa, 0xef, 0x31,
	0xb4, 0x6e, 0xd0, 0x95, 0xe8, 0x26, 0xc1, 0xcd, 0xf6, 0x61, 0x83, 0xe6, 0x0b, 0xb4, 0x05, 0x25,
	0x51, 0x2f, 0x1a, 0xb4, 0xc7, 0x71, 0x12, 0x89, 0x15, 0x62, 0xd7, 0xd0, 0x90, 0x4a, 0xd9, 0xc2,
	0x50, 0x9e, 0xc4, 0xbd, 0xda, 0x60, 0x73, 0x78, 0xcc, 0xd7, 0x06, 0xf3, 0xdf, 0xa1, 0xfc, 0xfc,
	0xeb, 0xce, 0x85, 0x21, 0x57, 0x89, 0xef, 0x08, 0x76, 0x04, 0xdb, 0x52, 0x6b, 0xfb, 0x3a, 0x23,
	0x27, 0x4d, 0x9e, 0x49, 0x87, 0x86, 0x92, 0x5a, 0x2f, 0x1a, 0x34, 0x44, 0x27, 0x08, 0xb7, 0x3f,
	0x9c, 0x1d, 0x00, 0x14, 0x39, 0xba, 0x99, 0xc6, 0x12, 0x75, 0x52, 0xf7, 0xc3, 0x89, 0xa6, 0x27,
	0x57, 0x1e, 0xec, 0x8d, 0xa0, 0xfd, 0xe7, 0x19, 0xd6, 0x81, 0xda, 0x0b, 0x56, 0xa1, 0x45, 0x53,
	0xf8, 0x2d, 0xdb, 0x81, 0x7f, 0xa5, 0xd4, 0x05, 0x26, 0x71, 0x60, 0xcb, 0xc3, 0x59, 0x7c, 0x1a,
	0xf5, 0xb7, 0xa0, 0x35, 0xd1, 0x73, 0x34, 0xb4, 0x1c, 0x78, 0x3c, 0x82, 0x5d, 0x65, 0x17, 0xeb,
	0xab, 0x4d, 0xa3, 0xfb, 0xba, 0x5f, 0x3f, 0xe2, 0xee, 0xdd, 0x50, 0xc8, 0x8a, 0x4f, 0xbc, 0x3e,
	0x0d, 0xfa, 0x25, 0x51, 0xf6, 0xf0, 0x3f, 0xfc, 0xf8, 0xc9, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x69, 0x94, 0x9f, 0xa7, 0x9b, 0x01, 0x00, 0x00,
}
