// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

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

type MessageType int32

const (
	MessageType_Message_         MessageType = 0
	MessageType_Message_Request  MessageType = 1
	MessageType_Message_Response MessageType = 2
)

var MessageType_name = map[int32]string{
	0: "Message_",
	1: "Message_Request",
	2: "Message_Response",
}

var MessageType_value = map[string]int32{
	"Message_":         0,
	"Message_Request":  1,
	"Message_Response": 2,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}

func (MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

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
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
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
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

type Record struct {
	Type                 RecordType `protobuf:"varint,1,opt,name=type,proto3,enum=tracer.RecordType" json:"type,omitempty"`
	Timestamp            int64      `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	MessageName          string     `protobuf:"bytes,3,opt,name=message_name,json=messageName,proto3" json:"message_name,omitempty"`
	Uuid                 string     `protobuf:"bytes,4,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Service              string     `protobuf:"bytes,5,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
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

func (m *Record) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Record) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type Trace struct {
	Id                   int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Records              []*Record `protobuf:"bytes,2,rep,name=records,proto3" json:"records,omitempty"`
	Rlfis                []*RLFI   `protobuf:"bytes,20,rep,name=rlfis,proto3" json:"rlfis,omitempty"`
	Tfis                 []*TFI    `protobuf:"bytes,21,rep,name=tfis,proto3" json:"tfis,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Trace) Reset()         { *m = Trace{} }
func (m *Trace) String() string { return proto.CompactTextString(m) }
func (*Trace) ProtoMessage()    {}
func (*Trace) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
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

func (m *Trace) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

func (m *Trace) GetRlfis() []*RLFI {
	if m != nil {
		return m.Rlfis
	}
	return nil
}

func (m *Trace) GetTfis() []*TFI {
	if m != nil {
		return m.Tfis
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
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
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

type TFIMeta struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Times                int64    `protobuf:"varint,2,opt,name=times,proto3" json:"times,omitempty"`
	Already              int64    `protobuf:"varint,3,opt,name=already,proto3" json:"already,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TFIMeta) Reset()         { *m = TFIMeta{} }
func (m *TFIMeta) String() string { return proto.CompactTextString(m) }
func (*TFIMeta) ProtoMessage()    {}
func (*TFIMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{3}
}

func (m *TFIMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TFIMeta.Unmarshal(m, b)
}
func (m *TFIMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TFIMeta.Marshal(b, m, deterministic)
}
func (m *TFIMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TFIMeta.Merge(m, src)
}
func (m *TFIMeta) XXX_Size() int {
	return xxx_messageInfo_TFIMeta.Size(m)
}
func (m *TFIMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_TFIMeta.DiscardUnknown(m)
}

var xxx_messageInfo_TFIMeta proto.InternalMessageInfo

func (m *TFIMeta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TFIMeta) GetTimes() int64 {
	if m != nil {
		return m.Times
	}
	return 0
}

func (m *TFIMeta) GetAlready() int64 {
	if m != nil {
		return m.Already
	}
	return 0
}

type TFI struct {
	Type                 FaultType  `protobuf:"varint,1,opt,name=type,proto3,enum=tracer.FaultType" json:"type,omitempty"`
	Name                 []string   `protobuf:"bytes,2,rep,name=name,proto3" json:"name,omitempty"`
	Delay                int64      `protobuf:"varint,3,opt,name=delay,proto3" json:"delay,omitempty"`
	After                []*TFIMeta `protobuf:"bytes,4,rep,name=after,proto3" json:"after,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TFI) Reset()         { *m = TFI{} }
func (m *TFI) String() string { return proto.CompactTextString(m) }
func (*TFI) ProtoMessage()    {}
func (*TFI) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{4}
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

func (m *TFI) GetName() []string {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *TFI) GetDelay() int64 {
	if m != nil {
		return m.Delay
	}
	return 0
}

func (m *TFI) GetAfter() []*TFIMeta {
	if m != nil {
		return m.After
	}
	return nil
}

type Request struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{5}
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

func (m *Request) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterEnum("tracer.MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("tracer.RecordType", RecordType_name, RecordType_value)
	proto.RegisterEnum("tracer.FaultType", FaultType_name, FaultType_value)
	proto.RegisterType((*Record)(nil), "tracer.Record")
	proto.RegisterType((*Trace)(nil), "tracer.Trace")
	proto.RegisterType((*RLFI)(nil), "tracer.RLFI")
	proto.RegisterType((*TFIMeta)(nil), "tracer.TFIMeta")
	proto.RegisterType((*TFI)(nil), "tracer.TFI")
	proto.RegisterType((*Request)(nil), "tracer.Request")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0xc5, 0xf9, 0x68, 0xe9, 0xa4, 0xdb, 0xcd, 0x0e, 0x45, 0x32, 0x12, 0x12, 0x25, 0xd2, 0xa2,
	0x68, 0x0f, 0x3d, 0x2c, 0x07, 0x2e, 0xdc, 0x40, 0x91, 0x2a, 0x51, 0x0e, 0x26, 0x12, 0xc7, 0x95,
	0x69, 0x66, 0x21, 0x52, 0xd3, 0x06, 0xdb, 0x5d, 0xa9, 0x07, 0xae, 0xfc, 0x0d, 0xfe, 0x2a, 0xb2,
	0x13, 0xf7, 0x83, 0x03, 0x87, 0xbd, 0xcd, 0x9b, 0xf7, 0xc6, 0x7e, 0x7e, 0x99, 0xc0, 0x45, 0x43,
	0x5a, 0xcb, 0xef, 0x34, 0x6f, 0xd5, 0xd6, 0x6c, 0x71, 0x60, 0x94, 0x5c, 0x91, 0xca, 0xfe, 0x30,
	0x18, 0x08, 0x5a, 0x6d, 0x55, 0x85, 0x6f, 0x20, 0x32, 0xfb, 0x96, 0x38, 0x9b, 0xb1, 0x7c, 0x72,
	0x8b, 0xf3, 0x4e, 0x31, 0xef, 0xd8, 0x72, 0xdf, 0x92, 0x70, 0x3c, 0xbe, 0x84, 0x91, 0xa9, 0x1b,
	0xd2, 0x46, 0x36, 0x2d, 0x0f, 0x66, 0x2c, 0x0f, 0xc5, 0xb1, 0x81, 0xaf, 0x61, 0xdc, 0xdf, 0x74,
	0xb7, 0x91, 0x0d, 0xf1, 0x70, 0xc6, 0xf2, 0x91, 0x48, 0xfa, 0xde, 0x67, 0xd9, 0x10, 0x22, 0x44,
	0xbb, 0x5d, 0x5d, 0xf1, 0xc8, 0x51, 0xae, 0x46, 0x0e, 0x43, 0x4d, 0xea, 0xa1, 0x5e, 0x11, 0x8f,
	0x5d, 0xdb, 0xc3, 0xec, 0x37, 0x83, 0xb8, 0xb4, 0x56, 0x70, 0x02, 0x41, 0x5d, 0x39, 0x7b, 0xa1,
	0x08, 0xea, 0x0a, 0x73, 0x18, 0x2a, 0x67, 0x4e, 0xf3, 0x60, 0x16, 0xe6, 0xc9, 0xed, 0xe4, 0xdc,
	0xb3, 0xf0, 0x34, 0x66, 0x10, 0xab, 0xf5, 0x7d, 0xad, 0xf9, 0xd4, 0xe9, 0xc6, 0x07, 0xdd, 0xa7,
	0x62, 0x21, 0x3a, 0x0a, 0x5f, 0x41, 0x64, 0xac, 0xe4, 0xb9, 0x93, 0x24, 0x5e, 0x52, 0x16, 0x0b,
	0xe1, 0x88, 0xec, 0x2b, 0x44, 0x56, 0x8f, 0xd7, 0x67, 0x39, 0x5d, 0x79, 0x61, 0x21, 0x77, 0x6b,
	0x73, 0x12, 0x13, 0x42, 0xe4, 0x02, 0x08, 0xba, 0x57, 0xda, 0x1a, 0xa7, 0x10, 0x57, 0xb4, 0x96,
	0x7b, 0x97, 0x4a, 0x28, 0x3a, 0x90, 0x2d, 0x61, 0x58, 0x16, 0x8b, 0x25, 0x19, 0x79, 0x18, 0x62,
	0xe7, 0x43, 0x2e, 0xde, 0x3e, 0xeb, 0x0e, 0xd8, 0xc0, 0xe4, 0x5a, 0x91, 0xac, 0xfc, 0x61, 0x1e,
	0x66, 0xbf, 0x20, 0x2c, 0x1f, 0x63, 0x33, 0xfc, 0xbf, 0x4d, 0xbc, 0x86, 0x58, 0xde, 0x1b, 0x52,
	0x3c, 0x72, 0x09, 0x5d, 0x9e, 0x24, 0x64, 0xbd, 0x8b, 0x8e, 0xcd, 0x5e, 0xc0, 0x50, 0xd0, 0xcf,
	0x1d, 0x69, 0xf3, 0xef, 0x07, 0xbb, 0x29, 0x20, 0x59, 0x76, 0x7b, 0x60, 0x0d, 0xe0, 0x18, 0x9e,
	0xf6, 0xf0, 0x2e, 0x7d, 0x82, 0xcf, 0xe0, 0xd2, 0xa3, 0x7e, 0x3e, 0x65, 0x38, 0x85, 0xf4, 0xd8,
	0xd4, 0xed, 0x76, 0xa3, 0x29, 0x0d, 0x6e, 0xde, 0x03, 0x1c, 0xb7, 0x12, 0x13, 0x7b, 0xa1, 0x45,
	0xf6, 0x94, 0x89, 0xa7, 0xbe, 0xd0, 0xa6, 0x4a, 0x19, 0x5e, 0xc1, 0x45, 0xbf, 0x0c, 0xb4, 0xa2,
	0xfa, 0xc1, 0x4e, 0xbf, 0x83, 0xd1, 0x21, 0x04, 0x04, 0x18, 0x38, 0xd0, 0xcf, 0xba, 0xfa, 0x83,
	0x92, 0xfa, 0x47, 0xca, 0x0e, 0xf8, 0xa3, 0x7d, 0x7e, 0x1a, 0x7c, 0x1b, 0xb8, 0x5f, 0xe7, 0xed,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x38, 0x90, 0x7c, 0x4b, 0x03, 0x00, 0x00,
}
