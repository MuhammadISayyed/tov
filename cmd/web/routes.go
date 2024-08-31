package main

import (
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/MuhammadISayyed/tov.git/internal/handlers"
)

func initializeRoutes(db *sql.DB) *httprouter.Router {
	router := httprouter.New()
	router.GET("/", homeHandler)
	
	// User routes
    router.POST("/users", wrapHandler(handlers.CreateUser(db)))
    
    // Add more routes here
    // router.GET("/users/:id", wrapHandler(handlers.GetUser(db)))
    // router.PUT("/users/:id", wrapHandler(handlers.UpdateUser(db)))
    // router.DELETE("/users/:id", wrapHandler(handlers.DeleteUser(db)))

	return router
}

func homeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./ui/html/pages/index.html")
}

// wrapHandler converts http.HandlerFunc to httprouter.Handle
func wrapHandler(h http.HandlerFunc) httprouter.Handle {
    return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        h(w, r)
    }
}