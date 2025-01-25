package elasticsearch

import (
	"context"
	"fmt"
	"github.com/rishusahu23/fam-youtube/config"
	"github.com/rishusahu23/fam-youtube/external/ohttp"
	"github.com/rishusahu23/fam-youtube/external/pkg"
	"github.com/rishusahu23/fam-youtube/gen/api/external/elasticsearch"
)

type ElkClient interface {
	FeedToElasticSearch(ctx context.Context, request *elasticsearch.FeedToElasticSearchRequest) (*elasticsearch.FeedToElasticSearchResponse, error)
	GetRecordsFromElk(ctx context.Context, request *elasticsearch.GetRecordsFromElkRequest) (*elasticsearch.GetRecordsFromElkResponse, error)
}

type ElkClientImpl struct {
	httpRequestHandler ohttp.IHttpRequestHandler
	conf               *config.Config
}

func NewElkClient(httpRequestHandler *ohttp.HttpRequestHandler, conf *config.Config) *ElkClientImpl {
	return &ElkClientImpl{
		httpRequestHandler: httpRequestHandler,
		conf:               conf,
	}
}

func (e *ElkClientImpl) requestFactoryMap() pkg.SyncRequestFactory {
	return e.NewElasticSearchRequest
}

var _ ElkClient = &ElkClientImpl{}

func (e *ElkClientImpl) FeedToElasticSearch(ctx context.Context, req *elasticsearch.FeedToElasticSearchRequest) (*elasticsearch.FeedToElasticSearchResponse, error) {
	venReq, err := pkg.NewVendorRequest(req, e.requestFactoryMap())
	if err != nil {
		return nil, err
	}
	fmt.Println(venReq)
	resp, err := e.httpRequestHandler.MakeHttpRequest(ctx, venReq)
	if err != nil {
		return nil, err
	}

	return resp.(*elasticsearch.FeedToElasticSearchResponse), nil
}

func (e *ElkClientImpl) GetRecordsFromElk(ctx context.Context, req *elasticsearch.GetRecordsFromElkRequest) (*elasticsearch.GetRecordsFromElkResponse, error) {
	venReq, err := pkg.NewVendorRequest(req, e.requestFactoryMap())
	if err != nil {
		return nil, err
	}
	fmt.Println(venReq)
	resp, err := e.httpRequestHandler.MakeHttpRequest(ctx, venReq)
	if err != nil {
		return nil, err
	}

	return resp.(*elasticsearch.GetRecordsFromElkResponse), nil
}
