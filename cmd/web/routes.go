package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func initializeRoutes(router *httprouter.Router) {
	router.GET("/", homeHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to the home page!")
}