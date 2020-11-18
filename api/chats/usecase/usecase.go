package usecase

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	domain "park_2020/2020_2_tmp_name/api/chats"
	"park_2020/2020_2_tmp_name/models"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type chatUsecase struct {
	chatRepo       domain.ChatRepository
	Hub            *models.Hub
	client         *models.Client
	contextTimeout time.Duration
}

func NewChatUsecase(ch domain.ChatRepository) domain.ChatUsecase {
	return &chatUsecase{
		chatRepo: ch,
	}
}

func (ch *chatUsecase) Chat(chat models.Chat) error {
	err := ch.chatRepo.InsertChat(chat)
	if err != nil {
		return models.ErrInternalServerError
	}
	return nil
}

func (ch *chatUsecase) Message(user models.User, message models.Message) error {
	err := ch.chatRepo.InsertMessage(message.Text, message.ChatID, user.ID)
	if err != nil {
		return models.ErrInternalServerError
	}
	return nil
}

func (ch *chatUsecase) Chats(user models.User) (models.ChatModel, error) {
	var chatModel models.ChatModel
	chats, err := ch.chatRepo.SelectChatsByID(user.ID)
	if err != nil {
		return chatModel, models.ErrInternalServerError
	}

	chatModel.Data = chats
	return chatModel, nil
}

func (ch *chatUsecase) ChatID(user models.User, chid int) (models.ChatData, error) {
	var chat models.ChatData
	chat, err := ch.chatRepo.SelectChatByID(user.ID, chid)
	if err != nil {
		return chat, models.ErrInternalServerError
	}

	return chat, nil
}

func (ch *chatUsecase) User(cookie string) (models.User, error) {
	telephone := ch.chatRepo.CheckUserBySession(cookie)
	user, err := ch.chatRepo.SelectUser(telephone)
	if err != nil {
		return user, models.ErrNotFound
	}
	return user, nil
}

func (ch *chatUsecase) ServeWs(hub *models.Hub, w http.ResponseWriter, r *http.Request, uid int) {
	conn, err := models.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &models.Client{ID: uid, Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	_, message, err := client.Conn.ReadMessage()
	_, ok := err.(*websocket.CloseError)

	if err != nil && !ok {
		log.Println(err)

	} else if (err != nil && ok) || err == nil {
		var msg models.Msg
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Println(err)
		}

		err = ch.chatRepo.InsertMessage(msg.Message, msg.ChatID, msg.UserID)
		if err != nil {
			log.Println(err)
		}

	}
	if err == nil {
		ch.Hub.Register <- client
	}

	client.Hub.Register <- client
	go ch.writePump()
	go ch.readPump()
}

// readPump pumps messages from the websocket connection to the hub.
func (ch *chatUsecase) readPump() {
	defer func() {
		ch.client.Hub.Unregister <- ch.client
		ch.client.Conn.Close()
	}()
	for {
		_, message, err := ch.client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		ch.client.Hub.Broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
func (ch *chatUsecase) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		ch.client.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-ch.client.Send:
			if !ok {
				ch.client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := ch.client.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(ch.client.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-ch.client.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			ch.client.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ch.client.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
