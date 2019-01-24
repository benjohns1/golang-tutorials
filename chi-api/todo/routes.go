package todo

import (
	"github.com/benjohns1/golang-tutorials/chi-api/middleware"

	"github.com/go-chi/chi"
)

// Routes returns all todo API routes
func Routes(u UnitOfWork, router *chi.Mux, responseFormatter middleware.ResponseFormatter) *chi.Mux {
	router.Get("/{todoID}", responseFormatter(getTodo(u)))
	router.Delete("/{todoID}", responseFormatter(deleteTodo(u)))
	router.Post("/", responseFormatter(createTodo(u)))
	router.Get("/", responseFormatter(getAllTodos(u)))
	return router
}
