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
	Info(context.Context, int32) (*User, error)
	Followers(context.Context, *UserSearch) ([]*User, int32, error)
	Following(context.Context, *UserSearch) ([]*User, int32, error)
	Relation(context.Context, *UserRelation) error
}

type UserSystemUsecase struct {
	repo UserSystemRepo
	log  *log.Helper
}

func NewUserSystemUsecase(repo UserSystemRepo, logger log.Logger) *UserSystemUsecase {
	return &UserSystemUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserSystemUsecase) Register(ctx context.Context, u *User) error {
	return uc.repo.Register(ctx, u)
}

func (uc *UserSystemUsecase) Info(ctx context.Context, uid int32) (*User, error) {
	return uc.repo.Info(ctx, uid)
}

func (uc *UserSystemUsecase) Followers(ctx context.Context, us *UserSearch) ([]*User, int32, error) {
	return uc.repo.Followers(ctx, us)
}

func (uc *UserSystemUsecase) Following(ctx context.Context, us *UserSearch) ([]*User, int32, error) {
	return uc.repo.Following(ctx, us)
}

func (uc *UserSystemUsecase) Relation(ctx context.Context, ur *UserRelation) error {
	return uc.repo.Relation(ctx, ur)
}
