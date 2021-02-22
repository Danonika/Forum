package router

import (
	"errors"
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"

	utils "../utils"
	data "github.com/Danonika/Forum/data"
)

//ViewData - struct to template html page
type ViewData struct {
	Result               string
	Authorised           bool
	ID                   int
	Title                string
	Time                 string
	Data                 []data.Thread
	Data2                []data.Thread
	Data3                []data.Thread
	Error                []byte
	CountOfPosts         int
	CountOfLikedThreads  int
	CountOfLikedComments int
	Me                   bool
	Image                string
}

// Wow - value to template html page
var Wow ViewData

// DefaultHandler - Default Request Handler
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	Authorised(r)
	tmpl, _ := template.ParseFiles("./public/html/login.html")
	if r.URL.Path[1:] == "register" {
		if Wow.Authorised == true {
			http.Redirect(w, r, "/articles", 302)
		}
		tmpl, _ = template.ParseFiles("./public/html/register.html")
		if r.Method == "GET" {
			tmpl.Execute(w, Wow)
		}
	} else if r.URL.Path[1:] == "login" || r.URL.Path == "/" {
		if Wow.Authorised == true {
			http.Redirect(w, r, "/articles", 302)
		}
		tmpl, _ = template.ParseFiles("./public/html/login.html")
		if r.Method == "GET" {
			tmpl.Execute(w, Wow)
		} else {
			r.ParseForm()
			_, ok := r.Form["checkbox"]
			if len(r.Form) == 3 && !ok {
				utils.AddUser(r.Form["username"][0], r.Form["mail"][0], r.Form["psw"][0])
			}
			user, err := data.GetUserByUsername(r.Form["username"][0])
			if err != nil {
				user, err = data.GetUserByMail(r.Form["username"][0])
			}
			if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Psw), []byte(r.Form["psw"][0])) != nil {
				Wow.Result = "Sorry the username or password is not correct\n"
				defer Reset(&Wow)
			} else {
				user.CreateAndSetSession(w, r)
				Wow.Authorised = true
				Wow.ID = user.ID
				http.Redirect(w, r, "/articles", 302)
			}
			tmpl.Execute(w, Wow)
		}
	} else {
		ErrorHandler(w, r, errors.New("no such page"), 2)
	}
}
