package datetime

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func TimestampToString(t *timestamp.Timestamp, format string) (string, error) {
	return "", nil
}

func StringToTimestamp(timeStr string) (*timestamp.Timestamp, error) {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		fmt.Printf("Error parsing time: %v\n", err)
		return nil, err
	}
	return timestamppb.New(t), nil
}
