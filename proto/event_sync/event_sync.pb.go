// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/event_sync/event_sync.proto

// import "proto/include/google/protobuf/any.proto";

package event_sync

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

type FindRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindRequest) Reset()         { *m = FindRequest{} }
func (m *FindRequest) String() string { return proto.CompactTextString(m) }
func (*FindRequest) ProtoMessage()    {}
func (*FindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cff40cac977f7df, []int{0}
}

func (m *FindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindRequest.Unmarshal(m, b)
}
func (m *FindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindRequest.Marshal(b, m, deterministic)
}
func (m *FindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRequest.Merge(m, src)
}
func (m *FindRequest) XXX_Size() int {
	return xxx_messageInfo_FindRequest.Size(m)
}
func (m *FindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindRequest proto.InternalMessageInfo

func (m *FindRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type GetRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Query                string   `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cff40cac977f7df, []int{1}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type CreateEventSyncRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEventSyncRequest) Reset()         { *m = CreateEventSyncRequest{} }
func (m *CreateEventSyncRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEventSyncRequest) ProtoMessage()    {}
func (*CreateEventSyncRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cff40cac977f7df, []int{2}
}

func (m *CreateEventSyncRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEventSyncRequest.Unmarshal(m, b)
}
func (m *CreateEventSyncRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEventSyncRequest.Marshal(b, m, deterministic)
}
func (m *CreateEventSyncRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEventSyncRequest.Merge(m, src)
}
func (m *CreateEventSyncRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEventSyncRequest.Size(m)
}
func (m *CreateEventSyncRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEventSyncRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEventSyncRequest proto.InternalMessageInfo

func (m *CreateEventSyncRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UpdateEventSyncRequest struct {
	Query                string                  `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Obj                  *CreateEventSyncRequest `protobuf:"bytes,2,opt,name=obj,proto3" json:"obj,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *UpdateEventSyncRequest) Reset()         { *m = UpdateEventSyncRequest{} }
func (m *UpdateEventSyncRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateEventSyncRequest) ProtoMessage()    {}
func (*UpdateEventSyncRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cff40cac977f7df, []int{3}
}

func (m *UpdateEventSyncRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEventSyncRequest.Unmarshal(m, b)
}
func (m *UpdateEventSyncRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEventSyncRequest.Marshal(b, m, deterministic)
}
func (m *UpdateEventSyncRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEventSyncRequest.Merge(m, src)
}
func (m *UpdateEventSyncRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateEventSyncRequest.Size(m)
}
func (m *UpdateEventSyncRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEventSyncRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEventSyncRequest proto.InternalMessageInfo

func (m *UpdateEventSyncRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *UpdateEventSyncRequest) GetObj() *CreateEventSyncRequest {
	if m != nil {
		return m.Obj
	}
	return nil
}

type UpdateEventSyncByIdRequest struct {
	Id                   string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Obj                  *CreateEventSyncRequest `protobuf:"bytes,2,opt,name=obj,proto3" json:"obj,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *UpdateEventSyncByIdRequest) Reset()         { *m = UpdateEventSyncByIdRequest{} }
func (m *UpdateEventSyncByIdRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateEventSyncByIdRequest) ProtoMessage()    {}
func (*UpdateEventSyncByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cff40cac977f7df, []int{4}
}

func (m *UpdateEventSyncByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEventSyncByIdRequest.Unmarshal(m, b)
}
func (m *UpdateEventSyncByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEventSyncByIdRequest.Marshal(b, m, deterministic)
}
func (m *UpdateEventSyncByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEventSyncByIdRequest.Merge(m, src)
}
func (m *UpdateEventSyncByIdRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateEventSyncByIdRequest.Size(m)
}
func (m *UpdateEventSyncByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEventSyncByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEventSyncByIdRequest proto.InternalMessageInfo

func (m *UpdateEventSyncByIdRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateEventSyncByIdRequest) GetObj() *CreateEventSyncRequest {
	if m != nil {
		return m.Obj
	}
	return nil
}

type DeleteResult struct {
	State                bool     `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
	Num                  int32    `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResult) Reset()         { *m = DeleteResult{} }
func (m *DeleteResult) String() string { return proto.CompactTextString(m) }
func (*DeleteResult) ProtoMessage()    {}
func (*DeleteResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cff40cac977f7df, []int{5}
}

func (m *DeleteResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResult.Unmarshal(m, b)
}
func (m *DeleteResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResult.Marshal(b, m, deterministic)
}
func (m *DeleteResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResult.Merge(m, src)
}
func (m *DeleteResult) XXX_Size() int {
	return xxx_messageInfo_DeleteResult.Size(m)
}
func (m *DeleteResult) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResult.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResult proto.InternalMessageInfo

func (m *DeleteResult) GetState() bool {
	if m != nil {
		return m.State
	}
	return false
}

func (m *DeleteResult) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type EventSync struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt            string   `protobuf:"bytes,13,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,14,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt            string   `protobuf:"bytes,15,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventSync) Reset()         { *m = EventSync{} }
func (m *EventSync) String() string { return proto.CompactTextString(m) }
func (*EventSync) ProtoMessage()    {}
func (*EventSync) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cff40cac977f7df, []int{6}
}

func (m *EventSync) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventSync.Unmarshal(m, b)
}
func (m *EventSync) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventSync.Marshal(b, m, deterministic)
}
func (m *EventSync) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSync.Merge(m, src)
}
func (m *EventSync) XXX_Size() int {
	return xxx_messageInfo_EventSync.Size(m)
}
func (m *EventSync) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSync.DiscardUnknown(m)
}

var xxx_messageInfo_EventSync proto.InternalMessageInfo

func (m *EventSync) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EventSync) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EventSync) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *EventSync) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *EventSync) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

type FindEventSyncResult struct {
	Data                 []*EventSync `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Limit                int32        `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Skip                 int64        `protobuf:"varint,3,opt,name=skip,proto3" json:"skip,omitempty"`
	Total                int64        `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FindEventSyncResult) Reset()         { *m = FindEventSyncResult{} }
func (m *FindEventSyncResult) String() string { return proto.CompactTextString(m) }
func (*FindEventSyncResult) ProtoMessage()    {}
func (*FindEventSyncResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cff40cac977f7df, []int{7}
}

func (m *FindEventSyncResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindEventSyncResult.Unmarshal(m, b)
}
func (m *FindEventSyncResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindEventSyncResult.Marshal(b, m, deterministic)
}
func (m *FindEventSyncResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindEventSyncResult.Merge(m, src)
}
func (m *FindEventSyncResult) XXX_Size() int {
	return xxx_messageInfo_FindEventSyncResult.Size(m)
}
func (m *FindEventSyncResult) XXX_DiscardUnknown() {
	xxx_messageInfo_FindEventSyncResult.DiscardUnknown(m)
}

var xxx_messageInfo_FindEventSyncResult proto.InternalMessageInfo

func (m *FindEventSyncResult) GetData() []*EventSync {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *FindEventSyncResult) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *FindEventSyncResult) GetSkip() int64 {
	if m != nil {
		return m.Skip
	}
	return 0
}

func (m *FindEventSyncResult) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*FindRequest)(nil), "event_sync.FindRequest")
	proto.RegisterType((*GetRequest)(nil), "event_sync.GetRequest")
	proto.RegisterType((*CreateEventSyncRequest)(nil), "event_sync.CreateEventSyncRequest")
	proto.RegisterType((*UpdateEventSyncRequest)(nil), "event_sync.UpdateEventSyncRequest")
	proto.RegisterType((*UpdateEventSyncByIdRequest)(nil), "event_sync.UpdateEventSyncByIdRequest")
	proto.RegisterType((*DeleteResult)(nil), "event_sync.DeleteResult")
	proto.RegisterType((*EventSync)(nil), "event_sync.EventSync")
	proto.RegisterType((*FindEventSyncResult)(nil), "event_sync.FindEventSyncResult")
}

func init() { proto.RegisterFile("proto/event_sync/event_sync.proto", fileDescriptor_5cff40cac977f7df) }

var fileDescriptor_5cff40cac977f7df = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x6d, 0x3e, 0x56, 0xec, 0xad, 0x6e, 0xeb, 0x54, 0x6b, 0x28, 0x82, 0x75, 0x04, 0xa9, 0x20,
	0x2b, 0x44, 0xf1, 0x5d, 0x57, 0x77, 0x59, 0x10, 0x85, 0x14, 0x7d, 0x95, 0x34, 0x73, 0x1f, 0x46,
	0xd3, 0xa4, 0x4d, 0x26, 0x85, 0xbc, 0xf9, 0xea, 0xef, 0xf1, 0x0f, 0xca, 0xcc, 0xb4, 0xc9, 0x34,
	0x26, 0x41, 0x7c, 0x9b, 0xb9, 0xe7, 0xce, 0xb9, 0xe7, 0x4c, 0xce, 0x04, 0x9e, 0x6c, 0xb3, 0x54,
	0xa4, 0x2f, 0x71, 0x8f, 0x89, 0xf8, 0x96, 0x97, 0x49, 0x64, 0x2c, 0x2f, 0x14, 0x46, 0xa0, 0xae,
	0xd0, 0xa7, 0x30, 0xba, 0xe2, 0x09, 0x0b, 0x70, 0x57, 0x60, 0x2e, 0xc8, 0x7d, 0x38, 0xdb, 0x15,
	0x98, 0x95, 0x9e, 0xb5, 0xb0, 0x96, 0xc3, 0x40, 0x6f, 0xa8, 0x0f, 0x70, 0x8d, 0xe2, 0xd8, 0x73,
	0x0e, 0x36, 0x67, 0x87, 0x06, 0x9b, 0xb3, 0xfa, 0x8c, 0x6d, 0x9e, 0x79, 0x01, 0xb3, 0xcb, 0x0c,
	0x43, 0x81, 0x1f, 0xe4, 0xb0, 0x55, 0x99, 0x44, 0xc7, 0xf3, 0x04, 0xdc, 0x24, 0xdc, 0xe0, 0x81,
	0x41, 0xad, 0x29, 0x83, 0xd9, 0x97, 0x2d, 0x6b, 0xeb, 0x6e, 0x55, 0x44, 0x5e, 0x83, 0x93, 0xae,
	0xbf, 0xab, 0x89, 0x23, 0x9f, 0x5e, 0x18, 0x16, 0xdb, 0x87, 0x06, 0xb2, 0x9d, 0xae, 0x61, 0xde,
	0x98, 0xf2, 0xae, 0xbc, 0x61, 0x5d, 0xbe, 0xfe, 0x6f, 0xc6, 0x1b, 0xb8, 0xf3, 0x1e, 0x63, 0x14,
	0x18, 0x60, 0x5e, 0xc4, 0x4a, 0x7f, 0x2e, 0x42, 0xa1, 0xed, 0xde, 0x0e, 0xf4, 0x86, 0x4c, 0xc0,
	0x49, 0x8a, 0x8d, 0xe2, 0x3e, 0x0b, 0xe4, 0x92, 0xfe, 0xb2, 0x60, 0x58, 0x31, 0xfe, 0xa5, 0xe5,
	0x78, 0x67, 0x76, 0x7d, 0x67, 0xe4, 0x11, 0x0c, 0x23, 0x25, 0x84, 0xbd, 0x15, 0xde, 0x5d, 0x05,
	0xd4, 0x05, 0x89, 0x16, 0xca, 0xab, 0x44, 0xcf, 0x35, 0x5a, 0x15, 0x24, 0xca, 0x94, 0x4a, 0x89,
	0x8e, 0x35, 0x5a, 0x15, 0xe8, 0x4f, 0x0b, 0xa6, 0x32, 0x15, 0x86, 0x43, 0xe5, 0xe5, 0x39, 0xb8,
	0x2c, 0x14, 0xa1, 0x67, 0x2d, 0x9c, 0xe5, 0xc8, 0x7f, 0x60, 0x5e, 0x49, 0xdd, 0xaa, 0x5a, 0xa4,
	0xed, 0x98, 0x6f, 0xb8, 0x38, 0x58, 0xd4, 0x1b, 0x69, 0x23, 0xff, 0xc1, 0xb7, 0x9e, 0xb3, 0xb0,
	0x96, 0x4e, 0xa0, 0xd6, 0xb2, 0x53, 0xa4, 0x22, 0x8c, 0x3d, 0x57, 0x15, 0xf5, 0xc6, 0xff, 0xed,
	0xc2, 0xa4, 0xe2, 0x5c, 0x61, 0xb6, 0xe7, 0x11, 0x92, 0xcf, 0x70, 0xef, 0x44, 0xd6, 0x47, 0x9e,
	0x0b, 0xf2, 0xd0, 0x94, 0x61, 0x64, 0x79, 0xfe, 0xb8, 0x09, 0x34, 0xec, 0xd0, 0x01, 0xb9, 0x84,
	0xc9, 0x35, 0x8a, 0x93, 0x34, 0x90, 0x99, 0x79, 0xac, 0x8e, 0xfd, 0xbc, 0xdd, 0x2e, 0x1d, 0x90,
	0x4f, 0x30, 0x6e, 0x04, 0x82, 0xfc, 0x43, 0x5a, 0x7a, 0xf9, 0x1a, 0x29, 0x3d, 0xe5, 0x6b, 0x7f,
	0x28, 0xdd, 0x7c, 0x5f, 0x61, 0xda, 0x92, 0x7a, 0xf2, 0xac, 0x87, 0xd3, 0x78, 0x16, 0xdd, 0xbc,
	0x57, 0x30, 0xd6, 0x49, 0xaf, 0x75, 0x76, 0x7e, 0x0b, 0xcf, 0x04, 0xcc, 0xf7, 0x41, 0x07, 0xe4,
	0x06, 0xa6, 0x0d, 0x9e, 0xde, 0xef, 0xd0, 0x43, 0xb5, 0xbe, 0xa5, 0x7e, 0x70, 0xaf, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xb3, 0xc5, 0x56, 0x1a, 0x05, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventSyncServiceClient is the client API for EventSyncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventSyncServiceClient interface {
	// your model Requests
	FindEventSyncList(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindEventSyncResult, error)
	GetEventSyncById(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*EventSync, error)
	CreateEventSync(ctx context.Context, in *CreateEventSyncRequest, opts ...grpc.CallOption) (*EventSync, error)
	UpdateEventSync(ctx context.Context, in *UpdateEventSyncRequest, opts ...grpc.CallOption) (*EventSync, error)
	UpdateEventSyncById(ctx context.Context, in *UpdateEventSyncByIdRequest, opts ...grpc.CallOption) (*EventSync, error)
	DeleteEventSync(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*DeleteResult, error)
	DeleteEventSyncById(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*DeleteResult, error)
}

type eventSyncServiceClient struct {
	cc *grpc.ClientConn
}

func NewEventSyncServiceClient(cc *grpc.ClientConn) EventSyncServiceClient {
	return &eventSyncServiceClient{cc}
}

func (c *eventSyncServiceClient) FindEventSyncList(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindEventSyncResult, error) {
	out := new(FindEventSyncResult)
	err := c.cc.Invoke(ctx, "/event_sync.EventSyncService/FindEventSyncList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventSyncServiceClient) GetEventSyncById(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*EventSync, error) {
	out := new(EventSync)
	err := c.cc.Invoke(ctx, "/event_sync.EventSyncService/GetEventSyncById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventSyncServiceClient) CreateEventSync(ctx context.Context, in *CreateEventSyncRequest, opts ...grpc.CallOption) (*EventSync, error) {
	out := new(EventSync)
	err := c.cc.Invoke(ctx, "/event_sync.EventSyncService/CreateEventSync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventSyncServiceClient) UpdateEventSync(ctx context.Context, in *UpdateEventSyncRequest, opts ...grpc.CallOption) (*EventSync, error) {
	out := new(EventSync)
	err := c.cc.Invoke(ctx, "/event_sync.EventSyncService/UpdateEventSync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventSyncServiceClient) UpdateEventSyncById(ctx context.Context, in *UpdateEventSyncByIdRequest, opts ...grpc.CallOption) (*EventSync, error) {
	out := new(EventSync)
	err := c.cc.Invoke(ctx, "/event_sync.EventSyncService/UpdateEventSyncById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventSyncServiceClient) DeleteEventSync(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*DeleteResult, error) {
	out := new(DeleteResult)
	err := c.cc.Invoke(ctx, "/event_sync.EventSyncService/DeleteEventSync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventSyncServiceClient) DeleteEventSyncById(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*DeleteResult, error) {
	out := new(DeleteResult)
	err := c.cc.Invoke(ctx, "/event_sync.EventSyncService/DeleteEventSyncById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventSyncServiceServer is the server API for EventSyncService service.
type EventSyncServiceServer interface {
	// your model Requests
	FindEventSyncList(context.Context, *FindRequest) (*FindEventSyncResult, error)
	GetEventSyncById(context.Context, *GetRequest) (*EventSync, error)
	CreateEventSync(context.Context, *CreateEventSyncRequest) (*EventSync, error)
	UpdateEventSync(context.Context, *UpdateEventSyncRequest) (*EventSync, error)
	UpdateEventSyncById(context.Context, *UpdateEventSyncByIdRequest) (*EventSync, error)
	DeleteEventSync(context.Context, *FindRequest) (*DeleteResult, error)
	DeleteEventSyncById(context.Context, *GetRequest) (*DeleteResult, error)
}

// UnimplementedEventSyncServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEventSyncServiceServer struct {
}

func (*UnimplementedEventSyncServiceServer) FindEventSyncList(ctx context.Context, req *FindRequest) (*FindEventSyncResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEventSyncList not implemented")
}
func (*UnimplementedEventSyncServiceServer) GetEventSyncById(ctx context.Context, req *GetRequest) (*EventSync, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventSyncById not implemented")
}
func (*UnimplementedEventSyncServiceServer) CreateEventSync(ctx context.Context, req *CreateEventSyncRequest) (*EventSync, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEventSync not implemented")
}
func (*UnimplementedEventSyncServiceServer) UpdateEventSync(ctx context.Context, req *UpdateEventSyncRequest) (*EventSync, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEventSync not implemented")
}
func (*UnimplementedEventSyncServiceServer) UpdateEventSyncById(ctx context.Context, req *UpdateEventSyncByIdRequest) (*EventSync, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEventSyncById not implemented")
}
func (*UnimplementedEventSyncServiceServer) DeleteEventSync(ctx context.Context, req *FindRequest) (*DeleteResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEventSync not implemented")
}
func (*UnimplementedEventSyncServiceServer) DeleteEventSyncById(ctx context.Context, req *GetRequest) (*DeleteResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEventSyncById not implemented")
}

func RegisterEventSyncServiceServer(s *grpc.Server, srv EventSyncServiceServer) {
	s.RegisterService(&_EventSyncService_serviceDesc, srv)
}

func _EventSyncService_FindEventSyncList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventSyncServiceServer).FindEventSyncList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event_sync.EventSyncService/FindEventSyncList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventSyncServiceServer).FindEventSyncList(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventSyncService_GetEventSyncById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventSyncServiceServer).GetEventSyncById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event_sync.EventSyncService/GetEventSyncById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventSyncServiceServer).GetEventSyncById(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventSyncService_CreateEventSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventSyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventSyncServiceServer).CreateEventSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event_sync.EventSyncService/CreateEventSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventSyncServiceServer).CreateEventSync(ctx, req.(*CreateEventSyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventSyncService_UpdateEventSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventSyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventSyncServiceServer).UpdateEventSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event_sync.EventSyncService/UpdateEventSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventSyncServiceServer).UpdateEventSync(ctx, req.(*UpdateEventSyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventSyncService_UpdateEventSyncById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventSyncByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventSyncServiceServer).UpdateEventSyncById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event_sync.EventSyncService/UpdateEventSyncById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventSyncServiceServer).UpdateEventSyncById(ctx, req.(*UpdateEventSyncByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventSyncService_DeleteEventSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventSyncServiceServer).DeleteEventSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event_sync.EventSyncService/DeleteEventSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventSyncServiceServer).DeleteEventSync(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventSyncService_DeleteEventSyncById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventSyncServiceServer).DeleteEventSyncById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event_sync.EventSyncService/DeleteEventSyncById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventSyncServiceServer).DeleteEventSyncById(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventSyncService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "event_sync.EventSyncService",
	HandlerType: (*EventSyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindEventSyncList",
			Handler:    _EventSyncService_FindEventSyncList_Handler,
		},
		{
			MethodName: "GetEventSyncById",
			Handler:    _EventSyncService_GetEventSyncById_Handler,
		},
		{
			MethodName: "CreateEventSync",
			Handler:    _EventSyncService_CreateEventSync_Handler,
		},
		{
			MethodName: "UpdateEventSync",
			Handler:    _EventSyncService_UpdateEventSync_Handler,
		},
		{
			MethodName: "UpdateEventSyncById",
			Handler:    _EventSyncService_UpdateEventSyncById_Handler,
		},
		{
			MethodName: "DeleteEventSync",
			Handler:    _EventSyncService_DeleteEventSync_Handler,
		},
		{
			MethodName: "DeleteEventSyncById",
			Handler:    _EventSyncService_DeleteEventSyncById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/event_sync/event_sync.proto",
}
