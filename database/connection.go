package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	dbUrl := os.ExpandEnv("${MYSQL_USER}:${MYSQL_PASSWORD}@${MYSQL_HOST}:${MYSQL_PORT}/${MYSQL_DBNAME}?parseTime=true")
	log.Println(dbUrl)
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err)
	}

	return db
}
