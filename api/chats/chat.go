package domain

import (
	"park_2020/2020_2_tmp_name/models"
)

//go:generate mockgen -destination=./mock/mock_usecase.go -package=mock park_2020/2020_2_tmp_name/api/chats ChatUsecase
//go:generate mockgen -destination=./mock/mock_repo.go -package=mock park_2020/2020_2_tmp_name/api/chats ChatRepository

type ChatUsecase interface {
	Chat(chat models.Chat) error
	Message(user models.User, message models.Message) error
	Msg(user models.User, message models.Msg) error
	Chats(user models.User) (models.ChatModel, error)
	ChatID(user models.User, chid int) (models.ChatData, error)
	Partner(user models.User, chid int) (models.UserFeed, error)
	Sessions(uid int) ([]string, error)
	User(cookie string) (models.User, error)
}

type ChatRepository interface {
	SelectUserFeed(telephone string) (models.UserFeed, error) // Tested
	SelectUserFeedByID(uid int) (models.UserFeed, error)      // Tested
	SelectUser(telephone string) (models.User, error)         // Tested
	SelectImages(uid int) ([]string, error)                   // Tested
	CheckChat(chat models.Chat) bool
	InsertChat(chat models.Chat) error                // Tested
	InsertMessage(text string, chatID, uid int) error // Tested
	SelectMessage(uid, chid int) (models.Msg, error)  // Tested
	SelectMessages(chid int) ([]models.Msg, error)    // Tested
	SelectChatsByID(uid int) ([]models.ChatData, error)
	SelectChatByID(uid, chid int) (models.ChatData, error)
	SelectUserByChat(uid, chid int) (models.UserFeed, error)
	SelectUserByID(uid int) (models.User, error)
	SelectSessions(uid int) ([]string, error)
	CheckUserBySession(sid string) string
}
