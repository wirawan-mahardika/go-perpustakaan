package app

import (
	"database/sql"
	"perpustakaan/helper"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:wm050604@tcp(localhost:3306)/perpustakaan?parseTime=true")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(time.Minute * 5)
	db.SetConnMaxLifetime(time.Minute * 25)

	return db
}
