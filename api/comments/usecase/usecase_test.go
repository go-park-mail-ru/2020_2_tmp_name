package usecase

import (
	"errors"
	"park_2020/2020_2_tmp_name/domain"
	"park_2020/2020_2_tmp_name/domain/mock"
	"park_2020/2020_2_tmp_name/models"

	"github.com/golang/mock/gomock"

	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewUserUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var u domain.UserRepository
	uu := NewUserUsecase(u, time.Duration(10*time.Second))
	require.NotEmpty(t, uu)
}

func TestUserUsecase_CommentSuccess(t *testing.T) {
	cookie := "Something-like-uuid"
	telephone := "909-277-47-21"

	userFeed := models.UserFeed{
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

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(cookie).Times(1).Return(telephone)
	mock.EXPECT().SelectUserFeed(telephone).Return(userFeed, nil)
	mock.EXPECT().InsertComment(comment, userFeed.ID).Return(nil)

	us := userUsecase{
		userRepo: mock,
	}

	err := us.Comment(cookie, comment)

	require.NoError(t, err)
	require.Equal(t, nil, err)
}

func TestUserUsecase_CommentFail(t *testing.T) {
	cookie := "Something-like-uuid"
	telephone := "909-277-47-21"

	userFeed := models.UserFeed{
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

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(cookie).Times(2).Return(telephone)
	gomock.InOrder(
		mock.EXPECT().SelectUserFeed(telephone).Return(userFeed, errors.New("error")),
		mock.EXPECT().SelectUserFeed(telephone).Return(userFeed, nil),
		mock.EXPECT().InsertComment(comment, userFeed.ID).Return(errors.New("error")),
	)

	us := userUsecase{
		userRepo: mock,
	}

	for i := 0; i < 2; i++ {
		err := us.Comment(cookie, comment)
		require.Equal(t, models.ErrInternalServerError, err)
	}

}

func TestUserUsecase_CommentsByIDSuccess(t *testing.T) {
	comments := models.CommentsById{}
	Data := models.CommentsData{}
	Data.Data = comments
	id := 2

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().SelectComments(id).Return(comments, nil)

	us := userUsecase{
		userRepo: mock,
	}

	data, err := us.CommentsByID(id)

	require.NoError(t, err)
	require.Equal(t, Data, data)
}

func TestUserUsecase_CommentsByIDFail(t *testing.T) {
	comments := models.CommentsById{}
	Data := models.CommentsData{}
	Data.Data = comments
	id := 2

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().SelectComments(id).Return(comments, errors.New("error"))

	us := userUsecase{
		userRepo: mock,
	}

	_, err := us.CommentsByID(id)

	require.Equal(t, models.ErrInternalServerError, err)
}
