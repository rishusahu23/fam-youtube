package elasticsearch

import (
	"fmt"
	"github.com/rishusahu23/fam-youtube/external/contants"
	"github.com/rishusahu23/fam-youtube/external/elasticsearch/elk"
	"github.com/rishusahu23/fam-youtube/external/pkg"
	"github.com/rishusahu23/fam-youtube/gen/api/external/elasticsearch"
)

func (e *ElkClientImpl) NewElasticSearchRequest(req interface{}) pkg.SyncRequest {
	switch v := req.(type) {
	case *elasticsearch.FeedToElasticSearchRequest:
		feedReq := req.(*elasticsearch.FeedToElasticSearchRequest)
		return &elk.FeedToElasticSearchRequest{
			Method: contants.PostMethod,
			Req:    feedReq,
			Url:    "http://elasticsearch:9200/videos/_doc/1",
		}
	case *elasticsearch.GetRecordsFromElkRequest:
		fetchReq := req.(*elasticsearch.GetRecordsFromElkRequest)
		return &elk.GetRecordsFromElkRequest{
			Method: contants.PostMethod,
			Req:    fetchReq,
			Url:    "http://elasticsearch:9200/videos/_search",
		}
	default:
		fmt.Println(v)
		return nil
	}
}
