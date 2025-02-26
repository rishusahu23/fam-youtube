// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/youtube/service.proto

package youtube

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	YoutubeService_TriggerJob_FullMethodName                    = "/youtube.YoutubeService/TriggerJob"
	YoutubeService_GetPaginatedRecords_FullMethodName           = "/youtube.YoutubeService/GetPaginatedRecords"
	YoutubeService_GetFilteredRecords_FullMethodName            = "/youtube.YoutubeService/GetFilteredRecords"
	YoutubeService_GetPartialMatchRecords_FullMethodName        = "/youtube.YoutubeService/GetPartialMatchRecords"
	YoutubeService_GetPartialMatchRecordsFromElk_FullMethodName = "/youtube.YoutubeService/GetPartialMatchRecordsFromElk"
)

// YoutubeServiceClient is the client API for YoutubeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YoutubeServiceClient interface {
	// api to fetch records from youtube api,
	// api will called every 10 seconds from server in goroutine to fetch new data from google youtube api
	TriggerJob(ctx context.Context, in *TriggerJobRequest, opts ...grpc.CallOption) (*TriggerJobResponse, error)
	// GetPaginatedRecords gets paginated records specify pagesize,
	// for first time before and after token, for second time use after token form first response
	GetPaginatedRecords(ctx context.Context, in *GetPaginatedRecordsRequest, opts ...grpc.CallOption) (*GetPaginatedRecordsResponse, error)
	// GetFilteredRecords will return the videos on the basis of title and description.
	// if title or description is empty it will be ignored.
	// it will do full text search and case should match.
	// copy the title and description from db itself.
	// generate url from /Users/rishusahu/go/src/github.com/rishusahu23/fam-youtube/main/main.go
	// by copying title and description, generating url is necessary else it won't work
	GetFilteredRecords(ctx context.Context, in *GetFilteredRecordsRequest, opts ...grpc.CallOption) (*GetFilteredRecordsResponse, error)
	// GetPartialMatchRecords is Optimise version of GetFilteredRecords search api, so that it's able to search videos containing partial match for the search query in either video title or description.
	//   - Ex 1: A video with title `*How to make tea?*` should match for the search query `tea how`
	GetPartialMatchRecords(ctx context.Context, in *GetPartialMatchRecordsRequest, opts ...grpc.CallOption) (*GetPartialMatchRecordsResponse, error)
	// GetPartialMatchRecordsFromElk is Optimise version of GetFilteredRecords search api, so that it's able to search videos containing partial match for the search query in either video title or description.
	//   - Ex 1: A video with title `*How to make tea?*` should match for the search query `tea how`
	GetPartialMatchRecordsFromElk(ctx context.Context, in *GetPartialMatchRecordsFromElkRequest, opts ...grpc.CallOption) (*GetPartialMatchRecordsFromElkResponse, error)
}

type youtubeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewYoutubeServiceClient(cc grpc.ClientConnInterface) YoutubeServiceClient {
	return &youtubeServiceClient{cc}
}

func (c *youtubeServiceClient) TriggerJob(ctx context.Context, in *TriggerJobRequest, opts ...grpc.CallOption) (*TriggerJobResponse, error) {
	out := new(TriggerJobResponse)
	err := c.cc.Invoke(ctx, YoutubeService_TriggerJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *youtubeServiceClient) GetPaginatedRecords(ctx context.Context, in *GetPaginatedRecordsRequest, opts ...grpc.CallOption) (*GetPaginatedRecordsResponse, error) {
	out := new(GetPaginatedRecordsResponse)
	err := c.cc.Invoke(ctx, YoutubeService_GetPaginatedRecords_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *youtubeServiceClient) GetFilteredRecords(ctx context.Context, in *GetFilteredRecordsRequest, opts ...grpc.CallOption) (*GetFilteredRecordsResponse, error) {
	out := new(GetFilteredRecordsResponse)
	err := c.cc.Invoke(ctx, YoutubeService_GetFilteredRecords_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *youtubeServiceClient) GetPartialMatchRecords(ctx context.Context, in *GetPartialMatchRecordsRequest, opts ...grpc.CallOption) (*GetPartialMatchRecordsResponse, error) {
	out := new(GetPartialMatchRecordsResponse)
	err := c.cc.Invoke(ctx, YoutubeService_GetPartialMatchRecords_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *youtubeServiceClient) GetPartialMatchRecordsFromElk(ctx context.Context, in *GetPartialMatchRecordsFromElkRequest, opts ...grpc.CallOption) (*GetPartialMatchRecordsFromElkResponse, error) {
	out := new(GetPartialMatchRecordsFromElkResponse)
	err := c.cc.Invoke(ctx, YoutubeService_GetPartialMatchRecordsFromElk_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YoutubeServiceServer is the server API for YoutubeService service.
// All implementations must embed UnimplementedYoutubeServiceServer
// for forward compatibility
type YoutubeServiceServer interface {
	// api to fetch records from youtube api,
	// api will called every 10 seconds from server in goroutine to fetch new data from google youtube api
	TriggerJob(context.Context, *TriggerJobRequest) (*TriggerJobResponse, error)
	// GetPaginatedRecords gets paginated records specify pagesize,
	// for first time before and after token, for second time use after token form first response
	GetPaginatedRecords(context.Context, *GetPaginatedRecordsRequest) (*GetPaginatedRecordsResponse, error)
	// GetFilteredRecords will return the videos on the basis of title and description.
	// if title or description is empty it will be ignored.
	// it will do full text search and case should match.
	// copy the title and description from db itself.
	// generate url from /Users/rishusahu/go/src/github.com/rishusahu23/fam-youtube/main/main.go
	// by copying title and description, generating url is necessary else it won't work
	GetFilteredRecords(context.Context, *GetFilteredRecordsRequest) (*GetFilteredRecordsResponse, error)
	// GetPartialMatchRecords is Optimise version of GetFilteredRecords search api, so that it's able to search videos containing partial match for the search query in either video title or description.
	//   - Ex 1: A video with title `*How to make tea?*` should match for the search query `tea how`
	GetPartialMatchRecords(context.Context, *GetPartialMatchRecordsRequest) (*GetPartialMatchRecordsResponse, error)
	// GetPartialMatchRecordsFromElk is Optimise version of GetFilteredRecords search api, so that it's able to search videos containing partial match for the search query in either video title or description.
	//   - Ex 1: A video with title `*How to make tea?*` should match for the search query `tea how`
	GetPartialMatchRecordsFromElk(context.Context, *GetPartialMatchRecordsFromElkRequest) (*GetPartialMatchRecordsFromElkResponse, error)
	mustEmbedUnimplementedYoutubeServiceServer()
}

// UnimplementedYoutubeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedYoutubeServiceServer struct {
}

func (UnimplementedYoutubeServiceServer) TriggerJob(context.Context, *TriggerJobRequest) (*TriggerJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggerJob not implemented")
}
func (UnimplementedYoutubeServiceServer) GetPaginatedRecords(context.Context, *GetPaginatedRecordsRequest) (*GetPaginatedRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaginatedRecords not implemented")
}
func (UnimplementedYoutubeServiceServer) GetFilteredRecords(context.Context, *GetFilteredRecordsRequest) (*GetFilteredRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilteredRecords not implemented")
}
func (UnimplementedYoutubeServiceServer) GetPartialMatchRecords(context.Context, *GetPartialMatchRecordsRequest) (*GetPartialMatchRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPartialMatchRecords not implemented")
}
func (UnimplementedYoutubeServiceServer) GetPartialMatchRecordsFromElk(context.Context, *GetPartialMatchRecordsFromElkRequest) (*GetPartialMatchRecordsFromElkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPartialMatchRecordsFromElk not implemented")
}
func (UnimplementedYoutubeServiceServer) mustEmbedUnimplementedYoutubeServiceServer() {}

// UnsafeYoutubeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YoutubeServiceServer will
// result in compilation errors.
type UnsafeYoutubeServiceServer interface {
	mustEmbedUnimplementedYoutubeServiceServer()
}

func RegisterYoutubeServiceServer(s grpc.ServiceRegistrar, srv YoutubeServiceServer) {
	s.RegisterService(&YoutubeService_ServiceDesc, srv)
}

func _YoutubeService_TriggerJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YoutubeServiceServer).TriggerJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YoutubeService_TriggerJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YoutubeServiceServer).TriggerJob(ctx, req.(*TriggerJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YoutubeService_GetPaginatedRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaginatedRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YoutubeServiceServer).GetPaginatedRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YoutubeService_GetPaginatedRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YoutubeServiceServer).GetPaginatedRecords(ctx, req.(*GetPaginatedRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YoutubeService_GetFilteredRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFilteredRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YoutubeServiceServer).GetFilteredRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YoutubeService_GetFilteredRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YoutubeServiceServer).GetFilteredRecords(ctx, req.(*GetFilteredRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YoutubeService_GetPartialMatchRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPartialMatchRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YoutubeServiceServer).GetPartialMatchRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YoutubeService_GetPartialMatchRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YoutubeServiceServer).GetPartialMatchRecords(ctx, req.(*GetPartialMatchRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YoutubeService_GetPartialMatchRecordsFromElk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPartialMatchRecordsFromElkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YoutubeServiceServer).GetPartialMatchRecordsFromElk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YoutubeService_GetPartialMatchRecordsFromElk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YoutubeServiceServer).GetPartialMatchRecordsFromElk(ctx, req.(*GetPartialMatchRecordsFromElkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// YoutubeService_ServiceDesc is the grpc.ServiceDesc for YoutubeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YoutubeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "youtube.YoutubeService",
	HandlerType: (*YoutubeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TriggerJob",
			Handler:    _YoutubeService_TriggerJob_Handler,
		},
		{
			MethodName: "GetPaginatedRecords",
			Handler:    _YoutubeService_GetPaginatedRecords_Handler,
		},
		{
			MethodName: "GetFilteredRecords",
			Handler:    _YoutubeService_GetFilteredRecords_Handler,
		},
		{
			MethodName: "GetPartialMatchRecords",
			Handler:    _YoutubeService_GetPartialMatchRecords_Handler,
		},
		{
			MethodName: "GetPartialMatchRecordsFromElk",
			Handler:    _YoutubeService_GetPartialMatchRecordsFromElk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/youtube/service.proto",
}
