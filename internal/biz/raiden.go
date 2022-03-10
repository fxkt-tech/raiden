package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Message struct {
	SendUid uint32
	Body    string
	RecvUid uint32
}

type MessageSearch struct {
	MyUid     uint32
	FriendUid uint32
	Page      int32
	Count     int32
}

type Friend struct {
	Uid  uint32
	Nick string
}

type FriendSearch struct {
	Uid   uint32
	Page  int32
	Count int32
}

type MessageSystemRepo interface {
	SendMessage(context.Context, *Message) error
	ChatHistory(context.Context, *MessageSearch) ([]*Message, error)
	FriendList(context.Context, *FriendSearch) ([]*Friend, error)
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

func (uc *MessageSystemUsecase) ChatHistory(ctx context.Context, ms *MessageSearch) ([]*Message, error) {
	return uc.repo.ChatHistory(ctx, ms)
}

func (uc *MessageSystemUsecase) FriendList(ctx context.Context, fs *FriendSearch) ([]*Friend, error) {
	return uc.repo.FriendList(ctx, fs)
}
