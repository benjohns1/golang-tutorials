package todo

import (
	"net/http"

	"github.com/go-chi/chi"
)

type todo struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type response = interface{}

func getTodo(r *http.Request) response {
	todoID := chi.URLParam(r, "todoID")
	todos := todo{
		Slug:  todoID,
		Title: "Hello world",
		Body:  "Heloo world from planet earth",
	}
	return todos
}

func deleteTodo(r *http.Request) response {
	response := make(map[string]string)
	response["message"] = "Deleted TODO successfully"
	return response
}

func createTodo(r *http.Request) response {
	response := make(map[string]string)
	response["message"] = "Created TODO successfully"
	return response
}

func getAllTodos(r *http.Request) response {
	todos := []todo{
		{
			Slug:  "slug",
			Title: "Hello world",
			Body:  "Heloo world from planet earth",
		},
	}
	return todos
}
