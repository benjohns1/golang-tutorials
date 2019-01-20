package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:asdf1234@tcp(localhost:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE test")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE test")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE userinfo (
        uid INT(10) NOT NULL AUTO_INCREMENT,
        username VARCHAR(64) NULL DEFAULT NULL,
        departname VARCHAR(64) NULL DEFAULT NULL,
        created DATE NULL DEFAULT NULL,
        PRIMARY KEY (uid)
    )`)
	if err != nil {
		panic(err)
	}
}
