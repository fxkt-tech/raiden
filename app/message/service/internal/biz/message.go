package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Message struct {
	SenderUid int32
	RecverUid int32
	Content   *Content
	CTime     int64 // create time
}

type Content struct {
	Text string
}

type MessageSearch struct {
	SenderUid int32
	RecverUid int32
	Page      int32
	Count     int32
}

type MessageSystemRepo interface {
	SendMessage(context.Context, *Message) error
	MessageHistory(context.Context, *MessageSearch) ([]*Message, error)
	RecallMessage(context.Context, *MessageSearch) error
}

type MessageSystemUsecase struct {
	repo MessageSystemRepo
	log  *log.Helper
}

func NewMessageSystemUsecase(repo MessageSystemRepo, logger log.Logger) *MessageSystemUsecase {
	return &MessageSystemUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *MessageSystemUsecase) SendMessage(ctx context.Context, m *Message) error {
	return uc.repo.SendMessage(ctx, m)
}

func (uc *MessageSystemUsecase) MessageHistory(ctx context.Context, ms *MessageSearch) ([]*Message, error) {
	return uc.repo.MessageHistory(ctx, ms)
}

func (uc *MessageSystemUsecase) RecallMessage(ctx context.Context, ms *MessageSearch) error {
	return uc.repo.RecallMessage(ctx, ms)
}
