package router

import (
	"net/http"
	"strconv"

	data "../data"
)

//ErrorHandler - Allow us to handle errors in forum
func ErrorHandler(w http.ResponseWriter, r *http.Request, err error, status int) {
	if status == 0 {
		if Wow.Authorised == false {
			http.Redirect(w, r, "/login", 302)
		}
	}
	if err != nil {
		switch status {
		case 1:
			Wow.Error = []byte("No such thread")
		case 2:
			Wow.Error = []byte("No such page")
		case 3:
			Wow.Error = []byte("No such user")
		case 4:
			Wow.Error = []byte("You don't have permission")
		case 5:
			Wow.Error = []byte("Internal error has been occured")
		}
		http.Redirect(w, r, "/error", 302)
	}
}

//Comment - Allow authorised users to comment Threads
func Comment(w http.ResponseWriter, r *http.Request) {
	Authorised(r)
	ErrorHandler(w, r, nil, 0)
	x, err := strconv.Atoi(r.URL.Path[8:])
	ErrorHandler(w, r, err, 1)
	CurThread, err := data.GetThreadByID(x)
	ErrorHandler(w, r, err, 1)
	user, err := data.GetUserByID(Wow.ID)
	data.CreateCommentToThread(r, Wow.ID, CurThread.ThreadID, user.Username)
}
