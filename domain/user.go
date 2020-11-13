package domain

import (
	"net/http"
	"park_2020/2020_2_tmp_name/models"

	"github.com/google/uuid"
)

//go:generate mockgen -destination=./mock/mock_usecase.go -package=mock park_2020/2020_2_tmp_name/domain UserUsecase
//go:generate mockgen -destination=./mock/mock_repo.go -package=mock park_2020/2020_2_tmp_name/domain UserRepository

type UserUsecase interface {
	Login(data models.LoginData) (string, error)
	Logout(session string) error
	Signup(user models.User) error
	Settings(cookie string, user models.User) error
	Me(cookie string) (models.UserFeed, error)
	Feed(cookie string) ([]models.UserFeed, error)
	AddPhoto(photo models.Photo) error
	UploadAvatar() (uuid.UUID, error)
	UserID(uid int) (models.UserFeed, error)
	Like(cookie string, like models.Like) error
	Dislike(cookie string, dislike models.Dislike) error
	Comment(cookie string, comment models.Comment) error
	CommentsByID(id int) (models.CommentsData, error)
	Chat(chat models.Chat) error
	Message(cookie string, message models.Message) error
	Chats(cookie string) (models.ChatModel, error)
	ChatID(cookie string, chid int) (models.ChatData, error)
	Gochat(cookie string) (models.UserFeed, error)
	ServeWs(hub *models.Hub, w http.ResponseWriter, r *http.Request, uid int)
}

type UserRepository interface {
	CheckUser(telephone string) bool
	InsertUser(user models.User) error
	SelectUser(telephone string) (models.User, error)
	SelectUserMe(telephone string) (models.UserMe, error)
	SelectUserFeed(telephone string) (models.UserFeed, error)
	SelectUserByID(uid int) (models.User, error)
	Match(uid1, uid2 int) bool
	SelectUsers(user models.User) ([]models.UserFeed, error)
	SelectImages(uid int) ([]string, error)
	UpdateUser(user models.User, uid int) error
	InsertSession(sid, telephone string) error
	DeleteSession(sid string) error
	CheckUserBySession(sid string) string
	InsertLike(uid1, uid2 int) error
	InsertDislike(uid1, uid2 int) error
	InsertComment(comment models.Comment, uid int) error
	CheckChat(chat models.Chat) bool
	InsertChat(chat models.Chat) error
	InsertMessage(text string, chatID, uid int) error
	InsertPhoto(path string, uid int) error
	SelectMessage(uid, chid int) (models.Msg, error)
	SelectMessages(chid int) ([]models.Msg, error)
	SelectChatsByID(uid int) ([]models.ChatData, error)
	SelectChatByID(uid, chid int) (models.ChatData, error)
	SelectUserByChat(uid, chid int) (models.UserFeed, error)
	SelectComments(userId int) (models.CommentsById, error)
	SelectUserFeedByID(uid int) (models.UserFeed, error)
}
