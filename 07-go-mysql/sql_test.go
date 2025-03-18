package gomysql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetMysqlConnection()
	defer db.Close()

	ctx := context.Background()
	sql_script := "insert into customers (id, name) values ('preact', 'Preact')"
	result, err := db.ExecContext(ctx, sql_script)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println(result.RowsAffected())
	fmt.Println(result.LastInsertId())
	fmt.Println(err)
	fmt.Println("Success insert a new customer")
}

func TestQuerySql(t *testing.T)  {
	db := GetMysqlConnection()
	defer db.Close()

	ctx := context.Background()
	sql_script := "select * from customers"
	rows, err := db.QueryContext(ctx, sql_script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("id:", id)
		fmt.Println("name:", name)
	}

	fmt.Println(rows)
	fmt.Println(err)
	fmt.Println("Success get customers")
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetMysqlConnection()
	defer db.Close()

	ctx := context.Background()
	sql_script := "select * from customers"
	rows, err := db.QueryContext(ctx, sql_script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name, email sql.NullString
		var balance int32
		var rating float64
		var created_at time.Time
		var birth_date sql.NullTime
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &created_at, &birth_date, &married)
		if err != nil {
			panic(err)
		}
		fmt.Println("==============================================")
		if email.Valid {
			fmt.Println("id:", id.String)
			fmt.Println("name:", name.String)
			fmt.Println("email:", email.String)
		}
		fmt.Println("balance:", balance)
		fmt.Println("rating:", rating)
		fmt.Println("created_at:", created_at)
		if birth_date.Valid {
			fmt.Println("birth_date:", birth_date.Time)
		}
		fmt.Println("married:", married)
	}
}

func TestQueryInjection(t *testing.T) {
	db := GetMysqlConnection()
	defer db.Close()

	username := "admin'; #"
	password := "wrong"

	ctx := context.Background()
	sqlQuery := "select username from users where username = '" + username +
							"' and password = '" + password + "' limit 1"
	fmt.Println(sqlQuery)
	rows, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login success:", username)
	} else {
		fmt.Println("Login failed")
	}
}

func TestQueryInjectionSafe(t *testing.T) {
	db := GetMysqlConnection()
	defer db.Close()

	username := "admin"
	password := "admin"

	ctx := context.Background()
	sqlQuery := "select username from users where username = ? and password = ? limit 1"
	fmt.Println(sqlQuery)
	rows, err := db.QueryContext(ctx, sqlQuery, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login success:", username)
	} else {
		fmt.Println("Login failed")
	}
}

func TestExecSqlSafe(t *testing.T) {
	db := GetMysqlConnection()
	defer db.Close()

	a := "alpinejs; drop table users #"
	b := "Alpinejs"

	ctx := context.Background()
	sql_script := "insert into customers (id, name) values (?, ?)"
	result, err := db.ExecContext(ctx, sql_script, a, b)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println(result.RowsAffected())
	fmt.Println(result.LastInsertId())
	fmt.Println(err)
	fmt.Println("Success insert a new customer")
}

func TestAutoIncrement(t *testing.T) {
	db := GetMysqlConnection()
	defer db.Close()

	comment := "Hell world"
	email := "react@example.cm"

	ctx := context.Background()
	sql_script := "insert into comments(comment, email) values(?, ?)"
	result, err := db.ExecContext(ctx, sql_script, comment, email)
	if err != nil {
		panic(err)
	}
	insertedId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
 
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
	fmt.Println("Success insert a new comment with id:", insertedId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetMysqlConnection()
	defer db.Close()

	ctx := context.Background()
	stmt, err := db.PrepareContext(ctx, "insert into comments(email, comment) values(?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := range 10 {
		email := "svelte" + strconv.Itoa(i + 1) + "@example.com"
		comment := "Comment - " + strconv.Itoa(i + 1)
		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println(result)
		fmt.Println(result.LastInsertId())
		fmt.Println("comment id:", id)
	}
}

func TestTransaction(t *testing.T)  {
	db := GetMysqlConnection()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	sql_script := "insert into comments(email, comment) values(?, ?)"
	for i := range 10 {
		email := "vue" + strconv.Itoa(i + 1) + "@example.com"
		comment := "Comment - " + strconv.Itoa(i + 1)
		result, err := tx.ExecContext(ctx, sql_script, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("comment id:", id)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}