package main

import (
	"backend/internals/auth"
	"backend/internals/data"
	"context"
	"errors"
	"net/http"
)

type ctxKey string

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLogger.Printf("%s %s", r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}

func (app *application) isAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("auth_token")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				app.notAuthorizedResponse(w, r)
				return
			}
			app.serverErrorResponse(w, r, err)
			return
		}

		id, err := auth.ValidateJWT(cookie.Value)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		_, err = app.usermodel.Exists(id)

		if err != nil {
			if errors.Is(err, data.ErrRecordNotFound) {
				app.forbiddenResponse(w, r)
				return
			}
			app.serverErrorResponse(w, r, err)
			return
		}

		const userIDKey ctxKey = "userID"

		ctx := context.WithValue(r.Context(), userIDKey, id)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

	})
}
