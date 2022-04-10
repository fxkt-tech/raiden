package service

import (
	"context"

	v1 "fxkt.tech/raiden/api/feed/service/v1"
	userv1 "fxkt.tech/raiden/api/user/service/v1"
	"fxkt.tech/raiden/app/feed/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type FeedSystemService struct {
	v1.UnimplementedFeedSystemServer

	uc  *biz.FeedSystemUsecase
	log *log.Helper
}

func NewFeedSystemService(uc *biz.FeedSystemUsecase, logger log.Logger) *FeedSystemService {
	return &FeedSystemService{uc: uc, log: log.NewHelper(logger)}
}

func (s *FeedSystemService) Publish(ctx context.Context, in *v1.PublishRequest) (*v1.PublishReply, error) {
	s.log.WithContext(ctx).Infof("Publish.Request: %s", in.String())

	dmc := &biz.Dynamic{
		User: &biz.User{
			Uid: in.Uid,
		},
		DmcType: in.DmcType,
		Text:    in.Text,
		Imgs:    in.Imgs,
	}
	err := s.uc.Publish(ctx, dmc)
	if err != nil {
		return nil, err
	}

	return &v1.PublishReply{DmcId: dmc.DmcId}, nil
}

func (s *FeedSystemService) Following(ctx context.Context, in *v1.FollowingRequest) (*v1.FollowingReply, error) {
	s.log.WithContext(ctx).Infof("Following.Request: %s", in.String())

	us := &biz.DynamicSearch{
		Uid:      in.Uid,
		LastId:   in.LastId,
		PageSize: in.PageSize,
	}
	dmcs, err := s.uc.Following(ctx, us)
	if err != nil {
		return nil, err
	}

	replyDmcs := make([]*v1.Dynamic, len(dmcs))
	for i, dmc := range dmcs {
		replyDmcs[i] = &v1.Dynamic{
			DmcId:   dmc.DmcId,
			DmcType: dmc.DmcType,
			User: &userv1.User{
				Uid: dmc.User.Uid,
			},
			Text: dmc.Text,
			Imgs: dmc.Imgs,
		}
	}

	return &v1.FollowingReply{
		Dmcs: replyDmcs,
	}, nil
}

func (s *FeedSystemService) History(ctx context.Context, in *v1.HistoryRequest) (*v1.HistoryReply, error) {
	s.log.WithContext(ctx).Infof("History.Request: %s", in.String())

	us := &biz.DynamicSearch{
		Uid:      in.Uid,
		LastId:   in.LastId,
		PageSize: in.PageSize,
	}
	dmcs, err := s.uc.History(ctx, us)
	if err != nil {
		return nil, err
	}

	replyDmcs := make([]*v1.Dynamic, len(dmcs))
	for i, dmc := range dmcs {
		replyDmcs[i] = &v1.Dynamic{
			DmcId:   dmc.DmcId,
			DmcType: dmc.DmcType,
			User: &userv1.User{
				Uid: dmc.User.Uid,
			},
			Text: dmc.Text,
			Imgs: dmc.Imgs,
		}
	}

	return &v1.HistoryReply{
		Dmcs: replyDmcs,
	}, nil
}
