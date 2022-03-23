package data

import (
	"fxkt.tech/raiden/app/message/service/internal/conf"
	"fxkt.tech/raiden/app/message/service/internal/data/db/query"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMessageSystemRepo)

type Data struct {
	db *query.Query
}

func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	dbclient, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
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
