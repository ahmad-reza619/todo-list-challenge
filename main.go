package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Server started at port 8081")
	r := setupRouter()

	http.ListenAndServe(":8081", r)
}
