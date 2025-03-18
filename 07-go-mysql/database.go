package gomysql

import (
	"database/sql"
	"time"
)

func GetMysqlConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:n3txt.vml1@tcp(localhost:3306)/learn_go_database?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}