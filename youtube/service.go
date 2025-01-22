package youtube

import (
	"context"
	"fmt"
	"github.com/rishusahu23/fam-youtube/external/youtube"
	vgPb "github.com/rishusahu23/fam-youtube/gen/api/external/youtube"
	"github.com/rishusahu23/fam-youtube/gen/api/rpc"
	youtubePb "github.com/rishusahu23/fam-youtube/gen/api/youtube"
	"github.com/rishusahu23/fam-youtube/gen/api/youtube/record"
	"github.com/rishusahu23/fam-youtube/pkg/pagination"
	"github.com/rishusahu23/fam-youtube/youtube/dao"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Service struct {
	youtubePb.UnimplementedYoutubeServiceServer
	googleClient youtube.Client
	dao          dao.Dao
}

func NewService(googleClient youtube.Client, dao dao.Dao) *Service {
	return &Service{
		googleClient: googleClient,
		dao:          dao,
	}
}

var _ youtubePb.YoutubeServiceServer = (*Service)(nil)

func (s *Service) TriggerJob(ctx context.Context, request *youtubePb.TriggerJobRequest) (*youtubePb.TriggerJobResponse, error) {
	resp, err := s.googleClient.FetchYoutubeData(ctx, &vgPb.FetchYoutubeDataListRequest{})
	if err != nil {
		return nil, err
	}
	records := getRecord(resp.GetItems())
	for _, item := range records {
		if err = s.dao.Create(ctx, item); err != nil {
			return nil, err
		}
	}

	fmt.Println(resp)
	return &youtubePb.TriggerJobResponse{
		Status: rpc.StatusOk(),
	}, nil
}

func getRecord(items []*vgPb.Item) []*record.Record {
	records := make([]*record.Record, len(items))
	timeStr := "2025-01-21T17:56:19Z"

	for i, item := range items {
		t, err := time.Parse(time.RFC3339, timeStr)
		if err != nil {
			fmt.Printf("Error parsing time: %v\n", err)
			continue
		}

		thumbnails := make(map[string]*record.Thumbnail)
		for key, val := range item.GetSnippet().GetThumbnails() {
			thumbnails[key] = &record.Thumbnail{
				Url:    val.GetUrl(),
				Width:  val.GetWidth(),
				Height: val.GetHeight(),
			}
		}

		records[i] = &record.Record{
			Id:          item.GetRawId().GetVideoId(),
			Title:       item.GetSnippet().GetTitle(),
			Description: item.GetSnippet().GetDescription(),
			PublishedAt: timestamppb.New(t),
			Metadata: &record.Metadata{
				Thumbnails: thumbnails,
			},
		}
	}
	return records
}

func (s *Service) GetPaginatedRecords(ctx context.Context, req *youtubePb.GetPaginatedRecordsRequest) (*youtubePb.GetPaginatedRecordsResponse, error) {
	var (
		errRes = func(status *rpc.Status) (*youtubePb.GetPaginatedRecordsResponse, error) {
			return &youtubePb.GetPaginatedRecordsResponse{
				Status: status,
			}, nil
		}
	)
	pageToken, err := pagination.GetPageToken(req.GetPageContext())
	if err != nil {
		return errRes(rpc.StatusInternal(err.Error()))
	}

	records, pageContextResp, err := s.dao.GetPaginatedRecords(ctx, pageToken, req.GetPageContext().GetPageSize())
	if err != nil {
		return errRes(rpc.StatusInternal(err.Error()))
	}

	return &youtubePb.GetPaginatedRecordsResponse{
		Status:      rpc.StatusOk(),
		PageContext: pageContextResp,
		Records:     records,
	}, nil
}

func (s *Service) GetFilteredRecords(ctx context.Context, req *youtubePb.GetFilteredRecordsRequest) (*youtubePb.GetFilteredRecordsResponse, error) {
	resp, err := s.dao.GetByTitleAndDescription(ctx, req.GetTitle(), req.GetDescription())
	if err != nil {
		return nil, err
	}
	return &youtubePb.GetFilteredRecordsResponse{
		Status:  rpc.StatusOk(),
		Records: resp,
	}, nil
}
