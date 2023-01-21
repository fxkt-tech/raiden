package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type Moment struct {
	MomentId    int64
	MomentType  int32
	User        *User
	Text        string
	Imgs        []string
	PublishTime time.Time
}

type MomentsSearch struct {
	Uid      int32
	LastId   int64
	PageSize int32
}

type FeedRepo interface {
	// 发布动态
	Publish(context.Context, *Moment) error
	// 关注的人的朋友圈动态
	Moments(context.Context, *MomentsSearch) ([]*Moment, error)
	// 我发表的朋友圈动态
	History(context.Context, *MomentsSearch) ([]*Moment, error)
}

type FeedUsecase struct {
	repo FeedRepo
	log  *log.Helper
}

func NewFeedUsecase(repo FeedRepo, logger log.Logger) *FeedUsecase {
	return &FeedUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *FeedUsecase) Publish(ctx context.Context, d *Moment) error {
	return uc.repo.Publish(ctx, d)
}

func (uc *FeedUsecase) Moments(ctx context.Context, ds *MomentsSearch) ([]*Moment, error) {
	return uc.repo.Moments(ctx, ds)
}

func (uc *FeedUsecase) History(ctx context.Context, ds *MomentsSearch) ([]*Moment, error) {
	return uc.repo.History(ctx, ds)
}
