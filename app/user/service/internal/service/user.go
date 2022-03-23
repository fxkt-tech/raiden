package service

import (
	"context"

	v1 "fxkt.tech/raiden/api/user/service/v1"
	"fxkt.tech/raiden/app/user/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type MessageSystemService struct {
	v1.UnimplementedUserSystemServer

	uc  *biz.MessageSystemUsecase
	log *log.Helper
}

func NewMessageSystemService(uc *biz.MessageSystemUsecase, logger log.Logger) *MessageSystemService {
	return &MessageSystemService{uc: uc, log: log.NewHelper(logger)}
}

func (s *MessageSystemService) Register(ctx context.Context, in *v1.RegisterRequest) (*v1.RegisterReply, error) {
	s.log.WithContext(ctx).Infof("SendMessage.Request: %s", in.String())

	msg := &biz.User{
		Nick: in.Nick,
	}
	err := s.uc.Register(ctx, msg)
	if err != nil {
		return nil, err
	}

	return &v1.RegisterReply{}, nil
}

func (s *MessageSystemService) Followers(ctx context.Context, in *v1.FollowersRequest) (*v1.FollowersReply, error) {
	s.log.WithContext(ctx).Infof("Followers.Request: %s", in.String())

	ms := &biz.UserSearch{
		Uid:   in.Uid,
		Page:  in.Page,
		Count: in.Count,
	}
	users, err := s.uc.Followers(ctx, ms)
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

	return &v1.FollowersReply{Users: replyUsers}, nil
}

func (s *MessageSystemService) Following(ctx context.Context, in *v1.FollowingRequest) (*v1.FollowingReply, error) {
	s.log.WithContext(ctx).Infof("Following.Request: %s", in.String())

	us := &biz.UserSearch{
		Uid:   in.Uid,
		Page:  in.Page,
		Count: in.Count,
	}
	users, err := s.uc.Following(ctx, us)
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

	return &v1.FollowingReply{Users: replyUsers}, nil
}

func (s *MessageSystemService) Relation(ctx context.Context, in *v1.RelationRequest) (*v1.RelationReply, error) {
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
