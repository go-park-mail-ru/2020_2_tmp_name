package client

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
	comments "park_2020/2020_2_tmp_name/microservices/comments"
	proto "park_2020/2020_2_tmp_name/microservices/comments/delivery/grpc/protobuf"
	"time"
)

type server struct {
	commentsUseCase comments.CommentUsecase
}

func NewCommentsServerGRPC(gServer *grpc.Server, commentsUCase comments.CommentUsecase) {
	articleServer := &server{
		commentsUseCase: commentsUCase,
	}
	proto.RegisterCommentsGRPCHandlerServer(gServer, articleServer)
	reflection.Register(gServer)
}

func StartCommentsGRPCServer(commentsUCase comments.CommentUsecase, url string) {
	list, err := net.Listen("tcp", url)
	if err != nil {
		logrus.Error(err)
	}

	server := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
		}),
	)

	NewCommentsServerGRPC(server, commentsUCase)

	_ = server.Serve(list)
}

func (s *server) Comment(ctx context.Context, userComment *proto.UserComment) (*proto.Empty, error) {

}

func (s *server) CommentsById(ctx context.Context, id *proto.Id) (*proto.CommentsData, error) {

}