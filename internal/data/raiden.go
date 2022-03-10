package data

import (
	"context"

	v1 "fxkt.tech/raiden/api/raiden/v1"
	"fxkt.tech/raiden/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type messageSystemRepo struct {
	data *Data
	log  *log.Helper
}

// NewMessageSystemRepo .
func NewMessageSystemRepo(data *Data, logger log.Logger) biz.MessageSystemRepo {
	return &messageSystemRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *messageSystemRepo) SendMessage(ctx context.Context, g *biz.Message) error {
	err := r.data.db.User.WithContext(ctx).Create()
	if err != nil {
		return v1.ErrorUserNotFound(err.Error())
	}
	return nil
}
