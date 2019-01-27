package restapi

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/benjohns1/golang-tutorials/chi-api/core"
	"github.com/go-chi/chi"
)

type todo struct {
	ID    int    `json:"id"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func todoRoutes(u core.UnitOfWork, router *chi.Mux, responseFormatter ResponseFormatter, createReqDecoder RequestDecoder) *chi.Mux {
	router.Get("/{todoID}", responseFormatter(getTodo(u)))
	router.Delete("/{todoID}", responseFormatter(deleteTodo(u)))
	router.Post("/", responseFormatter(createTodo(u, createReqDecoder)))
	router.Get("/", responseFormatter(getAllTodos(u)))
	return router
}

func getTodo(u core.UnitOfWork) func(*http.Request) interface{} {
	return func(r *http.Request) interface{} {
		todoID, err := strconv.Atoi(chi.URLParam(r, "todoID"))
		if err != nil {
			return errorResponse(err)
		}
		t, err := core.GetTodo(u, todoID)
		if err != nil {
			return errorResponse(err)
		}
		return todo(*t) // convert to json struct
	}
}

func deleteTodo(u core.UnitOfWork) func(*http.Request) interface{} {
	return func(r *http.Request) interface{} {
		todoID, err := strconv.Atoi(chi.URLParam(r, "todoID"))
		if err != nil {
			return errorResponse(err)
		}
		err = u.Todo().Remove(todoID)
		if err != nil {
			return errorResponse(err)
		}
		return successMessage(fmt.Sprintf("Todo ID %v deleted successfully", todoID))
	}
}

func createTodo(u core.UnitOfWork, createDecoder RequestDecoder) func(*http.Request) interface{} {
	return func(r *http.Request) interface{} {
		var t todo
		err := createDecoder.Decode(r.Body, &t)
		if err != nil {
			return errorResponse(err)
		}
		todo := core.Todo(t) // convert to entity

		todoID, err := u.Todo().Add(&todo)
		if err != nil {
			return errorResponse(err)
		}
		return successMessage(fmt.Sprintf("Todo ID %v created successfully", todoID))
	}
}

func getAllTodos(u core.UnitOfWork) func(*http.Request) interface{} {
	return func(r *http.Request) interface{} {
		ts, err := u.Todo().GetAll()
		if err != nil {
			return errorResponse(err)
		}

		// Convert to json
		todos := []todo{}
		for _, t := range ts {
			todos = append(todos, todo(*t))
		}

		return todos
	}
}
