package pagination

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/rishusahu23/fam-youtube/gen/api/rpc"
)

type Rows interface {
	Slice(start, end int) Rows
	GetTimestamp(index int) time.Time
	Size() int
}

func NewPageCtxResp(pageToken *PageToken, pageSize int, rows Rows) (Rows, *rpc.PageContextResponse, error) {
	pageCtxResp := &rpc.PageContextResponse{
		BeforeToken: "",
		HasBefore:   false,
		AfterToken:  "",
		HasAfter:    false,
	}
	rows = setBeforeAndAfter(pageToken, pageSize, rows, pageCtxResp)

	if rows.Size() > 0 {
		// update offset
		n := rows.Size()
		firstRowTS := rows.GetTimestamp(0)
		lastRowTS := rows.GetTimestamp(rows.Size() - 1)
		beforeOffset := uint32(0)
		afterOffset := uint32(0)
		if pageToken != nil && pageToken.Timestamp.AsTime().Equal(lastRowTS) {
			afterOffset = pageToken.Offset
		}
		for i := 0; i < n; i++ {
			if rows.GetTimestamp(i).Equal(firstRowTS) {
				beforeOffset++
			}
			if rows.GetTimestamp(i).Equal(lastRowTS) {
				afterOffset++
			}
		}
		beforeToken := &PageToken{
			Timestamp: timestamppb.New(firstRowTS),
			Offset:    beforeOffset,
			IsReverse: true,
		}
		afterToken := &PageToken{
			Timestamp: timestamppb.New(lastRowTS),
			Offset:    afterOffset,
			IsReverse: false,
		}
		if pageCtxResp.HasBefore {
			var err error
			pageCtxResp.BeforeToken, err = beforeToken.Marshal()
			if err != nil {
				return nil, nil, err
			}
		}
		if pageCtxResp.HasAfter {
			var err error
			pageCtxResp.AfterToken, err = afterToken.Marshal()
			if err != nil {
				return nil, nil, fmt.Errorf("couldn't marshal page token: %w", err)
			}
		}
	}
	return rows, pageCtxResp, nil
}

func setBeforeAndAfter(pageToken *PageToken, pageSize int, rows Rows, pageCtxResp *rpc.PageContextResponse) Rows {
	if pageToken == nil {
		if rows.Size() > pageSize {
			pageCtxResp.HasAfter = true
		}
		pageCtxResp.HasBefore = false
	} else {
		if pageToken.IsReverse {
			if rows.Size() > pageSize {
				pageCtxResp.HasBefore = true
			}
			if pageToken.Offset > 0 {
				pageCtxResp.HasAfter = true
			}
		} else {
			if rows.Size() > pageSize {
				pageCtxResp.HasAfter = true
			}
			if pageToken.Offset > 0 {
				pageCtxResp.HasBefore = true
			}
		}
	}
	if rows.Size() > pageSize {
		if pageToken != nil && pageToken.IsReverse {
			rows = rows.Slice(1, rows.Size())
		} else {
			rows = rows.Slice(0, rows.Size()-1)
		}
	}
	return rows
}
