package gomysql_test

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {
	
}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:n3txt.vml1@tcp(localhost:3306)/learn_go_database")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("TestOpenConnection function done!")
}