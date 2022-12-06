package controllers

import (
	"encoding/json"
	"hungrydev39/todo-list-challenge/database"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDB()
	defer db.Close()

	activity_group_id_str := r.URL.Query().Get("activity_group_id")
	activity_group_id, err := strconv.ParseInt(activity_group_id_str, 10, 64)

	var todos []database.Todo

	if err != nil {
		todos = database.FindAllTodo(db)
	} else {
		todos = database.FindTodoByActivityId(db, activity_group_id)
	}

	type Response struct {
		Status  string          `json:"status"`
		Message string          `json:"message"`
		Data    []database.Todo `json:"data"`
	}

	toJson, err := json.Marshal(Response{
		"Success",
		"Success",
		todos,
	})
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(toJson)
}

func ShowTodo(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDB()
	defer db.Close()

	idStr := chi.URLParam(r, "id")
	id, errParam := strconv.ParseInt(idStr, 10, 64)
	if errParam != nil {
		http.Error(w, "Id should be number", 400)
		return
	}

	todo, err := database.FindTodoById(db, id)
	if err != nil && err.Error() == "No Records Found" {
		json, _ := json.Marshal(FailedResponse{
			"Not Found",
			"Todo with ID " + idStr + " Not Found",
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(json)
		return
	}
	type Response struct {
		Status  string        `json:"status"`
		Message string        `json:"message"`
		Data    database.Todo `json:"data"`
	}

	toJson, errEncode := json.Marshal(Response{
		"Success",
		"Success",
		todo,
	})
	if errEncode != nil {
		http.Error(w, "Internal Server Error", 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(toJson)
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDB()
	defer db.Close()

	type RBody struct {
		ActivityGroupId int64  `json:"activity_group_id"`
		Title           string `json:"title"`
	}
	var requestBody RBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if requestBody.ActivityGroupId == int64(0) {
		json, _ := json.Marshal(FailedResponse{
			"Bad Request",
			"activity_group_id cannot be null",
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)
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

	lastId := database.AddTodo(db, requestBody.Title, requestBody.ActivityGroupId)
	todo, errDb := database.FindTodoById(db, lastId)
	if errDb != nil && errDb.Error() == "No Records Found" {
		json, _ := json.Marshal(FailedResponse{
			"Not Found",
			"Todo with ID " + strconv.FormatInt(lastId, 10) + " Not Found",
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(json)
		return
	}

	type Response struct {
		Status  string        `json:"status"`
		Message string        `json:"message"`
		Data    database.Todo `json:"data"`
	}

	toJson, err := json.Marshal(Response{
		"Success",
		"Success",
		todo,
	})
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(toJson)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDB()
	defer db.Close()

	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		http.Error(w, "Id should be present", 400)
	}

	errDb := database.DeleteTodoById(db, id)
	if errDb != nil && errDb.Error() == "No Records Found" {
		json, _ := json.Marshal(FailedResponse{
			"Not Found",
			"Todo with ID " + idStr + " Not Found",
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(json)
		return
	}

	type Response struct {
		Status  string   `json:"status"`
		Message string   `json:"message"`
		Data    struct{} `json:"data"`
	}

	toJson, err := json.Marshal(Response{
		"Success",
		"Success",
		struct{}{},
	})
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(toJson)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDB()
	defer db.Close()

	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Id should be present", 400)
		return
	}

	type RBody struct {
		Title    string `json:"title,omitempty"`
		IsActive *bool  `json:"is_active,omitempty"`
	}
	var requestBody RBody
	errBody := json.NewDecoder(r.Body).Decode(&requestBody)
	if errBody != nil {
		http.Error(w, errBody.Error(), 400)
		return
	}
	database.UpdateTodoById(db, id, requestBody.Title, requestBody.IsActive)
	todo, errDb := database.FindTodoById(db, id)
	if errDb != nil && errDb.Error() == "No Records Found" {
		json, _ := json.Marshal(FailedResponse{
			"Not Found",
			"Todo with ID " + idStr + " Not Found",
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(json)
		return
	}

	type Response struct {
		Status  string        `json:"status"`
		Message string        `json:"message"`
		Data    database.Todo `json:"data"`
	}
	toJson, errEncode := json.Marshal(Response{
		"Success",
		"Success",
		todo,
	})

	if errEncode != nil {
		http.Error(w, "Internal Server Error", 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(toJson)
}
