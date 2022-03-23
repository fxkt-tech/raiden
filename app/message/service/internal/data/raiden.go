package data

import (
	"context"

	v1 "fxkt.tech/raiden/api/raiden/v1"
	"fxkt.tech/raiden/internal/biz"
	"fxkt.tech/raiden/internal/data/db/model"
	"fxkt.tech/raiden/pkg/json"
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

func (r *messageSystemRepo) SendMessage(ctx context.Context, msg *biz.Message) error {
	po := &model.Message{
		SenderUID: msg.SendUid,
		RecverUID: msg.RecvUid,
		Content: json.ToString(&biz.Content{
			Text: msg.Content.Text,
		}),
	}
	err := r.data.db.Message.WithContext(ctx).Create(po)
	if err != nil {
		return v1.ErrorDatabase(err.Error())
	}
	return nil
}

func (r *messageSystemRepo) ChatHistory(ctx context.Context, ms *biz.MessageSearch) ([]*biz.Message, error) {
	var msgs []*model.Message
	msgs, msglen, err := r.data.db.Message.WithContext(ctx).
		Where(r.data.db.Message.SenderUID.Eq(ms.MyUid)).
		Where(r.data.db.Message.RecverUID.Eq(ms.FriendUid)).
		FindByPage(int((ms.Page-1)*ms.Count), int(ms.Count))
	if err != nil {
		return nil, v1.ErrorDatabase(err.Error())
	}

	retmsgs := make([]*biz.Message, msglen)
	for i, msg := range msgs {
		var content *biz.Content
		if msg.Content != "" {
			content = &biz.Content{}
			json.ToObject(msg.Content, content)
		}
		retmsgs[i] = &biz.Message{
			SendUid: msg.SenderUID,
			RecvUid: msg.RecverUID,
			Content: content,
		}
	}
	return retmsgs, nil
}

func (r *messageSystemRepo) FriendList(ctx context.Context, fs *biz.UserSearch) ([]*biz.User, error) {
	users, userlen, err := r.data.db.User.WithContext(ctx).
		Where(r.data.db.User.UID.Eq(fs.Uid)).
		FindByPage(int((fs.Page-1)*fs.Count), int(fs.Count))
	if err != nil {
		return nil, v1.ErrorDatabase(err.Error())
	}
	friends := make([]*biz.User, userlen)
	for i, user := range users {
		friends[i] = &biz.User{
			Uid:  user.UID,
			Nick: user.Nick,
		}
	}
	return friends, nil
}

func (r *messageSystemRepo) RegisterUser(ctx context.Context, f *biz.User) error {
	po := &model.User{
		Nick: f.Nick,
	}
	err := r.data.db.User.WithContext(ctx).Create(po)
	if err != nil {
		return v1.ErrorDatabase(err.Error())
	}
	f.Uid = po.UID
	return nil
}
