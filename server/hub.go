package server

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

func (s *Service) Run() {
	for {
		var client *Client
		select {
		case client = <-s.Hub.register:
			s.Hub.clients[client] = true
		case client = <-s.Hub.unregister:
			if _, ok := s.Hub.clients[client]; ok {
				delete(s.Hub.clients, client)
				close(client.send)
			}
		case message := <-s.Hub.broadcast:
			for client := range s.Hub.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(s.Hub.clients, client)
				}
			}
		}
	}
}
