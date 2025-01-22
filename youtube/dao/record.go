package dao

import (
	"context"
	"github.com/google/wire"
	"github.com/rishusahu23/fam-youtube/gen/api/youtube/record"
	"github.com/rishusahu23/fam-youtube/youtube/model"
	"gorm.io/gorm"
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
			return nil
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
