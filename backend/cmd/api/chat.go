package main

import (
	hub "backend/cmd/api/websocket"
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

	params := httprouter.ParamsFromContext(r.Context())
	id := params.ByName("roomID")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// database operation to check if room exists & if user belongs in room needs to go here

	room := app.hubManager.GetRoom(id)
	if room == nil {
		fmt.Println("creating chat room")
		room = app.hubManager.AddRoom(id)
	}
	go room.Run()
	client := &hub.Client{Hub: room, Conn: conn, Send: make(chan []byte, 256)}
	room.Register <- client

	go client.WritePump()
	go client.ReadPump()
}

func (app *application) createChatHandler(w http.ResponseWriter, r *http.Request) {

	id, ok := r.Context().Value(userIDKey).(uuid.UUID)

	if !ok {
		app.serverErrorResponse(w, r, errors.New("issue with authentication"))
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
