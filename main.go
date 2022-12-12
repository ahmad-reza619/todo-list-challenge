package main

import (
	"hungrydev39/todo-list-challenge/database"
	"log"
	"net/http"
)

func main() {
	database.MigrateDb()
	log.Println("Server started at port 8081")
	r := setupRouter()

	http.ListenAndServe(":3030", r)
}
