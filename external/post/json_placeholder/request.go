package placeholder

import "github.com/rishusahu23/fam-youtube/external/enums"

type FetchPostRequest struct {
	Vendor enums.Vendor
	PostId string
}

func (f *FetchPostRequest) GetHeader() enums.Vendor {
	return f.Vendor
}
