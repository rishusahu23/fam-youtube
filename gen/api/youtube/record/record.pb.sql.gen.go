package record

import (
	"database/sql/driver"
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
)

// Scanner interface implementation for parsing ProfileUpdateDetails while reading from DB
func (a *Metadata) Scan(src interface{}) error {
	marshalledData, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed, src: %T", src)
	}
	unmarshalOptions := &protojson.UnmarshalOptions{}
	unmarshalOptions.DiscardUnknown = true
	return unmarshalOptions.Unmarshal(marshalledData, a)
}

// Valuer interface implementation for storing the ProfileUpdateDetails in string format in DB
func (a *Metadata) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}
	return protojson.Marshal(a)
}
