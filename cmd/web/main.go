package main

import (
	"log"
	"net/http"

	"github.com/MuhammadISayyed/tov.git/internal/database"
)

func main() {
	db, err := database.Connect()
	if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

	err = database.Initialize(db)
    if err != nil {
        log.Fatal(err)
    }

	router := initializeRoutes(db)

	log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))

}