package server

import (
	"context"
	"net"
	auth "park_2020/2020_2_tmp_name/microservices/authorization"
	proto "park_2020/2020_2_tmp_name/microservices/authorization/delivery/grpc/protobuf"
	"park_2020/2020_2_tmp_name/models"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

type server struct {
	authUseCase auth.UserUsecase
}

func NewAuthServerGRPC(gServer *grpc.Server, authUCase auth.UserUsecase) {
	articleServer := &server{
		authUseCase: authUCase,
	}
	proto.RegisterAuthGRPCHandlerServer(gServer, articleServer)
	reflection.Register(gServer)
}

func StartAuthGRPCServer(authUCase auth.UserUsecase, url string) {
	list, err := net.Listen("tcp", url)
	if err != nil {
		logrus.Error(err)
	}

	server := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
		}),
	)

	NewAuthServerGRPC(server, authUCase)

	_ = server.Serve(list)
}

func (s *server) Login(ctx context.Context, data *proto.LoginData) (*proto.Session, error) {
	var err error
	var session *proto.Session
	var loginData models.LoginData
	loginData.Password = data.Password
	loginData.Telephone = data.Telephone
	session.Sess, err = s.authUseCase.Login(ctx, loginData)
	return session, err
}

func (s *server) Logout(ctx context.Context, session *proto.Session) (*proto.Nothing, error) {
	var nothing *proto.Nothing
	err := s.authUseCase.Logout(ctx, session.Sess)
	return nothing, err
}

func (s *server) CheckSession(ctx context.Context, session *proto.Session) (*proto.User, error) {
	user, err := s.authUseCase.CheckSession(ctx, session.Sess)
	userProto := &proto.User{
		Id:         int32(user.ID),
		Name:       user.Name,
		Telephone:  user.Telephone,
		Password:   user.Password,
		DateBirth:  int32(user.DateBirth),
		Day:        user.Day,
		Month:      user.Month,
		Year:       user.Year,
		Sex:        user.Sex,
		LinkImages: user.LinkImages,
		Job:        user.Job,
		Education:  user.Education,
		AboutMe:    user.AboutMe,
	}
	return userProto, err
}
