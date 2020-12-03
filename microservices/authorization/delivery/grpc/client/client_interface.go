package client

import (
	"context"
	"net/http"
	"park_2020/2020_2_tmp_name/models"
)

//go:generate mockgen -destination=./mock/mock.go -package=mock park_2020/2020_2_tmp_name/microservices/authorization/delivery/grpc/client AuthClientInterface

type AuthClientInterface interface {
	Login(ctx context.Context, in *models.LoginData) (string, error)
	Logout(ctx context.Context, in string) error
	CheckSession(ctx context.Context, in []*http.Cookie) (models.User, error)
}
