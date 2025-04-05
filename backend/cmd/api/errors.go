package main

import "net/http"

func (app *application) serverError(w http.ResponseWriter, r *http.Request) {

}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {

}
