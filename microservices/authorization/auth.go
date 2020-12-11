package auth

import (
	"context"
	"park_2020/2020_2_tmp_name/models"
)

//go:generate mockgen -destination=./mock/mock_usecase.go -package=mock park_2020/2020_2_tmp_name/microservices/authorization AuthUsecase
//go:generate mockgen -destination=./mock/mock_repo.go -package=mock park_2020/2020_2_tmp_name/microservices/authorization AuthRepository

type AuthUsecase interface {
	Login(ctx context.Context, data models.LoginData) (string, error)
	Logout(ctx context.Context, session string) error
	CheckSession(ctx context.Context, cookie string) (models.User, error)
}

type AuthRepository interface {
	SelectUser(telephone string) (models.User, error) // Tested
	InsertSession(sid, telephone string) error        // Tested
	SelectImages(uid int) ([]string, error)           // Tested
	DeleteSession(sid string) error                   // Tested
	SelectUserBySession(sid string) (string, error)   // Tested
	CheckUser(telephone string) bool
	CheckUserBySession(sid string) string
}
