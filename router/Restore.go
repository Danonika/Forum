package router

import (
	"errors"
	"net/http"
	"text/template"

	"github.com/Danonika/Forum/data"
)

//Restore password
func Restore(w http.ResponseWriter, r *http.Request) {
	Authorised(r)
	if r.URL.Path[8:] != "" {
		ErrorHandler(w, r, errors.New("no such page"), 2)
	}
	Wow.Result = ""
	tmpl, _ := template.ParseFiles("./public/html/restore.html")
	if r.Method == "GET" {
		tmpl.Execute(w, Wow)
	} else {
		r.ParseForm()
		user, err := data.GetUserByUsername(r.Form["username"][0])
		if err != nil || user.Code != r.Form["code"][0] {
			Wow.Result = "Sorry wrong code or username"
		} else {
			Wow.Result = "Your password successfully changed"
			user.Update(w, r, r.Form["psw"][0])
			if Wow.Authorised == true {
				http.Redirect(w, r, "/logout", 302)
			} else {
				http.Redirect(w, r, "/", 302)
			}
		}
		tmpl.Execute(w, Wow)
	}
	Reset(&Wow)
}
