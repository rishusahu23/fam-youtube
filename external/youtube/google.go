package youtube

import (
	"fmt"
	"github.com/rishusahu23/fam-youtube/external/contants"
	"github.com/rishusahu23/fam-youtube/external/pkg"
	"github.com/rishusahu23/fam-youtube/external/youtube/google"
	"github.com/rishusahu23/fam-youtube/gen/api/external/youtube"
)

func (c *ClientImpl) NewGoogleRequest(req interface{}) pkg.SyncRequest {
	switch v := req.(type) {
	case *youtube.FetchYoutubeDataListRequest:
		fetchReq := req.(*youtube.FetchYoutubeDataListRequest)
		return &google.FetchYoutubeDataRequest{
			Method: contants.GetMethod,
			Req:    fetchReq,
			Url:    "https://www.googleapis.com/youtube/v3/search?part=snippet&type=video&order=date&publishedAfter=2025-01-01T00:00:00Z&q=politics&key=%v",
		}
	default:
		fmt.Println(v)
		return nil
	}
}
