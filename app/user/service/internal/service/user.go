package service

import (
	"context"

	v1 "fxkt.tech/raiden/api/user/service/v1"
	"fxkt.tech/raiden/app/user/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type UserSystemService struct {
	v1.UnimplementedUserSystemServer

	uc  *biz.UserSystemUsecase
	log *log.Helper
}

func NewUserSystemService(uc *biz.UserSystemUsecase, logger log.Logger) *UserSystemService {
	return &UserSystemService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserSystemService) Register(ctx context.Context, in *v1.RegisterRequest) (*v1.RegisterReply, error) {
	s.log.WithContext(ctx).Infof("Register.Request: %s", in.String())

	user := &biz.User{
		Nick: in.Nick,
	}
	err := s.uc.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	replyUser := &v1.User{
		Uid:  user.Uid,
		Nick: user.Nick,
	}

	return &v1.RegisterReply{User: replyUser}, nil
}

func (s *UserSystemService) Info(ctx context.Context, in *v1.InfoRequest) (*v1.InfoReply, error) {
	s.log.WithContext(ctx).Infof("Info.Request: %s", in.String())

	user, err := s.uc.Info(ctx, in.Uid)
	if err != nil {
		return nil, err
	}

	replyUser := &v1.User{
		Uid:  user.Uid,
		Nick: user.Nick,
	}

	return &v1.InfoReply{User: replyUser}, nil
}

func (s *UserSystemService) Followers(ctx context.Context, in *v1.FollowersRequest) (*v1.FollowersReply, error) {
	s.log.WithContext(ctx).Infof("Followers.Request: %s", in.String())

	ms := &biz.UserSearch{
		Uid:   in.Uid,
		Page:  in.Page,
		Count: in.Count,
	}
	users, total, err := s.uc.Followers(ctx, ms)
	if err != nil {
		return nil, err
	}

	replyUsers := make([]*v1.User, len(users))
	for i, user := range users {
		replyUsers[i] = &v1.User{
			Uid:  user.Uid,
			Nick: user.Nick,
		}
	}

	return &v1.FollowersReply{
		Users: replyUsers,
		Total: total,
	}, nil
}

func (s *UserSystemService) Following(ctx context.Context, in *v1.FollowingRequest) (*v1.FollowingReply, error) {
	s.log.WithContext(ctx).Infof("Following.Request: %s", in.String())

	us := &biz.UserSearch{
		Uid:   in.Uid,
		Page:  in.Page,
		Count: in.Count,
	}
	users, total, err := s.uc.Following(ctx, us)
	if err != nil {
		return nil, err
	}

	replyUsers := make([]*v1.User, len(users))
	for i, user := range users {
		replyUsers[i] = &v1.User{
			Uid:  user.Uid,
			Nick: user.Nick,
		}
	}

	return &v1.FollowingReply{
		Users: replyUsers,
		Total: total,
	}, nil
}

func (s *UserSystemService) Relation(ctx context.Context, in *v1.RelationRequest) (*v1.RelationReply, error) {
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

	return &v1.RelationReply{}, nil
}
