package usecase

import (
	"context"
	domain "park_2020/2020_2_tmp_name/microservices/comments"
	"park_2020/2020_2_tmp_name/microservices/comments/mock"
	"park_2020/2020_2_tmp_name/models"

	"github.com/golang/mock/gomock"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCommentUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var c domain.CommentRepository
	cu := NewCommentUsecase(c)
	require.Empty(t, cu)
}

func TestCommentUsecase_CommentSuccess(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	comment := models.Comment{
		ID:           0,
		Uid1:         1,
		Uid2:         2,
		TimeDelivery: "7:23",
		CommentText:  "I love tests very much",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockCommentRepository(ctrl)
	mock.EXPECT().InsertComment(comment, user.ID).Return(nil)

	cs := commentUsecase{
		commentRepo: mock,
	}

	err := cs.Comment(context.Background(), user, comment)

	require.NoError(t, err)
	require.Equal(t, nil, err)
}

func TestCommentUsecase_CommentFail(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	comment := models.Comment{
		ID:           0,
		Uid1:         1,
		Uid2:         2,
		TimeDelivery: "7:23",
		CommentText:  "I love tests very much",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockCommentRepository(ctrl)
	mock.EXPECT().InsertComment(comment, user.ID).Return(models.ErrInternalServerError)

	cs := commentUsecase{
		commentRepo: mock,
	}

	err := cs.Comment(context.Background(), user, comment)
	require.Equal(t, models.ErrInternalServerError, err)

}

func TestCommentUsecase_CommentsByIDSuccess(t *testing.T) {
	comments := models.CommentsById{}
	Data := models.CommentsData{}
	Data.Data = comments
	id := 2

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockCommentRepository(ctrl)
	mock.EXPECT().SelectComments(id).Return(comments, nil)

	cs := commentUsecase{
		commentRepo: mock,
	}

	data, err := cs.CommentsByID(context.Background(), id)

	require.NoError(t, err)
	require.Equal(t, Data, data)
}

func TestCommentUsecase_CommentsByIDFail(t *testing.T) {
	comments := models.CommentsById{}
	Data := models.CommentsData{}
	Data.Data = comments
	id := 2

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockCommentRepository(ctrl)
	mock.EXPECT().SelectComments(id).Return(comments, models.ErrNotFound)

	cs := commentUsecase{
		commentRepo: mock,
	}

	_, err := cs.CommentsByID(context.Background(), id)

	require.Equal(t, models.ErrNotFound, err)
}

func TestCommentUsecase_UserSuccess(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	telephone := "(944) 546 98 24"
	sid := "something-like-this"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockCommentRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Return(telephone)
	mock.EXPECT().SelectUser(telephone).Return(user, nil)

	chs := commentUsecase{
		commentRepo: mock,
	}

	result, err := chs.User(context.Background(), sid)

	require.NoError(t, err)
	require.Equal(t, result, user)
}

func TestCommentUsecase_UserFail(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	telephone := "(944) 546 98 24"
	sid := "something-like-this"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockCommentRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Return(telephone)
	mock.EXPECT().SelectUser(telephone).Return(user, models.ErrNotFound)

	chs := commentUsecase{
		commentRepo: mock,
	}

	_, err := chs.User(context.Background(), sid)

	require.Equal(t, err, models.ErrNotFound)
}
