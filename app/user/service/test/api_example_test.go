package test

import (
	"context"
	"fmt"
	"testing"

	v1 "fxkt.tech/raiden/api/raiden/v1"
	"fxkt.tech/raiden/pkg/json"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

var service, close = func() (v1.MessageSystemHTTPClient, func()) {
	client, err := http.NewClient(context.Background(),
		http.WithEndpoint("0.0.0.0:8000"),
		http.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return v1.NewMessageSystemHTTPClient(client), func() {
		client.Close()
	}
}()

func TestSendMessage(t *testing.T) {
	req := &v1.SendMessageRequest{
		SenderUid: 1,
		Content: &v1.Content{
			Text: "xx",
		},
		RecverUid: 2,
	}
	rsp, err := service.SendMessage(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(json.Pretty(json.ToBytes(rsp)))
	close()
}

func TestChatHistory(t *testing.T) {
	req := &v1.ChatHistoryRequest{
		MyUid:     1,
		FriendUid: 2,
		Page:      1,
		Count:     10,
	}
	rsp, err := service.ChatHistory(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(json.Pretty(json.ToBytes(rsp)))
	close()
}

// func TestRegisterUser(t *testing.T) {
// 	req := &v1.RegisterUserRequest{
// 		Nick: "Mei",
// 	}
// 	rsp, err := service.RegisterUser(context.Background(), req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Println(json.Pretty(json.ToBytes(rsp)))
// 	close()
// }

// func TestFriends(t *testing.T) {
// 	req := &v1.FriendsRequest{
// 		Uid:   1,
// 		Page:  1,
// 		Count: 10,
// 	}
// 	rsp, err := service.Friends(context.Background(), req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Println(json.Pretty(json.ToBytes(rsp)))
// 	close()
// }
