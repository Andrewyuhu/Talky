package main

import (
	hub "backend/cmd/api/websocket"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
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
