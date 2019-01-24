package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	apimiddle "github.com/benjohns1/golang-tutorials/chi-api/middleware"
	"github.com/benjohns1/golang-tutorials/chi-api/persistence"
	"github.com/benjohns1/golang-tutorials/chi-api/todo"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	_ "github.com/lib/pq"
)

const (
	dbUser     = "postgres"
	dbPassword = "asdf1234"
	dbName     = "postgres"
)

func routes(unitOfWork todo.UnitOfWork) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/todo", todo.Routes(unitOfWork, chi.NewRouter(), apimiddle.NewResponseFormatter(render.JSON)))
	})

	return router
}

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	pingErr := db.Ping()
	if pingErr != nil {
		panic(pingErr)
	}
	todoUnitOfWork := persistence.NewTodoUnitOfWork(db)

	router := routes(todoUnitOfWork)

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("logging err: %s\n", err.Error())
	}
	log.Fatal(http.ListenAndServe(":8080", router))
}
