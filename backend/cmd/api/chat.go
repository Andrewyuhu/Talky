package main

import (
	hub "backend/cmd/api/websocket"
	"backend/internals/auth"
	"backend/internals/data"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (app *application) serveWs(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(userIDKey).(uuid.UUID)

	if !ok {
		auth.ClearAuthCookie(w)
		app.notAuthorizedResponse(w, r)
		return
	}

	params := httprouter.ParamsFromContext(r.Context())
	stringChatID := params.ByName("chatID")
	chatID, err := ConvertParamInt(stringChatID)

	if err != nil {
		app.badRequestResponse(w, r, err)
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

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		app.serverErrorResponse(w, r, err)
		return
	}

	room := app.hubManager.GetRoom(stringChatID)
	if room == nil {
		fmt.Println("creating chat room")
		room = app.hubManager.AddRoom(stringChatID)
	}
	go room.Run()
	client := &hub.Client{Hub: room, DB: app.messagemodel, Conn: conn, Send: make(chan []byte, 256)}
	room.Register <- client

	go client.WritePump()
	go client.ReadPump()
}

func (app *application) createChatHandler(w http.ResponseWriter, r *http.Request) {

	id, ok := r.Context().Value(userIDKey).(uuid.UUID)

	if !ok {
		app.serverErrorResponse(w, r, errors.New("issue with authentication"))
		return
	}

	input := struct {
		Username string `json:"username"`
	}{}

	err := app.readJSON(r, &input)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	recipientID, err := app.usermodel.GetByUsername(input.Username)

	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			app.badRequestResponse(w, r, err)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.chatmodel.InsertNewChat(id, recipientID)

	if err != nil {
		if errors.Is(err, data.ErrDuplicateChat) {
			app.errorResponse(w, r, http.StatusUnprocessableEntity, err.Error())
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	w.Write([]byte("Chat created!"))

}

func (app *application) getChatsHandler(w http.ResponseWriter, r *http.Request) {

	id, ok := r.Context().Value(userIDKey).(uuid.UUID)

	if !ok {
		app.serverErrorResponse(w, r, errors.New("issue with authentication"))
		return
	}

	chats, err := app.chatmodel.GetChats(id)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"chats": chats}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}
