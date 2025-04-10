package hub

import (
	"fmt"
	"sync"
)

type HubManager struct {
	hubs       map[string]*Hub
	hubsLock   *sync.Mutex
	unregister chan (string)
}

type Hub struct {
	id            string
	clients       map[*Client]bool
	broadcast     chan []byte
	Register      chan *Client
	unregister    chan *Client
	unregisterHub chan string
	hubsLock      *sync.Mutex
}

func NewHubManager() *HubManager {
	var hubsMu sync.Mutex
	return &HubManager{
		hubs:       make(map[string]*Hub),
		hubsLock:   &hubsMu,
		unregister: make(chan string),
	}
}

func (hm *HubManager) GetRoom(id string) *Hub {
	hm.hubsLock.Lock()
	defer hm.hubsLock.Unlock()
	return hm.hubs[id]
}

func (hm *HubManager) AddRoom(id string) *Hub {
	hm.hubsLock.Lock()
	defer hm.hubsLock.Unlock()
	room := NewHub(hm.hubsLock, hm.unregister, id)
	hm.hubs[id] = room
	return room
}

func (hm *HubManager) Run() {
	fmt.Println("hub manager started")
	for id := range hm.unregister {
		fmt.Println("removing room")
		hm.hubsLock.Lock()
		fmt.Printf("Deleting active room connection : %s\n", id)
		delete(hm.hubs, id)
		hm.hubsLock.Unlock()
	}
}

func NewHub(hubMutex *sync.Mutex, unregisterHub chan string, id string) *Hub {
	return &Hub{
		id:            id,
		broadcast:     make(chan []byte),
		Register:      make(chan *Client),
		unregister:    make(chan *Client),
		unregisterHub: unregisterHub,
		clients:       make(map[*Client]bool),
		hubsLock:      hubMutex,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			fmt.Println("Registered client")
			h.clients[client] = true
		case client := <-h.unregister:

			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)
			}

			// close hub if no clients
			if len(h.clients) == 0 {
				fmt.Println("Last client left connection, remove room")
				h.unregisterHub <- h.id
				fmt.Println("successfully removed room")
				return
			}

		case message := <-h.broadcast:
			fmt.Println("Broadcasting")
			fmt.Println(message)
			for client := range h.clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.clients, client)
				}
			}

		}
	}

}
