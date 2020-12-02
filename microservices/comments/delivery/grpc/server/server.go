package comments

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
	comments "park_2020/2020_2_tmp_name/microservices/comments"
	proto "park_2020/2020_2_tmp_name/microservices/comments/delivery/grpc/protobuf"
	"park_2020/2020_2_tmp_name/models"
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
	user, comment := transformIntoUserComment(userComment)
	err := s.commentsUseCase.Comment(ctx, user, comment)
	if err != nil {
		return nil, err
	}
	empty := proto.Empty{}
	return &empty, nil
}

func (s *server) CommentsById(ctx context.Context, id *proto.Id) (*proto.CommentsData, error) {
	userId := int(id.Id)
	commentsData, err := s.commentsUseCase.CommentsByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	res := transformIntoGRPCCommentsData(commentsData)

	return res, nil
}

func transformIntoUserComment(userComment *proto.UserComment) (models.User, models.Comment) {
	if userComment == nil {
		return models.User{}, models.Comment{}
	}

	user := models.User{
		ID:         int(userComment.User.ID),
		Name:       userComment.User.Name,
		Telephone:  userComment.User.Telephone,
		Password:   userComment.User.Password,
		DateBirth:  int(userComment.User.DateBirth),
		Day:        userComment.User.Day,
		Month:      userComment.User.Month,
		Year:       userComment.User.Year,
		Sex:        userComment.User.Sex,
		LinkImages: userComment.User.LinkImages,
		Job:       	userComment.User.Job,
		Education:  userComment.User.Education,
		AboutMe:    userComment.User.AboutMe,
	}

	comment := models.Comment{
		ID:           int(userComment.Comment.ID),
		Uid1:         int(userComment.Comment.Uid1),
		Uid2:         int(userComment.Comment.Uid2),
		TimeDelivery: userComment.Comment.TimeDelivery,
		CommentText:  userComment.Comment.CommentText,
	}

	return user, comment
}

func transformIntoGRPCCommentsData(commentsData models.CommentsData) *proto.CommentsData {
	commentsById := make([]*proto.CommentId, 0, 1)
	for _, item := range commentsData.Data.Comments {
		user := &proto.UserFeed{
			ID:          int32(item.User.ID),
			Name:        item.User.Name,
			DateBirth:   int32(item.User.DateBirth),
			LinkImages:  item.User.LinkImages,
			Job:         item.User.Job,
			Education:   item.User.Education,
			AboutMe:     item.User.AboutMe,
			IsSuperLike: item.User.IsSuperlike,
		}
		commentId := &proto.CommentId{
			User: user,
			CommentText: item.CommentText,
			TimeDelivery: item.TimeDelivery,
		}
		commentsById = append(commentsById, commentId)
	}
	commentsId := &proto.CommentsById{CommentById: commentsById}
	comments := &proto.CommentsData{Data: commentsId}

	return comments
}