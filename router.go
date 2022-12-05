package main

import (
	"hungrydev39/todo-list-challenge/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func setupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/activity-groups", controllers.AllActivity)
	r.Get("/activity-groups/{id}", controllers.ShowActivity)
	r.Post("/activity-groups", controllers.AddActivity)
	r.Patch("/activity-groups/{id}", controllers.UpdateActivity)
	r.Delete("/activity-groups/{id}", controllers.DeleteActivity)

	r.Get("/todo-items", controllers.GetTodos)
	r.Get("/todo-items/{id}", controllers.ShowTodo)
	r.Patch("/todo-items/{id}", controllers.UpdateTodo)
	r.Delete("/todo-items/{id}", controllers.DeleteTodo)
	r.Post("/todo-items", controllers.AddTodo)

	return r
}
