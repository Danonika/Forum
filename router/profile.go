package router

import (
	"net/http"
	"strconv"
	"text/template"

	data "../data"
)

//Profile - Profile page of user
func Profile(w http.ResponseWriter, r *http.Request) {
	Authorised(r)
	tmpl, _ := template.ParseFiles("./public/html/profile.html")
	x, err2 := strconv.Atoi(r.URL.Path[4:])
	user, err := data.GetUserByID(x)
	if err2 != nil {
		user, err = data.GetUserByUsername(r.URL.Path[4:])
	}
	if err != nil {
		ErrorHandler(w, r, err, 3)
	} else {
		Wow.Title = user.Username
		Wow.Result = user.Mail
		Wow.Data = data.GetAllUserCreatedPosts(user.ID)
		Wow.Data2 = data.GetAllUserLikedThread(user.ID)
		Wow.Data3 = data.GetAllUserLikedComments(user.ID)
		Wow.CountOfPosts = len(Wow.Data)
		Wow.CountOfLikedThreads = len(Wow.Data2)
		Wow.CountOfLikedComments = len(Wow.Data3)
		tmpl.Execute(w, Wow)
		Reset(&Wow)
	}
}
