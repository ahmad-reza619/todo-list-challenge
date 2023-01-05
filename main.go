package main

import (
	"hungrydev39/todo-list-challenge/database"
	"hungrydev39/todo-list-challenge/router"
	"log"
	"net/http"
)

func main() {
	database.MigrateDb()
	log.Println("Server started at port 3030")
	r := router.SetupRouter()

	http.ListenAndServe(":3030", r)
}

