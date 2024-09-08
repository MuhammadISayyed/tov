package main

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func (app *app) storeSignup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = app.users.Insert(
		r.PostForm.Get("name"),
		r.PostForm.Get("email"),
		r.PostForm.Get("password"),
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return 
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}

func (app *app) storeGuide(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = app.guides.Insert(
		r.PostForm.Get("title"),
		r.PostForm.Get("content"),
	)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func (app *app) createGuide(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles("./ui/html/pages/createGuide.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	t.Execute(w, nil)
}

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

func (app *app) signupHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./ui/html/pages/signup.html")	
}