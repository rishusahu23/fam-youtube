package elk

import (
	"fmt"
	"github.com/rishusahu23/fam-youtube/external/pkg"
	vgPb "github.com/rishusahu23/fam-youtube/gen/api/external/elasticsearch"
	"google.golang.org/protobuf/encoding/protojson"
)

type GetRecordsFromElkRequest struct {
	Method string
	Req    *vgPb.GetRecordsFromElkRequest
	Url    string
}

func (f *GetRecordsFromElkRequest) GetMethod() string {
	return f.Method
}

func (f *GetRecordsFromElkRequest) GetURL() string {
	return f.Url
}

func (f *GetRecordsFromElkRequest) GetResponse() pkg.Response {
	return &GetRecordsFromElkResponse{}
}

func (f *GetRecordsFromElkRequest) Marshal() ([]byte, error) {
	return protojson.Marshal(f.Req)
}

type GetRecordsFromElkResponse struct {
}

func (f *GetRecordsFromElkResponse) Unmarshal(b []byte) (interface{}, error) {
	vgRes := &vgPb.GetRecordsFromElkResponse{}
	um := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}
	err := um.Unmarshal(b, vgRes)
	if err != nil {
		return nil, err
	}
	fmt.Println("vgRes", string(b))
	return vgRes, nil
}
