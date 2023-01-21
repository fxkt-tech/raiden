package service

import (
	"context"

	mainsitev1 "github.com/fxkt-tech/raiden/api/mainsite/service/v1"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type MessageService struct {
	mainsitev1.UnimplementedMessageServiceServer

	uc  *biz.MessageUsecase
	log *log.Helper
}

func NewMessageService(uc *biz.MessageUsecase, logger log.Logger) *MessageService {
	return &MessageService{uc: uc, log: log.NewHelper(logger)}
}

func (s *MessageService) SendMessage(ctx context.Context, in *mainsitev1.SendMessageRequest) (*mainsitev1.SendMessageResponse, error) {
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

	return &mainsitev1.SendMessageResponse{}, nil
}

func (s *MessageService) MessageHistory(ctx context.Context, in *mainsitev1.MessageHistoryRequest) (*mainsitev1.MessageHistoryResponse, error) {
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

	replyMsgs := make([]*mainsitev1.Message, len(msgs))
	for i, msg := range msgs {
		var content *mainsitev1.Content
		if msg.Content != nil {
			content = &mainsitev1.Content{
				Text: msg.Content.Text,
			}
		}
		replyMsgs[i] = &mainsitev1.Message{
			SenderUid: msg.SenderUid,
			RecverUid: msg.RecverUid,
			Content:   content,
			Ctime:     msg.CTime,
		}
	}

	return &mainsitev1.MessageHistoryResponse{Msgs: replyMsgs}, nil
}
