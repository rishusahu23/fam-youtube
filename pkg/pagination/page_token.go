package pagination

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	errors2 "github.com/rishusahu23/fam-youtube/pkg/errors"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/pkg/errors"
	gormv2 "gorm.io/gorm"

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

func AddPaginationOnCreatedAtColumn(db *gormv2.DB, pageToken *PageToken, pageSize uint32, tableName, colName string) *gormv2.DB {
	return AddPaginationOnGivenColumn(db, pageToken, pageSize, tableName, colName)
}

func AddPaginationOnGivenColumn(db *gormv2.DB, pageToken *PageToken, pageSize uint32, tableName, colName string) *gormv2.DB {
	columnName := fmt.Sprintf("%s.%s", tableName, colName)
	if tableName == "" {
		columnName = colName
	}

	if pageToken != nil {
		if pageToken.IsReverse {
			db = db.Order(columnName)
			if pageToken.Timestamp != nil {
				db = db.Where(columnName+" >= ?", pageToken.Timestamp.AsTime())
			}
		} else {
			db = db.Order(columnName + " desc")
			if pageToken.Timestamp != nil {
				db = db.Where(columnName+" <= ?", pageToken.Timestamp.AsTime())
			}
		}
		db = db.Offset(int(pageToken.Offset))
	} else {
		db = db.Order(columnName + " desc")
	}
	// fetch pageSize + 1 extra row to compute next page availability.
	db = db.Limit(int(pageSize + 1))
	return db
}

func AddPaginationOnGivenColumns(db *gormv2.DB, pageToken *PageToken, pageSize uint32, tableName string, colNames ...string) (*gormv2.DB, error) {
	if len(colNames) == 0 {
		return nil, errors.Wrap(errors2.ErrInvalidArgument, "no column to paginate on")
	}
	db = AddPaginationOnGivenColumn(db, pageToken, pageSize, tableName, colNames[0])
	for i := 1; i < len(colNames); i++ {
		db = db.Order(colNames[i])
	}
	return db, nil
}
