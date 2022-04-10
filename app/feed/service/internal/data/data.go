package data

import (
	"context"
	"os"
	"time"

	v1 "fxkt.tech/raiden/api/user/service/v1"
	"fxkt.tech/raiden/app/feed/service/internal/conf"
	"fxkt.tech/raiden/app/feed/service/internal/data/db/query"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewfeedSystemRepo)

type Data struct {
	db *query.Query

	userService v1.UserSystemHTTPClient
}

func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	var logMode gormlogger.LogLevel
	if os.Getenv("RAIDEN_LOG_LEVEL") == "DEBUG" {
		logMode = gormlogger.Info
	} else {
		logMode = gormlogger.Silent
	}
	dbclient, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		Logger: gormlogger.Default.LogMode(logMode),
	})
	if err != nil {
		return nil, nil, err
	}
	db := query.Use(dbclient)

	// ext_api: lms client
	userClient, err := http.NewClient(context.Background(),
		http.WithEndpoint(c.ExtApi.UserEp),
		http.WithTimeout(1*time.Second),
		http.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		return nil, nil, err
	}
	userService := v1.NewUserSystemHTTPClient(userClient)

	return &Data{
			db:          db,
			userService: userService,
		}, func() {
			log.NewHelper(logger).Info("closing the data resources")
		}, nil
}
