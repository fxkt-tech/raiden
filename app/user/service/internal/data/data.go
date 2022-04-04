package data

import (
	"os"

	"fxkt.tech/raiden/app/user/service/internal/conf"
	"fxkt.tech/raiden/app/user/service/internal/data/db/query"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserSystemRepo)

type Data struct {
	db *query.Query
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

	return &Data{
			db: query.Use(dbclient),
		}, func() {
			log.NewHelper(logger).Info("closing the data resources")
		}, nil
}
