// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/youtube/service.proto

package youtube

import (
	rpc "github.com/rishusahu23/fam-youtube/gen/api/rpc"
	record "github.com/rishusahu23/fam-youtube/gen/api/youtube/record"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TriggerJobResponse_Status int32

const (
	TriggerJobResponse_OK               TriggerJobResponse_Status = 0
	TriggerJobResponse_INVALID_ARGUMENT TriggerJobResponse_Status = 3
	TriggerJobResponse_NO_RECORD_FOUND  TriggerJobResponse_Status = 5
	TriggerJobResponse_INTERNAL         TriggerJobResponse_Status = 13
)

// Enum value maps for TriggerJobResponse_Status.
var (
	TriggerJobResponse_Status_name = map[int32]string{
		0:  "OK",
		3:  "INVALID_ARGUMENT",
		5:  "NO_RECORD_FOUND",
		13: "INTERNAL",
	}
	TriggerJobResponse_Status_value = map[string]int32{
		"OK":               0,
		"INVALID_ARGUMENT": 3,
		"NO_RECORD_FOUND":  5,
		"INTERNAL":         13,
	}
)

func (x TriggerJobResponse_Status) Enum() *TriggerJobResponse_Status {
	p := new(TriggerJobResponse_Status)
	*p = x
	return p
}

func (x TriggerJobResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TriggerJobResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_api_youtube_service_proto_enumTypes[0].Descriptor()
}

func (TriggerJobResponse_Status) Type() protoreflect.EnumType {
	return &file_api_youtube_service_proto_enumTypes[0]
}

func (x TriggerJobResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TriggerJobResponse_Status.Descriptor instead.
func (TriggerJobResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_api_youtube_service_proto_rawDescGZIP(), []int{1, 0}
}

type GetPaginatedRecordsResponse_Status int32

const (
	GetPaginatedRecordsResponse_OK               GetPaginatedRecordsResponse_Status = 0
	GetPaginatedRecordsResponse_INVALID_ARGUMENT GetPaginatedRecordsResponse_Status = 3
	GetPaginatedRecordsResponse_NO_RECORD_FOUND  GetPaginatedRecordsResponse_Status = 5
	GetPaginatedRecordsResponse_INTERNAL         GetPaginatedRecordsResponse_Status = 13
)

// Enum value maps for GetPaginatedRecordsResponse_Status.
var (
	GetPaginatedRecordsResponse_Status_name = map[int32]string{
		0:  "OK",
		3:  "INVALID_ARGUMENT",
		5:  "NO_RECORD_FOUND",
		13: "INTERNAL",
	}
	GetPaginatedRecordsResponse_Status_value = map[string]int32{
		"OK":               0,
		"INVALID_ARGUMENT": 3,
		"NO_RECORD_FOUND":  5,
		"INTERNAL":         13,
	}
)

func (x GetPaginatedRecordsResponse_Status) Enum() *GetPaginatedRecordsResponse_Status {
	p := new(GetPaginatedRecordsResponse_Status)
	*p = x
	return p
}

func (x GetPaginatedRecordsResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetPaginatedRecordsResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_api_youtube_service_proto_enumTypes[1].Descriptor()
}

func (GetPaginatedRecordsResponse_Status) Type() protoreflect.EnumType {
	return &file_api_youtube_service_proto_enumTypes[1]
}

func (x GetPaginatedRecordsResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetPaginatedRecordsResponse_Status.Descriptor instead.
func (GetPaginatedRecordsResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_api_youtube_service_proto_rawDescGZIP(), []int{3, 0}
}

type TriggerJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interval int32 `protobuf:"varint,1,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *TriggerJobRequest) Reset() {
	*x = TriggerJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_youtube_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerJobRequest) ProtoMessage() {}

func (x *TriggerJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_youtube_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerJobRequest.ProtoReflect.Descriptor instead.
func (*TriggerJobRequest) Descriptor() ([]byte, []int) {
	return file_api_youtube_service_proto_rawDescGZIP(), []int{0}
}

func (x *TriggerJobRequest) GetInterval() int32 {
	if x != nil {
		return x.Interval
	}
	return 0
}

type TriggerJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *rpc.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *TriggerJobResponse) Reset() {
	*x = TriggerJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_youtube_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerJobResponse) ProtoMessage() {}

func (x *TriggerJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_youtube_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerJobResponse.ProtoReflect.Descriptor instead.
func (*TriggerJobResponse) Descriptor() ([]byte, []int) {
	return file_api_youtube_service_proto_rawDescGZIP(), []int{1}
}

func (x *TriggerJobResponse) GetStatus() *rpc.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type GetPaginatedRecordsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageContext *rpc.PageContextRequest `protobuf:"bytes,1,opt,name=page_context,json=pageContext,proto3" json:"page_context,omitempty"`
}

func (x *GetPaginatedRecordsRequest) Reset() {
	*x = GetPaginatedRecordsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_youtube_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaginatedRecordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaginatedRecordsRequest) ProtoMessage() {}

func (x *GetPaginatedRecordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_youtube_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaginatedRecordsRequest.ProtoReflect.Descriptor instead.
func (*GetPaginatedRecordsRequest) Descriptor() ([]byte, []int) {
	return file_api_youtube_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetPaginatedRecordsRequest) GetPageContext() *rpc.PageContextRequest {
	if x != nil {
		return x.PageContext
	}
	return nil
}

type GetPaginatedRecordsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      *rpc.Status              `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	PageContext *rpc.PageContextResponse `protobuf:"bytes,2,opt,name=page_context,json=pageContext,proto3" json:"page_context,omitempty"`
	Records     []*record.Record         `protobuf:"bytes,3,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *GetPaginatedRecordsResponse) Reset() {
	*x = GetPaginatedRecordsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_youtube_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaginatedRecordsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaginatedRecordsResponse) ProtoMessage() {}

func (x *GetPaginatedRecordsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_youtube_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaginatedRecordsResponse.ProtoReflect.Descriptor instead.
func (*GetPaginatedRecordsResponse) Descriptor() ([]byte, []int) {
	return file_api_youtube_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetPaginatedRecordsResponse) GetStatus() *rpc.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *GetPaginatedRecordsResponse) GetPageContext() *rpc.PageContextResponse {
	if x != nil {
		return x.PageContext
	}
	return nil
}

func (x *GetPaginatedRecordsResponse) GetRecords() []*record.Record {
	if x != nil {
		return x.Records
	}
	return nil
}

var File_api_youtube_service_proto protoreflect.FileDescriptor

var file_api_youtube_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x79, 0x6f, 0x75,
	0x74, 0x75, 0x62, 0x65, 0x1a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x61, 0x70, 0x69, 0x2f, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x2f, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a,
	0x11, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0x84,
	0x01, 0x0a, 0x12, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x49, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54,
	0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x4f, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f,
	0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x54, 0x45, 0x52,
	0x4e, 0x41, 0x4c, 0x10, 0x0d, 0x22, 0x58, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0c, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22,
	0xfc, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x23, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x3b, 0x0a, 0x0c, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x30, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x22, 0x49, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a,
	0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x4e,
	0x4f, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x05,
	0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x0d, 0x32, 0x84,
	0x02, 0x0a, 0x0e, 0x59, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x65, 0x0a, 0x0a, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x12,
	0x1a, 0x2e, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x79, 0x6f,
	0x75, 0x74, 0x75, 0x62, 0x65, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18,
	0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x2d, 0x6a, 0x6f, 0x62, 0x12, 0x8a, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x12, 0x23, 0x2e, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x65, 0x74, 0x2d, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x73, 0x68, 0x75, 0x73, 0x61, 0x68, 0x75, 0x32, 0x33, 0x2f,
	0x66, 0x61, 0x6d, 0x2d, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_youtube_service_proto_rawDescOnce sync.Once
	file_api_youtube_service_proto_rawDescData = file_api_youtube_service_proto_rawDesc
)

func file_api_youtube_service_proto_rawDescGZIP() []byte {
	file_api_youtube_service_proto_rawDescOnce.Do(func() {
		file_api_youtube_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_youtube_service_proto_rawDescData)
	})
	return file_api_youtube_service_proto_rawDescData
}

var file_api_youtube_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_youtube_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_youtube_service_proto_goTypes = []interface{}{
	(TriggerJobResponse_Status)(0),          // 0: youtube.TriggerJobResponse.Status
	(GetPaginatedRecordsResponse_Status)(0), // 1: youtube.GetPaginatedRecordsResponse.Status
	(*TriggerJobRequest)(nil),               // 2: youtube.TriggerJobRequest
	(*TriggerJobResponse)(nil),              // 3: youtube.TriggerJobResponse
	(*GetPaginatedRecordsRequest)(nil),      // 4: youtube.GetPaginatedRecordsRequest
	(*GetPaginatedRecordsResponse)(nil),     // 5: youtube.GetPaginatedRecordsResponse
	(*rpc.Status)(nil),                      // 6: rpc.Status
	(*rpc.PageContextRequest)(nil),          // 7: rpc.PageContextRequest
	(*rpc.PageContextResponse)(nil),         // 8: rpc.PageContextResponse
	(*record.Record)(nil),                   // 9: youtube.record.Record
}
var file_api_youtube_service_proto_depIdxs = []int32{
	6, // 0: youtube.TriggerJobResponse.status:type_name -> rpc.Status
	7, // 1: youtube.GetPaginatedRecordsRequest.page_context:type_name -> rpc.PageContextRequest
	6, // 2: youtube.GetPaginatedRecordsResponse.status:type_name -> rpc.Status
	8, // 3: youtube.GetPaginatedRecordsResponse.page_context:type_name -> rpc.PageContextResponse
	9, // 4: youtube.GetPaginatedRecordsResponse.records:type_name -> youtube.record.Record
	2, // 5: youtube.YoutubeService.TriggerJob:input_type -> youtube.TriggerJobRequest
	4, // 6: youtube.YoutubeService.GetPaginatedRecords:input_type -> youtube.GetPaginatedRecordsRequest
	3, // 7: youtube.YoutubeService.TriggerJob:output_type -> youtube.TriggerJobResponse
	5, // 8: youtube.YoutubeService.GetPaginatedRecords:output_type -> youtube.GetPaginatedRecordsResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_youtube_service_proto_init() }
func file_api_youtube_service_proto_init() {
	if File_api_youtube_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_youtube_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerJobRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_youtube_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerJobResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_youtube_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPaginatedRecordsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_youtube_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPaginatedRecordsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_youtube_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_youtube_service_proto_goTypes,
		DependencyIndexes: file_api_youtube_service_proto_depIdxs,
		EnumInfos:         file_api_youtube_service_proto_enumTypes,
		MessageInfos:      file_api_youtube_service_proto_msgTypes,
	}.Build()
	File_api_youtube_service_proto = out.File
	file_api_youtube_service_proto_rawDesc = nil
	file_api_youtube_service_proto_goTypes = nil
	file_api_youtube_service_proto_depIdxs = nil
}
