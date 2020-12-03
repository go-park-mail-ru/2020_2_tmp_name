package http

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	domain "park_2020/2020_2_tmp_name/api/chats"
	authClient "park_2020/2020_2_tmp_name/microservices/authorization/delivery/grpc/client"
	"park_2020/2020_2_tmp_name/models"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type ChatHandlerType struct {
	ChUsecase  domain.ChatUsecase
	AuthClient authClient.AuthClientInterface
	Hub        Hub
}

func (h Hub) run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client.Session] = client
		case client := <-h.Unregister:
			if _, ok := h.Clients[client.Session]; ok {
				delete(h.Clients, client.Session)
				close(client.Send)
			}
		}
	}
}

func NewChatHandler(r *mux.Router, chs domain.ChatUsecase, ac authClient.AuthClientInterface) {
	handler := &ChatHandlerType{
		ChUsecase:  chs,
		AuthClient: ac,
		Hub:        *NewHub(),
	}

	go handler.Hub.run()

	r.HandleFunc("/api/v1/chat", handler.ChatHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/message", handler.MessageHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/chats", handler.ChatsHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/chats/{chat_id}", handler.ChatIDHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/like", handler.LikeHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/dislike", handler.DislikeHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/superlike", handler.SuperlikeHandler).Methods(http.MethodPost)

	r.HandleFunc("/api/v1/gochat", handler.GochatHandler).Methods(http.MethodGet)
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

	user, err := ch.AuthClient.CheckSession(context.Background(), r.Cookies())
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
	user, err := ch.AuthClient.CheckSession(context.Background(), r.Cookies())
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

	user, err := ch.AuthClient.CheckSession(context.Background(), r.Cookies())
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

func (ch *ChatHandlerType) LikeHandler(w http.ResponseWriter, r *http.Request) {
	like := models.Like{}
	err := json.NewDecoder(r.Body).Decode(&like)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	user, err := ch.AuthClient.CheckSession(context.Background(), r.Cookies())
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	err = ch.ChUsecase.Like(user, like)
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	chat, match, err := ch.ChUsecase.MatchUser(user, like)
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	if match {
		var chatData models.ChatData
		chatData.ID = chat.ID
		chatData.Partner, err = ch.ChUsecase.Partner(user, chat.ID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(JSONError(err.Error()))
			return
		}

		msg := models.Msg{Message: ""}
		chatData.Messages = []models.Msg{msg}

		body, err := json.Marshal(chatData)
		if err != nil {
			logrus.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(JSONError(err.Error()))
			return
		}

		clients := ch.Hub.Clients

		clientsMe := make(map[string]*Client)
		clientsPartner := make(map[string]*Client)

		for _, client := range clients {
			if client.ID == user.ID {
				clientsMe[client.Session] = client
			} else if client.ID == chatData.Partner.ID {
				clientsPartner[client.Session] = client
			}
		}

		for _, client := range clientsMe {
			client.Send <- body
		}

		var myChatData models.ChatData
		myChatData.ID = chatData.ID
		myChatData.Partner, err = ch.ChUsecase.UserFeed(r.Cookies()[0].Value)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(JSONError(err.Error()))
			return
		}
		myChatData.Messages = []models.Msg{msg}

		bodyMe, err := json.Marshal(myChatData)
		if err != nil {
			logrus.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(JSONError(err.Error()))
			return
		}

		for _, client := range clientsPartner {

			client.Send <- bodyMe
		}
	}

	body, err := json.Marshal(like)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (ch *ChatHandlerType) DislikeHandler(w http.ResponseWriter, r *http.Request) {
	dislike := models.Dislike{}
	err := json.NewDecoder(r.Body).Decode(&dislike)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	user, err := ch.AuthClient.CheckSession(context.Background(), r.Cookies())
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	err = ch.ChUsecase.Dislike(user, dislike)
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(dislike)
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
	user, err := ch.AuthClient.CheckSession(context.Background(), r.Cookies())
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

	client := &Client{ID: user.ID, Session: r.Cookies()[0].Value, Hub: &ch.Hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Register <- client

	go client.writePump(ch, user)
	go client.readPump(ch, user)
}

// readPump pumps messages from the websocket connection to the hub.
func (c *Client) readPump(ch *ChatHandlerType, user models.User) {
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

		var msg models.Msg
		err = json.Unmarshal(message, &msg)
		if err != nil {
			logrus.Error(err)
			return
		}

		err = ch.ChUsecase.Msg(user, msg)
		if err != nil {
			logrus.Error(err)
			return
		}

		var chatData models.ChatData
		chatData.ID = msg.ChatID
		chatData.Partner, err = ch.ChUsecase.Partner(user, msg.ChatID)
		if err != nil {
			logrus.Error(err)
			return
		}
		chatData.Messages = append(chatData.Messages, msg)

		data, err := json.Marshal(chatData)
		if err != nil {
			logrus.Error(err)
			return
		}

		sessions, err := ch.ChUsecase.Sessions(chatData.Partner.ID)
		clients := ch.Hub.Clients
		for _, session := range sessions {
			client, ok := clients[session]
			if ok {
				client.Send <- data
			}
		}

	}
}

// writePump pumps messages from the hub to the websocket connection.
func (c *Client) writePump(ch *ChatHandlerType, user models.User) {
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

func (ch *ChatHandlerType) SuperlikeHandler(w http.ResponseWriter, r *http.Request) {
	superlike := models.Superlike{}
	err := json.NewDecoder(r.Body).Decode(&superlike)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	user, err := ch.AuthClient.CheckSession(context.Background(), r.Cookies())
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	err = ch.ChUsecase.Superlike(user, superlike)
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(superlike)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
