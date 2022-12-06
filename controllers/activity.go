package controllers

import (
	"encoding/json"
	"hungrydev39/todo-list-challenge/database"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type FailedResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func AllActivity(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDB()
	defer db.Close()

	activities := database.FindAllActivity(db)
	type ResponseActivity struct {
		Status  string              `json:"status"`
		Message string              `json:"message"`
		Data    []database.Activity `json:"data"`
	}

	json, err := json.Marshal(ResponseActivity{
		"Success",
		"Success",
		activities,
	})

	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func ShowActivity(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDB()
	defer db.Close()

	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		panic(err.Error())
	}
	activity, err := database.FindByActivityId(db, id)
	if err != nil && err.Error() == "No Record Found" {
		json, _ := json.Marshal(FailedResponse{
			"Not Found",
			"Activity with ID " + idStr + " Not Found",
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(json)
		return
	}
	type ResponseActivity struct {
		Status  string            `json:"status"`
		Message string            `json:"message"`
		Data    database.Activity `json:"data"`
	}

	toJson := ResponseActivity{
		"Success",
		"Success",
		activity,
	}

	response, err := json.Marshal(toJson)

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func AddActivity(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDB()
	defer db.Close()
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

	lastId := database.AddActivity(db, requestBody.Title, requestBody.Email)

	activity, err := database.FindByActivityId(db, lastId)

	type DataActivity struct {
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Id        int64     `json:"id"`
		Title     string    `json:"title"`
		Email     string    `json:"email"`
	}

	type ResponseActivity struct {
		Status  string       `json:"status"`
		Message string       `json:"message"`
		Data    DataActivity `json:"data"`
	}

	data := DataActivity{
		activity.CreatedAt,
		activity.UpdatedAt,
		activity.Id,
		activity.Title,
		activity.Email,
	}

	toJson := ResponseActivity{
		"Success",
		"Success",
		data,
	}

	response, err := json.Marshal(toJson)
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func sendJson(w http.ResponseWriter, v any) {
	response, err := json.Marshal(v)
	if err != nil {
		panic(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func UpdateActivity(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDB()
	defer db.Close()

	idStr := chi.URLParam(r, "id")
	type RBody struct {
		Title *string `json:"title,omitempty"`
	}
	var requestBody RBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		panic(err.Error())
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		panic(err.Error())
	}

	errTrans := database.UpdateActivityById(db, id, *requestBody.Title)
	if errTrans != nil {
		json, _ := json.Marshal(FailedResponse{
			"Not Found",
			"Activity with ID " + idStr + " Not Found",
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(json)
		return
	}

	type ResponseActivity struct {
		Status  string            `json:"status"`
		Message string            `json:"message"`
		Data    database.Activity `json:"data"`
	}

	data, err := database.FindByActivityId(db, id)

	response := ResponseActivity{
		"Success",
		"Success",
		data,
	}
	json, err := json.Marshal(response)
	if err != nil {
		panic(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func DeleteActivity(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDB()
	defer db.Close()

	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		panic(err.Error())
	}

	errDb := database.DeleteActivityById(db, id)
	if errDb != nil && errDb.Error() == "No Records Found" {
		json, _ := json.Marshal(FailedResponse{
			"Not Found",
			"Activity with ID " + idStr + " Not Found",
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(json)
		return
	}

	type ResponseActivity struct {
		Status  string   `json:"status"`
		Message string   `json:"message"`
		Data    struct{} `json:"data"`
	}

	toJson := ResponseActivity{
		"Success",
		"Success",
		struct{}{},
	}
	sendJson(w, toJson)
}
