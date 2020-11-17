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

func TestUserUsecase_LikeSuccess(t *testing.T) {
	like := models.Like{
		ID:   0,
		Uid1: 1,
		Uid2: 2,
	}

	userFeed := models.UserFeed{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	chat := models.Chat{
		ID:      0,
		Uid1:    userFeed.ID,
		Uid2:    like.Uid2,
		LastMsg: "",
	}
	cookie := "Something-like-uuid"
	telephone := "909-277-47-21"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	firstCall := mock.EXPECT().CheckUserBySession(cookie).Return(telephone)
	mock.EXPECT().SelectUserFeed(telephone).After(firstCall).Return(userFeed, nil)
	mock.EXPECT().InsertLike(userFeed.ID, like.Uid2).Return(nil)
	mock.EXPECT().Match(userFeed.ID, like.Uid2).Return(true)
	mock.EXPECT().CheckChat(chat).Return(false)
	mock.EXPECT().InsertChat(chat).Return(nil)

	us := userUsecase{
		userRepo: mock,
	}

	err := us.Like(cookie, like)

	require.NoError(t, err)
	require.Equal(t, nil, err)

}

func TestUserUsecase_LikeFail(t *testing.T) {
	like := models.Like{
		ID:   0,
		Uid1: 1,
		Uid2: 2,
	}

	userFeed := models.UserFeed{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}
	chat := models.Chat{
		ID:      0,
		Uid1:    userFeed.ID,
		Uid2:    like.Uid2,
		LastMsg: "",
	}

	cookie := "Something-like-uuid"
	telephone := "909-277-47-21"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(cookie).Times(3).Return(telephone)
	gomock.InOrder(
		mock.EXPECT().SelectUserFeed(telephone).Return(userFeed, errors.New("error select user")),
		mock.EXPECT().SelectUserFeed(telephone).Return(userFeed, nil),
		mock.EXPECT().InsertLike(userFeed.ID, like.Uid2).Return(errors.New("error of insert")),
		mock.EXPECT().SelectUserFeed(telephone).Return(userFeed, nil),
		mock.EXPECT().InsertLike(userFeed.ID, like.Uid2).Return(nil),
		mock.EXPECT().Match(userFeed.ID, like.Uid2).Return(true),
		mock.EXPECT().CheckChat(chat).Return(false),
		mock.EXPECT().InsertChat(chat).Return(errors.New("error of insert")),
	)

	us := userUsecase{
		userRepo: mock,
	}

	for i := 0; i < 3; i++ {
		err := us.Like(cookie, like)
		require.Equal(t, models.ErrInternalServerError, err)
	}
}

func TestUserUsecase_DislikeSuccess(t *testing.T) {
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

	dislike := models.Dislike{
		ID:   0,
		Uid1: userFeed.ID,
		Uid2: 2,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(cookie).Times(1).Return(telephone)
	mock.EXPECT().SelectUserFeed(telephone).Return(userFeed, nil)
	mock.EXPECT().InsertDislike(userFeed.ID, dislike.Uid2).Return(nil)

	us := userUsecase{
		userRepo: mock,
	}

	err := us.Dislike(cookie, dislike)

	require.NoError(t, err)
	require.Equal(t, nil, err)
}

func TestUserUsecase_DislikeFail(t *testing.T) {
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

	dislike := models.Dislike{
		ID:   0,
		Uid1: userFeed.ID,
		Uid2: 2,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(cookie).Times(2).Return(telephone)
	gomock.InOrder(
		mock.EXPECT().SelectUserFeed(telephone).Return(userFeed, errors.New("error")),
		mock.EXPECT().SelectUserFeed(telephone).Return(userFeed, nil),
		mock.EXPECT().InsertDislike(userFeed.ID, dislike.Uid2).Return(errors.New("error")),
	)

	us := userUsecase{
		userRepo: mock,
	}

	for i := 0; i < 2; i++ {
		err := us.Dislike(cookie, dislike)
		require.Equal(t, models.ErrInternalServerError, err)
	}

}
