package todo

import (
	"github.com/benjohns1/golang-tutorials/chi-api/middleware"

	"github.com/go-chi/chi"
)

// Routes returns all todo API routes
func Routes(router *chi.Mux, responseFormatter middleware.ResponseFormatter) *chi.Mux {
	router.Get("/{todoID}", responseFormatter(getTodo))
	router.Delete("/{todoID}", responseFormatter(deleteTodo))
	router.Post("/", responseFormatter(createTodo))
	router.Get("/", responseFormatter(getAllTodos))
	return router
}
