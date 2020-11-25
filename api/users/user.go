package domain

import (
	"park_2020/2020_2_tmp_name/models"
)

//go:generate mockgen -destination=./mock/mock_usecase.go -package=mock park_2020/2020_2_tmp_name/api/users UserUsecase
//go:generate mockgen -destination=./mock/mock_repo.go -package=mock park_2020/2020_2_tmp_name/api/users UserRepository

type UserUsecase interface {
	Login(data models.LoginData) (string, error)
	Logout(session string) error
	Signup(user models.User) error
	Settings(uid int, user models.User) error
	Me(cookie string) (models.UserFeed, error)
	Feed(user models.User) ([]models.UserFeed, error)
	UserID(uid int) (models.UserFeed, error)
	User(cookie string) (models.User, error)
	Telephone(telephone string) bool
}

type UserRepository interface {
	CheckUser(telephone string) bool
	InsertUser(user models.User) error                        // Tested
	SelectUser(telephone string) (models.User, error)         // Tested
	SelectUserMe(telephone string) (models.UserMe, error)     // Tested
	SelectUserFeed(telephone string) (models.UserFeed, error) // Tested
	SelectUserByID(uid int) (models.User, error)              // Tested
	SelectUserFeedByID(uid int) (models.UserFeed, error)      // Tested
	Match(uid1, uid2 int) bool
	SelectUsers(user models.User) ([]models.UserFeed, error)
	UpdateUser(user models.User, uid int) error // Tested
	InsertSession(sid, telephone string) error  // Tested
	DeleteSession(sid string) error
	CheckUserBySession(sid string) string
	SelectImages(uid int) ([]string, error) // Tested
}
