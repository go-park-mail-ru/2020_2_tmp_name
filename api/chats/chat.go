package domain

import (
	"park_2020/2020_2_tmp_name/models"
)

//go:generate mockgen -destination=./mock/mock_usecase.go -package=mock park_2020/2020_2_tmp_name/api/chats ChatUsecase
//go:generate mockgen -destination=./mock/mock_repo.go -package=mock park_2020/2020_2_tmp_name/api/chats ChatRepository

type ChatUsecase interface {
	Chat(chat models.Chat) error
	Message(user models.User, message models.Message) error
	Chats(user models.User) (models.ChatModel, error)
	ChatID(user models.User, chid int) (models.ChatData, error)
	User(cookie string) (models.User, error)
}

type ChatRepository interface {
	SelectUserFeed(telephone string) (models.UserFeed, error)
	SelectUserFeedByID(uid int) (models.UserFeed, error)
	SelectUser(telephone string) (models.User, error)
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
