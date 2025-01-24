// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/external/elasticsearch/feed.proto

package elasticsearch

import (
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

type FeedToElasticSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                    // Video ID
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`               // Video title
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`   // Video description
	PublishedAt string `protobuf:"bytes,4,opt,name=published_at,proto3" json:"published_at,omitempty"` // Video published timestamp (ISO 8601 format)
}

func (x *FeedToElasticSearchRequest) Reset() {
	*x = FeedToElasticSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_external_elasticsearch_feed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedToElasticSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedToElasticSearchRequest) ProtoMessage() {}

func (x *FeedToElasticSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_external_elasticsearch_feed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedToElasticSearchRequest.ProtoReflect.Descriptor instead.
func (*FeedToElasticSearchRequest) Descriptor() ([]byte, []int) {
	return file_api_external_elasticsearch_feed_proto_rawDescGZIP(), []int{0}
}

func (x *FeedToElasticSearchRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FeedToElasticSearchRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *FeedToElasticSearchRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FeedToElasticSearchRequest) GetPublishedAt() string {
	if x != nil {
		return x.PublishedAt
	}
	return ""
}

type FeedToElasticSearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"` // Video title
}

func (x *FeedToElasticSearchResponse) Reset() {
	*x = FeedToElasticSearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_external_elasticsearch_feed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedToElasticSearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedToElasticSearchResponse) ProtoMessage() {}

func (x *FeedToElasticSearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_external_elasticsearch_feed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedToElasticSearchResponse.ProtoReflect.Descriptor instead.
func (*FeedToElasticSearchResponse) Descriptor() ([]byte, []int) {
	return file_api_external_elasticsearch_feed_proto_rawDescGZIP(), []int{1}
}

func (x *FeedToElasticSearchResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_api_external_elasticsearch_feed_proto protoreflect.FileDescriptor

var file_api_external_elasticsearch_feed_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65,
	0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x66, 0x65, 0x65,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22,
	0x88, 0x01, 0x0a, 0x1a, 0x46, 0x65, 0x65, 0x64, 0x54, 0x6f, 0x45, 0x6c, 0x61, 0x73, 0x74, 0x69,
	0x63, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x35, 0x0a, 0x1b, 0x46, 0x65,
	0x65, 0x64, 0x54, 0x6f, 0x45, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x72, 0x69, 0x73, 0x68, 0x75, 0x73, 0x61, 0x68, 0x75, 0x32, 0x33, 0x2f, 0x66, 0x61, 0x6d, 0x2d,
	0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_external_elasticsearch_feed_proto_rawDescOnce sync.Once
	file_api_external_elasticsearch_feed_proto_rawDescData = file_api_external_elasticsearch_feed_proto_rawDesc
)

func file_api_external_elasticsearch_feed_proto_rawDescGZIP() []byte {
	file_api_external_elasticsearch_feed_proto_rawDescOnce.Do(func() {
		file_api_external_elasticsearch_feed_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_external_elasticsearch_feed_proto_rawDescData)
	})
	return file_api_external_elasticsearch_feed_proto_rawDescData
}

var file_api_external_elasticsearch_feed_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_external_elasticsearch_feed_proto_goTypes = []interface{}{
	(*FeedToElasticSearchRequest)(nil),  // 0: external.elasticsearch.FeedToElasticSearchRequest
	(*FeedToElasticSearchResponse)(nil), // 1: external.elasticsearch.FeedToElasticSearchResponse
}
var file_api_external_elasticsearch_feed_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_external_elasticsearch_feed_proto_init() }
func file_api_external_elasticsearch_feed_proto_init() {
	if File_api_external_elasticsearch_feed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_external_elasticsearch_feed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedToElasticSearchRequest); i {
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
		file_api_external_elasticsearch_feed_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedToElasticSearchResponse); i {
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
			RawDescriptor: file_api_external_elasticsearch_feed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_external_elasticsearch_feed_proto_goTypes,
		DependencyIndexes: file_api_external_elasticsearch_feed_proto_depIdxs,
		MessageInfos:      file_api_external_elasticsearch_feed_proto_msgTypes,
	}.Build()
	File_api_external_elasticsearch_feed_proto = out.File
	file_api_external_elasticsearch_feed_proto_rawDesc = nil
	file_api_external_elasticsearch_feed_proto_goTypes = nil
	file_api_external_elasticsearch_feed_proto_depIdxs = nil
}
