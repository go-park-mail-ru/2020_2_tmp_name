package client

import (
	"context"
	proto "park_2020/2020_2_tmp_name/microservices/comments/delivery/grpc/protobuf"
	"park_2020/2020_2_tmp_name/models"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type CommentClient struct {
	client proto.CommentsGRPCHandlerClient
}

func NewCommentsClientGRPC(conn *grpc.ClientConn) *CommentClient {
	c := proto.NewCommentsGRPCHandlerClient(conn)
	return &CommentClient{
		client: c,
	}
}

func transformIntoUserComment(user models.User, comment models.Comment) *proto.UserComment {
	userProto := &proto.User{
		ID:         int32(user.ID),
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

	commentProto := &proto.Comment{
		ID:           int32(comment.ID),
		Uid1:         int32(comment.Uid1),
		Uid2:         int32(comment.Uid2),
		TimeDelivery: comment.TimeDelivery,
		CommentText:  comment.CommentText,
	}

	userComment := &proto.UserComment{
		User:    userProto,
		Comment: commentProto,
	}

	return userComment
}

func transformFromCommentsData(data *proto.CommentsData) models.CommentsData {
	if data == nil {
		return models.CommentsData{}
	}

	commentsById := make([]models.CommentById, 0, 1)
	for _, comment := range data.Data.CommentById {
		user := models.UserFeed{
			ID:          int(comment.User.ID),
			Name:        comment.User.Name,
			DateBirth:   int(comment.User.DateBirth),
			LinkImages:  comment.User.LinkImages,
			Job:         comment.User.Job,
			Education:   comment.User.Education,
			AboutMe:     comment.User.AboutMe,
			IsSuperlike: comment.User.IsSuperLike,
		}

		commentId := models.CommentById{
			User:         user,
			CommentText:  comment.CommentText,
			TimeDelivery: comment.TimeDelivery,
		}

		commentsById = append(commentsById, commentId)
	}

	commentsId := models.CommentsById{Comments: commentsById}
	comments := models.CommentsData{Data: commentsId}

	return comments
}

func (c *CommentClient) Comment(ctx context.Context, user models.User, comment models.Comment) error {
	userComment := transformIntoUserComment(user, comment)

	_, err := c.client.Comment(ctx, userComment)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (c *CommentClient) CommentsByID(ctx context.Context, id int) (models.CommentsData, error) {
	idProto := &proto.Id{Id: int32(id)}
	commentsData, err := c.client.CommentsById(ctx, idProto)
	if err != nil {
		logrus.Error(err)
		return models.CommentsData{}, nil
	}

	return transformFromCommentsData(commentsData), nil
}
