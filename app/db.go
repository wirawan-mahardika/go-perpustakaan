package app

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() *gorm.DB {
	dialectMysql := mysql.Open("root:wm050604@tcp(localhost:3306)/perpustakaan?parseTime=true")
	db, err := gorm.Open(dialectMysql, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	return db
}
