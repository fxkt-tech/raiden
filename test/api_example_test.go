package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	v1 "fxkt.tech/raiden/api/raiden/v1"
	"fxkt.tech/raiden/pkg/json"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

var ctx, service, close = func() (context.Context, v1.MessageSystemHTTPClient, func()) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	client, err := http.NewClient(ctx,
		http.WithEndpoint("0.0.0.0:8000"),
		http.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return ctx, v1.NewMessageSystemHTTPClient(client), func() {
		cancel()
		client.Close()
	}
}()

func TestSendMessage(t *testing.T) {
	req := &v1.SendMessageRequest{
		SenderUid: 1,
		Msg:       "xx",
		RecvUid:   2,
	}
	rsp, err := service.SendMessage(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(json.Pretty(json.ToBytes(rsp)))
	close()
}
