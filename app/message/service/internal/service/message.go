package service

import (
	"context"

	v1 "fxkt.tech/raiden/api/message/service/v1"
	"fxkt.tech/raiden/app/message/service/internal/biz"
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
	s.log.WithContext(ctx).Infof("SendMessage.Request: %s", in.String())

	msg := &biz.Message{
		SenderUid: in.SenderUid,
		RecverUid: in.RecverUid,
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

func (s *MessageSystemService) MessageHistory(ctx context.Context, in *v1.MessageHistoryRequest) (*v1.MessageHistoryReply, error) {
	s.log.WithContext(ctx).Infof("MessageHistory.Request: %s", in.String())

	ms := &biz.MessageSearch{
		SenderUid: in.SenderUid,
		RecverUid: in.RecverUid,
		Page:      in.Page,
		Count:     in.Count,
	}
	msgs, err := s.uc.MessageHistory(ctx, ms)
	if err != nil {
		return nil, err
	}

	replyMsgs := make([]*v1.Message, len(msgs))
	for i, msg := range msgs {
		var content *v1.Content
		if msg.Content != nil {
			content = &v1.Content{
				Text: msg.Content.Text,
			}
		}
		replyMsgs[i] = &v1.Message{
			SenderUid: msg.SenderUid,
			RecverUid: msg.RecverUid,
			Content:   content,
			Ctime:     msg.CTime,
		}
	}

	return &v1.MessageHistoryReply{Msgs: replyMsgs}, nil
}
