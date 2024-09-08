package main

import (
	"github.com/julienschmidt/httprouter"
)

func (app *app) initializeRoutes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", app.homeHandler)
	router.GET("/guides/create", app.createGuide)
	router.POST("/guides/create", app.storeGuide)
	router.GET("/signup", app.signupHandler)
	return router
}
