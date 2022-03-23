package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Message struct {
	SendUid int32
	Content *Content
	RecvUid int32
}

type Content struct {
	Text  string
	Image string
	Video string
	Emoji string
}

type MessageSearch struct {
	MyUid     int32
	FriendUid int32
	Page      int32
	Count     int32
}

type User struct {
	Uid  int32
	Nick string
}

type UserSearch struct {
	Uid   int32
	Page  int32
	Count int32
}

type MessageSystemRepo interface {
	// user
	//
	RegisterUser(context.Context, *User) error
	FriendList(context.Context, *UserSearch) ([]*User, error)

	// message
	//
	SendMessage(context.Context, *Message) error
	ChatHistory(context.Context, *MessageSearch) ([]*Message, error)
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

func (uc *MessageSystemUsecase) FriendList(ctx context.Context, fs *UserSearch) ([]*User, error) {
	return uc.repo.FriendList(ctx, fs)
}

func (uc *MessageSystemUsecase) RegisterUser(ctx context.Context, f *User) error {
	return uc.repo.RegisterUser(ctx, f)
}
