package client

import (
	"context"
	"park_2020/2020_2_tmp_name/models"
)

//go:generate mockgen -destination=./mock/mock.go -package=mock park_2020/2020_2_tmp_name/microservices/comments/delivery/grpc/client CommentClientInterface

type CommentClientInterface interface {
	Comment(ctx context.Context, user models.User, comment models.Comment) error
	CommentsByID(ctx context.Context, id int) (models.CommentsData, error)
}
