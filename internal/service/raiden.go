package service

import (
	"context"

	v1 "fxkt.tech/raiden/api/raiden/v1"
	"fxkt.tech/raiden/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type MessageSystemService struct {
	v1.UnimplementedMessageSystemServer

	uc  *biz.MessageSystemUsecase
	log *log.Helper
}

func NewMessageSystemService(uc *biz.MessageSystemUsecase, logger log.Logger) *MessageSystemService {
	return &MessageSystemService{uc: uc, log: log.NewHelper(logger)}
}

func (s *MessageSystemService) SendMessage(ctx context.Context, in *v1.SendMessageRequest) (*v1.SendMessageReply, error) {
	s.log.WithContext(ctx).Infof("SendMessage Received: %s", in.String())

	msg := &biz.Message{
		SendUid: in.SenderUid,
		RecvUid: in.RecverUid,
		Content: &biz.Content{
			Text: in.Content.Text,
		},
	}
	err := s.uc.SendMessage(ctx, msg)
	if err != nil {
		return nil, err
	}

	return &v1.SendMessageReply{}, nil
}

func (s *MessageSystemService) ChatHistory(ctx context.Context, in *v1.ChatHistoryRequest) (*v1.ChatHistoryReply, error) {
	s.log.WithContext(ctx).Infof("SendMessage Received: %s", in.String())

	ms := &biz.MessageSearch{
		MyUid:     in.MyUid,
		FriendUid: in.FriendUid,
		Page:      in.Page,
		Count:     in.Count,
	}
	messages, err := s.uc.ChatHistory(ctx, ms)
	if err != nil {
		return nil, err
	}

	replymsgs := make([]*v1.Message, len(messages))
	for i, message := range messages {
		var content *v1.Content
		if message.Content != nil {
			content = &v1.Content{
				Text: message.Content.Text,
			}
		}
		replymsgs[i] = &v1.Message{
			Sender: &v1.User{
				Uid: message.SendUid,
			},
			Recver: &v1.User{
				Uid: message.RecvUid,
			},
			Content: content,
		}
	}

	return &v1.ChatHistoryReply{Msgs: replymsgs}, nil
}

func (s *MessageSystemService) RegisterUser(ctx context.Context, in *v1.RegisterUserRequest) (*v1.RegisterUserReply, error) {
	s.log.WithContext(ctx).Infof("SendMessage Received: %s", in.String())

	user := &biz.User{
		Nick: in.Nick,
	}
	err := s.uc.RegisterUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &v1.RegisterUserReply{User: &v1.User{Uid: user.Uid, Nick: user.Nick}}, nil
}

func (s *MessageSystemService) Friends(ctx context.Context, in *v1.FriendsRequest) (*v1.FriendsReply, error) {
	s.log.WithContext(ctx).Infof("Friends Received: %s", in.String())

	us := &biz.UserSearch{
		Uid:   in.Uid,
		Page:  in.Page,
		Count: in.Count,
	}
	friends, err := s.uc.FriendList(ctx, us)
	if err != nil {
		return nil, err
	}

	replyfriends := make([]*v1.User, len(friends))
	for i, friend := range friends {
		replyfriends[i] = &v1.User{
			Uid:  friend.Uid,
			Nick: friend.Nick,
		}
	}

	return &v1.FriendsReply{Friends: replyfriends}, nil
}
