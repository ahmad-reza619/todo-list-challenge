package database

import (
	"database/sql"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	dbUrl := os.ExpandEnv("${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DBNAME}?parseTime=true")
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)

	return db
}
func MigrateDb() {
	path := filepath.Join("database", "setup.sql")
	c, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		panic(ioErr.Error())
	}
	setup := strings.Split(string(c), ";")
	db := ConnectDB()
	for _, set := range setup {
		_, errDb := db.Exec(set)
		if errDb != nil {
			panic(errDb.Error())
		}
	}
	db.Close()
}
