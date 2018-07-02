// Code generated by protoc-gen-go. DO NOT EDIT.
// source: steammessages_cloud.steamworkssdk.proto

package protocol

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

type CCloud_GetUploadServerInfo_Request struct {
	Appid                *uint32  `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CCloud_GetUploadServerInfo_Request) Reset()         { *m = CCloud_GetUploadServerInfo_Request{} }
func (m *CCloud_GetUploadServerInfo_Request) String() string { return proto.CompactTextString(m) }
func (*CCloud_GetUploadServerInfo_Request) ProtoMessage()    {}
func (*CCloud_GetUploadServerInfo_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_cloud_steamworkssdk_7285ed0412ef3056, []int{0}
}
func (m *CCloud_GetUploadServerInfo_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CCloud_GetUploadServerInfo_Request.Unmarshal(m, b)
}
func (m *CCloud_GetUploadServerInfo_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CCloud_GetUploadServerInfo_Request.Marshal(b, m, deterministic)
}
func (dst *CCloud_GetUploadServerInfo_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CCloud_GetUploadServerInfo_Request.Merge(dst, src)
}
func (m *CCloud_GetUploadServerInfo_Request) XXX_Size() int {
	return xxx_messageInfo_CCloud_GetUploadServerInfo_Request.Size(m)
}
func (m *CCloud_GetUploadServerInfo_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_CCloud_GetUploadServerInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_CCloud_GetUploadServerInfo_Request proto.InternalMessageInfo

func (m *CCloud_GetUploadServerInfo_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

type CCloud_GetUploadServerInfo_Response struct {
	ServerUrl            *string  `protobuf:"bytes,1,opt,name=server_url,json=serverUrl" json:"server_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CCloud_GetUploadServerInfo_Response) Reset()         { *m = CCloud_GetUploadServerInfo_Response{} }
func (m *CCloud_GetUploadServerInfo_Response) String() string { return proto.CompactTextString(m) }
func (*CCloud_GetUploadServerInfo_Response) ProtoMessage()    {}
func (*CCloud_GetUploadServerInfo_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_cloud_steamworkssdk_7285ed0412ef3056, []int{1}
}
func (m *CCloud_GetUploadServerInfo_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CCloud_GetUploadServerInfo_Response.Unmarshal(m, b)
}
func (m *CCloud_GetUploadServerInfo_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CCloud_GetUploadServerInfo_Response.Marshal(b, m, deterministic)
}
func (dst *CCloud_GetUploadServerInfo_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CCloud_GetUploadServerInfo_Response.Merge(dst, src)
}
func (m *CCloud_GetUploadServerInfo_Response) XXX_Size() int {
	return xxx_messageInfo_CCloud_GetUploadServerInfo_Response.Size(m)
}
func (m *CCloud_GetUploadServerInfo_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_CCloud_GetUploadServerInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_CCloud_GetUploadServerInfo_Response proto.InternalMessageInfo

func (m *CCloud_GetUploadServerInfo_Response) GetServerUrl() string {
	if m != nil && m.ServerUrl != nil {
		return *m.ServerUrl
	}
	return ""
}

type CCloud_GetFileDetails_Request struct {
	Ugcid                *uint64  `protobuf:"varint,1,opt,name=ugcid" json:"ugcid,omitempty"`
	Appid                *uint32  `protobuf:"varint,2,opt,name=appid" json:"appid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CCloud_GetFileDetails_Request) Reset()         { *m = CCloud_GetFileDetails_Request{} }
func (m *CCloud_GetFileDetails_Request) String() string { return proto.CompactTextString(m) }
func (*CCloud_GetFileDetails_Request) ProtoMessage()    {}
func (*CCloud_GetFileDetails_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_cloud_steamworkssdk_7285ed0412ef3056, []int{2}
}
func (m *CCloud_GetFileDetails_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CCloud_GetFileDetails_Request.Unmarshal(m, b)
}
func (m *CCloud_GetFileDetails_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CCloud_GetFileDetails_Request.Marshal(b, m, deterministic)
}
func (dst *CCloud_GetFileDetails_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CCloud_GetFileDetails_Request.Merge(dst, src)
}
func (m *CCloud_GetFileDetails_Request) XXX_Size() int {
	return xxx_messageInfo_CCloud_GetFileDetails_Request.Size(m)
}
func (m *CCloud_GetFileDetails_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_CCloud_GetFileDetails_Request.DiscardUnknown(m)
}

var xxx_messageInfo_CCloud_GetFileDetails_Request proto.InternalMessageInfo

func (m *CCloud_GetFileDetails_Request) GetUgcid() uint64 {
	if m != nil && m.Ugcid != nil {
		return *m.Ugcid
	}
	return 0
}

func (m *CCloud_GetFileDetails_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

type CCloud_UserFile struct {
	Appid                *uint32  `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	Ugcid                *uint64  `protobuf:"varint,2,opt,name=ugcid" json:"ugcid,omitempty"`
	Filename             *string  `protobuf:"bytes,3,opt,name=filename" json:"filename,omitempty"`
	Timestamp            *uint64  `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	FileSize             *uint32  `protobuf:"varint,5,opt,name=file_size,json=fileSize" json:"file_size,omitempty"`
	Url                  *string  `protobuf:"bytes,6,opt,name=url" json:"url,omitempty"`
	SteamidCreator       *uint64  `protobuf:"fixed64,7,opt,name=steamid_creator,json=steamidCreator" json:"steamid_creator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CCloud_UserFile) Reset()         { *m = CCloud_UserFile{} }
func (m *CCloud_UserFile) String() string { return proto.CompactTextString(m) }
func (*CCloud_UserFile) ProtoMessage()    {}
func (*CCloud_UserFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_cloud_steamworkssdk_7285ed0412ef3056, []int{3}
}
func (m *CCloud_UserFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CCloud_UserFile.Unmarshal(m, b)
}
func (m *CCloud_UserFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CCloud_UserFile.Marshal(b, m, deterministic)
}
func (dst *CCloud_UserFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CCloud_UserFile.Merge(dst, src)
}
func (m *CCloud_UserFile) XXX_Size() int {
	return xxx_messageInfo_CCloud_UserFile.Size(m)
}
func (m *CCloud_UserFile) XXX_DiscardUnknown() {
	xxx_messageInfo_CCloud_UserFile.DiscardUnknown(m)
}

var xxx_messageInfo_CCloud_UserFile proto.InternalMessageInfo

func (m *CCloud_UserFile) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CCloud_UserFile) GetUgcid() uint64 {
	if m != nil && m.Ugcid != nil {
		return *m.Ugcid
	}
	return 0
}

func (m *CCloud_UserFile) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *CCloud_UserFile) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *CCloud_UserFile) GetFileSize() uint32 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

func (m *CCloud_UserFile) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *CCloud_UserFile) GetSteamidCreator() uint64 {
	if m != nil && m.SteamidCreator != nil {
		return *m.SteamidCreator
	}
	return 0
}

type CCloud_GetFileDetails_Response struct {
	Details              *CCloud_UserFile `protobuf:"bytes,1,opt,name=details" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CCloud_GetFileDetails_Response) Reset()         { *m = CCloud_GetFileDetails_Response{} }
func (m *CCloud_GetFileDetails_Response) String() string { return proto.CompactTextString(m) }
func (*CCloud_GetFileDetails_Response) ProtoMessage()    {}
func (*CCloud_GetFileDetails_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_cloud_steamworkssdk_7285ed0412ef3056, []int{4}
}
func (m *CCloud_GetFileDetails_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CCloud_GetFileDetails_Response.Unmarshal(m, b)
}
func (m *CCloud_GetFileDetails_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CCloud_GetFileDetails_Response.Marshal(b, m, deterministic)
}
func (dst *CCloud_GetFileDetails_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CCloud_GetFileDetails_Response.Merge(dst, src)
}
func (m *CCloud_GetFileDetails_Response) XXX_Size() int {
	return xxx_messageInfo_CCloud_GetFileDetails_Response.Size(m)
}
func (m *CCloud_GetFileDetails_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_CCloud_GetFileDetails_Response.DiscardUnknown(m)
}

var xxx_messageInfo_CCloud_GetFileDetails_Response proto.InternalMessageInfo

func (m *CCloud_GetFileDetails_Response) GetDetails() *CCloud_UserFile {
	if m != nil {
		return m.Details
	}
	return nil
}

type CCloud_EnumerateUserFiles_Request struct {
	Appid                *uint32  `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	ExtendedDetails      *bool    `protobuf:"varint,2,opt,name=extended_details,json=extendedDetails" json:"extended_details,omitempty"`
	Count                *uint32  `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	StartIndex           *uint32  `protobuf:"varint,4,opt,name=start_index,json=startIndex" json:"start_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CCloud_EnumerateUserFiles_Request) Reset()         { *m = CCloud_EnumerateUserFiles_Request{} }
func (m *CCloud_EnumerateUserFiles_Request) String() string { return proto.CompactTextString(m) }
func (*CCloud_EnumerateUserFiles_Request) ProtoMessage()    {}
func (*CCloud_EnumerateUserFiles_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_cloud_steamworkssdk_7285ed0412ef3056, []int{5}
}
func (m *CCloud_EnumerateUserFiles_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CCloud_EnumerateUserFiles_Request.Unmarshal(m, b)
}
func (m *CCloud_EnumerateUserFiles_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CCloud_EnumerateUserFiles_Request.Marshal(b, m, deterministic)
}
func (dst *CCloud_EnumerateUserFiles_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CCloud_EnumerateUserFiles_Request.Merge(dst, src)
}
func (m *CCloud_EnumerateUserFiles_Request) XXX_Size() int {
	return xxx_messageInfo_CCloud_EnumerateUserFiles_Request.Size(m)
}
func (m *CCloud_EnumerateUserFiles_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_CCloud_EnumerateUserFiles_Request.DiscardUnknown(m)
}

var xxx_messageInfo_CCloud_EnumerateUserFiles_Request proto.InternalMessageInfo

func (m *CCloud_EnumerateUserFiles_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CCloud_EnumerateUserFiles_Request) GetExtendedDetails() bool {
	if m != nil && m.ExtendedDetails != nil {
		return *m.ExtendedDetails
	}
	return false
}

func (m *CCloud_EnumerateUserFiles_Request) GetCount() uint32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *CCloud_EnumerateUserFiles_Request) GetStartIndex() uint32 {
	if m != nil && m.StartIndex != nil {
		return *m.StartIndex
	}
	return 0
}

type CCloud_EnumerateUserFiles_Response struct {
	Files                []*CCloud_UserFile `protobuf:"bytes,1,rep,name=files" json:"files,omitempty"`
	TotalFiles           *uint32            `protobuf:"varint,2,opt,name=total_files,json=totalFiles" json:"total_files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CCloud_EnumerateUserFiles_Response) Reset()         { *m = CCloud_EnumerateUserFiles_Response{} }
func (m *CCloud_EnumerateUserFiles_Response) String() string { return proto.CompactTextString(m) }
func (*CCloud_EnumerateUserFiles_Response) ProtoMessage()    {}
func (*CCloud_EnumerateUserFiles_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_cloud_steamworkssdk_7285ed0412ef3056, []int{6}
}
func (m *CCloud_EnumerateUserFiles_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CCloud_EnumerateUserFiles_Response.Unmarshal(m, b)
}
func (m *CCloud_EnumerateUserFiles_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CCloud_EnumerateUserFiles_Response.Marshal(b, m, deterministic)
}
func (dst *CCloud_EnumerateUserFiles_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CCloud_EnumerateUserFiles_Response.Merge(dst, src)
}
func (m *CCloud_EnumerateUserFiles_Response) XXX_Size() int {
	return xxx_messageInfo_CCloud_EnumerateUserFiles_Response.Size(m)
}
func (m *CCloud_EnumerateUserFiles_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_CCloud_EnumerateUserFiles_Response.DiscardUnknown(m)
}

var xxx_messageInfo_CCloud_EnumerateUserFiles_Response proto.InternalMessageInfo

func (m *CCloud_EnumerateUserFiles_Response) GetFiles() []*CCloud_UserFile {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *CCloud_EnumerateUserFiles_Response) GetTotalFiles() uint32 {
	if m != nil && m.TotalFiles != nil {
		return *m.TotalFiles
	}
	return 0
}

type CCloud_Delete_Request struct {
	Filename             *string  `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	Appid                *uint32  `protobuf:"varint,2,opt,name=appid" json:"appid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CCloud_Delete_Request) Reset()         { *m = CCloud_Delete_Request{} }
func (m *CCloud_Delete_Request) String() string { return proto.CompactTextString(m) }
func (*CCloud_Delete_Request) ProtoMessage()    {}
func (*CCloud_Delete_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_cloud_steamworkssdk_7285ed0412ef3056, []int{7}
}
func (m *CCloud_Delete_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CCloud_Delete_Request.Unmarshal(m, b)
}
func (m *CCloud_Delete_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CCloud_Delete_Request.Marshal(b, m, deterministic)
}
func (dst *CCloud_Delete_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CCloud_Delete_Request.Merge(dst, src)
}
func (m *CCloud_Delete_Request) XXX_Size() int {
	return xxx_messageInfo_CCloud_Delete_Request.Size(m)
}
func (m *CCloud_Delete_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_CCloud_Delete_Request.DiscardUnknown(m)
}

var xxx_messageInfo_CCloud_Delete_Request proto.InternalMessageInfo

func (m *CCloud_Delete_Request) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *CCloud_Delete_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

type CCloud_Delete_Response struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CCloud_Delete_Response) Reset()         { *m = CCloud_Delete_Response{} }
func (m *CCloud_Delete_Response) String() string { return proto.CompactTextString(m) }
func (*CCloud_Delete_Response) ProtoMessage()    {}
func (*CCloud_Delete_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_cloud_steamworkssdk_7285ed0412ef3056, []int{8}
}
func (m *CCloud_Delete_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CCloud_Delete_Response.Unmarshal(m, b)
}
func (m *CCloud_Delete_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CCloud_Delete_Response.Marshal(b, m, deterministic)
}
func (dst *CCloud_Delete_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CCloud_Delete_Response.Merge(dst, src)
}
func (m *CCloud_Delete_Response) XXX_Size() int {
	return xxx_messageInfo_CCloud_Delete_Response.Size(m)
}
func (m *CCloud_Delete_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_CCloud_Delete_Response.DiscardUnknown(m)
}

var xxx_messageInfo_CCloud_Delete_Response proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CCloud_GetUploadServerInfo_Request)(nil), "CCloud_GetUploadServerInfo_Request")
	proto.RegisterType((*CCloud_GetUploadServerInfo_Response)(nil), "CCloud_GetUploadServerInfo_Response")
	proto.RegisterType((*CCloud_GetFileDetails_Request)(nil), "CCloud_GetFileDetails_Request")
	proto.RegisterType((*CCloud_UserFile)(nil), "CCloud_UserFile")
	proto.RegisterType((*CCloud_GetFileDetails_Response)(nil), "CCloud_GetFileDetails_Response")
	proto.RegisterType((*CCloud_EnumerateUserFiles_Request)(nil), "CCloud_EnumerateUserFiles_Request")
	proto.RegisterType((*CCloud_EnumerateUserFiles_Response)(nil), "CCloud_EnumerateUserFiles_Response")
	proto.RegisterType((*CCloud_Delete_Request)(nil), "CCloud_Delete_Request")
	proto.RegisterType((*CCloud_Delete_Response)(nil), "CCloud_Delete_Response")
}

func init() {
	proto.RegisterFile("steammessages_cloud.steamworkssdk.proto", fileDescriptor_steammessages_cloud_steamworkssdk_7285ed0412ef3056)
}

var fileDescriptor_steammessages_cloud_steamworkssdk_7285ed0412ef3056 = []byte{
	// 957 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x6e, 0xdc, 0x44,
	0x14, 0x96, 0x69, 0x36, 0x6d, 0x26, 0x4a, 0x13, 0x0d, 0x50, 0xac, 0x2d, 0x6d, 0x4e, 0x9d, 0x96,
	0xa4, 0x80, 0x86, 0xaa, 0x52, 0xb9, 0x80, 0xab, 0x6e, 0xb6, 0x44, 0x91, 0x8a, 0x90, 0x1c, 0x45,
	0x08, 0x21, 0xb4, 0x9a, 0x5d, 0x1f, 0x6f, 0x86, 0xd8, 0x33, 0x66, 0x66, 0xdc, 0x84, 0x0a, 0x21,
	0x64, 0x9e, 0x81, 0x1b, 0xc4, 0x23, 0x70, 0xeb, 0x1b, 0x5e, 0x82, 0x27, 0xe1, 0x1d, 0xd0, 0x8c,
	0xed, 0xcd, 0x6e, 0x58, 0x12, 0xc4, 0xdd, 0xee, 0x99, 0xf3, 0xf3, 0x9d, 0xf3, 0x7d, 0xe7, 0x98,
	0xec, 0x1a, 0x8b, 0x3c, 0xcf, 0xd1, 0x18, 0x3e, 0x45, 0x33, 0x9a, 0x64, 0xaa, 0x4c, 0x98, 0xb7,
	0x9d, 0x29, 0x7d, 0x6a, 0x4c, 0x72, 0xca, 0x0a, 0xad, 0xac, 0xea, 0xb3, 0x45, 0xc7, 0x52, 0x8a,
	0x54, 0x60, 0x32, 0x1a, 0x73, 0x83, 0xcb, 0xfc, 0xa3, 0x53, 0x12, 0xed, 0xef, 0xbb, 0x6c, 0xa3,
	0x03, 0xb4, 0xc7, 0x45, 0xa6, 0x78, 0x72, 0x84, 0xfa, 0x15, 0xea, 0x43, 0x99, 0xaa, 0x51, 0x8c,
	0xdf, 0x95, 0x68, 0x2c, 0x7d, 0x41, 0x7a, 0xbc, 0x28, 0x44, 0x12, 0x06, 0x10, 0xec, 0x6d, 0x0c,
	0x3e, 0xaa, 0xea, 0xf0, 0x83, 0xe7, 0x45, 0x01, 0x87, 0x43, 0xb0, 0x0a, 0xce, 0x4e, 0xc4, 0xe4,
	0x04, 0x38, 0xa4, 0x22, 0x43, 0x38, 0x13, 0x59, 0x06, 0x63, 0x84, 0xd2, 0xe7, 0xc2, 0x04, 0xac,
	0x62, 0x71, 0x13, 0x1d, 0x0d, 0xc9, 0xce, 0x95, 0xc5, 0x4c, 0xa1, 0xa4, 0x41, 0x7a, 0x8f, 0x10,
	0xe3, 0xcd, 0xa3, 0x52, 0x67, 0xbe, 0xe4, 0x5a, 0xbc, 0xd6, 0x58, 0x8e, 0x75, 0x16, 0xfd, 0x1a,
	0x90, 0x7b, 0x17, 0x69, 0x3e, 0x13, 0x19, 0x0e, 0xd1, 0x72, 0x91, 0x99, 0x19, 0xdc, 0x01, 0xe9,
	0x95, 0xd3, 0x49, 0x0b, 0x77, 0x65, 0xf0, 0x61, 0x55, 0x87, 0x7b, 0x87, 0x43, 0x50, 0x29, 0xd8,
	0x13, 0x04, 0x1f, 0xda, 0x80, 0xb5, 0x0a, 0xa6, 0x68, 0x21, 0x69, 0xe2, 0x21, 0x55, 0x9a, 0xc5,
	0x4d, 0x28, 0x7d, 0xd6, 0xb5, 0xfc, 0x86, 0x6f, 0x79, 0xbb, 0xaa, 0xc3, 0xbb, 0x5d, 0xcb, 0x27,
	0xd8, 0x84, 0x8f, 0x31, 0x53, 0x72, 0x6a, 0xe6, 0x5b, 0xfc, 0x33, 0x20, 0x9b, 0x2d, 0xb8, 0x63,
	0x83, 0xda, 0xa1, 0xa3, 0x6f, 0x2d, 0x4c, 0xaf, 0xf5, 0x74, 0xd6, 0x06, 0xa4, 0x2b, 0xb0, 0xd2,
	0x95, 0xed, 0x93, 0x5b, 0x2e, 0xb5, 0xe4, 0x39, 0x86, 0x37, 0x7c, 0xe7, 0xb3, 0xff, 0xf4, 0x5d,
	0xb2, 0x66, 0x45, 0x8e, 0xc6, 0xf2, 0xbc, 0x08, 0x57, 0x7c, 0xd4, 0x85, 0x81, 0xde, 0x25, 0x6b,
	0xce, 0x73, 0x64, 0xc4, 0x6b, 0x0c, 0x7b, 0xbe, 0x92, 0x0f, 0x3d, 0x12, 0xaf, 0x91, 0x6e, 0x91,
	0x1b, 0x6e, 0x96, 0xab, 0x3e, 0xa3, 0xfb, 0x49, 0x77, 0xc9, 0xa6, 0xd7, 0x83, 0x48, 0x46, 0x13,
	0x8d, 0xdc, 0x2a, 0x1d, 0xde, 0x84, 0x60, 0x6f, 0x35, 0xbe, 0xdd, 0x9a, 0xf7, 0x1b, 0x6b, 0xf4,
	0x92, 0xdc, 0xff, 0xb7, 0x69, 0xb7, 0x7c, 0xbd, 0x4f, 0x6e, 0xb6, 0x13, 0xf4, 0x1d, 0xae, 0x3f,
	0xdd, 0x62, 0x97, 0x46, 0x10, 0x77, 0x0e, 0xd1, 0x2f, 0x2b, 0xe4, 0x41, 0xfb, 0xf8, 0x42, 0x96,
	0x39, 0x6a, 0x6e, 0xb1, 0xf3, 0xba, 0x20, 0xf0, 0xd3, 0x45, 0xbd, 0x3d, 0xaa, 0xea, 0xf0, 0xc1,
	0x85, 0xde, 0xb0, 0x0b, 0x9c, 0x31, 0x61, 0x40, 0xa5, 0x1d, 0x05, 0xf4, 0xf7, 0x80, 0x6c, 0xe1,
	0xb9, 0x45, 0x99, 0x60, 0x32, 0xea, 0x80, 0xb9, 0x21, 0xdf, 0x1a, 0xfc, 0x14, 0x54, 0x75, 0xf8,
	0xc3, 0xde, 0x17, 0x85, 0x15, 0x4a, 0xf2, 0xec, 0x31, 0x1c, 0xa0, 0x85, 0xce, 0x77, 0x26, 0x83,
	0x31, 0x9f, 0x9c, 0x82, 0x92, 0x73, 0xd9, 0x53, 0x55, 0xca, 0x84, 0xc1, 0x10, 0x53, 0x5e, 0x66,
	0xd6, 0x11, 0x0e, 0x4a, 0x66, 0xdf, 0x83, 0x46, 0x5b, 0x6a, 0xe9, 0x44, 0x7e, 0x82, 0xc0, 0x1d,
	0xc0, 0x04, 0xb8, 0x4c, 0xe0, 0xf8, 0x60, 0xdf, 0xfd, 0x6c, 0x05, 0x37, 0x9f, 0x23, 0xde, 0xec,
	0xca, 0xb5, 0x63, 0xa4, 0x3f, 0x07, 0xa4, 0x37, 0x51, 0xa5, 0xb4, 0x9e, 0xef, 0x8d, 0x41, 0x5e,
	0xd5, 0xa1, 0x98, 0x83, 0xf8, 0x39, 0x3f, 0x17, 0x79, 0x99, 0x83, 0x2c, 0xf3, 0x31, 0x6a, 0x97,
	0x53, 0xa3, 0xe9, 0x50, 0x34, 0x00, 0x1a, 0xb0, 0xc2, 0xc0, 0x84, 0x67, 0xd9, 0x22, 0x4c, 0x0e,
	0x79, 0x1b, 0xaf, 0x52, 0x78, 0xf6, 0xe4, 0x49, 0x0b, 0xa6, 0x03, 0xce, 0xe2, 0xa6, 0x36, 0xfd,
	0x91, 0xac, 0x1b, 0xcb, 0xb5, 0x1d, 0x09, 0x99, 0xe0, 0xb9, 0x57, 0xd7, 0xc6, 0xe0, 0x9b, 0xaa,
	0x0e, 0xbf, 0x9a, 0x83, 0x72, 0xe4, 0x3c, 0x84, 0x9c, 0x82, 0x77, 0x72, 0xe9, 0xc7, 0x38, 0x15,
	0x72, 0x46, 0x88, 0x50, 0x12, 0xb8, 0x5d, 0x04, 0xe0, 0x06, 0xe0, 0xbd, 0xa4, 0x0b, 0x6c, 0x27,
	0x92, 0x09, 0x63, 0x59, 0x4c, 0x7c, 0xc5, 0x43, 0x97, 0x2b, 0xca, 0x67, 0x77, 0x68, 0xa9, 0x2c,
	0x5a, 0xa5, 0xbd, 0x47, 0x7a, 0x1e, 0x7f, 0x18, 0xc0, 0x8d, 0xa5, 0x3a, 0x6b, 0x9e, 0xe9, 0x36,
	0x59, 0xb7, 0xca, 0xf2, 0x6c, 0xd4, 0x78, 0xfb, 0x15, 0x8e, 0x89, 0x37, 0xf9, 0x8c, 0xd1, 0xb7,
	0xe4, 0xed, 0x36, 0x74, 0x88, 0x19, 0x5a, 0x9c, 0x29, 0x6f, 0x7e, 0xff, 0x82, 0x4b, 0xfb, 0xf7,
	0x3f, 0x4f, 0x42, 0x48, 0xee, 0x5c, 0xae, 0xd5, 0xb4, 0xf3, 0xf4, 0xaf, 0x15, 0xd2, 0xf3, 0x2f,
	0xf4, 0xb7, 0x80, 0xbc, 0xb9, 0xe4, 0x26, 0xd2, 0x1d, 0x76, 0xfd, 0x75, 0xee, 0x3f, 0x64, 0xff,
	0xe1, 0xaa, 0x46, 0x9f, 0x54, 0x75, 0xf8, 0x71, 0xec, 0x79, 0x37, 0x1e, 0xef, 0x71, 0xfc, 0xb2,
	0xe3, 0xa3, 0xd0, 0xaa, 0x40, 0x0d, 0xfe, 0xb3, 0x02, 0xcd, 0xa5, 0x75, 0xc7, 0x10, 0x38, 0x94,
	0x06, 0x35, 0xa3, 0x67, 0xe4, 0xf6, 0xe2, 0xf2, 0xd3, 0xfb, 0xec, 0xca, 0x13, 0xdc, 0xdf, 0x66,
	0x57, 0x1f, 0x8d, 0xe8, 0x61, 0x55, 0x87, 0xd0, 0xc1, 0xe9, 0x36, 0xcf, 0xe9, 0x67, 0xee, 0x3e,
	0x33, 0xfa, 0x47, 0x40, 0xe8, 0x3f, 0x05, 0x41, 0x23, 0x76, 0xed, 0x0d, 0xe9, 0xef, 0xb0, 0xeb,
	0x05, 0x15, 0x7d, 0x5d, 0xd5, 0xe1, 0x97, 0x33, 0x07, 0x33, 0x57, 0xdb, 0xcc, 0x0d, 0xc0, 0x8d,
	0x89, 0xc3, 0x54, 0xbc, 0x42, 0xd9, 0x6c, 0xf9, 0x90, 0xcd, 0xa0, 0x97, 0x85, 0x53, 0xf9, 0xc5,
	0x66, 0x71, 0x0b, 0x1c, 0xdc, 0x55, 0x66, 0x34, 0x21, 0xab, 0x0d, 0xe3, 0xf4, 0x0e, 0x5b, 0xaa,
	0xb6, 0xfe, 0x3b, 0x6c, 0xb9, 0x32, 0xa2, 0xc7, 0x55, 0x1d, 0x3e, 0x6a, 0x8c, 0xa6, 0xfb, 0xb2,
	0xa6, 0x5a, 0xe5, 0x9e, 0x2d, 0x07, 0x69, 0xd7, 0x34, 0x6c, 0xb1, 0xbe, 0x77, 0x7d, 0xee, 0x59,
	0x13, 0x13, 0xf4, 0xa8, 0x8f, 0xdc, 0x05, 0x6f, 0x7b, 0x71, 0xc4, 0xfa, 0xcd, 0x34, 0xec, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x33, 0x02, 0x51, 0x46, 0x08, 0x00, 0x00,
}
