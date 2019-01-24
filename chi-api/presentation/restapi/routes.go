package restapi

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/benjohns1/golang-tutorials/chi-api/core"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Routes builds and returns the HTTP router
func Routes(unitOfWork core.UnitOfWork) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/todo", todoRoutes(unitOfWork, chi.NewRouter(), NewResponseFormatter(render.JSON)))
	})

	return router
}

func errorResponse(err error) map[string]string {
	response := make(map[string]string)
	response["error"] = err.Error()
	return response
}

func successMessage(message string) map[string]string {
	response := make(map[string]string)
	response["message"] = message
	return response
}

func todoRoutes(u core.UnitOfWork, router *chi.Mux, responseFormatter ResponseFormatter) *chi.Mux {
	router.Get("/{todoID}", responseFormatter(getTodo(u)))
	router.Delete("/{todoID}", responseFormatter(deleteTodo(u)))
	router.Post("/", responseFormatter(createTodo(u)))
	router.Get("/", responseFormatter(getAllTodos(u)))
	return router
}

func getTodo(u core.UnitOfWork) func(*http.Request) interface{} {
	return func(r *http.Request) interface{} {
		todoID, convErr := strconv.Atoi(chi.URLParam(r, "todoID"))
		if convErr != nil {
			return errorResponse(convErr)
		}
		todo, getErr := core.GetTodo(u, todoID)
		if getErr != nil {
			return errorResponse(getErr)
		}
		return todo
	}
}

func deleteTodo(u core.UnitOfWork) func(*http.Request) interface{} {
	return func(r *http.Request) interface{} {
		todoID, convErr := strconv.Atoi(chi.URLParam(r, "todoID"))
		if convErr != nil {
			return errorResponse(convErr)
		}
		removeErr := u.Todo().Remove(todoID)
		if removeErr != nil {
			return errorResponse(removeErr)
		}
		return successMessage(fmt.Sprintf("Todo ID %v deleted successfully", todoID))
	}
}

func createTodo(u core.UnitOfWork) func(*http.Request) interface{} {
	return func(r *http.Request) interface{} {
		// @TODO: pull data from request
		todo := core.Todo{
			Slug:  "slug",
			Title: "Hello world",
			Body:  "Heloo world from planet earth",
		}
		todoID, addErr := u.Todo().Add(&todo)
		if addErr != nil {
			return errorResponse(addErr)
		}
		return successMessage(fmt.Sprintf("Todo ID %v created successfully", todoID))
	}
}

func getAllTodos(u core.UnitOfWork) func(*http.Request) interface{} {
	return func(r *http.Request) interface{} {
		todos, err := u.Todo().GetAll()
		if err != nil {
			return errorResponse(err)
		}
		return todos
	}
}
