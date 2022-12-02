package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(struct {
		Text string `json:"text"`
	}{
		"Hello World",
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(data)
}

func main() {
	log.Println("Server started at port 8081")
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8081", nil)
}
