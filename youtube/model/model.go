package model

import (
	"github.com/rishusahu23/fam-youtube/gen/api/youtube/record"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Record struct {
	Id          string
	Title       string
	Description string
	Metadata    *record.Metadata
	PublishedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (r *Record) ConvertToProto() *record.Record {
	return &record.Record{
		Id:          r.Id,
		Title:       r.Title,
		Description: r.Description,
		PublishedAt: timestamppb.New(r.PublishedAt),
		CreatedAt:   timestamppb.New(r.CreatedAt),
		UpdatedAt:   timestamppb.New(r.UpdatedAt),
		Metadata:    r.Metadata,
	}
}

func ConvertToModel(rec *record.Record) *Record {
	return &Record{
		Id:          rec.GetId(),
		Title:       rec.GetTitle(),
		Description: rec.GetDescription(),
		Metadata:    rec.GetMetadata(),
		PublishedAt: rec.GetPublishedAt().AsTime(),
	}
}
