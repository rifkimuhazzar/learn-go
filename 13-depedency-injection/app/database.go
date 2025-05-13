package app

import (
	"database/sql"
	"time"

	"go_restful_api/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:n3txt.vml1@tcp(localhost:3306)/learn_go_restful_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(25)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}