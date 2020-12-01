package client

import (
	"context"
	auth "park_2020/2020_2_tmp_name/microservices/authorization/delivery/grpc/protobuf"

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

func (ac *AuthClient) Login(ctx context.Context, in *auth.LoginData) (*auth.Session, error) {
	session, err := ac.client.Login(ctx, in, grpc.EmptyCallOption{})
	return session, err
}

func (ac *AuthClient) Logout(ctx context.Context, in *auth.Session) error {
	_, err := ac.client.Logout(ctx, in, grpc.EmptyCallOption{})
	return err
}
