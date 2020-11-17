package usecase

import (
	domain "park_2020/2020_2_tmp_name/api/comments"
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

func (c *commentUsecase) Comment(cookie string, comment models.Comment) error {
	telephone := c.commentRepo.CheckUserBySession(cookie)
	user, err := c.commentRepo.SelectUserFeed(telephone)
	if err != nil {
		return models.ErrInternalServerError
	}

	err = c.commentRepo.InsertComment(comment, user.ID)
	if err != nil {
		return models.ErrInternalServerError
	}
	return nil
}

func (c *commentUsecase) CommentsByID(id int) (models.CommentsData, error) {
	comments, err := c.commentRepo.SelectComments(id)
	var data models.CommentsData
	data.Data = comments

	if err != nil {
		return data, models.ErrInternalServerError
	}

	return data, nil
}
