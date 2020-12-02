package client

import (
	"context"
	auth "park_2020/2020_2_tmp_name/microservices/authorization/delivery/grpc/protobuf"
	"park_2020/2020_2_tmp_name/models"

	"google.golang.org/grpc"
)

type AuthClient struct {
	client auth.AuthGRPCHandlerClient
}

func NewAuthClient(conn *grpc.ClientConn) *AuthClient {
	c := auth.NewAuthGRPCHandlerClient(conn)
	return &AuthClient{
		client: c,
	}
}

type AuthClientInterface interface {
	Login(data models.LoginData) (string, error)
	Logout(session string) error
	CheckSession(cookie string) (auth.User, error)
}

func (ac *AuthClient) Login(ctx context.Context, in *auth.LoginData) (*auth.Session, error) {
	session, err := ac.client.Login(ctx, in, grpc.EmptyCallOption{})
	return session, err
}

func (ac *AuthClient) Logout(ctx context.Context, in *auth.Session) error {
	_, err := ac.client.Logout(ctx, in, grpc.EmptyCallOption{})
	return err
}

func (ac *AuthClient) CheckSession(ctx context.Context, in *auth.Session) (*auth.User, error) {
	user, err := ac.client.CheckSession(ctx, in, grpc.EmptyCallOption{})
	return user, err
}
