package router

import (
	"errors"
	"net/http"
	"text/template"
)

//About - A little information about our forum
func About(w http.ResponseWriter, r *http.Request) {
	Authorised(r)
	tmpl, _ := template.ParseFiles("./public/html/about.html")
	if r.URL.Path[6:] != "" {
		ErrorHandler(w, r, errors.New("no such page"), 2)
	}
	tmpl.Execute(w, Wow)
}
