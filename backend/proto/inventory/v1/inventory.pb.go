// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inventory/v1/inventory.proto

package inventoryv1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetEntityRequest struct {
	Entity               *Entity  `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	IncludeFacts         []string `protobuf:"bytes,2,rep,name=include_facts,json=includeFacts,proto3" json:"include_facts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEntityRequest) Reset()         { *m = GetEntityRequest{} }
func (m *GetEntityRequest) String() string { return proto.CompactTextString(m) }
func (*GetEntityRequest) ProtoMessage()    {}
func (*GetEntityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_70be9028e322f9d8, []int{0}
}

func (m *GetEntityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEntityRequest.Unmarshal(m, b)
}
func (m *GetEntityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEntityRequest.Marshal(b, m, deterministic)
}
func (m *GetEntityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEntityRequest.Merge(m, src)
}
func (m *GetEntityRequest) XXX_Size() int {
	return xxx_messageInfo_GetEntityRequest.Size(m)
}
func (m *GetEntityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEntityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEntityRequest proto.InternalMessageInfo

func (m *GetEntityRequest) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *GetEntityRequest) GetIncludeFacts() []string {
	if m != nil {
		return m.IncludeFacts
	}
	return nil
}

type GetEntityReply struct {
	Entity               *Entity  `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Facts                []*Fact  `protobuf:"bytes,2,rep,name=facts,proto3" json:"facts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEntityReply) Reset()         { *m = GetEntityReply{} }
func (m *GetEntityReply) String() string { return proto.CompactTextString(m) }
func (*GetEntityReply) ProtoMessage()    {}
func (*GetEntityReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_70be9028e322f9d8, []int{1}
}

func (m *GetEntityReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEntityReply.Unmarshal(m, b)
}
func (m *GetEntityReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEntityReply.Marshal(b, m, deterministic)
}
func (m *GetEntityReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEntityReply.Merge(m, src)
}
func (m *GetEntityReply) XXX_Size() int {
	return xxx_messageInfo_GetEntityReply.Size(m)
}
func (m *GetEntityReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEntityReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetEntityReply proto.InternalMessageInfo

func (m *GetEntityReply) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *GetEntityReply) GetFacts() []*Fact {
	if m != nil {
		return m.Facts
	}
	return nil
}

type CreateEntityRequest struct {
	Entity               *Entity  `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEntityRequest) Reset()         { *m = CreateEntityRequest{} }
func (m *CreateEntityRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEntityRequest) ProtoMessage()    {}
func (*CreateEntityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_70be9028e322f9d8, []int{2}
}

func (m *CreateEntityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEntityRequest.Unmarshal(m, b)
}
func (m *CreateEntityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEntityRequest.Marshal(b, m, deterministic)
}
func (m *CreateEntityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEntityRequest.Merge(m, src)
}
func (m *CreateEntityRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEntityRequest.Size(m)
}
func (m *CreateEntityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEntityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEntityRequest proto.InternalMessageInfo

func (m *CreateEntityRequest) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

type CreateEntityReply struct {
	Entity               *Entity  `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEntityReply) Reset()         { *m = CreateEntityReply{} }
func (m *CreateEntityReply) String() string { return proto.CompactTextString(m) }
func (*CreateEntityReply) ProtoMessage()    {}
func (*CreateEntityReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_70be9028e322f9d8, []int{3}
}

func (m *CreateEntityReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEntityReply.Unmarshal(m, b)
}
func (m *CreateEntityReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEntityReply.Marshal(b, m, deterministic)
}
func (m *CreateEntityReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEntityReply.Merge(m, src)
}
func (m *CreateEntityReply) XXX_Size() int {
	return xxx_messageInfo_CreateEntityReply.Size(m)
}
func (m *CreateEntityReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEntityReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEntityReply proto.InternalMessageInfo

func (m *CreateEntityReply) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

type SetFactRequest struct {
	EntityUri            string   `protobuf:"bytes,1,opt,name=entityUri,proto3" json:"entityUri,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetFactRequest) Reset()         { *m = SetFactRequest{} }
func (m *SetFactRequest) String() string { return proto.CompactTextString(m) }
func (*SetFactRequest) ProtoMessage()    {}
func (*SetFactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_70be9028e322f9d8, []int{4}
}

func (m *SetFactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetFactRequest.Unmarshal(m, b)
}
func (m *SetFactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetFactRequest.Marshal(b, m, deterministic)
}
func (m *SetFactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetFactRequest.Merge(m, src)
}
func (m *SetFactRequest) XXX_Size() int {
	return xxx_messageInfo_SetFactRequest.Size(m)
}
func (m *SetFactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetFactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetFactRequest proto.InternalMessageInfo

func (m *SetFactRequest) GetEntityUri() string {
	if m != nil {
		return m.EntityUri
	}
	return ""
}

func (m *SetFactRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SetFactRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SetFactReply struct {
	Fact                 *Fact    `protobuf:"bytes,1,opt,name=fact,proto3" json:"fact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetFactReply) Reset()         { *m = SetFactReply{} }
func (m *SetFactReply) String() string { return proto.CompactTextString(m) }
func (*SetFactReply) ProtoMessage()    {}
func (*SetFactReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_70be9028e322f9d8, []int{5}
}

func (m *SetFactReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetFactReply.Unmarshal(m, b)
}
func (m *SetFactReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetFactReply.Marshal(b, m, deterministic)
}
func (m *SetFactReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetFactReply.Merge(m, src)
}
func (m *SetFactReply) XXX_Size() int {
	return xxx_messageInfo_SetFactReply.Size(m)
}
func (m *SetFactReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SetFactReply.DiscardUnknown(m)
}

var xxx_messageInfo_SetFactReply proto.InternalMessageInfo

func (m *SetFactReply) GetFact() *Fact {
	if m != nil {
		return m.Fact
	}
	return nil
}

type GetFactRequest struct {
	EntityUri            string   `protobuf:"bytes,1,opt,name=entityUri,proto3" json:"entityUri,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFactRequest) Reset()         { *m = GetFactRequest{} }
func (m *GetFactRequest) String() string { return proto.CompactTextString(m) }
func (*GetFactRequest) ProtoMessage()    {}
func (*GetFactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_70be9028e322f9d8, []int{6}
}

func (m *GetFactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFactRequest.Unmarshal(m, b)
}
func (m *GetFactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFactRequest.Marshal(b, m, deterministic)
}
func (m *GetFactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFactRequest.Merge(m, src)
}
func (m *GetFactRequest) XXX_Size() int {
	return xxx_messageInfo_GetFactRequest.Size(m)
}
func (m *GetFactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFactRequest proto.InternalMessageInfo

func (m *GetFactRequest) GetEntityUri() string {
	if m != nil {
		return m.EntityUri
	}
	return ""
}

func (m *GetFactRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetFactReply struct {
	Fact                 *Fact    `protobuf:"bytes,1,opt,name=fact,proto3" json:"fact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFactReply) Reset()         { *m = GetFactReply{} }
func (m *GetFactReply) String() string { return proto.CompactTextString(m) }
func (*GetFactReply) ProtoMessage()    {}
func (*GetFactReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_70be9028e322f9d8, []int{7}
}

func (m *GetFactReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFactReply.Unmarshal(m, b)
}
func (m *GetFactReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFactReply.Marshal(b, m, deterministic)
}
func (m *GetFactReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFactReply.Merge(m, src)
}
func (m *GetFactReply) XXX_Size() int {
	return xxx_messageInfo_GetFactReply.Size(m)
}
func (m *GetFactReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFactReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetFactReply proto.InternalMessageInfo

func (m *GetFactReply) GetFact() *Fact {
	if m != nil {
		return m.Fact
	}
	return nil
}

type Entity struct {
	Uri                  string   `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_70be9028e322f9d8, []int{8}
}

func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

type Fact struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Fact) Reset()         { *m = Fact{} }
func (m *Fact) String() string { return proto.CompactTextString(m) }
func (*Fact) ProtoMessage()    {}
func (*Fact) Descriptor() ([]byte, []int) {
	return fileDescriptor_70be9028e322f9d8, []int{9}
}

func (m *Fact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fact.Unmarshal(m, b)
}
func (m *Fact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fact.Marshal(b, m, deterministic)
}
func (m *Fact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fact.Merge(m, src)
}
func (m *Fact) XXX_Size() int {
	return xxx_messageInfo_Fact.Size(m)
}
func (m *Fact) XXX_DiscardUnknown() {
	xxx_messageInfo_Fact.DiscardUnknown(m)
}

var xxx_messageInfo_Fact proto.InternalMessageInfo

func (m *Fact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Fact) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*GetEntityRequest)(nil), "spotify.backstage.inventory.v1.GetEntityRequest")
	proto.RegisterType((*GetEntityReply)(nil), "spotify.backstage.inventory.v1.GetEntityReply")
	proto.RegisterType((*CreateEntityRequest)(nil), "spotify.backstage.inventory.v1.CreateEntityRequest")
	proto.RegisterType((*CreateEntityReply)(nil), "spotify.backstage.inventory.v1.CreateEntityReply")
	proto.RegisterType((*SetFactRequest)(nil), "spotify.backstage.inventory.v1.SetFactRequest")
	proto.RegisterType((*SetFactReply)(nil), "spotify.backstage.inventory.v1.SetFactReply")
	proto.RegisterType((*GetFactRequest)(nil), "spotify.backstage.inventory.v1.GetFactRequest")
	proto.RegisterType((*GetFactReply)(nil), "spotify.backstage.inventory.v1.GetFactReply")
	proto.RegisterType((*Entity)(nil), "spotify.backstage.inventory.v1.Entity")
	proto.RegisterType((*Fact)(nil), "spotify.backstage.inventory.v1.Fact")
}

func init() { proto.RegisterFile("inventory/v1/inventory.proto", fileDescriptor_70be9028e322f9d8) }

var fileDescriptor_70be9028e322f9d8 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0x4d, 0x29, 0x60, 0x3a, 0x7c, 0x04, 0x57, 0x0f, 0x4d, 0x43, 0x0c, 0x59, 0x8d, 0xe1, 0x60,
	0x0a, 0x85, 0x8b, 0xf1, 0xe0, 0x01, 0xa3, 0xd5, 0x6b, 0x09, 0x89, 0xf1, 0x62, 0x0a, 0x2e, 0xa4,
	0x11, 0x5a, 0x6c, 0xb7, 0x35, 0xfd, 0x0f, 0xfe, 0x2c, 0x7f, 0x98, 0xd9, 0xed, 0x5a, 0x8a, 0x21,
	0x16, 0x02, 0xb7, 0xdd, 0x99, 0x79, 0xf3, 0x5e, 0x5f, 0x5f, 0x0b, 0x4d, 0xc7, 0x8d, 0x88, 0x4b,
	0x3d, 0x3f, 0xee, 0x44, 0x46, 0x27, 0xbd, 0xe8, 0x4b, 0xdf, 0xa3, 0x1e, 0x3a, 0x0b, 0x96, 0x1e,
	0x75, 0xa6, 0xb1, 0x3e, 0xb6, 0x27, 0xef, 0x01, 0xb5, 0x67, 0x44, 0x5f, 0x8d, 0x44, 0x06, 0xfe,
	0x84, 0x86, 0x49, 0xe8, 0xbd, 0x4b, 0x1d, 0x1a, 0x5b, 0xe4, 0x23, 0x24, 0x01, 0x45, 0xb7, 0x50,
	0x26, 0xbc, 0xa0, 0x4a, 0x2d, 0xa9, 0x5d, 0xe9, 0x5d, 0xea, 0xff, 0x2f, 0xd1, 0x05, 0x5c, 0xa0,
	0xd0, 0x39, 0xd4, 0x1c, 0x77, 0x32, 0x0f, 0xdf, 0xc8, 0xeb, 0xd4, 0x9e, 0xd0, 0x40, 0x2d, 0xb4,
	0xe4, 0xb6, 0x62, 0x55, 0x45, 0xf1, 0x81, 0xd5, 0xf0, 0x97, 0x04, 0xf5, 0x0c, 0xf3, 0x72, 0x1e,
	0xef, 0xcd, 0x7b, 0x03, 0xa5, 0x15, 0x5f, 0xa5, 0x77, 0x91, 0x07, 0x67, 0x42, 0xac, 0x04, 0x82,
	0x47, 0x70, 0x72, 0xe7, 0x13, 0x9b, 0x92, 0x83, 0x5a, 0x81, 0x87, 0x70, 0xbc, 0xbe, 0xf6, 0x00,
	0xcf, 0x89, 0x9f, 0xa1, 0x3e, 0x24, 0x94, 0xab, 0x17, 0x32, 0x9b, 0xa0, 0x24, 0xbd, 0x91, 0xef,
	0xf0, 0xa5, 0x8a, 0xb5, 0x2a, 0x20, 0x04, 0x45, 0xd7, 0x5e, 0x10, 0xb5, 0xc0, 0x1b, 0xfc, 0x8c,
	0x4e, 0xa1, 0x14, 0xd9, 0xf3, 0x90, 0xa8, 0x32, 0x2f, 0x26, 0x17, 0xfc, 0x08, 0xd5, 0x74, 0x33,
	0x53, 0x7a, 0x0d, 0x45, 0x66, 0x8f, 0xd0, 0xb9, 0x9d, 0xa1, 0x1c, 0x81, 0x07, 0xfc, 0xed, 0xee,
	0xa5, 0x91, 0xa9, 0x31, 0x0f, 0xa3, 0x46, 0x83, 0x72, 0xe2, 0x21, 0x6a, 0x80, 0x1c, 0xa6, 0xfc,
	0xec, 0x88, 0xbb, 0x50, 0x64, 0x93, 0xa9, 0x02, 0x69, 0x93, 0x4b, 0x85, 0x8c, 0x4b, 0xbd, 0x6f,
	0x19, 0x94, 0xa7, 0x5f, 0x26, 0xb4, 0x00, 0x25, 0xcd, 0x31, 0xea, 0xe6, 0x89, 0xfa, 0xfb, 0xb1,
	0x69, 0xfa, 0x0e, 0x08, 0x66, 0x42, 0x04, 0xd5, 0x6c, 0xa2, 0x50, 0x3f, 0x0f, 0xbf, 0x21, 0xd6,
	0x9a, 0xb1, 0x1b, 0x88, 0xf1, 0xce, 0xe0, 0x48, 0x44, 0x03, 0xe5, 0x4a, 0x5e, 0x4f, 0xa7, 0x76,
	0xb5, 0xf5, 0xbc, 0x20, 0x32, 0xb7, 0x25, 0x32, 0x77, 0x24, 0xca, 0xc6, 0x69, 0x50, 0x7b, 0xa9,
	0xa4, 0xcd, 0xc8, 0x18, 0x97, 0xf9, 0x0f, 0xb3, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x4f, 0xab,
	0xd5, 0x54, 0x50, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// InventoryClient is the client API for Inventory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InventoryClient interface {
	GetEntity(ctx context.Context, in *GetEntityRequest, opts ...grpc.CallOption) (*GetEntityReply, error)
	CreateEntity(ctx context.Context, in *CreateEntityRequest, opts ...grpc.CallOption) (*CreateEntityReply, error)
	SetFact(ctx context.Context, in *SetFactRequest, opts ...grpc.CallOption) (*SetFactReply, error)
	GetFact(ctx context.Context, in *GetFactRequest, opts ...grpc.CallOption) (*GetFactReply, error)
}

type inventoryClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryClient(cc grpc.ClientConnInterface) InventoryClient {
	return &inventoryClient{cc}
}

func (c *inventoryClient) GetEntity(ctx context.Context, in *GetEntityRequest, opts ...grpc.CallOption) (*GetEntityReply, error) {
	out := new(GetEntityReply)
	err := c.cc.Invoke(ctx, "/spotify.backstage.inventory.v1.Inventory/GetEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) CreateEntity(ctx context.Context, in *CreateEntityRequest, opts ...grpc.CallOption) (*CreateEntityReply, error) {
	out := new(CreateEntityReply)
	err := c.cc.Invoke(ctx, "/spotify.backstage.inventory.v1.Inventory/CreateEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) SetFact(ctx context.Context, in *SetFactRequest, opts ...grpc.CallOption) (*SetFactReply, error) {
	out := new(SetFactReply)
	err := c.cc.Invoke(ctx, "/spotify.backstage.inventory.v1.Inventory/SetFact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) GetFact(ctx context.Context, in *GetFactRequest, opts ...grpc.CallOption) (*GetFactReply, error) {
	out := new(GetFactReply)
	err := c.cc.Invoke(ctx, "/spotify.backstage.inventory.v1.Inventory/GetFact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryServer is the server API for Inventory service.
type InventoryServer interface {
	GetEntity(context.Context, *GetEntityRequest) (*GetEntityReply, error)
	CreateEntity(context.Context, *CreateEntityRequest) (*CreateEntityReply, error)
	SetFact(context.Context, *SetFactRequest) (*SetFactReply, error)
	GetFact(context.Context, *GetFactRequest) (*GetFactReply, error)
}

// UnimplementedInventoryServer can be embedded to have forward compatible implementations.
type UnimplementedInventoryServer struct {
}

func (*UnimplementedInventoryServer) GetEntity(ctx context.Context, req *GetEntityRequest) (*GetEntityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEntity not implemented")
}
func (*UnimplementedInventoryServer) CreateEntity(ctx context.Context, req *CreateEntityRequest) (*CreateEntityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEntity not implemented")
}
func (*UnimplementedInventoryServer) SetFact(ctx context.Context, req *SetFactRequest) (*SetFactReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFact not implemented")
}
func (*UnimplementedInventoryServer) GetFact(ctx context.Context, req *GetFactRequest) (*GetFactReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFact not implemented")
}

func RegisterInventoryServer(s *grpc.Server, srv InventoryServer) {
	s.RegisterService(&_Inventory_serviceDesc, srv)
}

func _Inventory_GetEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).GetEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotify.backstage.inventory.v1.Inventory/GetEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).GetEntity(ctx, req.(*GetEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_CreateEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).CreateEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotify.backstage.inventory.v1.Inventory/CreateEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).CreateEntity(ctx, req.(*CreateEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_SetFact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).SetFact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotify.backstage.inventory.v1.Inventory/SetFact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).SetFact(ctx, req.(*SetFactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_GetFact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).GetFact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotify.backstage.inventory.v1.Inventory/GetFact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).GetFact(ctx, req.(*GetFactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Inventory_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spotify.backstage.inventory.v1.Inventory",
	HandlerType: (*InventoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEntity",
			Handler:    _Inventory_GetEntity_Handler,
		},
		{
			MethodName: "CreateEntity",
			Handler:    _Inventory_CreateEntity_Handler,
		},
		{
			MethodName: "SetFact",
			Handler:    _Inventory_SetFact_Handler,
		},
		{
			MethodName: "GetFact",
			Handler:    _Inventory_GetFact_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventory/v1/inventory.proto",
}