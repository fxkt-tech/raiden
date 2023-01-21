package service

import (
	"context"

	mainsitev1 "github.com/fxkt-tech/raiden/api/mainsite/service/v1"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type FeedServiceService struct {
	mainsitev1.UnimplementedFeedServiceServer

	uc  *biz.FeedUsecase
	log *log.Helper
}

func NewFeedServiceService(uc *biz.FeedUsecase, logger log.Logger) *FeedServiceService {
	return &FeedServiceService{uc: uc, log: log.NewHelper(logger)}
}

func (s *FeedServiceService) Publish(ctx context.Context, in *mainsitev1.PublishRequest) (*mainsitev1.PublishResponse, error) {
	s.log.WithContext(ctx).Infof("Publish.Request: %s", in.String())

	dmc := &biz.Moment{
		User: &biz.User{
			Uid: in.Uid,
		},
		MomentType: in.MomentType,
		Text:       in.Text,
		Imgs:       in.Imgs,
	}
	err := s.uc.Publish(ctx, dmc)
	if err != nil {
		return nil, err
	}

	return &mainsitev1.PublishResponse{MomentId: dmc.MomentId}, nil
}

func (s *FeedServiceService) Moments(ctx context.Context, in *mainsitev1.MomentsRequest) (*mainsitev1.MomentsResponse, error) {
	s.log.WithContext(ctx).Infof("Following.Request: %s", in.String())

	us := &biz.MomentsSearch{
		Uid:      in.Uid,
		LastId:   in.LastId,
		PageSize: in.PageSize,
	}
	moments, err := s.uc.Moments(ctx, us)
	if err != nil {
		return nil, err
	}

	resMoments := make([]*mainsitev1.Moment, len(moments))
	for i, moment := range moments {
		resMoments[i] = &mainsitev1.Moment{
			MomentId:   moment.MomentId,
			MomentType: moment.MomentType,
			User: &mainsitev1.User{
				Uid:  moment.User.Uid,
				Nick: moment.User.Nick,
			},
			Text: moment.Text,
			Imgs: moment.Imgs,
		}
	}

	return &mainsitev1.MomentsResponse{
		Moments: resMoments,
	}, nil
}

func (s *FeedServiceService) History(ctx context.Context, in *mainsitev1.HistoryRequest) (*mainsitev1.HistoryResponse, error) {
	s.log.WithContext(ctx).Infof("History.Request: %s", in.String())

	us := &biz.MomentsSearch{
		Uid:      in.Uid,
		LastId:   in.LastId,
		PageSize: in.PageSize,
	}
	moments, err := s.uc.History(ctx, us)
	if err != nil {
		return nil, err
	}

	resMoments := make([]*mainsitev1.Moment, len(moments))
	for i, moment := range moments {
		resMoments[i] = &mainsitev1.Moment{
			MomentId:   moment.MomentId,
			MomentType: moment.MomentType,
			User: &mainsitev1.User{
				Uid: moment.User.Uid,
			},
			Text: moment.Text,
			Imgs: moment.Imgs,
		}
	}

	return &mainsitev1.HistoryResponse{
		Moments: resMoments,
	}, nil
}
