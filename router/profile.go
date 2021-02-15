package router

import (
	"errors"
	"fmt"
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
		if Wow.Authorised {
			Wow.Data = data.GetAllUserCreatedPosts(user.ID)
			Wow.Data2 = data.GetAllUserLikedThread(user.ID)
			Wow.Data3 = data.GetAllUserLikedComments(user.ID)
		}
		Wow.CountOfPosts = len(Wow.Data)
		Wow.CountOfLikedThreads = len(Wow.Data2)
		Wow.CountOfLikedComments = len(Wow.Data3)
		if Wow.ID == user.ID {
			Wow.Me = true
		}
		Wow.Image = "User" + strconv.Itoa(user.ID)
		err = data.FindImage(Wow.Image)
		if err != nil {
			Wow.Image = "/static/images/default-avatar.jpg"
		} else {
			Wow.Image = "/static/images/User" + strconv.Itoa(user.ID)
		}
		tmpl.Execute(w, Wow)
		Reset(&Wow)
	}
}

//UpdateAva - Updating our avatar image
func UpdateAva(w http.ResponseWriter, r *http.Request) {
	Authorised(r)
	x, err2 := strconv.Atoi(r.URL.Path[20:])
	fmt.Println(x)
	user, err := data.GetUserByID(x)
	if err != nil || err2 != nil || r.Method == "GET" {
		ErrorHandler(w, r, err, 2)
		return
	}
	if user.ID != Wow.ID {
		ErrorHandler(w, r, errors.New("No permission"), 4)
		return
	}
	Wow.Image = "User" + strconv.Itoa(user.ID)
	err = data.AddImage(Wow.Image, 1, user.ID, r)
	if err != nil {
		ErrorHandler(w, r, err, 5)
		return
	}
	http.Redirect(w, r, "/id/"+strconv.Itoa(user.ID), 302)
}
