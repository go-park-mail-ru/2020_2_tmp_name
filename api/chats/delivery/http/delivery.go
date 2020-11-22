package http

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	domain "park_2020/2020_2_tmp_name/api/chats"
	"park_2020/2020_2_tmp_name/models"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type ChatHandlerType struct {
	ChUsecase domain.ChatUsecase
	Hub       Hub
}

func (h Hub) run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}

func NewChatHandler(r *mux.Router, chs domain.ChatUsecase) {
	handler := &ChatHandlerType{
		ChUsecase: chs,
		Hub:       *NewHub(),
	}

	go handler.Hub.run()

	r.HandleFunc("/api/v1/chat", handler.ChatHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/message", handler.MessageHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/chats", handler.ChatsHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/chats/{chat_id}", handler.ChatIDHandler).Methods(http.MethodGet)

	r.HandleFunc("/api/v1/gochat", handler.GochatHandler).Methods(http.MethodGet, http.MethodPost)
}

func JSONError(message string) []byte {
	jsonError, err := json.Marshal(models.Error{Message: message})
	if err != nil {
		return []byte("")
	}
	return jsonError
}

func (ch *ChatHandlerType) ChatHandler(w http.ResponseWriter, r *http.Request) {
	chat := models.Chat{}
	err := json.NewDecoder(r.Body).Decode(&chat)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	err = ch.ChUsecase.Chat(chat)
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(chat)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (ch *ChatHandlerType) MessageHandler(w http.ResponseWriter, r *http.Request) {
	message := models.Message{}
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	if len(r.Cookies()) == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(JSONError("User not authorized"))
		return
	}

	user, err := ch.ChUsecase.User(r.Cookies()[0].Value)
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	err = ch.ChUsecase.Message(user, message)
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(message)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (ch *ChatHandlerType) ChatsHandler(w http.ResponseWriter, r *http.Request) {
	if len(r.Cookies()) == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(JSONError("User not authorized"))
		return
	}

	user, err := ch.ChUsecase.User(r.Cookies()[0].Value)
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	chatModel, err := ch.ChUsecase.Chats(user)
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(chatModel)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (ch *ChatHandlerType) ChatIDHandler(w http.ResponseWriter, r *http.Request) {
	chid, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/v1/chats/"))
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	if len(r.Cookies()) == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(JSONError("User not authorized"))
		return
	}

	user, err := ch.ChUsecase.User(r.Cookies()[0].Value)
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	chat, err := ch.ChUsecase.ChatID(user, chid)
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(chat)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (ch *ChatHandlerType) GochatHandler(w http.ResponseWriter, r *http.Request) {
	if len(r.Cookies()) == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(JSONError("User not authorized"))
		return
	}

	user, err := ch.ChUsecase.User(r.Cookies()[0].Value)
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	client := &Client{ID: user.ID, Hub: &ch.Hub, Conn: conn, Send: make(chan []byte, 256)}
	for {
		_, message, err := client.Conn.ReadMessage()
		_, ok := err.(*websocket.CloseError)

		if err != nil && !ok {
			logrus.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(JSONError(err.Error()))
			return

		} else if (err != nil && ok) || err == nil {
			var msg models.Msg
			err = json.Unmarshal(message, &msg)
			if err != nil {
				logrus.Error(err)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(JSONError(err.Error()))
				return
			}

			err = ch.ChUsecase.Msg(user, msg)
			if err != nil {
				w.WriteHeader(models.GetStatusCode(err))
				w.Write(JSONError(err.Error()))
				return
			}

		}
		client.Hub.Register <- client
	}

	// go client.writePump()
	// go client.readPump()

}

// readPump pumps messages from the websocket connection to the hub.
func (c *Client) readPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.Hub.Broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
