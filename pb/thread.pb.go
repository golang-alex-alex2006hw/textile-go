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
	ThreadBlock_MERGE    ThreadBlock_Type = 0
	ThreadBlock_IGNORE   ThreadBlock_Type = 1
	ThreadBlock_FLAG     ThreadBlock_Type = 2
	ThreadBlock_JOIN     ThreadBlock_Type = 3
	ThreadBlock_ANNOUNCE ThreadBlock_Type = 4
	ThreadBlock_LEAVE    ThreadBlock_Type = 5
	ThreadBlock_MESSAGE  ThreadBlock_Type = 6
	ThreadBlock_FILES    ThreadBlock_Type = 7
	ThreadBlock_COMMENT  ThreadBlock_Type = 8
	ThreadBlock_LIKE     ThreadBlock_Type = 9
	ThreadBlock_INVITE   ThreadBlock_Type = 50
)

var ThreadBlock_Type_name = map[int32]string{
	0:  "MERGE",
	1:  "IGNORE",
	2:  "FLAG",
	3:  "JOIN",
	4:  "ANNOUNCE",
	5:  "LEAVE",
	6:  "MESSAGE",
	7:  "FILES",
	8:  "COMMENT",
	9:  "LIKE",
	50: "INVITE",
}
var ThreadBlock_Type_value = map[string]int32{
	"MERGE":    0,
	"IGNORE":   1,
	"FLAG":     2,
	"JOIN":     3,
	"ANNOUNCE": 4,
	"LEAVE":    5,
	"MESSAGE":  6,
	"FILES":    7,
	"COMMENT":  8,
	"LIKE":     9,
	"INVITE":   50,
}

func (x ThreadBlock_Type) String() string {
	return proto.EnumName(ThreadBlock_Type_name, int32(x))
}
func (ThreadBlock_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_thread_7287cd426a3ee648, []int{1, 0}
}

// for wire transport
type ThreadEnvelope struct {
	Thread               string   `protobuf:"bytes,1,opt,name=thread,proto3" json:"thread,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Ciphertext           []byte   `protobuf:"bytes,3,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadEnvelope) Reset()         { *m = ThreadEnvelope{} }
func (m *ThreadEnvelope) String() string { return proto.CompactTextString(m) }
func (*ThreadEnvelope) ProtoMessage()    {}
func (*ThreadEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_7287cd426a3ee648, []int{0}
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

func (m *ThreadEnvelope) GetCiphertext() []byte {
	if m != nil {
		return m.Ciphertext
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
	return fileDescriptor_thread_7287cd426a3ee648, []int{1}
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
	return fileDescriptor_thread_7287cd426a3ee648, []int{2}
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
	return fileDescriptor_thread_7287cd426a3ee648, []int{3}
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
	Target               string   `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadIgnore) Reset()         { *m = ThreadIgnore{} }
func (m *ThreadIgnore) String() string { return proto.CompactTextString(m) }
func (*ThreadIgnore) ProtoMessage()    {}
func (*ThreadIgnore) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_7287cd426a3ee648, []int{4}
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

func (m *ThreadIgnore) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

type ThreadFlag struct {
	Target               string   `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadFlag) Reset()         { *m = ThreadFlag{} }
func (m *ThreadFlag) String() string { return proto.CompactTextString(m) }
func (*ThreadFlag) ProtoMessage()    {}
func (*ThreadFlag) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_7287cd426a3ee648, []int{5}
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

func (m *ThreadFlag) GetTarget() string {
	if m != nil {
		return m.Target
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
	return fileDescriptor_thread_7287cd426a3ee648, []int{6}
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
	return fileDescriptor_thread_7287cd426a3ee648, []int{7}
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

type ThreadMessage struct {
	Body                 string   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadMessage) Reset()         { *m = ThreadMessage{} }
func (m *ThreadMessage) String() string { return proto.CompactTextString(m) }
func (*ThreadMessage) ProtoMessage()    {}
func (*ThreadMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_7287cd426a3ee648, []int{8}
}
func (m *ThreadMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadMessage.Unmarshal(m, b)
}
func (m *ThreadMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadMessage.Marshal(b, m, deterministic)
}
func (dst *ThreadMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadMessage.Merge(dst, src)
}
func (m *ThreadMessage) XXX_Size() int {
	return xxx_messageInfo_ThreadMessage.Size(m)
}
func (m *ThreadMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadMessage proto.InternalMessageInfo

func (m *ThreadMessage) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type ThreadFiles struct {
	Target               string            `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Body                 string            `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Keys                 map[string]string `protobuf:"bytes,3,rep,name=keys,proto3" json:"keys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ThreadFiles) Reset()         { *m = ThreadFiles{} }
func (m *ThreadFiles) String() string { return proto.CompactTextString(m) }
func (*ThreadFiles) ProtoMessage()    {}
func (*ThreadFiles) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_7287cd426a3ee648, []int{9}
}
func (m *ThreadFiles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadFiles.Unmarshal(m, b)
}
func (m *ThreadFiles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadFiles.Marshal(b, m, deterministic)
}
func (dst *ThreadFiles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadFiles.Merge(dst, src)
}
func (m *ThreadFiles) XXX_Size() int {
	return xxx_messageInfo_ThreadFiles.Size(m)
}
func (m *ThreadFiles) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadFiles.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadFiles proto.InternalMessageInfo

func (m *ThreadFiles) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *ThreadFiles) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *ThreadFiles) GetKeys() map[string]string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type ThreadComment struct {
	Target               string   `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadComment) Reset()         { *m = ThreadComment{} }
func (m *ThreadComment) String() string { return proto.CompactTextString(m) }
func (*ThreadComment) ProtoMessage()    {}
func (*ThreadComment) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_7287cd426a3ee648, []int{10}
}
func (m *ThreadComment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadComment.Unmarshal(m, b)
}
func (m *ThreadComment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadComment.Marshal(b, m, deterministic)
}
func (dst *ThreadComment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadComment.Merge(dst, src)
}
func (m *ThreadComment) XXX_Size() int {
	return xxx_messageInfo_ThreadComment.Size(m)
}
func (m *ThreadComment) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadComment.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadComment proto.InternalMessageInfo

func (m *ThreadComment) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *ThreadComment) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type ThreadLike struct {
	Target               string   `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadLike) Reset()         { *m = ThreadLike{} }
func (m *ThreadLike) String() string { return proto.CompactTextString(m) }
func (*ThreadLike) ProtoMessage()    {}
func (*ThreadLike) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_7287cd426a3ee648, []int{11}
}
func (m *ThreadLike) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadLike.Unmarshal(m, b)
}
func (m *ThreadLike) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadLike.Marshal(b, m, deterministic)
}
func (dst *ThreadLike) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadLike.Merge(dst, src)
}
func (m *ThreadLike) XXX_Size() int {
	return xxx_messageInfo_ThreadLike.Size(m)
}
func (m *ThreadLike) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadLike.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadLike proto.InternalMessageInfo

func (m *ThreadLike) GetTarget() string {
	if m != nil {
		return m.Target
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
	proto.RegisterType((*ThreadMessage)(nil), "ThreadMessage")
	proto.RegisterType((*ThreadFiles)(nil), "ThreadFiles")
	proto.RegisterMapType((map[string]string)(nil), "ThreadFiles.KeysEntry")
	proto.RegisterType((*ThreadComment)(nil), "ThreadComment")
	proto.RegisterType((*ThreadLike)(nil), "ThreadLike")
	proto.RegisterEnum("ThreadBlock_Type", ThreadBlock_Type_name, ThreadBlock_Type_value)
}

func init() { proto.RegisterFile("thread.proto", fileDescriptor_thread_7287cd426a3ee648) }

var fileDescriptor_thread_7287cd426a3ee648 = []byte{
	// 610 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0xad, 0x8d, 0x43, 0x60, 0xa0, 0x91, 0xb3, 0x8a, 0x22, 0xca, 0xa1, 0x8d, 0xdc, 0x0f, 0x45,
	0x39, 0x38, 0x12, 0x3d, 0xb4, 0x6a, 0x4f, 0x24, 0x5a, 0x28, 0x09, 0x18, 0xc9, 0xa1, 0x39, 0x54,
	0xb9, 0x2c, 0x61, 0x0a, 0x16, 0x66, 0xd7, 0xb2, 0xd7, 0x28, 0xfe, 0x01, 0xfd, 0x0f, 0x3d, 0xf5,
	0xb7, 0x56, 0xbb, 0xb6, 0x09, 0x69, 0xc4, 0xa1, 0xb7, 0x7d, 0x33, 0x6f, 0xde, 0xcc, 0xdb, 0x19,
	0x68, 0xca, 0x45, 0x8c, 0x6c, 0xe6, 0x46, 0xb1, 0x90, 0xa2, 0xfd, 0x6a, 0x2e, 0xc4, 0x3c, 0xc4,
	0x73, 0x8d, 0xa6, 0xe9, 0xcf, 0x73, 0xc6, 0xb3, 0x22, 0xf5, 0xe6, 0xdf, 0x94, 0x0c, 0x56, 0x98,
	0x48, 0xb6, 0x8a, 0x72, 0x82, 0x73, 0x07, 0x07, 0x13, 0xad, 0x45, 0xf9, 0x1a, 0x43, 0x11, 0x21,
	0x39, 0x86, 0x6a, 0xae, 0xde, 0x32, 0x4e, 0x8c, 0xd3, 0xba, 0x5f, 0x20, 0x42, 0xc0, 0x5a, 0xb0,
	0x64, 0xd1, 0x32, 0x75, 0x54, 0xbf, 0xc9, 0x6b, 0x80, 0xfb, 0x20, 0x5a, 0x60, 0x2c, 0xf1, 0x41,
	0xb6, 0x2a, 0x27, 0xc6, 0x69, 0xd3, 0xdf, 0x8a, 0x38, 0xbf, 0x4d, 0x68, 0xe4, 0xf2, 0x17, 0xa1,
	0xb8, 0x5f, 0x92, 0x33, 0xa8, 0x2e, 0x90, 0xcd, 0x30, 0xd6, 0xda, 0x8d, 0x0e, 0x71, 0xb7, 0xb2,
	0xdf, 0x74, 0xc6, 0x2f, 0x18, 0xe4, 0x3d, 0x58, 0x32, 0x8b, 0x50, 0xf7, 0x3b, 0xe8, 0x1c, 0x6e,
	0x33, 0xdd, 0x49, 0x16, 0xa1, 0xaf, 0xd3, 0xc4, 0x85, 0xfd, 0x88, 0x65, 0xa1, 0x60, 0x33, 0xdd,
	0xbf, 0xd1, 0x39, 0x72, 0x73, 0xcf, 0x6e, 0xe9, 0xd9, 0xed, 0xf2, 0xcc, 0x2f, 0x49, 0xce, 0x2f,
	0x03, 0x2c, 0x55, 0x4e, 0xea, 0xb0, 0x37, 0xa2, 0x7e, 0x9f, 0xda, 0x2f, 0x08, 0x40, 0x75, 0xd0,
	0xf7, 0xc6, 0x3e, 0xb5, 0x0d, 0x52, 0x03, 0xab, 0x37, 0xec, 0xf6, 0x6d, 0x53, 0xbd, 0xae, 0xc6,
	0x03, 0xcf, 0xae, 0x90, 0x26, 0xd4, 0xba, 0x9e, 0x37, 0xfe, 0xee, 0x5d, 0x52, 0xdb, 0x52, 0x85,
	0x43, 0xda, 0xbd, 0xa5, 0xf6, 0x1e, 0x69, 0xc0, 0xfe, 0x88, 0xde, 0xdc, 0x74, 0xfb, 0xd4, 0xae,
	0xaa, 0x78, 0x6f, 0x30, 0xa4, 0x37, 0xf6, 0xbe, 0x8a, 0x5f, 0x8e, 0x47, 0x23, 0xea, 0x4d, 0xec,
	0x9a, 0xd2, 0x19, 0x0e, 0xae, 0xa9, 0x5d, 0xd7, 0x7d, 0xbc, 0xdb, 0xc1, 0x84, 0xda, 0x1d, 0x27,
	0x85, 0xc3, 0x67, 0xde, 0x89, 0x0b, 0xd6, 0x8c, 0x49, 0x2c, 0x7e, 0xa7, 0xfd, 0xcc, 0xc9, 0xa4,
	0xdc, 0x9e, 0xaf, 0x79, 0xa4, 0xa5, 0xcc, 0xc7, 0xc8, 0x65, 0xd2, 0x32, 0x4f, 0x2a, 0xa7, 0x75,
	0xbf, 0x84, 0x6a, 0x8b, 0x2c, 0x95, 0x0b, 0x11, 0xeb, 0x5f, 0xa9, 0xfb, 0x05, 0x72, 0x3a, 0xd0,
	0xcc, 0xdb, 0x0e, 0xf8, 0x3a, 0x90, 0x48, 0x0e, 0xc0, 0x4c, 0x96, 0xba, 0x5f, 0xd3, 0x37, 0x93,
	0xa5, 0xda, 0x32, 0x67, 0x2b, 0x2c, 0xb7, 0xac, 0xde, 0xce, 0x87, 0x4d, 0xcd, 0x9c, 0x8b, 0x38,
	0xbf, 0x10, 0x16, 0xcf, 0x51, 0x6e, 0x2e, 0x44, 0x23, 0xe7, 0x1d, 0x40, 0xce, 0xeb, 0x85, 0x6c,
	0xbe, 0x93, 0x75, 0x57, 0xb2, 0xae, 0x44, 0xc0, 0x95, 0x83, 0x40, 0x4f, 0x12, 0x17, 0xb4, 0x12,
	0x92, 0x36, 0xd4, 0xd2, 0x04, 0xe3, 0xad, 0x69, 0x36, 0x38, 0xaf, 0x9a, 0x8a, 0x07, 0x4c, 0x5a,
	0x95, 0xdc, 0x77, 0x01, 0x9d, 0x5e, 0x79, 0xcf, 0x5d, 0xce, 0x45, 0xca, 0xef, 0xf1, 0x89, 0x8e,
	0xb1, 0x5b, 0xc7, 0x7c, 0xaa, 0xf3, 0x16, 0x5e, 0xe6, 0x3a, 0x23, 0x4c, 0x12, 0x36, 0x47, 0xf5,
	0x31, 0x53, 0x31, 0xcb, 0x0a, 0x09, 0xfd, 0x76, 0xfe, 0x18, 0xe5, 0x79, 0xf7, 0x82, 0x10, 0x93,
	0x5d, 0x96, 0x37, 0xb5, 0xe6, 0x63, 0x2d, 0x39, 0x03, 0x6b, 0x89, 0x59, 0x3e, 0x7f, 0xa3, 0x73,
	0xec, 0x6e, 0xe9, 0xb8, 0xd7, 0x98, 0x25, 0x94, 0xcb, 0x38, 0xf3, 0x35, 0xa7, 0xfd, 0x09, 0xea,
	0x9b, 0x10, 0xb1, 0xa1, 0xb2, 0xc4, 0x72, 0x0e, 0xf5, 0x24, 0x47, 0xb0, 0xb7, 0x66, 0x61, 0x5a,
	0x7e, 0x53, 0x0e, 0xbe, 0x98, 0x9f, 0x0d, 0xe7, 0x6b, 0xe9, 0xe2, 0x52, 0xac, 0x56, 0xc8, 0xe5,
	0xff, 0x4c, 0xf8, 0xb8, 0xce, 0x61, 0xb0, 0xdc, 0xb9, 0xf4, 0x0b, 0xeb, 0x87, 0x19, 0x4d, 0xa7,
	0x55, 0x7d, 0xa2, 0x1f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x58, 0x58, 0x5d, 0x54, 0x99, 0x04,
	0x00, 0x00,
}