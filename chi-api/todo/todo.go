package todo

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// Todo domain entity contains todo data
type Todo struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func errorResponse(err error) map[string]string {
	response := make(map[string]string)
	response["error"] = err.Error()
	return response
}

func getTodo(u UnitOfWork) func(*http.Request) interface{} {
	return func(r *http.Request) interface{} {
		todoID, err := strconv.Atoi(chi.URLParam(r, "todoID"))
		if err != nil {
			return errorResponse(err)
		}
		todo, err := u.Todo().Get(todoID)
		if err != nil {
			return errorResponse(err)
		}
		return todo
	}
}

func deleteTodo(u UnitOfWork) func(*http.Request) interface{} {
	return func(r *http.Request) interface{} {
		response := make(map[string]string)
		response["message"] = "Deleted TODO successfully"
		return response
	}
}

func createTodo(u UnitOfWork) func(*http.Request) interface{} {
	return func(r *http.Request) interface{} {
		response := make(map[string]string)
		response["message"] = "Created TODO successfully"
		return response
	}
}

func getAllTodos(u UnitOfWork) func(*http.Request) interface{} {
	return func(r *http.Request) interface{} {
		todos := []Todo{
			{
				Slug:  "slug",
				Title: "Hello world",
				Body:  "Heloo world from planet earth",
			},
		}
		return todos
	}
}
