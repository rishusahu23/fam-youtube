// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/external/youtube/google.proto

package youtube

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

type FetchYoutubeDataListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FetchYoutubeDataListRequest) Reset() {
	*x = FetchYoutubeDataListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_external_youtube_google_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchYoutubeDataListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchYoutubeDataListRequest) ProtoMessage() {}

func (x *FetchYoutubeDataListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_external_youtube_google_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchYoutubeDataListRequest.ProtoReflect.Descriptor instead.
func (*FetchYoutubeDataListRequest) Descriptor() ([]byte, []int) {
	return file_api_external_youtube_google_proto_rawDescGZIP(), []int{0}
}

type FetchYoutubeDataListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *FetchYoutubeDataListResponse) Reset() {
	*x = FetchYoutubeDataListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_external_youtube_google_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchYoutubeDataListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchYoutubeDataListResponse) ProtoMessage() {}

func (x *FetchYoutubeDataListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_external_youtube_google_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchYoutubeDataListResponse.ProtoReflect.Descriptor instead.
func (*FetchYoutubeDataListResponse) Descriptor() ([]byte, []int) {
	return file_api_external_youtube_google_proto_rawDescGZIP(), []int{1}
}

func (x *FetchYoutubeDataListResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawId   *RawId   `protobuf:"bytes,1,opt,name=raw_id,json=id,proto3" json:"raw_id,omitempty"`
	Snippet *Snippet `protobuf:"bytes,2,opt,name=snippet,proto3" json:"snippet,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_external_youtube_google_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_api_external_youtube_google_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_api_external_youtube_google_proto_rawDescGZIP(), []int{2}
}

func (x *Item) GetRawId() *RawId {
	if x != nil {
		return x.RawId
	}
	return nil
}

func (x *Item) GetSnippet() *Snippet {
	if x != nil {
		return x.Snippet
	}
	return nil
}

type Snippet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublishedAt string                `protobuf:"bytes,1,opt,name=published_at,json=publishedAt,proto3" json:"published_at,omitempty"`
	Title       string                `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string                `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Thumbnails  map[string]*Thumbnail `protobuf:"bytes,4,rep,name=thumbnails,proto3" json:"thumbnails,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Snippet) Reset() {
	*x = Snippet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_external_youtube_google_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Snippet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Snippet) ProtoMessage() {}

func (x *Snippet) ProtoReflect() protoreflect.Message {
	mi := &file_api_external_youtube_google_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Snippet.ProtoReflect.Descriptor instead.
func (*Snippet) Descriptor() ([]byte, []int) {
	return file_api_external_youtube_google_proto_rawDescGZIP(), []int{3}
}

func (x *Snippet) GetPublishedAt() string {
	if x != nil {
		return x.PublishedAt
	}
	return ""
}

func (x *Snippet) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Snippet) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Snippet) GetThumbnails() map[string]*Thumbnail {
	if x != nil {
		return x.Thumbnails
	}
	return nil
}

type Thumbnail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url    string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Width  int32  `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height int32  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *Thumbnail) Reset() {
	*x = Thumbnail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_external_youtube_google_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Thumbnail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Thumbnail) ProtoMessage() {}

func (x *Thumbnail) ProtoReflect() protoreflect.Message {
	mi := &file_api_external_youtube_google_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Thumbnail.ProtoReflect.Descriptor instead.
func (*Thumbnail) Descriptor() ([]byte, []int) {
	return file_api_external_youtube_google_proto_rawDescGZIP(), []int{4}
}

func (x *Thumbnail) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Thumbnail) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Thumbnail) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type RawId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId string `protobuf:"bytes,1,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
}

func (x *RawId) Reset() {
	*x = RawId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_external_youtube_google_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawId) ProtoMessage() {}

func (x *RawId) ProtoReflect() protoreflect.Message {
	mi := &file_api_external_youtube_google_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawId.ProtoReflect.Descriptor instead.
func (*RawId) Descriptor() ([]byte, []int) {
	return file_api_external_youtube_google_proto_rawDescGZIP(), []int{5}
}

func (x *RawId) GetVideoId() string {
	if x != nil {
		return x.VideoId
	}
	return ""
}

var File_api_external_youtube_google_proto protoreflect.FileDescriptor

var file_api_external_youtube_google_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x79,
	0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x79, 0x6f,
	0x75, 0x74, 0x75, 0x62, 0x65, 0x22, 0x1d, 0x0a, 0x1b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x59, 0x6f,
	0x75, 0x74, 0x75, 0x62, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x4c, 0x0a, 0x1c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x59, 0x6f, 0x75,
	0x74, 0x75, 0x62, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x79,
	0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x22, 0x68, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x2b, 0x0a, 0x06, 0x72, 0x61,
	0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x2e, 0x52, 0x61,
	0x77, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x07, 0x73, 0x6e, 0x69, 0x70, 0x70,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x2e, 0x53, 0x6e, 0x69, 0x70,
	0x70, 0x65, 0x74, 0x52, 0x07, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x22, 0x8b, 0x02, 0x0a,
	0x07, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x0a, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x2e, 0x53, 0x6e, 0x69, 0x70, 0x70,
	0x65, 0x74, 0x2e, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0a, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x5a,
	0x0a, 0x0f, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x79, 0x6f,
	0x75, 0x74, 0x75, 0x62, 0x65, 0x2e, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4b, 0x0a, 0x09, 0x54, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x22, 0x0a, 0x05, 0x52, 0x61, 0x77, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x42, 0x3d, 0x5a, 0x3b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x73, 0x68, 0x75, 0x73,
	0x61, 0x68, 0x75, 0x32, 0x33, 0x2f, 0x66, 0x61, 0x6d, 0x2d, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62,
	0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_external_youtube_google_proto_rawDescOnce sync.Once
	file_api_external_youtube_google_proto_rawDescData = file_api_external_youtube_google_proto_rawDesc
)

func file_api_external_youtube_google_proto_rawDescGZIP() []byte {
	file_api_external_youtube_google_proto_rawDescOnce.Do(func() {
		file_api_external_youtube_google_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_external_youtube_google_proto_rawDescData)
	})
	return file_api_external_youtube_google_proto_rawDescData
}

var file_api_external_youtube_google_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_external_youtube_google_proto_goTypes = []interface{}{
	(*FetchYoutubeDataListRequest)(nil),  // 0: external.youtube.FetchYoutubeDataListRequest
	(*FetchYoutubeDataListResponse)(nil), // 1: external.youtube.FetchYoutubeDataListResponse
	(*Item)(nil),                         // 2: external.youtube.Item
	(*Snippet)(nil),                      // 3: external.youtube.Snippet
	(*Thumbnail)(nil),                    // 4: external.youtube.Thumbnail
	(*RawId)(nil),                        // 5: external.youtube.RawId
	nil,                                  // 6: external.youtube.Snippet.ThumbnailsEntry
}
var file_api_external_youtube_google_proto_depIdxs = []int32{
	2, // 0: external.youtube.FetchYoutubeDataListResponse.items:type_name -> external.youtube.Item
	5, // 1: external.youtube.Item.raw_id:type_name -> external.youtube.RawId
	3, // 2: external.youtube.Item.snippet:type_name -> external.youtube.Snippet
	6, // 3: external.youtube.Snippet.thumbnails:type_name -> external.youtube.Snippet.ThumbnailsEntry
	4, // 4: external.youtube.Snippet.ThumbnailsEntry.value:type_name -> external.youtube.Thumbnail
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_external_youtube_google_proto_init() }
func file_api_external_youtube_google_proto_init() {
	if File_api_external_youtube_google_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_external_youtube_google_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchYoutubeDataListRequest); i {
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
		file_api_external_youtube_google_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchYoutubeDataListResponse); i {
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
		file_api_external_youtube_google_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_api_external_youtube_google_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Snippet); i {
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
		file_api_external_youtube_google_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Thumbnail); i {
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
		file_api_external_youtube_google_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawId); i {
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
			RawDescriptor: file_api_external_youtube_google_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_external_youtube_google_proto_goTypes,
		DependencyIndexes: file_api_external_youtube_google_proto_depIdxs,
		MessageInfos:      file_api_external_youtube_google_proto_msgTypes,
	}.Build()
	File_api_external_youtube_google_proto = out.File
	file_api_external_youtube_google_proto_rawDesc = nil
	file_api_external_youtube_google_proto_goTypes = nil
	file_api_external_youtube_google_proto_depIdxs = nil
}
