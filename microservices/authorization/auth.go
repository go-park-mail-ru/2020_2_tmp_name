package auth

import (
	"context"
	"park_2020/2020_2_tmp_name/models"
)

type UserUsecase interface {
	Login(ctx context.Context, data models.LoginData) (string, error)
	Logout(ctx context.Context, session string) error
	CheckSession(ctx context.Context, cookie string) (models.User, error)
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