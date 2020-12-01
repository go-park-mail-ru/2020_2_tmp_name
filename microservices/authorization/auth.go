package domain

import (
	"park_2020/2020_2_tmp_name/models"
)

type UserUsecase interface {
	Login(data models.LoginData) (string, error)
	Logout(session string) error
	CheckSession(cookie string) (models.User, error)
}

type UserRepository interface {
	CheckUser(telephone string) bool
	SelectUser(telephone string) (models.User, error) // Tested
	InsertSession(sid, telephone string) error        // Tested
	DeleteSession(sid string) error
	CheckUserBySession(sid string) string
	SelectUserBySession(sid string) (string, error)
	SelectImages(uid int) ([]string, error)
}
