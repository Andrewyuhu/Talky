package main

import "net/http"

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
	// RELEARN HOW TO PARSE INFO FROM A REQUEST
	// VERIFY CRED IN DB
	// JWT GEN
	w.Write([]byte("login end point\n"))
}

func (app *application) signUpHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("signup endpoint\n"))
}

func (app *application) logoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("logout endpoint\n"))
}
