package data

import (
	"context"

	"fxkt.tech/raiden/app/message/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type messageSystemRepo struct {
	data *Data
	log  *log.Helper
}

func NewMessageSystemRepo(data *Data, logger log.Logger) biz.MessageSystemRepo {
	return &messageSystemRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *messageSystemRepo) SendMessage(ctx context.Context, m *biz.Message) error {
	return nil
}
func (r *messageSystemRepo) ChatHistory(ctx context.Context, ms *biz.MessageSearch) ([]*biz.Message, error) {
	return nil, nil
}
