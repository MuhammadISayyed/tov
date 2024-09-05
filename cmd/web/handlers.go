package main

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func (app *app) homeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	guides, err := app.guides.All()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return 
	}

	// http.ServeFile(w, r, "./ui/html/pages/index.html")
	t, err := template.ParseFiles("./ui/html/pages/index.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	t.Execute(w, map[string]any{"Guides": guides})
}