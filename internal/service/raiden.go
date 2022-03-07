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
