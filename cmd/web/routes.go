package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func initializeRoutes(router *httprouter.Router) {
	router.GET("/", homeHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./ui/html/pages/index.html")
}