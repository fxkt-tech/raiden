//go:build wireinject
// +build wireinject

package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./query",
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})
	db, _ := gorm.Open(mysql.Open("root:qingchuan495@(127.0.0.1:3306)/db_raiden?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(db)
	g.ApplyBasic(
		g.GenerateModel("dynamic_following"),
		g.GenerateModel("dynamic_history"),
	)
	g.Execute()
}
