package elk

import (
	"fmt"
	"github.com/rishusahu23/fam-youtube/external/pkg"
	vgPb "github.com/rishusahu23/fam-youtube/gen/api/external/elasticsearch"
	"github.com/rishusahu23/fam-youtube/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
)

type FeedToElasticSearchRequest struct {
	Method string
	Req    *vgPb.FeedToElasticSearchRequest
	Url    string
}

func (f *FeedToElasticSearchRequest) GetMethod() string {
	return f.Method
}

func (f *FeedToElasticSearchRequest) GetURL() string {
	return f.Url
}

func (f *FeedToElasticSearchRequest) GetResponse() pkg.Response {
	return &FeedToElasticSearchResponse{}
}

func (f *FeedToElasticSearchRequest) Marshal() ([]byte, error) {
	return protojson.Marshal(f.Req)
}

type FeedToElasticSearchResponse struct {
}

func (f *FeedToElasticSearchResponse) Unmarshal(b []byte) (interface{}, error) {
	vgRes := &vgPb.FeedToElasticSearchResponse{}
	um := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}
	err := um.Unmarshal(b, vgRes)
	if err != nil {
		return nil, err
	}
	fmt.Println("vgRes", string(b))
	if !(vgRes.GetResult() == "created" || vgRes.GetResult() == "updated") {
		return nil, errors.ErrInternalServer
	}
	return vgRes, nil
}
