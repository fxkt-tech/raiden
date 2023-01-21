package server

import (
	mainsitev1 "github.com/fxkt-tech/raiden/api/mainsite/service/v1"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/conf"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/service"
	"github.com/fxkt-tech/raiden/pkg/server"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, impl *service.MessageService, logger log.Logger) *http.Server {
	opts := []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	server.RegisterPprof(srv)
	mainsitev1.RegisterMessageServiceHTTPServer(srv, impl)
	return srv
}
