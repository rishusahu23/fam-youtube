package youtube

import (
	"context"
	"fmt"
	"github.com/rishusahu23/fam-youtube/external/youtube"
	vgPb "github.com/rishusahu23/fam-youtube/gen/api/external/youtube"
	"github.com/rishusahu23/fam-youtube/gen/api/rpc"
	youtubePb "github.com/rishusahu23/fam-youtube/gen/api/youtube"
)

type Service struct {
	youtubePb.UnimplementedYoutubeServiceServer
	googleClient youtube.Client
}

func NewService(googleClient youtube.Client) *Service {
	return &Service{
		googleClient: googleClient,
	}
}

var _ youtubePb.YoutubeServiceServer = (*Service)(nil)

func (s *Service) TriggerJob(ctx context.Context, request *youtubePb.TriggerJobRequest) (*youtubePb.TriggerJobResponse, error) {
	resp, err := s.googleClient.FetchYoutubeData(ctx, &vgPb.FetchYoutubeDataListRequest{})
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	return &youtubePb.TriggerJobResponse{
		Status: rpc.StatusOk(),
	}, nil
}
