package domain

import (
	"net/http"
	"park_2020/2020_2_tmp_name/models"

	"github.com/google/uuid"
)

//go:generate mockgen -destination=./mock/mock_usecase.go -package=mock park_2020/2020_2_tmp_name/domain UserUsecase

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
