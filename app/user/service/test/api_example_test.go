package test

import (
	"context"
	"fmt"
	"testing"

	v1 "fxkt.tech/raiden/api/user/service/v1"
	"fxkt.tech/raiden/pkg/json"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

var service, close = func() (v1.UserSystemHTTPClient, func()) {
	client, err := http.NewClient(context.Background(),
		http.WithEndpoint("0.0.0.0:8000"),
		http.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return v1.NewUserSystemHTTPClient(client), func() {
		client.Close()
	}
}()

func TestRegister(t *testing.T) {
	req := &v1.RegisterRequest{
		Nick: "Myzh",
	}
	rsp, err := service.Register(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(json.Pretty(json.ToBytes(rsp)))
	close()
}

func TestFollowers(t *testing.T) {
	req := &v1.FollowersRequest{
		Uid:   1,
		Page:  1,
		Count: 10,
	}
	rsp, err := service.Followers(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(json.Pretty(json.ToBytes(rsp)))
	close()
}
func TestFollowing(t *testing.T) {
	req := &v1.FollowingRequest{
		Uid:   1,
		Page:  1,
		Count: 10,
	}
	rsp, err := service.Following(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(json.Pretty(json.ToBytes(rsp)))
	close()
}
func TestRelation(t *testing.T) {
	req := &v1.RelationRequest{
		ActiveUid:  1,
		PassiveUid: 2,
		Action:     1,
	}
	rsp, err := service.Relation(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(json.Pretty(json.ToBytes(rsp)))
	close()
}
