package main

import (
	"backend/internals/auth"
	"backend/internals/data"
	"backend/internals/validator"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
	input := struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	err := app.readJSON(r, &input)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	id, err := app.usermodel.Authenticate(input.Username, input.Password)

	if err != nil {
		if errors.Is(err, data.ErrInvalidCredentials) {
			app.notAuthorizedResponse(w, r)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	jwtToken, err := auth.GenerateJWT(id)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	auth.SetAuthCookie(w, jwtToken)

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "login successful"}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}

func (app *application) signUpHandler(w http.ResponseWriter, r *http.Request) {

	input := struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	err := app.readJSON(r, &input)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.NewValidator()

	validator.MinLength("username", input.Username, 3, v)
	validator.MaxLength("username", input.Username, 15, v)
	validator.MinLength("password", input.Password, 5, v)
	validator.MaxLength("password", input.Password, 50, v)
	validator.ValidEmail("email", input.Email, v)

	if !v.IsValid() {
		app.failedValidationResponse(w, r, v.InvalidProperties)
		return
	}

	err = app.usermodel.Insert(input.Username, input.Email, input.Password)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateEmail):
			app.errorResponse(w, r, http.StatusUnprocessableEntity, err.Error())
			return
		case errors.Is(err, data.ErrDuplicateUsername):
			app.errorResponse(w, r, http.StatusUnprocessableEntity, err.Error())
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusCreated, envelope{"message": "account created successfully"}, nil)

}

func (app *application) logoutHandler(w http.ResponseWriter, r *http.Request) {
	auth.ClearAuthCookie(w)

	err := app.writeJSON(w, http.StatusOK, envelope{"message": "logout successful"}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) meHandler(w http.ResponseWriter, r *http.Request) {
	// by here, the user should already have a valid JWT token
	id, ok := r.Context().Value(userIDKey).(uuid.UUID)

	if !ok {
		app.serverErrorResponse(w, r, errors.New("issue with authentication"))
		return
	}

	user, err := app.usermodel.Get(id)

	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			app.notAuthorizedResponse(w, r)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}
