package test

import (
	"context"
	"fmt"
	"testing"

	v1 "fxkt.tech/raiden/api/feed/service/v1"
	"fxkt.tech/raiden/pkg/json"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

var service, close = func() (v1.FeedSystemHTTPClient, func()) {
	client, err := http.NewClient(context.Background(),
		http.WithEndpoint("0.0.0.0:8100"),
		http.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return v1.NewFeedSystemHTTPClient(client), func() {
		client.Close()
	}
}()

func TestPublish(t *testing.T) {
	req := &v1.PublishRequest{
		Uid:     1,
		DmcType: 1,
		Text:    "xx me.",
	}
	rsp, err := service.Publish(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(json.Pretty(json.ToBytes(rsp)))
	close()
}

func TestFollowing(t *testing.T) {
	req := &v1.FollowingRequest{
		Uid:      1,
		PageSize: 20,
	}
	rsp, err := service.Following(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(json.Pretty(json.ToBytes(rsp)))
	close()
}
func TestHistory(t *testing.T) {
	req := &v1.HistoryRequest{
		Uid:      1,
		PageSize: 20,
	}
	rsp, err := service.History(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(json.Pretty(json.ToBytes(rsp)))
	close()
}
