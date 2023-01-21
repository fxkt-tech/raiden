package service

import (
	"context"

	mainsitev1 "github.com/fxkt-tech/raiden/api/mainsite/service/v1"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	mainsitev1.UnimplementedUserServiceServer

	uc  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserService) Register(ctx context.Context, in *mainsitev1.RegisterRequest) (*mainsitev1.RegisterResponse, error) {
	s.log.WithContext(ctx).Infof("Register.Request: %s", in.String())

	user := &biz.User{
		Nick: in.Nick,
	}
	err := s.uc.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	replyUser := &mainsitev1.User{
		Uid:  user.Uid,
		Nick: user.Nick,
	}

	return &mainsitev1.RegisterResponse{User: replyUser}, nil
}

func (s *UserService) Info(ctx context.Context, in *mainsitev1.InfoRequest) (*mainsitev1.InfoResponse, error) {
	s.log.WithContext(ctx).Infof("Info.Request: %s", in.String())

	user, err := s.uc.Info(ctx, in.Uid)
	if err != nil {
		return nil, err
	}

	replyUser := &mainsitev1.User{
		Uid:  user.Uid,
		Nick: user.Nick,
	}

	return &mainsitev1.InfoResponse{User: replyUser}, nil
}

func (s *UserService) Followers(ctx context.Context, in *mainsitev1.FollowersRequest) (*mainsitev1.FollowersResponse, error) {
	s.log.WithContext(ctx).Infof("Followers.Request: %s", in.String())

	ms := &biz.UserSearch{
		Uid:   in.Uid,
		Page:  in.PageNum,
		Count: in.PageSize,
	}
	users, total, err := s.uc.Followers(ctx, ms)
	if err != nil {
		return nil, err
	}

	replyUsers := make([]*mainsitev1.User, len(users))
	for i, user := range users {
		replyUsers[i] = &mainsitev1.User{
			Uid:  user.Uid,
			Nick: user.Nick,
		}
	}

	return &mainsitev1.FollowersResponse{
		Users: replyUsers,
		Total: total,
	}, nil
}

func (s *UserService) Following(ctx context.Context, in *mainsitev1.FollowingRequest) (*mainsitev1.FollowingResponse, error) {
	s.log.WithContext(ctx).Infof("Following.Request: %s", in.String())

	us := &biz.UserSearch{
		Uid:   in.Uid,
		Page:  in.PageNum,
		Count: in.PageSize,
	}
	users, total, err := s.uc.Following(ctx, us)
	if err != nil {
		return nil, err
	}

	replyUsers := make([]*mainsitev1.User, len(users))
	for i, user := range users {
		replyUsers[i] = &mainsitev1.User{
			Uid:  user.Uid,
			Nick: user.Nick,
		}
	}

	return &mainsitev1.FollowingResponse{
		Users: replyUsers,
		Total: total,
	}, nil
}

func (s *UserService) Relation(ctx context.Context, in *mainsitev1.RelationRequest) (*mainsitev1.RelationResponse, error) {
	s.log.WithContext(ctx).Infof("Relation.Request: %s", in.String())

	ur := &biz.UserRelation{
		ActiveUid:  in.ActiveUid,
		PassiveUid: in.PassiveUid,
		Action:     in.Action,
	}
	err := s.uc.Relation(ctx, ur)
	if err != nil {
		return nil, err
	}

	return &mainsitev1.RelationResponse{}, nil
}
