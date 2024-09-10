package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *app) initializeRoutes() http.Handler {
	router := httprouter.New()
	router.GET("/", app.homeHandler)
	router.GET("/guides/create", app.createGuide)
	router.POST("/guides/create", app.storeGuide)
	router.GET("/signup", app.signupHandler)
	router.POST("/signup", app.storeSignup)
	router.GET("/login", app.loginHandler)
	router.POST("/login", app.storeLogin)
	router.POST("/logout", app.storeLogout)
	return app.authenticate(router)
}
