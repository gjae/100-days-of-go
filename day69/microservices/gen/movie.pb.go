// Code generated by protoc-gen-go. DO NOT EDIT.
// source: movie.proto

package gen

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

type Metadata struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Director             string   `protobuf:"bytes,4,opt,name=director,proto3" json:"director,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{0}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Metadata) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Metadata) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Metadata) GetDirector() string {
	if m != nil {
		return m.Director
	}
	return ""
}

type MovieDetails struct {
	Rating               float32   `protobuf:"fixed32,1,opt,name=rating,proto3" json:"rating,omitempty"`
	Metadata             *Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MovieDetails) Reset()         { *m = MovieDetails{} }
func (m *MovieDetails) String() string { return proto.CompactTextString(m) }
func (*MovieDetails) ProtoMessage()    {}
func (*MovieDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{1}
}

func (m *MovieDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MovieDetails.Unmarshal(m, b)
}
func (m *MovieDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MovieDetails.Marshal(b, m, deterministic)
}
func (m *MovieDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MovieDetails.Merge(m, src)
}
func (m *MovieDetails) XXX_Size() int {
	return xxx_messageInfo_MovieDetails.Size(m)
}
func (m *MovieDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_MovieDetails.DiscardUnknown(m)
}

var xxx_messageInfo_MovieDetails proto.InternalMessageInfo

func (m *MovieDetails) GetRating() float32 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *MovieDetails) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*Metadata)(nil), "Metadata")
	proto.RegisterType((*MovieDetails)(nil), "MovieDetails")
}

func init() {
	proto.RegisterFile("movie.proto", fileDescriptor_fde087a4194eda75)
}

var fileDescriptor_fde087a4194eda75 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x41, 0x0b, 0x82, 0x40,
	0x10, 0x46, 0x71, 0x33, 0xd1, 0x31, 0x3a, 0x2c, 0x11, 0x4b, 0x27, 0x11, 0x82, 0x4e, 0x06, 0xf5,
	0x0f, 0xa2, 0xab, 0x17, 0x8f, 0xdd, 0x36, 0x77, 0x90, 0x01, 0xdd, 0x95, 0x75, 0xe8, 0xf7, 0x47,
	0x9b, 0x49, 0xc7, 0x37, 0xef, 0xf0, 0xbe, 0x81, 0x7c, 0x70, 0x2f, 0xc2, 0x6a, 0xf4, 0x8e, 0x5d,
	0x69, 0x21, 0xad, 0x91, 0xb5, 0xd1, 0xac, 0xe5, 0x16, 0x04, 0x19, 0x15, 0x15, 0xd1, 0x29, 0x6b,
	0x04, 0x19, 0xb9, 0x83, 0x35, 0x13, 0xf7, 0xa8, 0x44, 0x38, 0x7d, 0x41, 0x16, 0x90, 0x1b, 0x9c,
	0x5a, 0x4f, 0x23, 0x93, 0xb3, 0x6a, 0x15, 0xdc, 0xff, 0x49, 0x1e, 0x20, 0x35, 0xe4, 0xb1, 0x65,
	0xe7, 0x55, 0x1c, 0xf4, 0xc2, 0x65, 0x0d, 0x9b, 0xfa, 0x93, 0xbf, 0x23, 0x6b, 0xea, 0x27, 0xb9,
	0x87, 0xc4, 0x6b, 0x26, 0xdb, 0x85, 0xae, 0x68, 0x66, 0x92, 0x47, 0x48, 0x87, 0x79, 0x57, 0xc8,
	0xe7, 0x97, 0xac, 0xfa, 0x0d, 0x6d, 0x16, 0x75, 0x4b, 0x1e, 0xf1, 0xb9, 0x43, 0xfb, 0x4c, 0xc2,
	0x37, 0xd7, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4f, 0xdd, 0x04, 0x27, 0xdc, 0x00, 0x00, 0x00,
}