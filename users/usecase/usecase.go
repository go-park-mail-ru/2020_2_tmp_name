package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"park_2020/2020_2_tmp_name/domain"
	"park_2020/2020_2_tmp_name/models"
	"time"

	"github.com/google/uuid"
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

type userUsecase struct {
	userRepo       domain.UserRepository
	Hub            *models.Hub
	client         *models.Client
	contextTimeout time.Duration
}

func NewUserUsecase(u domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepo:       u,
		contextTimeout: timeout,
	}
}

func (u *userUsecase) Login(data models.LoginData) (string, error) {
	var check bool
	if check = u.userRepo.CheckUser(data.Telephone); !check {
		return "", domain.ErrUnauthorized
	}

	user, err := u.userRepo.SelectUser(data.Telephone)
	if err != nil {
		return "", domain.ErrInternalServerError
	}

	if !models.CheckPasswordHash(data.Password, user.Password) {
		return "", domain.ErrUnauthorized
	}

	SID, err := uuid.NewRandom()
	if err != nil {
		return "", domain.ErrInternalServerError
	}

	err = u.userRepo.InsertSession(SID.String(), data.Telephone)
	if err != nil {
		return "", domain.ErrInternalServerError
	}

	return SID.String(), nil
}

func (u *userUsecase) Logout(session string) error {
	return u.userRepo.DeleteSession(session)
}

func (u *userUsecase) Signup(user models.User) error {
	var check bool
	if check = u.userRepo.CheckUser(user.Telephone); check {
		return domain.ErrUnauthorized
	}

	err := u.userRepo.InsertUser(user)
	if err != nil {
		return domain.ErrInternalServerError
	}
	return nil
}

func (u *userUsecase) Settings(cookie string, userData models.User) error {
	telephone := u.userRepo.CheckUserBySession(cookie)
	user, err := u.userRepo.SelectUserFeed(telephone)
	if err != nil {
		return domain.ErrInternalServerError
	}

	err = u.userRepo.UpdateUser(userData, user.ID)
	if err != nil {
		return domain.ErrInternalServerError
	}
	return nil
}

func (u *userUsecase) Me(cookie string) (models.UserFeed, error) {
	telephone := u.userRepo.CheckUserBySession(cookie)
	user, err := u.userRepo.SelectUserFeed(telephone)
	if err != nil {
		return user, domain.ErrInternalServerError
	}
	return user, nil
}

func (u *userUsecase) Feed(cookie string) ([]models.UserFeed, error) {
	var data []models.UserFeed
	telephone := u.userRepo.CheckUserBySession(cookie)
	user, err := u.userRepo.SelectUser(telephone)
	if err != nil {
		return data, domain.ErrInternalServerError
	}

	data, err = u.userRepo.SelectUsers(user)
	if err != nil {
		return data, domain.ErrInternalServerError
	}
	return data, nil
}

func (u *userUsecase) AddPhoto(photo models.Photo) error {
	user, err := u.userRepo.SelectUserFeed(photo.Telephone)
	if err != nil {
		return domain.ErrInternalServerError
	}

	user.LinkImages = append(user.LinkImages, photo.Path)

	err = u.userRepo.InsertPhoto(photo.Path, user.ID)
	if err != nil {
		return domain.ErrInternalServerError
	}
	return nil
}

func (u *userUsecase) UploadAvatar() (uuid.UUID, error) {
	var photoID uuid.UUID
	str, err := os.Getwd()
	if err != nil {
		return photoID, domain.ErrInternalServerError
	}

	os.Chdir("/home/ubuntu/go/src/2020_2_tmp_name/static/avatars")

	photoID, err = uuid.NewRandom()
	if err != nil {
		return photoID, domain.ErrInternalServerError
	}

	os.Chdir(str)
	return photoID, nil
}

func (u *userUsecase) Like(cookie string, like models.Like) error {
	telephone := u.userRepo.CheckUserBySession(cookie)
	user, err := u.userRepo.SelectUserFeed(telephone)
	if err != nil {
		return domain.ErrInternalServerError
	}

	err = u.userRepo.InsertLike(user.ID, like.Uid2)
	if err != nil {
		return domain.ErrInternalServerError
	}

	if res := u.userRepo.Match(user.ID, like.Uid2); !res {
		fmt.Println("There is not match")
	} else {
		var chat models.Chat
		chat.Uid1 = user.ID
		chat.Uid2 = like.Uid2
		if !u.userRepo.CheckChat(chat) {
			err := u.userRepo.InsertChat(chat)
			if err != nil {
				return domain.ErrInternalServerError
			}
		}
	}
	return nil
}

func (u *userUsecase) Dislike(cookie string, dislike models.Dislike) error {
	telephone := u.userRepo.CheckUserBySession(cookie)
	user, err := u.userRepo.SelectUserFeed(telephone)
	if err != nil {
		return domain.ErrInternalServerError
	}

	err = u.userRepo.InsertDislike(user.ID, dislike.Uid2)
	if err != nil {
		return domain.ErrInternalServerError
	}
	return nil
}

func (u *userUsecase) Comment(cookie string, comment models.Comment) error {
	telephone := u.userRepo.CheckUserBySession(cookie)
	user, err := u.userRepo.SelectUserFeed(telephone)
	if err != nil {
		return domain.ErrInternalServerError
	}

	err = u.userRepo.InsertComment(comment, user.ID)
	if err != nil {
		return domain.ErrInternalServerError
	}
	return nil
}

func (u *userUsecase) CommentsByID(id int) (models.CommentsById, error) {
	comments, err := u.userRepo.SelectComments(id)
	if err != nil {
		return comments, domain.ErrInternalServerError
	}

	var data models.CommentsData
	data.Data = comments
	return comments, nil
}

func (u *userUsecase) Chat(chat models.Chat) error {
	err := u.userRepo.InsertChat(chat)
	if err != nil {
		return domain.ErrInternalServerError
	}
	return nil
}

func (u *userUsecase) Message(cookie string, message models.Message) error {
	telephone := u.userRepo.CheckUserBySession(cookie)
	user, err := u.userRepo.SelectUserFeed(telephone)
	if err != nil {
		return domain.ErrInternalServerError
	}

	err = u.userRepo.InsertMessage(message.Text, message.ChatID, user.ID)
	if err != nil {
		return domain.ErrInternalServerError
	}
	return nil
}

func (u *userUsecase) Chats(cookie string) (models.ChatModel, error) {
	var chatModel models.ChatModel
	telephone := u.userRepo.CheckUserBySession(cookie)
	user, err := u.userRepo.SelectUserFeed(telephone)
	if err != nil {
		return chatModel, domain.ErrInternalServerError
	}

	chats, err := u.userRepo.SelectChatsByID(user.ID)
	if err != nil {
		return chatModel, domain.ErrInternalServerError
	}

	chatModel.Data = chats
	return chatModel, nil
}

func (u *userUsecase) ChatID(cookie string, chid int) (models.ChatData, error) {
	var chat models.ChatData
	telephone := u.userRepo.CheckUserBySession(cookie)
	user, err := u.userRepo.SelectUserFeed(telephone)
	if err != nil {
		return chat, domain.ErrInternalServerError
	}

	chat, err = u.userRepo.SelectChatByID(user.ID, chid)
	if err != nil {
		return chat, domain.ErrInternalServerError
	}

	return chat, nil
}

func (u *userUsecase) Gochat(cookie string) (models.UserFeed, error) {
	telephone := u.userRepo.CheckUserBySession(cookie)
	user, err := u.userRepo.SelectUserFeed(telephone)
	if err != nil {
		return user, domain.ErrInternalServerError
	}
	return user, nil
}

func (u *userUsecase) ServeWs(hub *models.Hub, w http.ResponseWriter, r *http.Request, uid int) {
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

		err = u.userRepo.InsertMessage(msg.Message, msg.ChatID, msg.UserID)
		if err != nil {
			log.Println(err)
		}

	}
	if err == nil {
		u.Hub.Register <- client
	}

	client.Hub.Register <- client
	go u.writePump()
	go u.readPump()
}

// readPump pumps messages from the websocket connection to the hub.
func (u *userUsecase) readPump() {
	defer func() {
		u.client.Hub.Unregister <- u.client
		u.client.Conn.Close()
	}()
	u.client.Conn.SetReadLimit(maxMessageSize)
	u.client.Conn.SetReadDeadline(time.Now().Add(pongWait))
	u.client.Conn.SetPongHandler(func(string) error { u.client.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := u.client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		u.client.Hub.Broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
func (u *userUsecase) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		u.client.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-u.client.Send:
			u.client.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				u.client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := u.client.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(u.client.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-u.client.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			u.client.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := u.client.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
