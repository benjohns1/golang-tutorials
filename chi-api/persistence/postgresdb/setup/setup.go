package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	dbUser     = "postgres"
	dbPassword = "asdf1234"
	dbName     = "postgres"
)

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE todo (
        id serial NOT NULL,
		slug character varying(100) NOT NULL,
		title character varying(100) NOT NULL,
		body character varying(500) NOT NULL,
		CONSTRAINT todo_pkey PRIMARY KEY (id)
    )
	WITH (OIDS=FALSE)`)
	if err != nil {
		panic(err)
	}
}
