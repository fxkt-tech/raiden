package data

import (
	"context"

	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type messageRepo struct {
	data *Data
	log  *log.Helper
}

func NewMessageRepo(data *Data, logger log.Logger) biz.MessageRepo {
	return &messageRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *messageRepo) SendMessage(ctx context.Context, m *biz.Message) error {
	return nil
}

func (r *messageRepo) MessageHistory(ctx context.Context, ms *biz.MessageSearch) ([]*biz.Message, error) {
	return nil, nil
}

func (r *messageRepo) RecallMessage(ctx context.Context, ms *biz.MessageSearch) error {
	return nil
}
