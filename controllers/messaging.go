package controllers

import (
	"encoding/json"
	"hungrydev39/todo-list-challenge/utilities"
	"net/http"
    "time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func AddMessage(w http.ResponseWriter, r *http.Request) {
	type RBody struct {
		Title string `json:"title"`
		Email string `json:"email"`
	}
	var requestBody RBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if requestBody.Title == "" {
		json, _ := json.Marshal(FailedResponse{
			"Bad Request",
			"title cannot be null",
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)
		return
	}
    
    data, _ := json.Marshal(map[string]string {
        "title": requestBody.Title,
        "email": requestBody.Email,
    })

    utilities.PublishMessageToMqQueue("activities", amqp.Publishing{
        ContentType: "application/json",
        Body: data,
    })

	type DataActivity struct {
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Id        int64     `json:"id"`
		Title     string    `json:"title"`
		Email     string    `json:"email"`
	}

	toJson := ResponseActivity[interface{}]{
		"Success",
		"Success",
        map[string]string { "Success": "yes" },
	}

	response, err := json.Marshal(toJson)
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
