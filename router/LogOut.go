package router

import (
	"net/http"

	data "../data"
)

//LogOut - log out
func LogOut(w http.ResponseWriter, r *http.Request) {
	Authorised(r)
	cookie, err := r.Cookie("Cookie")
	if err != http.ErrNoCookie {
		UUID := cookie.Value
		data.DeleteByUUID(UUID)
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
	http.Redirect(w, r, "/", 302)
	Wow.Authorised = false
	Wow.ID = 0
}
