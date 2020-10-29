// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/api1.proto

package helloworld

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TestReq struct {
	Input                string   `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestReq) Reset()         { *m = TestReq{} }
func (m *TestReq) String() string { return proto.CompactTextString(m) }
func (*TestReq) ProtoMessage()    {}
func (*TestReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9adf9d64b96096c4, []int{0}
}

func (m *TestReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestReq.Unmarshal(m, b)
}
func (m *TestReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestReq.Marshal(b, m, deterministic)
}
func (m *TestReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestReq.Merge(m, src)
}
func (m *TestReq) XXX_Size() int {
	return xxx_messageInfo_TestReq.Size(m)
}
func (m *TestReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TestReq.DiscardUnknown(m)
}

var xxx_messageInfo_TestReq proto.InternalMessageInfo

func (m *TestReq) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

type TestApiData struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	SrvVersion           string   `protobuf:"bytes,2,opt,name=srvVersion,proto3" json:"srvVersion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestApiData) Reset()         { *m = TestApiData{} }
func (m *TestApiData) String() string { return proto.CompactTextString(m) }
func (*TestApiData) ProtoMessage()    {}
func (*TestApiData) Descriptor() ([]byte, []int) {
	return fileDescriptor_9adf9d64b96096c4, []int{1}
}

func (m *TestApiData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestApiData.Unmarshal(m, b)
}
func (m *TestApiData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestApiData.Marshal(b, m, deterministic)
}
func (m *TestApiData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestApiData.Merge(m, src)
}
func (m *TestApiData) XXX_Size() int {
	return xxx_messageInfo_TestApiData.Size(m)
}
func (m *TestApiData) XXX_DiscardUnknown() {
	xxx_messageInfo_TestApiData.DiscardUnknown(m)
}

var xxx_messageInfo_TestApiData proto.InternalMessageInfo

func (m *TestApiData) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *TestApiData) GetSrvVersion() string {
	if m != nil {
		return m.SrvVersion
	}
	return ""
}

type TestApiOutput struct {
	Code                 int32        `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string       `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	NowTime              int64        `protobuf:"varint,3,opt,name=nowTime,proto3" json:"nowTime,omitempty"`
	Data                 *TestApiData `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TestApiOutput) Reset()         { *m = TestApiOutput{} }
func (m *TestApiOutput) String() string { return proto.CompactTextString(m) }
func (*TestApiOutput) ProtoMessage()    {}
func (*TestApiOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_9adf9d64b96096c4, []int{2}
}

func (m *TestApiOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestApiOutput.Unmarshal(m, b)
}
func (m *TestApiOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestApiOutput.Marshal(b, m, deterministic)
}
func (m *TestApiOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestApiOutput.Merge(m, src)
}
func (m *TestApiOutput) XXX_Size() int {
	return xxx_messageInfo_TestApiOutput.Size(m)
}
func (m *TestApiOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_TestApiOutput.DiscardUnknown(m)
}

var xxx_messageInfo_TestApiOutput proto.InternalMessageInfo

func (m *TestApiOutput) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *TestApiOutput) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *TestApiOutput) GetNowTime() int64 {
	if m != nil {
		return m.NowTime
	}
	return 0
}

func (m *TestApiOutput) GetData() *TestApiData {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*TestReq)(nil), "helloworld.TestReq")
	proto.RegisterType((*TestApiData)(nil), "helloworld.TestApiData")
	proto.RegisterType((*TestApiOutput)(nil), "helloworld.TestApiOutput")
}

func init() {
	proto.RegisterFile("proto/api1.proto", fileDescriptor_9adf9d64b96096c4)
}

var fileDescriptor_9adf9d64b96096c4 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x50, 0xdb, 0x6a, 0x32, 0x31,
	0x10, 0x66, 0x7f, 0x0f, 0xfb, 0x3b, 0x52, 0x28, 0xb9, 0xa9, 0x15, 0x21, 0x22, 0x2d, 0x88, 0xd0,
	0x5d, 0xda, 0xd2, 0x07, 0x50, 0x7a, 0x5f, 0x48, 0xc5, 0xdb, 0x12, 0x35, 0x8d, 0x81, 0x35, 0x93,
	0xee, 0x66, 0xf5, 0x8d, 0xfa, 0x58, 0xfb, 0x00, 0xfb, 0x14, 0x25, 0xc9, 0x2e, 0x7a, 0x33, 0x99,
	0xf9, 0x0e, 0xf3, 0x91, 0x81, 0x5b, 0x93, 0xa3, 0xc5, 0x94, 0x1b, 0xf5, 0x9c, 0xf8, 0x96, 0xc0,
	0x41, 0x64, 0x19, 0x9e, 0x31, 0xcf, 0xf6, 0xe3, 0x27, 0xa9, 0xec, 0xa1, 0xdc, 0x26, 0x3b, 0x3c,
	0xa6, 0x12, 0x25, 0xa6, 0x5e, 0xb2, 0x2d, 0xbf, 0xfd, 0x14, 0xac, 0xae, 0x0b, 0xd6, 0xf1, 0x44,
	0x22, 0xca, 0x4c, 0xb8, 0x6d, 0x29, 0xd7, 0x1a, 0x2d, 0xb7, 0x0a, 0x75, 0x11, 0xd8, 0xd9, 0x02,
	0xe2, 0xb5, 0x28, 0x2c, 0x13, 0x3f, 0x84, 0x42, 0x4f, 0x69, 0x53, 0xda, 0x51, 0x34, 0x8d, 0xe6,
	0x83, 0xd5, 0xa0, 0xae, 0x68, 0x00, 0x58, 0x78, 0x66, 0x5f, 0x30, 0x74, 0xda, 0xa5, 0x51, 0xef,
	0xdc, 0x72, 0xf2, 0x00, 0xf1, 0x49, 0xe4, 0x85, 0x42, 0xdd, 0x38, 0xa0, 0xae, 0x68, 0x7f, 0x69,
	0xd4, 0x46, 0xe4, 0xac, 0xa5, 0xc8, 0x02, 0xa0, 0xc8, 0x4f, 0x9b, 0x46, 0xf8, 0xef, 0x22, 0xfc,
	0xf4, 0x28, 0xbb, 0x62, 0x67, 0xbf, 0x11, 0xdc, 0x34, 0x09, 0x1f, 0xa5, 0x35, 0xa5, 0x25, 0x13,
	0xe8, 0xee, 0x70, 0x2f, 0x7c, 0x40, 0x6f, 0xf5, 0xbf, 0xae, 0xa8, 0x9f, 0x99, 0xaf, 0xe4, 0x1e,
	0x3a, 0xc7, 0x42, 0x36, 0x4b, 0xe3, 0xba, 0xa2, 0x6e, 0x64, 0xae, 0x90, 0x47, 0x88, 0x35, 0x9e,
	0xd7, 0xea, 0x28, 0x46, 0x9d, 0x69, 0x34, 0xef, 0xac, 0x86, 0x75, 0x45, 0x5b, 0x88, 0xb5, 0x0d,
	0x79, 0x83, 0xee, 0x9e, 0x5b, 0x3e, 0xea, 0x4e, 0xa3, 0xf9, 0xf0, 0xe5, 0x2e, 0xb9, 0x9c, 0x39,
	0xb9, 0xfa, 0x6a, 0x08, 0x76, 0x42, 0xe6, 0xeb, 0xb6, 0xef, 0x8f, 0xf7, 0xfa, 0x17, 0x00, 0x00,
	0xff, 0xff, 0x07, 0x70, 0x54, 0x3d, 0xa9, 0x01, 0x00, 0x00,
}
