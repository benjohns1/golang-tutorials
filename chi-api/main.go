package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/benjohns1/golang-tutorials/chi-api/persistence/postgresdb"
	"github.com/benjohns1/golang-tutorials/chi-api/presentation/restapi"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

const (
	dbUser     = "postgres"
	dbPassword = "asdf1234"
	dbName     = "postgres"
)

func main() {

	// Persistence (PostgreSQL DB)
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
	todoUnitOfWork := postgresdb.NewUnitOfWork(db)

	// Presentation (REST API)
	router := restapi.Routes(todoUnitOfWork)

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("logging err: %s\n", err.Error())
	}
	log.Fatal(http.ListenAndServe(":8080", router))
}
