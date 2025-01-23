package dao

import (
	"context"
	"github.com/rishusahu23/fam-youtube/gen/api/rpc"
	"github.com/rishusahu23/fam-youtube/gen/api/youtube/record"
	"github.com/rishusahu23/fam-youtube/pkg/pagination"
)

type Dao interface {
	GetById(ctx context.Context, id string) (*record.Record, error)
	GetByTitle(ctx context.Context, title string) ([]*record.Record, error)
	GetByDescription(ctx context.Context, desc string) ([]*record.Record, error)
	GetByTitleAndDescription(ctx context.Context, title, desc string) ([]*record.Record, error)
	Create(ctx context.Context, rec *record.Record) error
	GetPaginatedRecords(ctx context.Context, pageToken *pagination.PageToken, pageSize uint32) ([]*record.Record, *rpc.PageContextResponse, error)
	GetPartialMatchRecords(ctx context.Context, query string) ([]*record.Record, error)
}
