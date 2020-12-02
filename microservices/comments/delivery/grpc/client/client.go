package client

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	proto "park_2020/2020_2_tmp_name/microservices/comments/delivery/grpc/protobuf"
	"park_2020/2020_2_tmp_name/models"
)

type CommentClient struct {
	client proto.CommentsGRPCHandlerClient
}

func NewCommentsClientGRPC(conn *grpc.ClientConn) *CommentClient{
	c := proto.NewCommentsGRPCHandlerClient(conn)
	return &CommentClient{
		client: c,
	}
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


func (c *CommentClient) Comment(ctx context.Context, userComment *proto.UserComment) error {
	_, err := c.client.Comment(ctx, userComment)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (c *CommentClient) CommentsById(ctx context.Context, id *proto.Id) (models.CommentsData, error) {
	commentsData, err := c.client.CommentsById(ctx, id)
	if err != nil {
		logrus.Error(err)
		return models.CommentsData{}, nil
	}

	return transformFromCommentsData(commentsData), nil
}