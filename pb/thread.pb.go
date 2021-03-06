// Code generated by protoc-gen-go. DO NOT EDIT.
// source: thread.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ThreadBlock_Type int32

const (
	ThreadBlock_MERGE      ThreadBlock_Type = 0
	ThreadBlock_IGNORE     ThreadBlock_Type = 1
	ThreadBlock_FLAG       ThreadBlock_Type = 2
	ThreadBlock_JOIN       ThreadBlock_Type = 3
	ThreadBlock_ANNOUNCE   ThreadBlock_Type = 4
	ThreadBlock_LEAVE      ThreadBlock_Type = 5
	ThreadBlock_DATA       ThreadBlock_Type = 6
	ThreadBlock_ANNOTATION ThreadBlock_Type = 7
	ThreadBlock_INVITE     ThreadBlock_Type = 50
)

var ThreadBlock_Type_name = map[int32]string{
	0:  "MERGE",
	1:  "IGNORE",
	2:  "FLAG",
	3:  "JOIN",
	4:  "ANNOUNCE",
	5:  "LEAVE",
	6:  "DATA",
	7:  "ANNOTATION",
	50: "INVITE",
}
var ThreadBlock_Type_value = map[string]int32{
	"MERGE":      0,
	"IGNORE":     1,
	"FLAG":       2,
	"JOIN":       3,
	"ANNOUNCE":   4,
	"LEAVE":      5,
	"DATA":       6,
	"ANNOTATION": 7,
	"INVITE":     50,
}

func (x ThreadBlock_Type) String() string {
	return proto.EnumName(ThreadBlock_Type_name, int32(x))
}
func (ThreadBlock_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_thread_890bcf6c9a2f6ba0, []int{1, 0}
}

type ThreadData_Type int32

const (
	ThreadData_FILE ThreadData_Type = 0
	ThreadData_TEXT ThreadData_Type = 1
)

var ThreadData_Type_name = map[int32]string{
	0: "FILE",
	1: "TEXT",
}
var ThreadData_Type_value = map[string]int32{
	"FILE": 0,
	"TEXT": 1,
}

func (x ThreadData_Type) String() string {
	return proto.EnumName(ThreadData_Type_name, int32(x))
}
func (ThreadData_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_thread_890bcf6c9a2f6ba0, []int{8, 0}
}

type ThreadAnnotation_Type int32

const (
	ThreadAnnotation_COMMENT ThreadAnnotation_Type = 0
	ThreadAnnotation_LIKE    ThreadAnnotation_Type = 1
)

var ThreadAnnotation_Type_name = map[int32]string{
	0: "COMMENT",
	1: "LIKE",
}
var ThreadAnnotation_Type_value = map[string]int32{
	"COMMENT": 0,
	"LIKE":    1,
}

func (x ThreadAnnotation_Type) String() string {
	return proto.EnumName(ThreadAnnotation_Type_name, int32(x))
}
func (ThreadAnnotation_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_thread_890bcf6c9a2f6ba0, []int{9, 0}
}

// for wire transport
type ThreadEnvelope struct {
	Thread               string   `protobuf:"bytes,1,opt,name=thread,proto3" json:"thread,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	CipherBlock          []byte   `protobuf:"bytes,3,opt,name=cipherBlock,proto3" json:"cipherBlock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadEnvelope) Reset()         { *m = ThreadEnvelope{} }
func (m *ThreadEnvelope) String() string { return proto.CompactTextString(m) }
func (*ThreadEnvelope) ProtoMessage()    {}
func (*ThreadEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_890bcf6c9a2f6ba0, []int{0}
}
func (m *ThreadEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadEnvelope.Unmarshal(m, b)
}
func (m *ThreadEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadEnvelope.Marshal(b, m, deterministic)
}
func (dst *ThreadEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadEnvelope.Merge(dst, src)
}
func (m *ThreadEnvelope) XXX_Size() int {
	return xxx_messageInfo_ThreadEnvelope.Size(m)
}
func (m *ThreadEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadEnvelope proto.InternalMessageInfo

func (m *ThreadEnvelope) GetThread() string {
	if m != nil {
		return m.Thread
	}
	return ""
}

func (m *ThreadEnvelope) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *ThreadEnvelope) GetCipherBlock() []byte {
	if m != nil {
		return m.CipherBlock
	}
	return nil
}

type ThreadBlock struct {
	Header               *ThreadBlockHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Type                 ThreadBlock_Type   `protobuf:"varint,2,opt,name=type,proto3,enum=ThreadBlock_Type" json:"type,omitempty"`
	Payload              *any.Any           `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ThreadBlock) Reset()         { *m = ThreadBlock{} }
func (m *ThreadBlock) String() string { return proto.CompactTextString(m) }
func (*ThreadBlock) ProtoMessage()    {}
func (*ThreadBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_890bcf6c9a2f6ba0, []int{1}
}
func (m *ThreadBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadBlock.Unmarshal(m, b)
}
func (m *ThreadBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadBlock.Marshal(b, m, deterministic)
}
func (dst *ThreadBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadBlock.Merge(dst, src)
}
func (m *ThreadBlock) XXX_Size() int {
	return xxx_messageInfo_ThreadBlock.Size(m)
}
func (m *ThreadBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadBlock.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadBlock proto.InternalMessageInfo

func (m *ThreadBlock) GetHeader() *ThreadBlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ThreadBlock) GetType() ThreadBlock_Type {
	if m != nil {
		return m.Type
	}
	return ThreadBlock_MERGE
}

func (m *ThreadBlock) GetPayload() *any.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

type ThreadBlockHeader struct {
	Date                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Parents              []string             `protobuf:"bytes,2,rep,name=parents,proto3" json:"parents,omitempty"`
	Author               string               `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ThreadBlockHeader) Reset()         { *m = ThreadBlockHeader{} }
func (m *ThreadBlockHeader) String() string { return proto.CompactTextString(m) }
func (*ThreadBlockHeader) ProtoMessage()    {}
func (*ThreadBlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_890bcf6c9a2f6ba0, []int{2}
}
func (m *ThreadBlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadBlockHeader.Unmarshal(m, b)
}
func (m *ThreadBlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadBlockHeader.Marshal(b, m, deterministic)
}
func (dst *ThreadBlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadBlockHeader.Merge(dst, src)
}
func (m *ThreadBlockHeader) XXX_Size() int {
	return xxx_messageInfo_ThreadBlockHeader.Size(m)
}
func (m *ThreadBlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadBlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadBlockHeader proto.InternalMessageInfo

func (m *ThreadBlockHeader) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *ThreadBlockHeader) GetParents() []string {
	if m != nil {
		return m.Parents
	}
	return nil
}

func (m *ThreadBlockHeader) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

type ThreadInvite struct {
	Sk                   []byte   `protobuf:"bytes,1,opt,name=sk,proto3" json:"sk,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadInvite) Reset()         { *m = ThreadInvite{} }
func (m *ThreadInvite) String() string { return proto.CompactTextString(m) }
func (*ThreadInvite) ProtoMessage()    {}
func (*ThreadInvite) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_890bcf6c9a2f6ba0, []int{3}
}
func (m *ThreadInvite) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadInvite.Unmarshal(m, b)
}
func (m *ThreadInvite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadInvite.Marshal(b, m, deterministic)
}
func (dst *ThreadInvite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadInvite.Merge(dst, src)
}
func (m *ThreadInvite) XXX_Size() int {
	return xxx_messageInfo_ThreadInvite.Size(m)
}
func (m *ThreadInvite) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadInvite.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadInvite proto.InternalMessageInfo

func (m *ThreadInvite) GetSk() []byte {
	if m != nil {
		return m.Sk
	}
	return nil
}

func (m *ThreadInvite) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ThreadIgnore struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadIgnore) Reset()         { *m = ThreadIgnore{} }
func (m *ThreadIgnore) String() string { return proto.CompactTextString(m) }
func (*ThreadIgnore) ProtoMessage()    {}
func (*ThreadIgnore) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_890bcf6c9a2f6ba0, []int{4}
}
func (m *ThreadIgnore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadIgnore.Unmarshal(m, b)
}
func (m *ThreadIgnore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadIgnore.Marshal(b, m, deterministic)
}
func (dst *ThreadIgnore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadIgnore.Merge(dst, src)
}
func (m *ThreadIgnore) XXX_Size() int {
	return xxx_messageInfo_ThreadIgnore.Size(m)
}
func (m *ThreadIgnore) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadIgnore.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadIgnore proto.InternalMessageInfo

func (m *ThreadIgnore) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type ThreadFlag struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadFlag) Reset()         { *m = ThreadFlag{} }
func (m *ThreadFlag) String() string { return proto.CompactTextString(m) }
func (*ThreadFlag) ProtoMessage()    {}
func (*ThreadFlag) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_890bcf6c9a2f6ba0, []int{5}
}
func (m *ThreadFlag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadFlag.Unmarshal(m, b)
}
func (m *ThreadFlag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadFlag.Marshal(b, m, deterministic)
}
func (dst *ThreadFlag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadFlag.Merge(dst, src)
}
func (m *ThreadFlag) XXX_Size() int {
	return xxx_messageInfo_ThreadFlag.Size(m)
}
func (m *ThreadFlag) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadFlag.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadFlag proto.InternalMessageInfo

func (m *ThreadFlag) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type ThreadJoin struct {
	Inviter              string   `protobuf:"bytes,1,opt,name=inviter,proto3" json:"inviter,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Inboxes              []string `protobuf:"bytes,3,rep,name=inboxes,proto3" json:"inboxes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadJoin) Reset()         { *m = ThreadJoin{} }
func (m *ThreadJoin) String() string { return proto.CompactTextString(m) }
func (*ThreadJoin) ProtoMessage()    {}
func (*ThreadJoin) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_890bcf6c9a2f6ba0, []int{6}
}
func (m *ThreadJoin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadJoin.Unmarshal(m, b)
}
func (m *ThreadJoin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadJoin.Marshal(b, m, deterministic)
}
func (dst *ThreadJoin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadJoin.Merge(dst, src)
}
func (m *ThreadJoin) XXX_Size() int {
	return xxx_messageInfo_ThreadJoin.Size(m)
}
func (m *ThreadJoin) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadJoin.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadJoin proto.InternalMessageInfo

func (m *ThreadJoin) GetInviter() string {
	if m != nil {
		return m.Inviter
	}
	return ""
}

func (m *ThreadJoin) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ThreadJoin) GetInboxes() []string {
	if m != nil {
		return m.Inboxes
	}
	return nil
}

type ThreadAnnounce struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Inboxes              []string `protobuf:"bytes,2,rep,name=inboxes,proto3" json:"inboxes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadAnnounce) Reset()         { *m = ThreadAnnounce{} }
func (m *ThreadAnnounce) String() string { return proto.CompactTextString(m) }
func (*ThreadAnnounce) ProtoMessage()    {}
func (*ThreadAnnounce) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_890bcf6c9a2f6ba0, []int{7}
}
func (m *ThreadAnnounce) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadAnnounce.Unmarshal(m, b)
}
func (m *ThreadAnnounce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadAnnounce.Marshal(b, m, deterministic)
}
func (dst *ThreadAnnounce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadAnnounce.Merge(dst, src)
}
func (m *ThreadAnnounce) XXX_Size() int {
	return xxx_messageInfo_ThreadAnnounce.Size(m)
}
func (m *ThreadAnnounce) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadAnnounce.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadAnnounce proto.InternalMessageInfo

func (m *ThreadAnnounce) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ThreadAnnounce) GetInboxes() []string {
	if m != nil {
		return m.Inboxes
	}
	return nil
}

type ThreadData struct {
	Type                 ThreadData_Type `protobuf:"varint,1,opt,name=type,proto3,enum=ThreadData_Type" json:"type,omitempty"`
	Data                 string          `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Key                  []byte          `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Caption              string          `protobuf:"bytes,4,opt,name=caption,proto3" json:"caption,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ThreadData) Reset()         { *m = ThreadData{} }
func (m *ThreadData) String() string { return proto.CompactTextString(m) }
func (*ThreadData) ProtoMessage()    {}
func (*ThreadData) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_890bcf6c9a2f6ba0, []int{8}
}
func (m *ThreadData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadData.Unmarshal(m, b)
}
func (m *ThreadData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadData.Marshal(b, m, deterministic)
}
func (dst *ThreadData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadData.Merge(dst, src)
}
func (m *ThreadData) XXX_Size() int {
	return xxx_messageInfo_ThreadData.Size(m)
}
func (m *ThreadData) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadData.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadData proto.InternalMessageInfo

func (m *ThreadData) GetType() ThreadData_Type {
	if m != nil {
		return m.Type
	}
	return ThreadData_FILE
}

func (m *ThreadData) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *ThreadData) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *ThreadData) GetCaption() string {
	if m != nil {
		return m.Caption
	}
	return ""
}

type ThreadAnnotation struct {
	Type                 ThreadAnnotation_Type `protobuf:"varint,1,opt,name=type,proto3,enum=ThreadAnnotation_Type" json:"type,omitempty"`
	Data                 string                `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Caption              string                `protobuf:"bytes,3,opt,name=caption,proto3" json:"caption,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ThreadAnnotation) Reset()         { *m = ThreadAnnotation{} }
func (m *ThreadAnnotation) String() string { return proto.CompactTextString(m) }
func (*ThreadAnnotation) ProtoMessage()    {}
func (*ThreadAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_890bcf6c9a2f6ba0, []int{9}
}
func (m *ThreadAnnotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadAnnotation.Unmarshal(m, b)
}
func (m *ThreadAnnotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadAnnotation.Marshal(b, m, deterministic)
}
func (dst *ThreadAnnotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadAnnotation.Merge(dst, src)
}
func (m *ThreadAnnotation) XXX_Size() int {
	return xxx_messageInfo_ThreadAnnotation.Size(m)
}
func (m *ThreadAnnotation) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadAnnotation.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadAnnotation proto.InternalMessageInfo

func (m *ThreadAnnotation) GetType() ThreadAnnotation_Type {
	if m != nil {
		return m.Type
	}
	return ThreadAnnotation_COMMENT
}

func (m *ThreadAnnotation) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *ThreadAnnotation) GetCaption() string {
	if m != nil {
		return m.Caption
	}
	return ""
}

func init() {
	proto.RegisterType((*ThreadEnvelope)(nil), "ThreadEnvelope")
	proto.RegisterType((*ThreadBlock)(nil), "ThreadBlock")
	proto.RegisterType((*ThreadBlockHeader)(nil), "ThreadBlockHeader")
	proto.RegisterType((*ThreadInvite)(nil), "ThreadInvite")
	proto.RegisterType((*ThreadIgnore)(nil), "ThreadIgnore")
	proto.RegisterType((*ThreadFlag)(nil), "ThreadFlag")
	proto.RegisterType((*ThreadJoin)(nil), "ThreadJoin")
	proto.RegisterType((*ThreadAnnounce)(nil), "ThreadAnnounce")
	proto.RegisterType((*ThreadData)(nil), "ThreadData")
	proto.RegisterType((*ThreadAnnotation)(nil), "ThreadAnnotation")
	proto.RegisterEnum("ThreadBlock_Type", ThreadBlock_Type_name, ThreadBlock_Type_value)
	proto.RegisterEnum("ThreadData_Type", ThreadData_Type_name, ThreadData_Type_value)
	proto.RegisterEnum("ThreadAnnotation_Type", ThreadAnnotation_Type_name, ThreadAnnotation_Type_value)
}

func init() { proto.RegisterFile("thread.proto", fileDescriptor_thread_890bcf6c9a2f6ba0) }

var fileDescriptor_thread_890bcf6c9a2f6ba0 = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x5d, 0x6b, 0xdb, 0x30,
	0x14, 0x9d, 0x3f, 0x9a, 0x26, 0x37, 0x21, 0xa8, 0x62, 0x14, 0x2f, 0x30, 0x16, 0xcc, 0x06, 0xa5,
	0x0f, 0x2e, 0x64, 0xbf, 0xc0, 0x6d, 0xdd, 0xce, 0x5d, 0xea, 0x80, 0xf1, 0xca, 0x18, 0x63, 0xa0,
	0x24, 0x5a, 0x6c, 0x92, 0x4a, 0xc6, 0x56, 0xca, 0xfc, 0xba, 0xd7, 0xc1, 0x7e, 0xf3, 0x90, 0x2c,
	0x37, 0x4e, 0xbb, 0xb1, 0xb7, 0x7b, 0x7c, 0xcf, 0xbd, 0xf7, 0x1c, 0xeb, 0xc0, 0x40, 0xa4, 0x05,
	0x25, 0x4b, 0x2f, 0x2f, 0xb8, 0xe0, 0xa3, 0x57, 0x2b, 0xce, 0x57, 0x1b, 0x7a, 0xa6, 0xd0, 0x7c,
	0xfb, 0xfd, 0x8c, 0xb0, 0x4a, 0xb7, 0xde, 0x3c, 0x6d, 0x89, 0xec, 0x9e, 0x96, 0x82, 0xdc, 0xe7,
	0x35, 0xc1, 0xfd, 0x06, 0xc3, 0x44, 0xed, 0x0a, 0xd8, 0x03, 0xdd, 0xf0, 0x9c, 0xe2, 0x63, 0xe8,
	0xd4, 0xdb, 0x1d, 0x63, 0x6c, 0x9c, 0xf4, 0x62, 0x8d, 0x30, 0x06, 0x3b, 0x25, 0x65, 0xea, 0x98,
	0xea, 0xab, 0xaa, 0xf1, 0x18, 0xfa, 0x8b, 0x2c, 0x4f, 0x69, 0x71, 0xbe, 0xe1, 0x8b, 0xb5, 0x63,
	0x8d, 0x8d, 0x93, 0x41, 0xdc, 0xfe, 0xe4, 0xfe, 0x34, 0xa1, 0x5f, 0x1f, 0x50, 0x18, 0x9f, 0x42,
	0x27, 0xa5, 0x64, 0x49, 0x0b, 0xb5, 0xbd, 0x3f, 0xc1, 0x5e, 0xab, 0xfb, 0x41, 0x75, 0x62, 0xcd,
	0xc0, 0xef, 0xc0, 0x16, 0x55, 0x4e, 0xd5, 0xc5, 0xe1, 0xe4, 0xa8, 0xcd, 0xf4, 0x92, 0x2a, 0xa7,
	0xb1, 0x6a, 0x63, 0x0f, 0x0e, 0x73, 0x52, 0x6d, 0x38, 0x59, 0x2a, 0x01, 0xfd, 0xc9, 0x4b, 0xaf,
	0x76, 0xed, 0x35, 0xae, 0x3d, 0x9f, 0x55, 0x71, 0x43, 0x72, 0x73, 0xb0, 0xe5, 0x34, 0xee, 0xc1,
	0xc1, 0x6d, 0x10, 0x5f, 0x07, 0xe8, 0x05, 0x06, 0xe8, 0x84, 0xd7, 0xd1, 0x2c, 0x0e, 0x90, 0x81,
	0xbb, 0x60, 0x5f, 0x4d, 0xfd, 0x6b, 0x64, 0xca, 0xea, 0x66, 0x16, 0x46, 0xc8, 0xc2, 0x03, 0xe8,
	0xfa, 0x51, 0x34, 0xfb, 0x14, 0x5d, 0x04, 0xc8, 0x96, 0x83, 0xd3, 0xc0, 0xbf, 0x0b, 0xd0, 0x81,
	0xa4, 0x5c, 0xfa, 0x89, 0x8f, 0x3a, 0x78, 0x08, 0x20, 0x29, 0x89, 0x9f, 0x84, 0xb3, 0x08, 0x1d,
	0xaa, 0x95, 0xd1, 0x5d, 0x98, 0x04, 0x68, 0xe2, 0x6e, 0xe1, 0xe8, 0x99, 0x4b, 0xec, 0x81, 0xbd,
	0x24, 0x82, 0xea, 0xff, 0x30, 0x7a, 0xa6, 0x39, 0x69, 0x5e, 0x2a, 0x56, 0x3c, 0xec, 0x48, 0x9b,
	0x05, 0x65, 0xa2, 0x74, 0xcc, 0xb1, 0x75, 0xd2, 0x8b, 0x1b, 0x28, 0x5f, 0x8c, 0x6c, 0x45, 0xca,
	0x0b, 0xe5, 0xbf, 0x17, 0x6b, 0xe4, 0x4e, 0x60, 0x50, 0x9f, 0x0d, 0xd9, 0x43, 0x26, 0x28, 0x1e,
	0x82, 0x59, 0xae, 0xd5, 0xbd, 0x41, 0x6c, 0x96, 0x6b, 0xf9, 0xa2, 0x8c, 0xdc, 0xd3, 0xe6, 0x45,
	0x65, 0xed, 0xba, 0x8f, 0x33, 0x2b, 0xc6, 0x0b, 0x2a, 0x39, 0x4b, 0x22, 0x88, 0xce, 0x82, 0xaa,
	0xdd, 0x31, 0x40, 0xcd, 0xb9, 0xda, 0x90, 0xd5, 0x5f, 0x19, 0x5f, 0x1b, 0xc6, 0x0d, 0xcf, 0x98,
	0x54, 0x9e, 0x29, 0x05, 0x85, 0x26, 0x35, 0x10, 0x8f, 0xa0, 0xbb, 0x2d, 0x69, 0xd1, 0x52, 0xf1,
	0x88, 0xeb, 0xa9, 0x39, 0xff, 0x41, 0x4b, 0xc7, 0xaa, 0xfd, 0x6a, 0xe8, 0x5e, 0x35, 0x99, 0xf5,
	0x19, 0xe3, 0x5b, 0xb6, 0xa0, 0x7b, 0x7b, 0x8c, 0x7f, 0xef, 0x31, 0xf7, 0xf7, 0xfc, 0x36, 0x1a,
	0x99, 0x97, 0x44, 0x10, 0xfc, 0x56, 0xc7, 0xcd, 0x50, 0x71, 0x43, 0xde, 0xae, 0xd5, 0x4e, 0x5b,
	0x63, 0xd7, 0xdc, 0xd9, 0xc5, 0x08, 0xac, 0x35, 0xad, 0x74, 0xfc, 0x65, 0x29, 0x8f, 0x2e, 0x48,
	0x2e, 0x32, 0xce, 0x1c, 0xbb, 0xb6, 0xac, 0xa1, 0x3b, 0xd2, 0xe9, 0x93, 0x31, 0x0b, 0xa7, 0x32,
	0x7c, 0x5d, 0xb0, 0x93, 0xe0, 0x73, 0x82, 0x0c, 0xf7, 0x97, 0x01, 0x68, 0xe7, 0x4c, 0x10, 0x39,
	0x80, 0x4f, 0xf7, 0x64, 0x1d, 0x7b, 0x4f, 0x09, 0xff, 0x13, 0xd7, 0x92, 0x62, 0xed, 0x4b, 0x79,
	0xad, 0xa5, 0xf4, 0xe1, 0xf0, 0x62, 0x76, 0x7b, 0x1b, 0x44, 0x49, 0xad, 0x66, 0x1a, 0x7e, 0x0c,
	0x90, 0x71, 0x6e, 0x7f, 0x31, 0xf3, 0xf9, 0xbc, 0xa3, 0x02, 0xf9, 0xfe, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xbb, 0x93, 0x8f, 0xb0, 0x73, 0x04, 0x00, 0x00,
}
