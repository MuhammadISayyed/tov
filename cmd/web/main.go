package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/MuhammadISayyed/tov.git/internal/models/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

type app struct {
	guides *sqlite.GuideModel
	users *sqlite.UserModel

}

func main() {
	db, err := sql.Open("sqlite3", "./app.db")
    if err != nil {
        log.Fatal(err)
    }
	log.Println(db)

	app := app{
		guides: &sqlite.GuideModel{
			DB: db,
		},
		users: &sqlite.UserModel{
			DB: db,
		},
	}

	router := app.initializeRoutes()

	log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))

}