package dao

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/rishusahu23/fam-youtube/gen/api/rpc"
	"github.com/rishusahu23/fam-youtube/gen/api/youtube/record"
	"github.com/rishusahu23/fam-youtube/pkg/errors"
	"github.com/rishusahu23/fam-youtube/pkg/pagination"
	"github.com/rishusahu23/fam-youtube/youtube/model"
	"gorm.io/gorm"
	"sort"
	"strings"
)

type RecordDaoImpl struct {
	db *gorm.DB
}

func (r *RecordDaoImpl) GetById(ctx context.Context, id string) (*record.Record, error) {
	var rec model.Record
	err := r.db.Where("id = ?", id).First(&rec).Error
	if err != nil {
		return nil, err
	}
	return rec.ConvertToProto(), nil
}

func (r *RecordDaoImpl) GetByTitle(ctx context.Context, title string) ([]*record.Record, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RecordDaoImpl) GetByDescription(ctx context.Context, desc string) ([]*record.Record, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RecordDaoImpl) Create(ctx context.Context, rec *record.Record) error {
	recModel := model.ConvertToModel(rec)
	if err := r.db.Create(&recModel).Error; err != nil {
		if isDuplicateEntry(err) {
			return errors.ErrDuplicateEntry
		}
		return err
	}
	return nil
}

func NewRecordDaoImpl(db *gorm.DB) *RecordDaoImpl {
	return &RecordDaoImpl{db: db}
}

var (
	_             Dao = &RecordDaoImpl{}
	RecordWireSet     = wire.NewSet(NewRecordDaoImpl, wire.Bind(new(Dao), new(*RecordDaoImpl)))
)

func isDuplicateEntry(err error) bool {
	return strings.Contains(err.Error(), "duplicate key value violates unique constraint \"records_pkey\"")
}

func (r *RecordDaoImpl) GetPaginatedRecords(ctx context.Context, pageToken *pagination.PageToken, pageSize uint32) ([]*record.Record, *rpc.PageContextResponse, error) {
	query := r.db.Model(&model.Record{})

	if pageToken != nil {
		if pageToken.IsReverse {
			query = query.Where("published_at >= ?", pageToken.Timestamp.AsTime()).Order("published_at" + " ASC")
		} else {
			query = query.Where("published_at <= ?", pageToken.Timestamp.AsTime()).Order("published_at" + " DESC")
		}
		query = query.Offset(int(pageToken.Offset))
	} else {
		query = query.Order("published_at" + " DESC")
	}

	query = query.Limit(int(pageSize + 1))

	var recModels []*model.Record
	res := query.Find(&recModels)
	if res.Error != nil {
		return nil, nil, fmt.Errorf("error fetching records from db, err : %w", res.Error)
	}

	if pageToken != nil && pageToken.IsReverse {
		sort.Slice(recModels, func(i, j int) bool {
			return recModels[i].PublishedAt.After(recModels[j].PublishedAt)
		})
	}
	rows, pageCtxResp, err := pagination.NewPageCtxResp(pageToken, int(pageSize), model.Records(recModels))
	if err != nil {
		return nil, nil, err
	}
	recModels = rows.(model.Records)
	recProtos := make([]*record.Record, 0)
	for _, recModel := range recModels {
		recProtos = append(recProtos, recModel.ConvertToProto())
	}
	return recProtos, pageCtxResp, nil
}

func (r *RecordDaoImpl) GetByTitleAndDescription(ctx context.Context, title, desc string) ([]*record.Record, error) {
	db := r.db
	if title != "" {
		db = db.Where("title = ?", title)
	}
	if desc != "" {
		db = db.Where("description = ?", desc)
	}

	var recModels []*model.Record
	res := db.Find(&recModels)
	if res.Error != nil {
		return nil, fmt.Errorf("error fetching records from db, err : %w", res.Error)
	}
	recProtos := make([]*record.Record, 0)
	for _, recModel := range recModels {
		recProtos = append(recProtos, recModel.ConvertToProto())
	}
	return recProtos, nil
}

func (r *RecordDaoImpl) GetPartialMatchRecords(ctx context.Context, query string) ([]*record.Record, error) {
	db := r.db.WithContext(ctx)

	tsQuery := strings.ReplaceAll(query, " ", " & ")

	// Perform the search with GORM
	var videoModels []*model.Record
	res := db.Raw(`
        SELECT *,
               ts_rank(to_tsvector('english', title || ' ' || description), to_tsquery(?)) AS rank
        FROM records
        WHERE to_tsvector('english', title || ' ' || description) @@ to_tsquery(?)
        ORDER BY rank DESC
    `, tsQuery, tsQuery).Scan(&videoModels)

	// Handle errors
	if res.Error != nil {
		return nil, fmt.Errorf("error performing full-text search, err: %w", res.Error)
	}

	recProtos := make([]*record.Record, 0)
	for _, recModel := range videoModels {
		recProtos = append(recProtos, recModel.ConvertToProto())
	}
	return recProtos, nil
}
