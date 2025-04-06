package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) router() http.Handler {
	router := httprouter.New()
	dynamic := alice.New(app.logRequest)
	router.Handler(http.MethodGet, "/v1/healthcheck", dynamic.Then(http.HandlerFunc(app.healthcheckHandler)))

	protected := dynamic.Append(app.isAuthenticated)

	router.Handler(http.MethodPost, "/v1/user/login", dynamic.Then(http.HandlerFunc(app.loginHandler)))
	router.Handler(http.MethodPost, "/v1/user/logout", dynamic.Then(http.HandlerFunc(app.logoutHandler)))
	router.Handler(http.MethodPost, "/v1/user/signup", dynamic.Then(http.HandlerFunc(app.signUpHandler)))

	router.Handler(http.MethodGet, "/v1/protected", protected.Then(http.HandlerFunc(app.healthcheckHandler)))

	return router
}
