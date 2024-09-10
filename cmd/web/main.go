package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/MuhammadISayyed/tov.git/internal/models/sqlite"
	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
)

type app struct {
	guides *sqlite.GuideModel
	users *sqlite.UserModel
	session *sessions.CookieStore
}

func main() {
	db, err := sql.Open("sqlite3", "./app.db")
    if err != nil {
        log.Fatal(err)
    }
	log.Println(db)

	session := sessions.NewCookieStore([]byte("6qIR6SJWVcBYukoqFgNztDeEmu1nwbYt"))
	session.Options.HttpOnly = true
	session.Options.SameSite = http.SameSiteLaxMode

	app := app{
		guides: &sqlite.GuideModel{
			DB: db,
		},
		users: &sqlite.UserModel{
			DB: db,
		},
		session: session,
	}

	router := app.initializeRoutes()

	log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))

}