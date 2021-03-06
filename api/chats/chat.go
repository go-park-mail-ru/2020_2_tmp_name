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
	UserFeed(cookie string) (models.UserFeed, error)
	Like(user models.User, like models.Like) error
	Dislike(user models.User, dislike models.Dislike) error
	MatchUser(user models.User, like models.Like) (models.Chat, bool, error)
	Superlike(user models.User, superlike models.Superlike) error
}

type ChatRepository interface {
	SelectUserFeed(telephone string) (models.UserFeed, error) // Tested
	SelectUserFeedByID(uid int) (models.UserFeed, error)      // Tested
	SelectImages(uid int) ([]string, error)                   // Tested
	InsertChat(chat models.Chat) error                        // Tested
	InsertMessage(text string, chatID, uid int) error         // Tested
	SelectMessages(chid int) ([]models.Msg, error)            // Tested
	SelectMessage(uid, chid int) (models.Msg, error)          // Tested
	InsertLike(uid1, uid2, fid int) error                     // Tested
	InsertDislike(uid1, uid2, fid int) error                  // Tested
	InsertSuperlike(uid1, uid2, fid int) error                // Tested
	DeleteLike(uid1, uid2 int) error                          // Tested
	DeleteDislike(uid1, uid2 int) error                       // Tested
	SelectUserByID(uid int) (models.User, error)              // Tested
	SelectChatID(uid1, uid2 int) (int, error)                 // Tested
	CheckUserBySession(sid string) string                     // Tested
	CheckLike(uid1, uid2 int) bool
	CheckDislike(uid1, uid2 int) bool
	CheckChat(chat models.Chat) bool
	SelectChatsByID(uid int) ([]models.ChatData, error)
	SelectChatByID(uid, chid int) (models.ChatData, error)
	SelectUserByChat(uid, chid int) (models.UserFeed, error)
	SelectSessions(uid int) ([]string, error)
	Match(uid1, uid2, fid int) bool
}
