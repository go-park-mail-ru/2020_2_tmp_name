package domain

import (
	"net/http"
	"park_2020/2020_2_tmp_name/models"
)

//go:generate mockgen -destination=./mock/mock_usecase.go -package=mock park_2020/2020_2_tmp_name/domain UserUsecase
//go:generate mockgen -destination=./mock/mock_repo.go -package=mock park_2020/2020_2_tmp_name/domain UserRepository

type ChatUsecase interface {
	Chat(chat models.Chat) error
	Message(cookie string, message models.Message) error
	Chats(cookie string) (models.ChatModel, error)
	ChatID(cookie string, chid int) (models.ChatData, error)
	Gochat(cookie string) (models.UserFeed, error)
	ServeWs(hub *models.Hub, w http.ResponseWriter, r *http.Request, uid int)
}

type ChatRepository interface {
	SelectUserFeed(telephone string) (models.UserFeed, error)
	SelectUserFeedByID(uid int) (models.UserFeed, error)
	SelectImages(uid int) ([]string, error)
	CheckChat(chat models.Chat) bool
	InsertChat(chat models.Chat) error
	InsertMessage(text string, chatID, uid int) error
	SelectMessage(uid, chid int) (models.Msg, error)
	SelectMessages(chid int) ([]models.Msg, error)
	SelectChatsByID(uid int) ([]models.ChatData, error)
	SelectChatByID(uid, chid int) (models.ChatData, error)
	SelectUserByChat(uid, chid int) (models.UserFeed, error)
	CheckUserBySession(sid string) string
}
