package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	data "../data"
)

//StatsTH - struct of Thread statistic
type StatsTH struct {
	ThreadID string
	Value    int
	Likes    string
	Dislikes string
	Liked    string
}

//Authorised - Is user authorised ?
func Authorised(r *http.Request) {
	id, err := data.CheckCookie(r)
	if err == nil {
		Wow.Authorised = true
		Wow.ID = id
	} else {
		Wow.Authorised = false
	}
}

//CreateThread - Html page to Create Thread
func CreateThread(w http.ResponseWriter, r *http.Request) {
	Authorised(r)
	ErrorHandler(w, r, nil, 0)
	user, _ := data.GetUserByID(Wow.ID)
	tmpl, _ := template.ParseFiles("./public/html/CreateThread.html")
	if r.Method == "POST" {
		num, err := data.CreateTH(r, Wow.ID, user.Username)
		fmt.Println(err)
		if err != nil {
			ErrorHandler(w, r, err, 5)
		}
		http.Redirect(w, r, "/thread/"+strconv.Itoa(num), 302)

	}
	tmpl.Execute(w, Wow)
}

//Post - Page to Thread
func Post(w http.ResponseWriter, r *http.Request) {
	Authorised(r)
	tmpl, _ := template.ParseFiles("./public/html/thread.html")
	if r.Method == "POST" {
		Comment(w, r)
	}
	x, err := strconv.Atoi(r.URL.Path[8:])
	if err != nil {
		ErrorHandler(w, r, err, 2)
	} else {
		CurThread, err := data.GetThreadByID(x)
		if err != nil {
			ErrorHandler(w, r, err, 1)
		} else {
			Wow.Title = CurThread.Title
			user, _ := data.GetUserByID(CurThread.UserID)
			Wow.Result = user.Username
			Wow.Time = CurThread.Date
			Wow.Data = data.GetAllToThreadByID(CurThread.ThreadID, Wow.ID)
			tmpl.Execute(w, Wow)
			Reset(&Wow)
		}
	}

}

// Stats - Ajax handler to online statistic
func Stats(w http.ResponseWriter, r *http.Request) {
	Authorised(r)
	if Wow.Authorised == false {
		return
	}
	ErrorHandler(w, r, nil, 0)
	if r.Method == "GET" {
		ErrorHandler(w, r, errors.New("page for ajax"), 2)
	} else {
		var x StatsTH
		err := json.NewDecoder(r.Body).Decode(&x)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		ID, err := strconv.Atoi(x.ThreadID)
		ErrorHandler(w, r, err, 1)
		data.AddNewValueToThread(Wow.ID, ID, x.Value)
		Thread, err := data.GetThreadByID(ID)
		x.Likes = strconv.Itoa(Thread.Likes)
		x.Dislikes = strconv.Itoa(Thread.Dislikes)
		x.Liked = strconv.Itoa(data.CheckUserLikedThread(Wow.ID, ID))
		a, err := json.Marshal(x)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(a)
	}
}

//Articles - List of articles
func Articles(w http.ResponseWriter, r *http.Request) {
	Authorised(r)
	Wow.Data = data.GetAll(Wow.ID)
	tmpl, _ := template.ParseFiles("./public/html/articles.html")
	tmpl.Execute(w, Wow)
}

//DeleteThread - Need Update After
func DeleteThread() {}

//EditThread - New Feature
func EditThread() {}
