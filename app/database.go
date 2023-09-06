package app

import (
	"database/sql"
	"time"

	"github.com/masrayfa/go-basic-rest/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/belajar_go_rest")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
