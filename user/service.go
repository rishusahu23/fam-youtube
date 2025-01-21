package user

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rishusahu23/fam-youtube/gen/api/rpc"
	userPb "github.com/rishusahu23/fam-youtube/gen/api/user"
	customerrors "github.com/rishusahu23/fam-youtube/pkg/errors"
	txn "github.com/rishusahu23/fam-youtube/pkg/transaction"
	"github.com/rishusahu23/fam-youtube/user/dao"
	"github.com/rishusahu23/fam-youtube/user/dao/mongo"
	strategy "github.com/rishusahu23/fam-youtube/user/getuserstrategy"
)

type Service struct {
	dao        dao.UserDao
	txnManager txn.TransactionManager
	userPb.UnimplementedUserServiceServer
	factory strategy.GetUserStrategyFactory
}

func NewService(dao dao.UserDao, txnManager txn.TransactionManager, factory strategy.GetUserStrategyFactory) *Service {
	return &Service{
		dao:        dao,
		txnManager: txnManager,
		factory:    factory,
	}
}

func (s *Service) GetUser(ctx context.Context, req *userPb.GetUserRequest) (*userPb.GetUserResponse, error) {
	strategyImpl, err := s.factory.GetStrategy(ctx, "db")
	if err != nil {
		return &userPb.GetUserResponse{
			Status: rpc.StatusInternal(err.Error()),
		}, nil
	}
	resp, err := strategyImpl.GetUser(ctx, &strategy.GetUserRequest{
		UserId: req.GetUserId(),
	})
	if err != nil {
		return &userPb.GetUserResponse{
			Status: rpc.StatusInternal(err.Error()),
		}, nil
	}
	return &userPb.GetUserResponse{
		Status: rpc.StatusOk(),
		User:   resp.User,
	}, nil
}

func (s *Service) CreateUser(ctx context.Context, req *userPb.CreateUserRequest) (*userPb.CreateUserResponse, error) {
	_, err := s.dao.Get(ctx, mongo.WithUserId(req.GetUser().GetId()))
	if err != nil && !errors.Is(err, customerrors.ErrRecordNotFound) {
		return &userPb.CreateUserResponse{
			Status: rpc.StatusInternal(""),
		}, nil
	}

	err = s.dao.Create(ctx, req.GetUser())

	if err != nil {
		return &userPb.CreateUserResponse{
			Status: rpc.StatusInternal(""),
		}, nil
	}
	return &userPb.CreateUserResponse{
		Status: rpc.StatusOk(),
	}, nil
}

func (s *Service) GetPost(ctx context.Context, req *userPb.GetPostRequest) (*userPb.GetPostResponse, error) {
	return &userPb.GetPostResponse{
		Status: rpc.StatusOk(),
	}, nil
}

var _ userPb.UserServiceServer = (*Service)(nil)
