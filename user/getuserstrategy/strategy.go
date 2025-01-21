package strategy

import (
	"context"
	"github.com/rishusahu23/fam-youtube/gen/api/user"
)

type GetUserStrategy interface {
	GetUser(ctx context.Context, request *GetUserRequest) (*GetUserResponse, error)
}

type GetUserRequest struct {
	UserId string
}

type GetUserResponse struct {
	User *user.User
}
