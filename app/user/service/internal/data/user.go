package data

import (
	"context"

	"fxkt.tech/raiden/app/user/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type userSystemRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserSystemRepo(data *Data, logger log.Logger) biz.UserSystemRepo {
	return &userSystemRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userSystemRepo) Register(ctx context.Context, u *biz.User) error {
	return nil
}
func (r *userSystemRepo) Followers(ctx context.Context, us *biz.UserSearch) ([]*biz.User, error) {
	return nil, nil
}
func (r *userSystemRepo) Following(ctx context.Context, us *biz.UserSearch) ([]*biz.User, error) {
	return nil, nil
}
func (r *userSystemRepo) Relation(ctx context.Context, ur *biz.UserRelation) error {
	return nil
}
