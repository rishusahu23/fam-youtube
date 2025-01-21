package post

import (
	"fmt"
	"github.com/rishusahu23/fam-youtube/external/contants"
	"github.com/rishusahu23/fam-youtube/external/pkg"
	placeholder "github.com/rishusahu23/fam-youtube/external/post/json_placeholder"
)

func (c *ClientImpl) NewPlaceholderRequest(req interface{}) pkg.SyncRequest {
	switch v := req.(type) {
	case *placeholder.FetchPostRequest:
		fetchReq := req.(*placeholder.FetchPostRequest)
		return &placeholder.VFetchPostRequest{
			Method: contants.GetMethod,
			Req:    fetchReq,
			Url:    c.conf.ExternalService.JsonPlaceholder.FetchPostUrl,
		}
	default:
		fmt.Println(v)
		return nil
	}
}
