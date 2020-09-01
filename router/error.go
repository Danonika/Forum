package router

import (
	"net/http"
	"text/template"
)

//SetAndDelete - Sets and clears two variables
func SetAndDelete(s *string, t *[]byte) {
	*s = string(*t)
	*t = make([]byte, 0)
}

//Reset - Reset all values after templates except main attributes
func Reset(x *ViewData) {
	*x = ViewData{Authorised: x.Authorised, ID: x.ID}
}

//Error - html page to handle errors
func Error(w http.ResponseWriter, r *http.Request) {
	Authorised(r)
	defer Reset(&Wow)
	SetAndDelete(&Wow.Result, &Wow.Error)
	if string(Wow.Result[:]) == "" {
		http.Redirect(w, r, "/articles", 302)
	}
	tmpl, _ := template.ParseFiles("./public/html/error.html")
	tmpl.Execute(w, Wow)
}

// if (response.Liked == 0){
//   $('#{{.ThreadID}}1').css('color: white');
//   $('#{{.ThreadID}}2').css('color: white');
// }
// if (response.Liked == 1){
//   $('#{{.ThreadID}}1').css('color: darkblue');
//   $('#{{.ThreadID}}2').css('color: white');
// }
// if (response.Liked == -1){
//   $('#{{.ThreadID}}1').css('color: white');
//   $('#{{.ThreadID}}2').css('color: darkblue');
// }
