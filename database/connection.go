package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "user:password@/todo-list?parseTime=true")
	if err != nil {
		panic(err)
	}

	return db
}
