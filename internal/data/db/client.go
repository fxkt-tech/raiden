package db

import (
	"fxkt.tech/raiden/internal/data/db/query"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Client struct {
	db    *gorm.DB
	Query *query.Query
}

func MustNew(dsn string) *Client {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	// u.WithContext(context.Background()).Where(u.UID.In()).Create()
	return &Client{
		db:    db,
		Query: query.Use(db),
	}
}
