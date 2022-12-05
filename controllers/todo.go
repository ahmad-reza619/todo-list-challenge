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

	todo := database.FindTodoById(db, id)
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
		ActivityGroupId *int64  `json:"activity_group_id"`
		Title           *string `json:"title"`
	}
	requestBody := RBody{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		panic(err.Error())
	}

	lastId := database.AddTodo(db, *requestBody.Title, *requestBody.ActivityGroupId)
	todo := database.FindTodoById(db, lastId)

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

	database.DeleteTodoById(db, id)

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
		Title *string `json:"title"`
	}
	requestBody := RBody{}
	errBody := json.NewDecoder(r.Body).Decode(&requestBody)
	if errBody != nil {
		http.Error(w, "Body should be present", 400)
		return
	}
	database.UpdateTodoById(db, id, *requestBody.Title)
	todo := database.FindTodoById(db, id)

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
