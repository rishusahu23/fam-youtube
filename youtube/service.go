package youtube

import (
	"context"
	"github.com/rishusahu23/fam-youtube/external/youtube"
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
	return &youtubePb.TriggerJobResponse{
		Status: rpc.StatusOk(),
	}, nil
}
