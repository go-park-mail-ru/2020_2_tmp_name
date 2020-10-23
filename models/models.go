package models

import "github.com/gorilla/websocket"

type Account struct {
	AccountID int    `json:"account_id"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	Telephone string `json:"telephone"`
}

type LoginData struct {
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
}

type Photo struct {
	Telephone string `json:"telephone"`
	LinkImage string `json:"link_image"`
}

type Error struct {
	Message string `json:"message"`
}

type Message struct {
	Author string `json:"author"`
	Body   string `json:"body"`
}

type Client struct {
	id     int
	ws     *websocket.Conn
	server *Server
	ch     chan *Message
	doneCh chan bool
}

type Server struct {
	pattern   string
	messages  []*Message
	clients   map[int]*Client
	addCh     chan *Client
	delCh     chan *Client
	sendAllCh chan *Message
	doneCh    chan bool
	errCh     chan error
}
