package usecase

import (
	"context"
	domain "park_2020/2020_2_tmp_name/microservices/comments"
	"park_2020/2020_2_tmp_name/models"
)

type commentUsecase struct {
	commentRepo domain.CommentRepository
}

func NewCommentUsecase(c domain.CommentRepository) domain.CommentUsecase {
	return &commentUsecase{
		commentRepo: c,
	}
}

func (c *commentUsecase) Comment(ctx context.Context, user models.User, comment models.Comment) error {
	err := c.commentRepo.InsertComment(comment, user.ID)
	if err != nil {
		return models.ErrInternalServerError
	}
	return nil
}

func (c *commentUsecase) CommentsByID(ctx context.Context, id int) (models.CommentsData, error) {
	var data models.CommentsData
	comments, err := c.commentRepo.SelectComments(id)
	if err != nil {
		return data, models.ErrNotFound
	}
	data.Data = comments

	return data, nil
}
