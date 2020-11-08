package server

import (
	"encoding/json"
	"log"
	"park_2020/2020_2_tmp_name/models"

	"github.com/gorilla/websocket"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

var MyHub *Hub

func (h *Hub) Run() {
	for {
		var client *Client
		select {
		case client = <-h.register:
			h.clients[client] = true
		case client = <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}

		s := NewServer()
		_, message, err := client.conn.ReadMessage()
		_, ok := err.(*websocket.CloseError)

		if err != nil && !ok {
			log.Println(err)

		} else if (err != nil && ok) || err == nil {
			var msg models.Message
			err = json.Unmarshal(message, &msg)
			if err != nil {
				log.Println(err)
			}

			err = s.InsertMessage(msg.Text, msg.ChatID, msg.UserID)
			if err != nil {
				log.Println(err)
			}

		} else {
			h.register <- client
		}
	}
}
