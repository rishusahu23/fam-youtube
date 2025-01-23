package pagination

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/rishusahu23/fam-youtube/gen/api/rpc"
)

type PageToken struct {
	Timestamp *timestamp.Timestamp
	Offset    uint32
	IsReverse bool
}

func (p *PageToken) GetTimestamp() *timestamp.Timestamp {
	if p == nil {
		return nil
	}
	return p.Timestamp
}
func (p *PageToken) GetOffset() uint32 {
	if p == nil {
		return 0
	}
	return p.Offset
}
func (p *PageToken) GetIsReverse() bool {
	if p == nil {
		return false
	}
	return p.IsReverse
}

func (p *PageToken) Marshal() (string, error) {
	if p == nil {
		return "", nil
	}
	s, err := json.Marshal(p)
	if err != nil {
		return "", fmt.Errorf("failed to marshal page token: %w", err)
	}
	return base64.StdEncoding.EncodeToString(s), nil
}

func (p *PageToken) Unmarshal(s string) error {
	if s == "" {
		return nil
	}
	ds, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return fmt.Errorf("failed to unmarshal page token: %w", err)
	}
	if p == nil {
		*p = PageToken{}
	}
	return json.Unmarshal(ds, p)
}

func GetPageToken(req *rpc.PageContextRequest) (*PageToken, error) {
	var tokenStr string
	if req.GetAfterToken() != "" {
		tokenStr = req.GetAfterToken()
	} else {
		tokenStr = req.GetBeforeToken()
	}
	if tokenStr == "" {
		return nil, nil
	}
	token := &PageToken{}
	err := token.Unmarshal(tokenStr)
	return token, err
}
