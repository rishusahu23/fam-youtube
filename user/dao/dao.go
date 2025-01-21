package dao

import (
	"context"
	"github.com/rishusahu23/fam-youtube/gen/api/user"
	"github.com/rishusahu23/fam-youtube/pkg/filters"
)

type UserDao interface {
	Get(ctx context.Context, filters ...filters.FilterOption) (*user.User, error)
	Create(ctx context.Context, user *user.User) error
	Update(ctx context.Context, user *user.User) error
}
