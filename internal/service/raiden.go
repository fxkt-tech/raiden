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

	msg := &biz.Message{}
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
		replymsgs[i] = &v1.Message{
			Content: &v1.Content{
				Text: message.Body,
			},
		}
	}

	return &v1.ChatHistoryReply{Msgs: replymsgs}, nil
}
