package main

import (
	"fmt"
	"net/http"
)

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
	input := struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	err := app.readJSON(r, &input)

	if err != nil {
		fmt.Println("read error")
		panic("dead")
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"data": input}, nil)

	if err != nil {
		fmt.Println("write error")
		panic("dead")
	}

}

func (app *application) signUpHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("signup endpoint\n"))
}

func (app *application) logoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("logout endpoint\n"))
}
