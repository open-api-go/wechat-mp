// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wx-miniprogram/login.proto

package wx_miniprogram // import "github.com/dev-openapi/wx-miniprogram"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Code2SessionReq struct {
	Appid                string   `protobuf:"bytes,1,opt,name=appid" json:"appid,omitempty"`
	Secret               string   `protobuf:"bytes,2,opt,name=secret" json:"secret,omitempty"`
	JsCode               string   `protobuf:"bytes,3,opt,name=js_code,json=jsCode" json:"js_code,omitempty"`
	GrantType            string   `protobuf:"bytes,4,opt,name=grant_type,json=grantType" json:"grant_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Code2SessionReq) Reset()         { *m = Code2SessionReq{} }
func (m *Code2SessionReq) String() string { return proto.CompactTextString(m) }
func (*Code2SessionReq) ProtoMessage()    {}
func (*Code2SessionReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_login_159e9bd379d0dbfb, []int{0}
}
func (m *Code2SessionReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Code2SessionReq.Unmarshal(m, b)
}
func (m *Code2SessionReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Code2SessionReq.Marshal(b, m, deterministic)
}
func (dst *Code2SessionReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Code2SessionReq.Merge(dst, src)
}
func (m *Code2SessionReq) XXX_Size() int {
	return xxx_messageInfo_Code2SessionReq.Size(m)
}
func (m *Code2SessionReq) XXX_DiscardUnknown() {
	xxx_messageInfo_Code2SessionReq.DiscardUnknown(m)
}

var xxx_messageInfo_Code2SessionReq proto.InternalMessageInfo

func (m *Code2SessionReq) GetAppid() string {
	if m != nil {
		return m.Appid
	}
	return ""
}

func (m *Code2SessionReq) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *Code2SessionReq) GetJsCode() string {
	if m != nil {
		return m.JsCode
	}
	return ""
}

func (m *Code2SessionReq) GetGrantType() string {
	if m != nil {
		return m.GrantType
	}
	return ""
}

type Code2SessionRes struct {
	Errcode              int32    `protobuf:"varint,1,opt,name=errcode" json:"errcode,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=errmsg" json:"errmsg,omitempty"`
	Openid               string   `protobuf:"bytes,3,opt,name=openid" json:"openid,omitempty"`
	Unionid              string   `protobuf:"bytes,4,opt,name=unionid" json:"unionid,omitempty"`
	SessionKey           string   `protobuf:"bytes,5,opt,name=session_key,json=sessionKey" json:"session_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Code2SessionRes) Reset()         { *m = Code2SessionRes{} }
func (m *Code2SessionRes) String() string { return proto.CompactTextString(m) }
func (*Code2SessionRes) ProtoMessage()    {}
func (*Code2SessionRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_login_159e9bd379d0dbfb, []int{1}
}
func (m *Code2SessionRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Code2SessionRes.Unmarshal(m, b)
}
func (m *Code2SessionRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Code2SessionRes.Marshal(b, m, deterministic)
}
func (dst *Code2SessionRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Code2SessionRes.Merge(dst, src)
}
func (m *Code2SessionRes) XXX_Size() int {
	return xxx_messageInfo_Code2SessionRes.Size(m)
}
func (m *Code2SessionRes) XXX_DiscardUnknown() {
	xxx_messageInfo_Code2SessionRes.DiscardUnknown(m)
}

var xxx_messageInfo_Code2SessionRes proto.InternalMessageInfo

func (m *Code2SessionRes) GetErrcode() int32 {
	if m != nil {
		return m.Errcode
	}
	return 0
}

func (m *Code2SessionRes) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *Code2SessionRes) GetOpenid() string {
	if m != nil {
		return m.Openid
	}
	return ""
}

func (m *Code2SessionRes) GetUnionid() string {
	if m != nil {
		return m.Unionid
	}
	return ""
}

func (m *Code2SessionRes) GetSessionKey() string {
	if m != nil {
		return m.SessionKey
	}
	return ""
}

func init() {
	proto.RegisterType((*Code2SessionReq)(nil), "api.weixin.qq.com.Code2SessionReq")
	proto.RegisterType((*Code2SessionRes)(nil), "api.weixin.qq.com.Code2SessionRes")
}

func init() { proto.RegisterFile("wx-miniprogram/login.proto", fileDescriptor_login_159e9bd379d0dbfb) }

var fileDescriptor_login_159e9bd379d0dbfb = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcb, 0x4e, 0xeb, 0x30,
	0x10, 0x86, 0x95, 0x9e, 0xd3, 0x56, 0xf5, 0xa9, 0x74, 0x84, 0xb9, 0x45, 0x05, 0x04, 0x8a, 0x84,
	0x60, 0xd3, 0x44, 0x2a, 0x6f, 0x00, 0x4b, 0x58, 0xb5, 0xac, 0xd8, 0x54, 0x69, 0x32, 0x0a, 0x53,
	0x1a, 0x8f, 0xeb, 0x49, 0x2f, 0xd9, 0xc2, 0x23, 0x20, 0xf1, 0x62, 0xbc, 0x02, 0x0f, 0x82, 0x1c,
	0x37, 0x12, 0x94, 0x05, 0xcb, 0xff, 0xff, 0xed, 0xf9, 0xe6, 0x22, 0x7a, 0xab, 0x75, 0x3f, 0x47,
	0x85, 0xda, 0x50, 0x66, 0xe2, 0x3c, 0x9a, 0x51, 0x86, 0x2a, 0xd4, 0x86, 0x0a, 0x92, 0x3b, 0xb1,
	0xc6, 0x70, 0x05, 0xb8, 0x46, 0x15, 0xce, 0xe7, 0x61, 0x42, 0x79, 0xef, 0x38, 0x23, 0xca, 0x66,
	0x10, 0xc5, 0x1a, 0xa3, 0x58, 0x29, 0x2a, 0xe2, 0x02, 0x49, 0xb1, 0xfb, 0x10, 0xac, 0xc4, 0xff,
	0x1b, 0x4a, 0x61, 0x30, 0x02, 0x66, 0x24, 0x35, 0x84, 0xb9, 0xdc, 0x13, 0xcd, 0x58, 0x6b, 0x4c,
	0x7d, 0xef, 0xcc, 0xbb, 0xec, 0x0c, 0x9d, 0x90, 0x07, 0xa2, 0xc5, 0x90, 0x18, 0x28, 0xfc, 0x46,
	0x65, 0x6f, 0x94, 0x3c, 0x14, 0xed, 0x29, 0x8f, 0x13, 0x4a, 0xc1, 0xff, 0xe3, 0x82, 0x29, 0xdb,
	0x8a, 0xf2, 0x44, 0x88, 0xcc, 0xc4, 0xaa, 0x18, 0x17, 0xa5, 0x06, 0xff, 0x6f, 0x95, 0x75, 0x2a,
	0xe7, 0xbe, 0xd4, 0x10, 0xbc, 0x79, 0xdb, 0x64, 0x96, 0xbe, 0x68, 0x83, 0x31, 0x55, 0x2d, 0xcb,
	0x6e, 0x0e, 0x6b, 0x69, 0xe9, 0x60, 0x4c, 0xce, 0x59, 0x4d, 0x77, 0xca, 0xfa, 0xa4, 0x41, 0x61,
	0x5a, 0xc3, 0x9d, 0xb2, 0x95, 0x16, 0x0a, 0xc9, 0x06, 0x8e, 0x5c, 0x4b, 0x79, 0x2a, 0xfe, 0xb1,
	0x23, 0x8e, 0x9f, 0xa0, 0xf4, 0x9b, 0x55, 0x2a, 0x36, 0xd6, 0x2d, 0x94, 0x83, 0x17, 0x4f, 0x74,
	0xef, 0xec, 0x4a, 0x47, 0x60, 0x96, 0x98, 0x80, 0x64, 0xd1, 0xfd, 0xda, 0xa8, 0x0c, 0xc2, 0x1f,
	0x4b, 0x0e, 0xb7, 0x76, 0xd8, 0xfb, 0xfd, 0x0d, 0x07, 0x47, 0xcf, 0xef, 0x1f, 0xaf, 0x8d, 0x7d,
	0xb9, 0x1b, 0xb1, 0xe2, 0x68, 0xca, 0x76, 0xd0, 0xc1, 0xa6, 0x91, 0xeb, 0x8b, 0x87, 0xf3, 0x0c,
	0x8b, 0xc7, 0xc5, 0xc4, 0x7e, 0x8d, 0x52, 0x58, 0xf6, 0xed, 0x64, 0xf6, 0x86, 0xdf, 0xaf, 0x3f,
	0x69, 0x55, 0x77, 0xbc, 0xfa, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x31, 0xf6, 0x33, 0x16, 0x02,
	0x00, 0x00,
}