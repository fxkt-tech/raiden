package server

import (
	"net/http/pprof"

	v1 "fxkt.tech/raiden/api/raiden/v1"
	"fxkt.tech/raiden/internal/conf"
	"fxkt.tech/raiden/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, impl *service.MessageSystemService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
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
	RegisterPprof(srv)
	v1.RegisterMessageSystemHTTPServer(srv, impl)
	return srv
}

func RegisterPprof(srv *http.Server) {
	srv.HandleFunc("/debug/pprof", pprof.Index)
	srv.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	srv.HandleFunc("/debug/pprof/profile", pprof.Profile)
	srv.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	srv.HandleFunc("/debug/pprof/trace", pprof.Trace)
	srv.HandleFunc("/debug/allocs", pprof.Handler("allocs").ServeHTTP)
	srv.HandleFunc("/debug/block", pprof.Handler("block").ServeHTTP)
	srv.HandleFunc("/debug/goroutine", pprof.Handler("goroutine").ServeHTTP)
	srv.HandleFunc("/debug/heap", pprof.Handler("heap").ServeHTTP)
	srv.HandleFunc("/debug/mutex", pprof.Handler("mutex").ServeHTTP)
	srv.HandleFunc("/debug/threadcreate", pprof.Handler("threadcreate").ServeHTTP)
}
