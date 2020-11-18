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

func (c *commentUsecase) Comment(user models.User, comment models.Comment) error {
	err := c.commentRepo.InsertComment(comment, user.ID)
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
		return data, models.ErrNotFound
	}

	return data, nil
}

func (c *commentUsecase) User(cookie string) (models.User, error) {
	telephone := c.commentRepo.CheckUserBySession(cookie)
	user, err := c.commentRepo.SelectUser(telephone)
	if err != nil {
		return user, models.ErrNotFound
	}
	return user, nil
}
