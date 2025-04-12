package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/rs/cors"
)

func (app *application) router() http.Handler {
	router := httprouter.New()
	dynamic := alice.New(app.logRequest)
	router.Handler(http.MethodGet, "/v1/healthcheck", dynamic.Then(http.HandlerFunc(app.healthcheckHandler)))

	protected := dynamic.Append(app.isAuthenticated)

	router.Handler(http.MethodPost, "/v1/user/login", dynamic.Then(http.HandlerFunc(app.loginHandler)))
	router.Handler(http.MethodPost, "/v1/user/logout", dynamic.Then(http.HandlerFunc(app.logoutHandler)))
	router.Handler(http.MethodPost, "/v1/user/signup", dynamic.Then(http.HandlerFunc(app.signUpHandler)))
	router.Handler(http.MethodGet, "/v1/user/me", protected.Then(http.HandlerFunc(app.meHandler)))

	router.Handler(http.MethodPost, "/v1/chat", protected.Then(http.HandlerFunc(app.createChatHandler)))
	router.Handler(http.MethodGet, "/v1/chat", protected.Then(http.HandlerFunc(app.getChatsHandler)))
	router.Handler(http.MethodGet, "/v1/ws/chat/:chatID", protected.Then(http.HandlerFunc(app.serveWs)))

	router.Handler(http.MethodGet, "/v1/message/:chatID", protected.Then(http.HandlerFunc(app.getMessagesHandler)))

	corsRouter := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Your frontend's origin
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	}).Handler(router)

	return corsRouter
}
