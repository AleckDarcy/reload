// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/tracer/message.proto

package tracer

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

type RecordType int32

const (
	RecordType_Record_       RecordType = 0
	RecordType_RecordSend    RecordType = 1
	RecordType_RecordReceive RecordType = 2
)

var RecordType_name = map[int32]string{
	0: "Record_",
	1: "RecordSend",
	2: "RecordReceive",
}

var RecordType_value = map[string]int32{
	"Record_":       0,
	"RecordSend":    1,
	"RecordReceive": 2,
}

func (x RecordType) String() string {
	return proto.EnumName(RecordType_name, int32(x))
}

func (RecordType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_30bd59f929831d10, []int{0}
}

type FaultType int32

const (
	FaultType_Fault_     FaultType = 0
	FaultType_FaultCrash FaultType = 1
	FaultType_FaultDelay FaultType = 2
)

var FaultType_name = map[int32]string{
	0: "Fault_",
	1: "FaultCrash",
	2: "FaultDelay",
}

var FaultType_value = map[string]int32{
	"Fault_":     0,
	"FaultCrash": 1,
	"FaultDelay": 2,
}

func (x FaultType) String() string {
	return proto.EnumName(FaultType_name, int32(x))
}

func (FaultType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_30bd59f929831d10, []int{1}
}

type Record struct {
	Type                 RecordType `protobuf:"varint,1,opt,name=type,proto3,enum=tracer.RecordType" json:"type,omitempty"`
	Timestamp            int64      `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	MessageName          string     `protobuf:"bytes,3,opt,name=message_name,json=messageName,proto3" json:"message_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_30bd59f929831d10, []int{0}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetType() RecordType {
	if m != nil {
		return m.Type
	}
	return RecordType_Record_
}

func (m *Record) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Record) GetMessageName() string {
	if m != nil {
		return m.MessageName
	}
	return ""
}

type Trace struct {
	Id                   int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Depth                int64     `protobuf:"varint,2,opt,name=depth,proto3" json:"depth,omitempty"`
	Records              []*Record `protobuf:"bytes,3,rep,name=records,proto3" json:"records,omitempty"`
	Rlfi                 *RLFI     `protobuf:"bytes,20,opt,name=rlfi,proto3" json:"rlfi,omitempty"`
	Tfi                  *TFI      `protobuf:"bytes,21,opt,name=tfi,proto3" json:"tfi,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Trace) Reset()         { *m = Trace{} }
func (m *Trace) String() string { return proto.CompactTextString(m) }
func (*Trace) ProtoMessage()    {}
func (*Trace) Descriptor() ([]byte, []int) {
	return fileDescriptor_30bd59f929831d10, []int{1}
}

func (m *Trace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trace.Unmarshal(m, b)
}
func (m *Trace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trace.Marshal(b, m, deterministic)
}
func (m *Trace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trace.Merge(m, src)
}
func (m *Trace) XXX_Size() int {
	return xxx_messageInfo_Trace.Size(m)
}
func (m *Trace) XXX_DiscardUnknown() {
	xxx_messageInfo_Trace.DiscardUnknown(m)
}

var xxx_messageInfo_Trace proto.InternalMessageInfo

func (m *Trace) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Trace) GetDepth() int64 {
	if m != nil {
		return m.Depth
	}
	return 0
}

func (m *Trace) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

func (m *Trace) GetRlfi() *RLFI {
	if m != nil {
		return m.Rlfi
	}
	return nil
}

func (m *Trace) GetTfi() *TFI {
	if m != nil {
		return m.Tfi
	}
	return nil
}

type RLFI struct {
	Type                 FaultType `protobuf:"varint,1,opt,name=type,proto3,enum=tracer.FaultType" json:"type,omitempty"`
	Name                 string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Delay                int64     `protobuf:"varint,3,opt,name=delay,proto3" json:"delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RLFI) Reset()         { *m = RLFI{} }
func (m *RLFI) String() string { return proto.CompactTextString(m) }
func (*RLFI) ProtoMessage()    {}
func (*RLFI) Descriptor() ([]byte, []int) {
	return fileDescriptor_30bd59f929831d10, []int{2}
}

func (m *RLFI) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RLFI.Unmarshal(m, b)
}
func (m *RLFI) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RLFI.Marshal(b, m, deterministic)
}
func (m *RLFI) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RLFI.Merge(m, src)
}
func (m *RLFI) XXX_Size() int {
	return xxx_messageInfo_RLFI.Size(m)
}
func (m *RLFI) XXX_DiscardUnknown() {
	xxx_messageInfo_RLFI.DiscardUnknown(m)
}

var xxx_messageInfo_RLFI proto.InternalMessageInfo

func (m *RLFI) GetType() FaultType {
	if m != nil {
		return m.Type
	}
	return FaultType_Fault_
}

func (m *RLFI) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RLFI) GetDelay() int64 {
	if m != nil {
		return m.Delay
	}
	return 0
}

type TFI struct {
	Type                 FaultType `protobuf:"varint,1,opt,name=type,proto3,enum=tracer.FaultType" json:"type,omitempty"`
	Name                 string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Delay                int64     `protobuf:"varint,3,opt,name=delay,proto3" json:"delay,omitempty"`
	After                []string  `protobuf:"bytes,4,rep,name=after,proto3" json:"after,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *TFI) Reset()         { *m = TFI{} }
func (m *TFI) String() string { return proto.CompactTextString(m) }
func (*TFI) ProtoMessage()    {}
func (*TFI) Descriptor() ([]byte, []int) {
	return fileDescriptor_30bd59f929831d10, []int{3}
}

func (m *TFI) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TFI.Unmarshal(m, b)
}
func (m *TFI) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TFI.Marshal(b, m, deterministic)
}
func (m *TFI) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TFI.Merge(m, src)
}
func (m *TFI) XXX_Size() int {
	return xxx_messageInfo_TFI.Size(m)
}
func (m *TFI) XXX_DiscardUnknown() {
	xxx_messageInfo_TFI.DiscardUnknown(m)
}

var xxx_messageInfo_TFI proto.InternalMessageInfo

func (m *TFI) GetType() FaultType {
	if m != nil {
		return m.Type
	}
	return FaultType_Fault_
}

func (m *TFI) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TFI) GetDelay() int64 {
	if m != nil {
		return m.Delay
	}
	return 0
}

func (m *TFI) GetAfter() []string {
	if m != nil {
		return m.After
	}
	return nil
}

func init() {
	proto.RegisterEnum("tracer.RecordType", RecordType_name, RecordType_value)
	proto.RegisterEnum("tracer.FaultType", FaultType_name, FaultType_value)
	proto.RegisterType((*Record)(nil), "tracer.Record")
	proto.RegisterType((*Trace)(nil), "tracer.Trace")
	proto.RegisterType((*RLFI)(nil), "tracer.RLFI")
	proto.RegisterType((*TFI)(nil), "tracer.TFI")
}

func init() { proto.RegisterFile("core/tracer/message.proto", fileDescriptor_30bd59f929831d10) }

var fileDescriptor_30bd59f929831d10 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xcd, 0xaa, 0x9b, 0x40,
	0x14, 0xee, 0x38, 0xc6, 0xe0, 0x31, 0x15, 0x73, 0x48, 0xc1, 0x42, 0x0b, 0x56, 0x68, 0x91, 0x2c,
	0x12, 0x48, 0x17, 0xdd, 0x74, 0xd7, 0x12, 0x08, 0x94, 0x2e, 0xa6, 0x42, 0x97, 0x61, 0xaa, 0x27,
	0x8d, 0xa0, 0xd1, 0x8e, 0xd3, 0x0b, 0xbe, 0xc9, 0x7d, 0xdc, 0x8b, 0xa3, 0x26, 0x37, 0x0f, 0x70,
	0x77, 0xf3, 0xfd, 0x78, 0xe6, 0xfb, 0x8e, 0x03, 0x6f, 0xb3, 0x5a, 0xd1, 0x56, 0x2b, 0x99, 0x91,
	0xda, 0x56, 0xd4, 0xb6, 0xf2, 0x2f, 0x6d, 0x1a, 0x55, 0xeb, 0x1a, 0x9d, 0x81, 0x8d, 0xff, 0x81,
	0x23, 0x28, 0xab, 0x55, 0x8e, 0x9f, 0xc0, 0xd6, 0x5d, 0x43, 0x21, 0x8b, 0x58, 0xe2, 0xef, 0x70,
	0x33, 0x18, 0x36, 0x83, 0x9a, 0x76, 0x0d, 0x09, 0xa3, 0xe3, 0x3b, 0x70, 0x75, 0x51, 0x51, 0xab,
	0x65, 0xd5, 0x84, 0x56, 0xc4, 0x12, 0x2e, 0x6e, 0x04, 0x7e, 0x80, 0xc5, 0x78, 0xd1, 0xf1, 0x22,
	0x2b, 0x0a, 0x79, 0xc4, 0x12, 0x57, 0x78, 0x23, 0xf7, 0x53, 0x56, 0x14, 0x3f, 0x32, 0x98, 0xa5,
	0xfd, 0x70, 0xf4, 0xc1, 0x2a, 0x72, 0x73, 0x21, 0x17, 0x56, 0x91, 0xe3, 0x0a, 0x66, 0x39, 0x35,
	0xfa, 0x3c, 0x8e, 0x1d, 0x00, 0x26, 0x30, 0x57, 0x26, 0x44, 0x1b, 0xf2, 0x88, 0x27, 0xde, 0xce,
	0xbf, 0xcf, 0x26, 0x26, 0x19, 0x23, 0xb0, 0x55, 0x79, 0x2a, 0xc2, 0x55, 0xc4, 0x12, 0x6f, 0xb7,
	0xb8, 0xda, 0x7e, 0xec, 0x0f, 0xc2, 0x28, 0xf8, 0x1e, 0xb8, 0x3e, 0x15, 0xe1, 0x1b, 0x63, 0xf0,
	0x26, 0x43, 0xba, 0x3f, 0x88, 0x9e, 0x8f, 0x7f, 0x83, 0xdd, 0x9b, 0xf1, 0xe3, 0xdd, 0x2e, 0x96,
	0x93, 0x6f, 0x2f, 0xff, 0x97, 0xfa, 0xd9, 0x2a, 0x10, 0x6c, 0x53, 0xd2, 0x32, 0x25, 0xcd, 0x79,
	0xe8, 0x50, 0xca, 0xce, 0x34, 0x37, 0x1d, 0x4a, 0xd9, 0xc5, 0x25, 0xf0, 0xf4, 0x05, 0xe6, 0xf6,
	0xac, 0x3c, 0x69, 0x52, 0xa1, 0x1d, 0xf1, 0xc4, 0x15, 0x03, 0x58, 0x7f, 0x05, 0xb8, 0xfd, 0x36,
	0xf4, 0x60, 0x3e, 0xa0, 0x63, 0xf0, 0x0a, 0xfd, 0x49, 0xfa, 0x45, 0x97, 0x3c, 0x60, 0xb8, 0x84,
	0xd7, 0xe3, 0x16, 0x29, 0xa3, 0xe2, 0x81, 0x02, 0x6b, 0xfd, 0x05, 0xdc, 0x6b, 0x20, 0x04, 0x70,
	0x0c, 0x18, 0xbf, 0x35, 0xe7, 0x6f, 0x4a, 0xb6, 0xe7, 0x80, 0x5d, 0xf1, 0xf7, 0x3e, 0x4a, 0x60,
	0xfd, 0x71, 0xcc, 0xd3, 0xfa, 0xfc, 0x14, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xe0, 0x88, 0x63, 0x77,
	0x02, 0x00, 0x00,
}