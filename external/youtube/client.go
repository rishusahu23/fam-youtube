package youtube

import (
	"context"
	"fmt"
	"github.com/rishusahu23/fam-youtube/config"
	"github.com/rishusahu23/fam-youtube/external/ohttp"
	"github.com/rishusahu23/fam-youtube/external/pkg"
	"github.com/rishusahu23/fam-youtube/gen/api/external/youtube"
)

type Client interface {
	FetchYoutubeData(ctx context.Context, req *youtube.FetchYoutubeDataListRequest) (*youtube.FetchYoutubeDataListResponse, error)
}

type ClientImpl struct {
	httpRequestHandler ohttp.IHttpRequestHandler
	conf               *config.Config
}

var _ Client = &ClientImpl{}

func NewYoutubeClientImpl(httpRequestHandler *ohttp.HttpRequestHandler, conf *config.Config) *ClientImpl {
	return &ClientImpl{
		httpRequestHandler: httpRequestHandler,
		conf:               conf,
	}
}

func (c *ClientImpl) requestFactoryMap() pkg.SyncRequestFactory {
	return c.NewGoogleRequest

}

func (c *ClientImpl) FetchYoutubeData(ctx context.Context, req *youtube.FetchYoutubeDataListRequest) (*youtube.FetchYoutubeDataListResponse, error) {
	venReq, err := pkg.NewVendorRequest(req, c.requestFactoryMap())
	if err != nil {
		return nil, err
	}
	fmt.Println(venReq)
	// will make api call to youtube api
	resp, err := c.httpRequestHandler.MakeHttpRequest(ctx, venReq)
	if err != nil {
		return nil, err
	}

	return resp.(*youtube.FetchYoutubeDataListResponse), nil
}
