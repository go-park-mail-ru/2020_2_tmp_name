package domain

import (
	"park_2020/2020_2_tmp_name/models"
)

//go:generate mockgen -destination=./mock/mock_usecase.go -package=mock park_2020/2020_2_tmp_name/api/comments CommentUsecase
//go:generate mockgen -destination=./mock/mock_repo.go -package=mock park_2020/2020_2_tmp_name/api/comments CommentRepository

type CommentUsecase interface {
	Comment(user models.User, comment models.Comment) error
	CommentsByID(id int) (models.CommentsData, error)
	User(cookie string) (models.User, error)
}

type CommentRepository interface {
	CheckUserBySession(sid string) string
	SelectUser(telephone string) (models.User, error) // Tested
	SelectUserFeed(telephone string) (models.UserFeed, error)
	SelectUserFeedByID(uid int) (models.UserFeed, error)
	InsertComment(comment models.Comment, uid int) error // Tested
	SelectComments(userId int) (models.CommentsById, error)
	SelectImages(uid int) ([]string, error) // Tested
}
