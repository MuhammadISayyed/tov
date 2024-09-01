package main

import (
	"github.com/julienschmidt/httprouter"
)

func initializeRoutes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", homeHandler)
	return router
}
