package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Message struct {
	SendUid int64
	Body    string
	RecvUid int64
}

type MessageSystemRepo interface {
	SendMessage(context.Context, *Message) error
}

type MessageSystemUsecase struct {
	repo MessageSystemRepo
	log  *log.Helper
}

func NewMessageSystemUsecase(repo MessageSystemRepo, logger log.Logger) *MessageSystemUsecase {
	return &MessageSystemUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *MessageSystemUsecase) SendMessage(ctx context.Context, msg *Message) error {
	return uc.repo.SendMessage(ctx, msg)
}
