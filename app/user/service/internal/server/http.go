package server

import (
	v1 "fxkt.tech/raiden/api/user/service/v1"
	"fxkt.tech/raiden/app/user/service/internal/conf"
	"fxkt.tech/raiden/app/user/service/internal/service"
	"fxkt.tech/raiden/pkg/server"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, impl *service.UserSystemService, logger log.Logger) *http.Server {
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
	v1.RegisterUserSystemHTTPServer(srv, impl)
	return srv
}
