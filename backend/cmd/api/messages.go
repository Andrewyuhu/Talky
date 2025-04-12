package main

import (
	"backend/internals/auth"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

func (app *application) getMessagesHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	chatID, err := ConvertParamInt(params.ByName("chatID"))

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	userID, ok := r.Context().Value(userIDKey).(uuid.UUID)

	if !ok {
		auth.ClearAuthCookie(w)
		app.notAuthorizedResponse(w, r)
		return
	}

	valid, err := app.chatmodel.ValidateChat(userID, chatID)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if !valid {
		app.badRequestResponse(w, r, err)
		return
	}

	messages, err := app.messagemodel.Get(chatID)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"messages": messages}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}
