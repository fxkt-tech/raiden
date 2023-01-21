// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/biz"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/conf"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/data"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/server"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/service"
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
	messageRepo := data.NewMessageRepo(dataData, logger)
	messageUsecase := biz.NewMessageUsecase(messageRepo, logger)
	messageService := service.NewMessageService(messageUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, messageService, logger)
	grpcServer := server.NewGRPCServer(confServer, messageService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
