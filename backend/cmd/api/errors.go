package main

import "net/http"

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorLogger.Println(err)
	message := "server was unable to process the request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) errorResponse(w http.ResponseWriter, _ *http.Request, status int, message any) {

	err := app.writeJSON(w, status, envelope{"error": message}, nil)
	if err != nil {
		app.errorLogger.Println(err)
		w.WriteHeader(500)
	}
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	message := "invalid request body format"
	app.errorLogger.Println(err)
	app.errorResponse(w, r, http.StatusBadRequest, message)
}

func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {

	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

func (app *application) notAuthorizedResponse(w http.ResponseWriter, r *http.Request) {

	app.errorResponse(w, r, http.StatusUnauthorized, "invalid credentials")
}

func (app *application) forbiddenResponse(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusForbidden, "user does not exist")
}
