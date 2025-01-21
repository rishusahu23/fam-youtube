package google

import (
	"fmt"
	"github.com/rishusahu23/fam-youtube/external/contants"
	"github.com/rishusahu23/fam-youtube/external/pkg"
	vgPb "github.com/rishusahu23/fam-youtube/gen/api/external/youtube"
	"google.golang.org/protobuf/encoding/protojson"
)

type FetchYoutubeDataRequest struct {
	Method string
	Req    *vgPb.FetchYoutubeDataListRequest
	Url    string
}

func (v *FetchYoutubeDataRequest) GetMethod() string {
	return contants.GetMethod
}

func (v *FetchYoutubeDataRequest) GetURL() string {
	return fmt.Sprintf("%v", v.Url)
}

func (v *FetchYoutubeDataRequest) GetResponse() pkg.Response {
	return &FetchYoutubeDataResponse{}
}

func (v *FetchYoutubeDataRequest) Marshal() ([]byte, error) {
	return nil, nil
}

type FetchYoutubeDataResponse struct {
}

func (V *FetchYoutubeDataResponse) Unmarshal(b []byte) (interface{}, error) {
	vendorRes := &vgPb.FetchYoutubeDataListResponse{}
	um := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}
	err := um.Unmarshal(b, vendorRes)
	if err != nil {
		return nil, err
	}
	return vendorRes, nil
}
