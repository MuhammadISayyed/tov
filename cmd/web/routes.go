package main

import (
	"github.com/julienschmidt/httprouter"
)

func (app *app) initializeRoutes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", app.homeHandler)
	return router
}
