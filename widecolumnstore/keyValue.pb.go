// Code generated by protoc-gen-go. DO NOT EDIT.
// source: keyValue.proto

package widecolumnstore

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type KeyValue struct {
	Key                  *Key     `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value                *any.Any `protobuf:"bytes,5,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyValue) Reset()         { *m = KeyValue{} }
func (m *KeyValue) String() string { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()    {}
func (*KeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_029fe9abc084c724, []int{0}
}

func (m *KeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValue.Unmarshal(m, b)
}
func (m *KeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValue.Marshal(b, m, deterministic)
}
func (m *KeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValue.Merge(m, src)
}
func (m *KeyValue) XXX_Size() int {
	return xxx_messageInfo_KeyValue.Size(m)
}
func (m *KeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValue proto.InternalMessageInfo

func (m *KeyValue) GetKey() *Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KeyValue) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

type Key struct {
	// RowKey Unique identifier for the row.
	RowKey []byte `protobuf:"bytes,1,opt,name=RowKey,proto3" json:"RowKey,omitempty"`
	// ColumnFamily This field can be used to partition data within a node.
	ColumnFamily []byte `protobuf:"bytes,2,opt,name=ColumnFamily,proto3" json:"ColumnFamily,omitempty"`
	// QualColumnQualifierifier More specific attribute of the key.
	ColumnQualifier      []byte               `protobuf:"bytes,3,opt,name=ColumnQualifier,proto3" json:"ColumnQualifier,omitempty"`
	Time                 *timestamp.Timestamp `protobuf:"bytes,4,opt,name=Time,proto3" json:"Time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_029fe9abc084c724, []int{1}
}

func (m *Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key.Unmarshal(m, b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key.Marshal(b, m, deterministic)
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return xxx_messageInfo_Key.Size(m)
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetRowKey() []byte {
	if m != nil {
		return m.RowKey
	}
	return nil
}

func (m *Key) GetColumnFamily() []byte {
	if m != nil {
		return m.ColumnFamily
	}
	return nil
}

func (m *Key) GetColumnQualifier() []byte {
	if m != nil {
		return m.ColumnQualifier
	}
	return nil
}

func (m *Key) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func init() {
	proto.RegisterType((*KeyValue)(nil), "widecolumnstore.KeyValue")
	proto.RegisterType((*Key)(nil), "widecolumnstore.Key")
}

func init() { proto.RegisterFile("keyValue.proto", fileDescriptor_029fe9abc084c724) }

var fileDescriptor_029fe9abc084c724 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4e, 0xad, 0x0c,
	0x4b, 0xcc, 0x29, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2f, 0xcf, 0x4c, 0x49,
	0x4d, 0xce, 0xcf, 0x29, 0xcd, 0xcd, 0x2b, 0x2e, 0xc9, 0x2f, 0x4a, 0x95, 0x92, 0x4f, 0xcf, 0xcf,
	0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x4b, 0x27, 0x95, 0xa6, 0xe9, 0x97, 0x64, 0xe6, 0xa6, 0x16, 0x97,
	0x24, 0xe6, 0x16, 0x40, 0x74, 0x48, 0x49, 0xa2, 0x2b, 0x48, 0xcc, 0xab, 0x84, 0x48, 0x29, 0xc5,
	0x71, 0x71, 0x78, 0x43, 0x8d, 0x17, 0x52, 0xe3, 0x62, 0xf6, 0x4e, 0xad, 0x94, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x36, 0x12, 0xd1, 0x43, 0xb3, 0x46, 0xcf, 0x3b, 0xb5, 0x32, 0x08, 0xa4, 0x40, 0x48,
	0x8b, 0x8b, 0x15, 0xac, 0x41, 0x82, 0x15, 0xaa, 0x12, 0x62, 0xbc, 0x1e, 0xcc, 0x78, 0x3d, 0xc7,
	0xbc, 0xca, 0x20, 0x88, 0x12, 0xa5, 0xd9, 0x8c, 0x60, 0x43, 0x85, 0xc4, 0xb8, 0xd8, 0x82, 0xf2,
	0xcb, 0x61, 0xc6, 0xf3, 0x04, 0x41, 0x79, 0x42, 0x4a, 0x5c, 0x3c, 0xce, 0x60, 0x3b, 0xdc, 0x12,
	0x73, 0x33, 0x73, 0x2a, 0x25, 0x98, 0xc0, 0xb2, 0x28, 0x62, 0x42, 0x1a, 0x5c, 0xfc, 0x10, 0x7e,
	0x60, 0x69, 0x62, 0x4e, 0x66, 0x5a, 0x66, 0x6a, 0x91, 0x04, 0x33, 0x58, 0x19, 0xba, 0xb0, 0x90,
	0x1e, 0x17, 0x4b, 0x48, 0x66, 0x6e, 0xaa, 0x04, 0x0b, 0xd8, 0x61, 0x52, 0x18, 0x0e, 0x0b, 0x81,
	0x05, 0x4c, 0x10, 0x58, 0x5d, 0x12, 0x1b, 0x58, 0xc6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xdc,
	0x1e, 0x4a, 0x08, 0x63, 0x01, 0x00, 0x00,
}
