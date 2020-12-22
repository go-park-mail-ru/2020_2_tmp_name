package domain

import (
	"context"
	"park_2020/2020_2_tmp_name/models"
)

//go:generate mockgen -destination=./mock/mock_usecase.go -package=mock park_2020/2020_2_tmp_name/microservices/comments CommentUsecase
//go:generate mockgen -destination=./mock/mock_repo.go -package=mock park_2020/2020_2_tmp_name/microservices/comments CommentRepository

type CommentUsecase interface {
	Comment(ctx context.Context, user models.User, comment models.Comment) error
	CommentsByID(ctx context.Context, id int) (models.CommentsData, error)
}

type CommentRepository interface {
	SelectUserFeed(telephone string) (models.UserFeed, error) // Tested
	SelectUserFeedByID(uid int) (models.UserFeed, error)      // Tested
	InsertComment(comment models.Comment, uid int) error      // Tested
	SelectImages(uid int) ([]string, error)                   // Tested
	SelectComments(userId int) (models.CommentsById, error)
}
