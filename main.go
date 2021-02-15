package main

import (
	"fmt"
	"net/http"
	"time"

	"./data"
	"./router"
	_ "github.com/mattn/go-sqlite3"
)

var Ok bool

func main() {
	data.Init()
	defer data.CloseDB()
	mux := http.NewServeMux()
	// // handle static assets
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mapper := map[string]func(http.ResponseWriter, *http.Request){
		"/":                    router.DefaultHandler, /*This login/register pages      Needs Session time auto remove */
		"/ajax":                router.AjaxHandler,    /*Checking mail/ username for validation*/
		"/logout":              router.LogOut,         /*Old school(Needs upgrade)*/
		"/id/":                 router.Profile,        /*NOTE!!!!!!: Needs update(changing pass/ pics of profile/ some basic info)*/
		"/about":               router.About,          /*Only for good present*/
		"/rules":               router.Rules,          /*Copy pasted thing(I didn't read it)*/
		"/restore":             router.Restore,        /*Good thing but works only for gmail*/
		"/thread/create":       router.CreateThread,   /*NOTE!!!!!!: Simple and stupid method(No pics/vids - Needs fix)*/
		"/thread/comment/":     router.Comment,        /*Old school method ajax(Needs upgrade)*/
		"/thread/":             router.Post,           /*NOTE!!!!!!: needs update to add some new feauters like editing/deleting thread*/
		"/error":               router.Error,          /*NOTE!!!!!!: Add some new errors*/
		"/stats":               router.Stats,          /*Perfect*/
		"/articles":            router.Articles,       /* NOTE!!!!!!: Needs update like sort by likes or Date*/
		"/updateProfileImage/": router.UpdateAva,
	}
	for pattern, handler := range mapper {
		mux.HandleFunc(pattern, handler)
	}

	fmt.Println("Server is listening...")
	server := &http.Server{
		Addr:           "0.0.0.0:8181",
		Handler:        mux,
		ReadTimeout:    time.Duration(10 * int64(time.Second)),
		WriteTimeout:   time.Duration(600 * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
