package server

import (
	v1 "fxkt.tech/raiden/api/feed/service/v1"
	"fxkt.tech/raiden/app/feed/service/internal/conf"
	"fxkt.tech/raiden/app/feed/service/internal/service"
	"fxkt.tech/raiden/pkg/server"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, impl *service.FeedSystemService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			validate.Validator(),
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
	v1.RegisterFeedSystemHTTPServer(srv, impl)
	return srv
}
