package router

import (
	"hungrydev39/todo-list-challenge/controllers"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

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

	r.Post("/messaging", controllers.AddMessage)

	return r
}
