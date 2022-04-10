package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type Dynamic struct {
	DmcId       int64
	DmcType     int32
	User        *User
	Text        string
	Imgs        []string
	PublishTime time.Time
}

type User struct {
	Uid  int32
	Nick string
}

type DynamicSearch struct {
	Uid      int32
	LastId   int64
	PageSize int32
}

type FeedSystemRepo interface {
	Publish(context.Context, *Dynamic) error
	Following(context.Context, *DynamicSearch) ([]*Dynamic, error)
	History(context.Context, *DynamicSearch) ([]*Dynamic, error)
}

type FeedSystemUsecase struct {
	repo FeedSystemRepo
	log  *log.Helper
}

func NewFeedSystemUsecase(repo FeedSystemRepo, logger log.Logger) *FeedSystemUsecase {
	return &FeedSystemUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *FeedSystemUsecase) Publish(ctx context.Context, d *Dynamic) error {
	return uc.repo.Publish(ctx, d)
}

func (uc *FeedSystemUsecase) Following(ctx context.Context, ds *DynamicSearch) ([]*Dynamic, error) {
	return uc.repo.Following(ctx, ds)
}

func (uc *FeedSystemUsecase) History(ctx context.Context, ds *DynamicSearch) ([]*Dynamic, error) {
	return uc.repo.History(ctx, ds)
}
