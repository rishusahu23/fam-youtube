package youtube

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rishusahu23/fam-youtube/config"
	"github.com/rishusahu23/fam-youtube/external/youtube"
	vgPb "github.com/rishusahu23/fam-youtube/gen/api/external/youtube"
	"github.com/rishusahu23/fam-youtube/gen/api/rpc"
	youtubePb "github.com/rishusahu23/fam-youtube/gen/api/youtube"
	"github.com/rishusahu23/fam-youtube/gen/api/youtube/record"
	custerr "github.com/rishusahu23/fam-youtube/pkg/errors"
	"github.com/rishusahu23/fam-youtube/pkg/pagination"
	"github.com/rishusahu23/fam-youtube/youtube/dao"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Service struct {
	youtubePb.UnimplementedYoutubeServiceServer
	googleClient youtube.Client
	dao          dao.Dao
	ind          int
	conf         *config.Config
}

func NewService(conf *config.Config, googleClient youtube.Client, dao dao.Dao) *Service {
	return &Service{
		googleClient: googleClient,
		dao:          dao,
		ind:          0,
		conf:         conf,
	}
}

var _ youtubePb.YoutubeServiceServer = (*Service)(nil)

// TriggerJob api will be used to fetch data from youtube google api.
// Do not need to trigger this, it will be called automatically from server every 10 seconds.
func (s *Service) TriggerJob(ctx context.Context, request *youtubePb.TriggerJobRequest) (*youtubePb.TriggerJobResponse, error) {
	for i := 0; i < len(s.conf.ApiKeys); i++ {
		err := s.triggerJob(ctx, &vgPb.FetchYoutubeDataListRequest{
			ApiKey: s.conf.ApiKeys[s.ind],
		})
		// If quota of allowed request finished, we try with new key until all keys got exhausted
		if err != nil && errors.Is(err, custerr.ErrQuotaExceeded) {
			// we update the index to get new keys
			s.ind = (s.ind + 1) % len(s.conf.ApiKeys)
			continue
		}
		// return internal if some other issue.
		if err != nil {
			return &youtubePb.TriggerJobResponse{
				Status: rpc.StatusInternal(err.Error()),
			}, status.Errorf(codes.Internal, err.Error())
		}
		return &youtubePb.TriggerJobResponse{
			Status: rpc.StatusOk(),
		}, nil
	}

	// return internal if all keys quota exhausted
	return &youtubePb.TriggerJobResponse{
		Status: rpc.StatusInternal(""),
	}, status.Errorf(codes.Internal, "all keys quota exhausted")

}

func (s *Service) triggerJob(ctx context.Context, request *vgPb.FetchYoutubeDataListRequest) error {
	resp, err := s.googleClient.FetchYoutubeData(ctx, request)
	if err != nil {
		return err
	}
	records := getRecord(resp.GetItems())
	for _, item := range records {
		if err = s.dao.Create(ctx, item); err != nil {
			return err
		}
	}

	return nil
}

func getRecord(items []*vgPb.Item) []*record.Record {
	records := make([]*record.Record, len(items))

	for i, item := range items {
		t, err := time.Parse(time.RFC3339, item.GetSnippet().GetPublishedAt())
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

func (s *Service) GetPartialMatchRecords(ctx context.Context, req *youtubePb.GetPartialMatchRecordsRequest) (*youtubePb.GetPartialMatchRecordsResponse, error) {
	resp, err := s.dao.GetPartialMatchRecords(ctx, req.GetQuery())
	if err != nil {
		return nil, err
	}
	return &youtubePb.GetPartialMatchRecordsResponse{
		Status:  rpc.StatusOk(),
		Records: resp,
	}, nil
}
