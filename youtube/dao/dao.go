package dao

import (
	"context"
	"github.com/rishusahu23/fam-youtube/gen/api/youtube/record"
)

type Dao interface {
	GetById(ctx context.Context, id string) (*record.Record, error)
	GetByTitle(ctx context.Context, title string) ([]*record.Record, error)
	GetByDescription(ctx context.Context, desc string) ([]*record.Record, error)
	Create(ctx context.Context, rec *record.Record) error
}
