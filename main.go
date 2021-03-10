package main

import (
	"net/http"

	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
	"golang.org/x/oauth2"
)

var googleOAuthConfig = oauth2.Config{

}

func googleLoginHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	mux := pat.New()
	mux.HandleFunc("/auth/google/login", googleLoginHandler)

	n := negroni.Classic()
	n.UseHandler(mux)
	http.ListenAndServe(":8000", n)
}