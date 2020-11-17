package domain

import (
	"park_2020/2020_2_tmp_name/models"
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
	UserID(uid int) (models.UserFeed, error)
}

type UserRepository interface {
	CheckUser(telephone string) bool
	InsertUser(user models.User) error
	SelectUser(telephone string) (models.User, error)
	SelectUserMe(telephone string) (models.UserMe, error)
	SelectUserFeed(telephone string) (models.UserFeed, error)
	SelectUserByID(uid int) (models.User, error)
	SelectUserFeedByID(uid int) (models.UserFeed, error)
	Match(uid1, uid2 int) bool
	SelectUsers(user models.User) ([]models.UserFeed, error)
	UpdateUser(user models.User, uid int) error
	InsertSession(sid, telephone string) error
	DeleteSession(sid string) error
	CheckUserBySession(sid string) string
	SelectImages(uid int) ([]string, error)
}
