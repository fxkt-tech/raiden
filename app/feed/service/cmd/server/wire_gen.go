// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"fxkt.tech/raiden/app/feed/service/internal/biz"
	"fxkt.tech/raiden/app/feed/service/internal/conf"
	"fxkt.tech/raiden/app/feed/service/internal/data"
	"fxkt.tech/raiden/app/feed/service/internal/server"
	"fxkt.tech/raiden/app/feed/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	feedSystemRepo := data.NewfeedSystemRepo(dataData, logger)
	feedSystemUsecase := biz.NewFeedSystemUsecase(feedSystemRepo, logger)
	feedSystemService := service.NewFeedSystemService(feedSystemUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, feedSystemService, logger)
	grpcServer := server.NewGRPCServer(confServer, feedSystemService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}