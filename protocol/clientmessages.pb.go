// Code generated by protoc-gen-go. DO NOT EDIT.
// source: clientmessages.proto

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

type EBaseClientMessages int32

const (
	EBaseClientMessages_CM_CustomGameEvent                  EBaseClientMessages = 280
	EBaseClientMessages_CM_CustomGameEventBounce            EBaseClientMessages = 281
	EBaseClientMessages_CM_ClientUIEvent                    EBaseClientMessages = 282
	EBaseClientMessages_CM_DevPaletteVisibilityChanged      EBaseClientMessages = 283
	EBaseClientMessages_CM_WorldUIControllerHasPanelChanged EBaseClientMessages = 284
	EBaseClientMessages_CM_RotateAnchor                     EBaseClientMessages = 285
	EBaseClientMessages_CM_MAX_BASE                         EBaseClientMessages = 300
)

var EBaseClientMessages_name = map[int32]string{
	280: "CM_CustomGameEvent",
	281: "CM_CustomGameEventBounce",
	282: "CM_ClientUIEvent",
	283: "CM_DevPaletteVisibilityChanged",
	284: "CM_WorldUIControllerHasPanelChanged",
	285: "CM_RotateAnchor",
	300: "CM_MAX_BASE",
}
var EBaseClientMessages_value = map[string]int32{
	"CM_CustomGameEvent":                  280,
	"CM_CustomGameEventBounce":            281,
	"CM_ClientUIEvent":                    282,
	"CM_DevPaletteVisibilityChanged":      283,
	"CM_WorldUIControllerHasPanelChanged": 284,
	"CM_RotateAnchor":                     285,
	"CM_MAX_BASE":                         300,
}

func (x EBaseClientMessages) Enum() *EBaseClientMessages {
	p := new(EBaseClientMessages)
	*p = x
	return p
}
func (x EBaseClientMessages) String() string {
	return proto.EnumName(EBaseClientMessages_name, int32(x))
}
func (x *EBaseClientMessages) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EBaseClientMessages_value, data, "EBaseClientMessages")
	if err != nil {
		return err
	}
	*x = EBaseClientMessages(value)
	return nil
}
func (EBaseClientMessages) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_clientmessages_36878fc012e94506, []int{0}
}

type EClientUIEvent int32

const (
	EClientUIEvent_EClientUIEvent_Invalid        EClientUIEvent = 0
	EClientUIEvent_EClientUIEvent_DialogFinished EClientUIEvent = 1
	EClientUIEvent_EClientUIEvent_FireOutput     EClientUIEvent = 2
)

var EClientUIEvent_name = map[int32]string{
	0: "EClientUIEvent_Invalid",
	1: "EClientUIEvent_DialogFinished",
	2: "EClientUIEvent_FireOutput",
}
var EClientUIEvent_value = map[string]int32{
	"EClientUIEvent_Invalid":        0,
	"EClientUIEvent_DialogFinished": 1,
	"EClientUIEvent_FireOutput":     2,
}

func (x EClientUIEvent) Enum() *EClientUIEvent {
	p := new(EClientUIEvent)
	*p = x
	return p
}
func (x EClientUIEvent) String() string {
	return proto.EnumName(EClientUIEvent_name, int32(x))
}
func (x *EClientUIEvent) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EClientUIEvent_value, data, "EClientUIEvent")
	if err != nil {
		return err
	}
	*x = EClientUIEvent(value)
	return nil
}
func (EClientUIEvent) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_clientmessages_36878fc012e94506, []int{1}
}

type CClientMsg_CustomGameEvent struct {
	EventName            *string  `protobuf:"bytes,1,opt,name=event_name,json=eventName" json:"event_name,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CClientMsg_CustomGameEvent) Reset()         { *m = CClientMsg_CustomGameEvent{} }
func (m *CClientMsg_CustomGameEvent) String() string { return proto.CompactTextString(m) }
func (*CClientMsg_CustomGameEvent) ProtoMessage()    {}
func (*CClientMsg_CustomGameEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_clientmessages_36878fc012e94506, []int{0}
}
func (m *CClientMsg_CustomGameEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CClientMsg_CustomGameEvent.Unmarshal(m, b)
}
func (m *CClientMsg_CustomGameEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CClientMsg_CustomGameEvent.Marshal(b, m, deterministic)
}
func (dst *CClientMsg_CustomGameEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CClientMsg_CustomGameEvent.Merge(dst, src)
}
func (m *CClientMsg_CustomGameEvent) XXX_Size() int {
	return xxx_messageInfo_CClientMsg_CustomGameEvent.Size(m)
}
func (m *CClientMsg_CustomGameEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_CClientMsg_CustomGameEvent.DiscardUnknown(m)
}

var xxx_messageInfo_CClientMsg_CustomGameEvent proto.InternalMessageInfo

func (m *CClientMsg_CustomGameEvent) GetEventName() string {
	if m != nil && m.EventName != nil {
		return *m.EventName
	}
	return ""
}

func (m *CClientMsg_CustomGameEvent) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type CClientMsg_CustomGameEventBounce struct {
	EventName            *string  `protobuf:"bytes,1,opt,name=event_name,json=eventName" json:"event_name,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	PlayerIndex          *int32   `protobuf:"varint,3,opt,name=player_index,json=playerIndex" json:"player_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CClientMsg_CustomGameEventBounce) Reset()         { *m = CClientMsg_CustomGameEventBounce{} }
func (m *CClientMsg_CustomGameEventBounce) String() string { return proto.CompactTextString(m) }
func (*CClientMsg_CustomGameEventBounce) ProtoMessage()    {}
func (*CClientMsg_CustomGameEventBounce) Descriptor() ([]byte, []int) {
	return fileDescriptor_clientmessages_36878fc012e94506, []int{1}
}
func (m *CClientMsg_CustomGameEventBounce) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CClientMsg_CustomGameEventBounce.Unmarshal(m, b)
}
func (m *CClientMsg_CustomGameEventBounce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CClientMsg_CustomGameEventBounce.Marshal(b, m, deterministic)
}
func (dst *CClientMsg_CustomGameEventBounce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CClientMsg_CustomGameEventBounce.Merge(dst, src)
}
func (m *CClientMsg_CustomGameEventBounce) XXX_Size() int {
	return xxx_messageInfo_CClientMsg_CustomGameEventBounce.Size(m)
}
func (m *CClientMsg_CustomGameEventBounce) XXX_DiscardUnknown() {
	xxx_messageInfo_CClientMsg_CustomGameEventBounce.DiscardUnknown(m)
}

var xxx_messageInfo_CClientMsg_CustomGameEventBounce proto.InternalMessageInfo

func (m *CClientMsg_CustomGameEventBounce) GetEventName() string {
	if m != nil && m.EventName != nil {
		return *m.EventName
	}
	return ""
}

func (m *CClientMsg_CustomGameEventBounce) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *CClientMsg_CustomGameEventBounce) GetPlayerIndex() int32 {
	if m != nil && m.PlayerIndex != nil {
		return *m.PlayerIndex
	}
	return 0
}

type CClientMsg_ClientUIEvent struct {
	Event                *EClientUIEvent `protobuf:"varint,1,opt,name=event,enum=EClientUIEvent,def=0" json:"event,omitempty"`
	EntEhandle           *uint32         `protobuf:"varint,2,opt,name=ent_ehandle,json=entEhandle" json:"ent_ehandle,omitempty"`
	ClientEhandle        *uint32         `protobuf:"varint,3,opt,name=client_ehandle,json=clientEhandle" json:"client_ehandle,omitempty"`
	Data1                *string         `protobuf:"bytes,4,opt,name=data1" json:"data1,omitempty"`
	Data2                *string         `protobuf:"bytes,5,opt,name=data2" json:"data2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CClientMsg_ClientUIEvent) Reset()         { *m = CClientMsg_ClientUIEvent{} }
func (m *CClientMsg_ClientUIEvent) String() string { return proto.CompactTextString(m) }
func (*CClientMsg_ClientUIEvent) ProtoMessage()    {}
func (*CClientMsg_ClientUIEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_clientmessages_36878fc012e94506, []int{2}
}
func (m *CClientMsg_ClientUIEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CClientMsg_ClientUIEvent.Unmarshal(m, b)
}
func (m *CClientMsg_ClientUIEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CClientMsg_ClientUIEvent.Marshal(b, m, deterministic)
}
func (dst *CClientMsg_ClientUIEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CClientMsg_ClientUIEvent.Merge(dst, src)
}
func (m *CClientMsg_ClientUIEvent) XXX_Size() int {
	return xxx_messageInfo_CClientMsg_ClientUIEvent.Size(m)
}
func (m *CClientMsg_ClientUIEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_CClientMsg_ClientUIEvent.DiscardUnknown(m)
}

var xxx_messageInfo_CClientMsg_ClientUIEvent proto.InternalMessageInfo

const Default_CClientMsg_ClientUIEvent_Event EClientUIEvent = EClientUIEvent_EClientUIEvent_Invalid

func (m *CClientMsg_ClientUIEvent) GetEvent() EClientUIEvent {
	if m != nil && m.Event != nil {
		return *m.Event
	}
	return Default_CClientMsg_ClientUIEvent_Event
}

func (m *CClientMsg_ClientUIEvent) GetEntEhandle() uint32 {
	if m != nil && m.EntEhandle != nil {
		return *m.EntEhandle
	}
	return 0
}

func (m *CClientMsg_ClientUIEvent) GetClientEhandle() uint32 {
	if m != nil && m.ClientEhandle != nil {
		return *m.ClientEhandle
	}
	return 0
}

func (m *CClientMsg_ClientUIEvent) GetData1() string {
	if m != nil && m.Data1 != nil {
		return *m.Data1
	}
	return ""
}

func (m *CClientMsg_ClientUIEvent) GetData2() string {
	if m != nil && m.Data2 != nil {
		return *m.Data2
	}
	return ""
}

type CClientMsg_DevPaletteVisibilityChangedEvent struct {
	Visible              *bool    `protobuf:"varint,1,opt,name=visible" json:"visible,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CClientMsg_DevPaletteVisibilityChangedEvent) Reset() {
	*m = CClientMsg_DevPaletteVisibilityChangedEvent{}
}
func (m *CClientMsg_DevPaletteVisibilityChangedEvent) String() string {
	return proto.CompactTextString(m)
}
func (*CClientMsg_DevPaletteVisibilityChangedEvent) ProtoMessage() {}
func (*CClientMsg_DevPaletteVisibilityChangedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_clientmessages_36878fc012e94506, []int{3}
}
func (m *CClientMsg_DevPaletteVisibilityChangedEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CClientMsg_DevPaletteVisibilityChangedEvent.Unmarshal(m, b)
}
func (m *CClientMsg_DevPaletteVisibilityChangedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CClientMsg_DevPaletteVisibilityChangedEvent.Marshal(b, m, deterministic)
}
func (dst *CClientMsg_DevPaletteVisibilityChangedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CClientMsg_DevPaletteVisibilityChangedEvent.Merge(dst, src)
}
func (m *CClientMsg_DevPaletteVisibilityChangedEvent) XXX_Size() int {
	return xxx_messageInfo_CClientMsg_DevPaletteVisibilityChangedEvent.Size(m)
}
func (m *CClientMsg_DevPaletteVisibilityChangedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_CClientMsg_DevPaletteVisibilityChangedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_CClientMsg_DevPaletteVisibilityChangedEvent proto.InternalMessageInfo

func (m *CClientMsg_DevPaletteVisibilityChangedEvent) GetVisible() bool {
	if m != nil && m.Visible != nil {
		return *m.Visible
	}
	return false
}

type CClientMsg_WorldUIControllerHasPanelChangedEvent struct {
	HasPanel             *bool    `protobuf:"varint,1,opt,name=has_panel,json=hasPanel" json:"has_panel,omitempty"`
	ClientEhandle        *uint32  `protobuf:"varint,2,opt,name=client_ehandle,json=clientEhandle" json:"client_ehandle,omitempty"`
	LiteralHandType      *uint32  `protobuf:"varint,3,opt,name=literal_hand_type,json=literalHandType" json:"literal_hand_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CClientMsg_WorldUIControllerHasPanelChangedEvent) Reset() {
	*m = CClientMsg_WorldUIControllerHasPanelChangedEvent{}
}
func (m *CClientMsg_WorldUIControllerHasPanelChangedEvent) String() string {
	return proto.CompactTextString(m)
}
func (*CClientMsg_WorldUIControllerHasPanelChangedEvent) ProtoMessage() {}
func (*CClientMsg_WorldUIControllerHasPanelChangedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_clientmessages_36878fc012e94506, []int{4}
}
func (m *CClientMsg_WorldUIControllerHasPanelChangedEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CClientMsg_WorldUIControllerHasPanelChangedEvent.Unmarshal(m, b)
}
func (m *CClientMsg_WorldUIControllerHasPanelChangedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CClientMsg_WorldUIControllerHasPanelChangedEvent.Marshal(b, m, deterministic)
}
func (dst *CClientMsg_WorldUIControllerHasPanelChangedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CClientMsg_WorldUIControllerHasPanelChangedEvent.Merge(dst, src)
}
func (m *CClientMsg_WorldUIControllerHasPanelChangedEvent) XXX_Size() int {
	return xxx_messageInfo_CClientMsg_WorldUIControllerHasPanelChangedEvent.Size(m)
}
func (m *CClientMsg_WorldUIControllerHasPanelChangedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_CClientMsg_WorldUIControllerHasPanelChangedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_CClientMsg_WorldUIControllerHasPanelChangedEvent proto.InternalMessageInfo

func (m *CClientMsg_WorldUIControllerHasPanelChangedEvent) GetHasPanel() bool {
	if m != nil && m.HasPanel != nil {
		return *m.HasPanel
	}
	return false
}

func (m *CClientMsg_WorldUIControllerHasPanelChangedEvent) GetClientEhandle() uint32 {
	if m != nil && m.ClientEhandle != nil {
		return *m.ClientEhandle
	}
	return 0
}

func (m *CClientMsg_WorldUIControllerHasPanelChangedEvent) GetLiteralHandType() uint32 {
	if m != nil && m.LiteralHandType != nil {
		return *m.LiteralHandType
	}
	return 0
}

type CClientMsg_RotateAnchor struct {
	Angle                *float32 `protobuf:"fixed32,1,opt,name=angle" json:"angle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CClientMsg_RotateAnchor) Reset()         { *m = CClientMsg_RotateAnchor{} }
func (m *CClientMsg_RotateAnchor) String() string { return proto.CompactTextString(m) }
func (*CClientMsg_RotateAnchor) ProtoMessage()    {}
func (*CClientMsg_RotateAnchor) Descriptor() ([]byte, []int) {
	return fileDescriptor_clientmessages_36878fc012e94506, []int{5}
}
func (m *CClientMsg_RotateAnchor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CClientMsg_RotateAnchor.Unmarshal(m, b)
}
func (m *CClientMsg_RotateAnchor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CClientMsg_RotateAnchor.Marshal(b, m, deterministic)
}
func (dst *CClientMsg_RotateAnchor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CClientMsg_RotateAnchor.Merge(dst, src)
}
func (m *CClientMsg_RotateAnchor) XXX_Size() int {
	return xxx_messageInfo_CClientMsg_RotateAnchor.Size(m)
}
func (m *CClientMsg_RotateAnchor) XXX_DiscardUnknown() {
	xxx_messageInfo_CClientMsg_RotateAnchor.DiscardUnknown(m)
}

var xxx_messageInfo_CClientMsg_RotateAnchor proto.InternalMessageInfo

func (m *CClientMsg_RotateAnchor) GetAngle() float32 {
	if m != nil && m.Angle != nil {
		return *m.Angle
	}
	return 0
}

func init() {
	proto.RegisterType((*CClientMsg_CustomGameEvent)(nil), "CClientMsg_CustomGameEvent")
	proto.RegisterType((*CClientMsg_CustomGameEventBounce)(nil), "CClientMsg_CustomGameEventBounce")
	proto.RegisterType((*CClientMsg_ClientUIEvent)(nil), "CClientMsg_ClientUIEvent")
	proto.RegisterType((*CClientMsg_DevPaletteVisibilityChangedEvent)(nil), "CClientMsg_DevPaletteVisibilityChangedEvent")
	proto.RegisterType((*CClientMsg_WorldUIControllerHasPanelChangedEvent)(nil), "CClientMsg_WorldUIControllerHasPanelChangedEvent")
	proto.RegisterType((*CClientMsg_RotateAnchor)(nil), "CClientMsg_RotateAnchor")
	proto.RegisterEnum("EBaseClientMessages", EBaseClientMessages_name, EBaseClientMessages_value)
	proto.RegisterEnum("EClientUIEvent", EClientUIEvent_name, EClientUIEvent_value)
}

func init() {
	proto.RegisterFile("clientmessages.proto", fileDescriptor_clientmessages_36878fc012e94506)
}

var fileDescriptor_clientmessages_36878fc012e94506 = []byte{
	// 546 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0xad, 0xdd, 0x46, 0xbf, 0x66, 0xd2, 0x26, 0xfb, 0x5b, 0x02, 0x35, 0x45, 0x81, 0xd4, 0x15,
	0x52, 0x14, 0xa4, 0x02, 0x39, 0x22, 0x71, 0x48, 0xdc, 0xb4, 0xc9, 0xc1, 0xb4, 0x32, 0x14, 0xb8,
	0x59, 0x4b, 0x3c, 0x8a, 0x57, 0xda, 0xac, 0x2d, 0x7b, 0x13, 0x35, 0x37, 0x3e, 0x06, 0x7f, 0x2f,
	0x9c, 0xf9, 0x2a, 0x7c, 0x02, 0xbe, 0x0c, 0xb2, 0xd7, 0x81, 0x24, 0xb4, 0x20, 0x71, 0x9b, 0x79,
	0x6f, 0x76, 0x66, 0xdf, 0xd3, 0x83, 0xfa, 0x48, 0x70, 0x94, 0x6a, 0x82, 0x69, 0xca, 0xc6, 0x98,
	0x1e, 0xc5, 0x49, 0xa4, 0x22, 0xfb, 0x0c, 0xf6, 0x1d, 0x27, 0x27, 0xdc, 0x74, 0xec, 0x3b, 0xd3,
	0x54, 0x45, 0x93, 0x53, 0x36, 0xc1, 0xfe, 0x0c, 0xa5, 0xa2, 0x0d, 0x00, 0xcc, 0x0a, 0x5f, 0xb2,
	0x09, 0x5a, 0x46, 0xd3, 0x68, 0x95, 0xbd, 0x72, 0x8e, 0x3c, 0x63, 0x13, 0xa4, 0x14, 0xb6, 0x02,
	0xa6, 0x98, 0x65, 0x36, 0x8d, 0xd6, 0x8e, 0x97, 0xd7, 0xf6, 0x25, 0x34, 0xaf, 0x5f, 0xd8, 0x8b,
	0xa6, 0x72, 0x84, 0xff, 0xb0, 0x96, 0x1e, 0xc0, 0x4e, 0x2c, 0xd8, 0x1c, 0x13, 0x9f, 0xcb, 0x00,
	0x2f, 0xad, 0xcd, 0xa6, 0xd1, 0x2a, 0x79, 0x15, 0x8d, 0x0d, 0x33, 0xc8, 0xfe, 0x66, 0x80, 0xb5,
	0x7c, 0x3a, 0xaf, 0x2e, 0x86, 0x5a, 0xc9, 0x53, 0x28, 0xe5, 0x07, 0xf2, 0x6b, 0xd5, 0x4e, 0xed,
	0xa8, 0xbf, 0xc2, 0x3f, 0xb9, 0xb5, 0xda, 0xfb, 0x43, 0x39, 0x63, 0x82, 0x07, 0x9e, 0x7e, 0x45,
	0xef, 0x41, 0x25, 0x43, 0x31, 0x64, 0x32, 0x10, 0x98, 0xff, 0x6c, 0xd7, 0x03, 0x94, 0xaa, 0xaf,
	0x11, 0x7a, 0x1f, 0xaa, 0xda, 0xdf, 0x9f, 0x33, 0x9b, 0xf9, 0xcc, 0xae, 0x46, 0x17, 0x63, 0x75,
	0x28, 0x65, 0x72, 0x1e, 0x5b, 0x5b, 0xb9, 0x68, 0xdd, 0x2c, 0xd0, 0x8e, 0x55, 0xfa, 0x85, 0x76,
	0xec, 0x53, 0x78, 0xb0, 0x24, 0xe7, 0x18, 0x67, 0xe7, 0x4c, 0xa0, 0x52, 0xf8, 0x92, 0xa7, 0xfc,
	0x0d, 0x17, 0x5c, 0xcd, 0x9d, 0x90, 0xc9, 0x31, 0x06, 0x5a, 0xa1, 0x05, 0xff, 0xcd, 0x32, 0x46,
	0x68, 0x47, 0xb7, 0xbd, 0x45, 0x6b, 0x7f, 0x31, 0xe0, 0xd1, 0xd2, 0xa6, 0x57, 0x51, 0x22, 0x82,
	0x8b, 0xa1, 0x13, 0x49, 0x95, 0x44, 0x42, 0x60, 0x32, 0x60, 0xe9, 0x39, 0x93, 0x28, 0x56, 0xd6,
	0xdd, 0x81, 0x72, 0xc8, 0x52, 0x3f, 0xce, 0x88, 0x62, 0xe1, 0x76, 0x58, 0x0c, 0x5e, 0xa1, 0xd6,
	0xbc, 0x4a, 0x6d, 0x1b, 0xfe, 0x17, 0x5c, 0x61, 0xc2, 0x84, 0x9f, 0x21, 0xbe, 0x9a, 0xc7, 0x0b,
	0x5f, 0x6a, 0x05, 0x31, 0x60, 0x32, 0x78, 0x31, 0x8f, 0xd1, 0x7e, 0x08, 0x7b, 0x4b, 0x7f, 0xf4,
	0x22, 0xc5, 0x14, 0x76, 0xe5, 0x28, 0x8c, 0x92, 0xcc, 0x1e, 0x26, 0xc7, 0x85, 0x2e, 0xd3, 0xd3,
	0x4d, 0xfb, 0xbb, 0x01, 0x37, 0xfa, 0x3d, 0x96, 0x62, 0xf1, 0xaa, 0xc8, 0x35, 0xdd, 0x03, 0xea,
	0xb8, 0xeb, 0xc1, 0x23, 0xef, 0x4c, 0xda, 0x00, 0xeb, 0x77, 0x42, 0x27, 0x92, 0xbc, 0x37, 0xe9,
	0x4d, 0x20, 0x19, 0xbd, 0x9c, 0x02, 0xf2, 0xc1, 0xa4, 0x87, 0x70, 0xd7, 0x71, 0xff, 0xe4, 0x3e,
	0xf9, 0x68, 0xd2, 0x16, 0x1c, 0x3a, 0xee, 0x5f, 0x8d, 0x25, 0x9f, 0x4c, 0x5a, 0x87, 0x9a, 0xe3,
	0xae, 0xc8, 0x23, 0x9f, 0x4d, 0x4a, 0xa0, 0xe2, 0xb8, 0xbe, 0xdb, 0x7d, 0xed, 0xf7, 0xba, 0xcf,
	0xfb, 0xe4, 0xab, 0xd9, 0x96, 0x50, 0x5d, 0x4d, 0x24, 0xdd, 0x87, 0x6b, 0x32, 0x4a, 0x36, 0xe8,
	0x01, 0x34, 0xd6, 0xb8, 0x63, 0xce, 0x44, 0x34, 0x3e, 0xe1, 0x92, 0xa7, 0x21, 0x06, 0xc4, 0xa0,
	0x0d, 0xb8, 0xbd, 0x36, 0x72, 0xc2, 0x13, 0x3c, 0x9b, 0xaa, 0x78, 0xaa, 0x88, 0xd9, 0x2b, 0x0d,
	0x8c, 0xb7, 0xc6, 0xc6, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x62, 0x5e, 0xb3, 0x25, 0x04,
	0x00, 0x00,
}
