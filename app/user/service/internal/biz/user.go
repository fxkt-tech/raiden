package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Uid  int32
	Nick string
}

type UserSearch struct {
	Uid   int32
	Page  int32
	Count int32
}

type UserRelation struct {
	ActiveUid  int32
	PassiveUid int32
	Action     int32
}

type UserSystemRepo interface {
	Register(context.Context, *User) error
	Followers(context.Context, *UserSearch) ([]*User, error)
	Following(context.Context, *UserSearch) ([]*User, error)
	Relation(context.Context, *UserRelation) error
}

type MessageSystemUsecase struct {
	repo UserSystemRepo
	log  *log.Helper
}

func NewMessageSystemUsecase(repo UserSystemRepo, logger log.Logger) *MessageSystemUsecase {
	return &MessageSystemUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *MessageSystemUsecase) Register(ctx context.Context, u *User) error {
	return uc.repo.Register(ctx, u)
}

func (uc *MessageSystemUsecase) Followers(ctx context.Context, us *UserSearch) ([]*User, error) {
	return uc.repo.Followers(ctx, us)
}

func (uc *MessageSystemUsecase) Following(ctx context.Context, us *UserSearch) ([]*User, error) {
	return uc.repo.Following(ctx, us)
}

func (uc *MessageSystemUsecase) Relation(ctx context.Context, ur *UserRelation) error {
	return uc.repo.Relation(ctx, ur)
}
