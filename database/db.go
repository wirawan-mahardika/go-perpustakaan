package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDb() *sql.DB {
	db, err := sql.Open("mysql", "root:wm050604@tcp(localhost:3306)/perpustakaan")
	if err != nil { panic(err) }

	db.SetConnMaxIdleTime(time.Minute * 5)
	db.SetConnMaxLifetime(time.Hour)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)

	return db
}