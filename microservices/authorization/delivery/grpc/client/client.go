package client

import (
	"context"
	"net/http"
	proto "park_2020/2020_2_tmp_name/microservices/authorization/delivery/grpc/protobuf"
	"park_2020/2020_2_tmp_name/models"

	"google.golang.org/grpc"
)

type AuthClient struct {
	client proto.AuthGRPCHandlerClient
}

func NewAuthClient(conn *grpc.ClientConn) *AuthClient {
	c := proto.NewAuthGRPCHandlerClient(conn)
	return &AuthClient{
		client: c,
	}
}

func (ac *AuthClient) Login(ctx context.Context, in *models.LoginData) (string, error) {
	if in == nil {
		return "", nil
	}

	loginData := transformIntoGRPCLoginData(in)

	session, err := ac.client.Login(ctx, loginData, grpc.EmptyCallOption{})
	if err != nil {
		return "", err
	}
	res := session.Sess
	return res, err
}

func (ac *AuthClient) Logout(ctx context.Context, in string) error {
	session := &proto.Session{
		Sess: in,
	}
	_, err := ac.client.Logout(ctx, session, grpc.EmptyCallOption{})
	return err
}

func (ac *AuthClient) CheckSession(ctx context.Context, in []*http.Cookie) (models.User, error) {
	if len(in) == 0 {
		return models.User{}, models.ErrUnauthorized
	}

	session := &proto.Session{Sess: in[0].Value}
	user, err := ac.client.CheckSession(ctx, session, grpc.EmptyCallOption{})
	return transformIntoUserModel(user), err
}

func transformIntoGRPCLoginData(data *models.LoginData) *proto.LoginData {
	if data == nil {
		return &proto.LoginData{}
	}

	loginDataProto := &proto.LoginData{
		Telephone: data.Telephone,
		Password:  data.Password,
	}
	return loginDataProto
}

func transformIntoLoginDataModel(data *proto.LoginData) models.LoginData {
	if data == nil {
		return models.LoginData{}
	}

	loginData := models.LoginData{
		Telephone: data.Telephone,
		Password:  data.Password,
	}

	return loginData
}

func transformIntoUserModel(user *proto.User) models.User {
	if user == nil {
		return models.User{}
	}
	userModel := models.User{
		ID:         int(user.Id),
		Name:       user.Name,
		Telephone:  user.Telephone,
		Password:   user.Password,
		DateBirth:  int(user.DateBirth),
		Day:        user.Day,
		Month:      user.Month,
		Year:       user.Year,
		Sex:        user.Sex,
		LinkImages: user.LinkImages,
		Job:        user.Job,
		Education:  user.Education,
		AboutMe:    user.AboutMe,
		Target: 	user.Target,
	}

	return userModel
}
