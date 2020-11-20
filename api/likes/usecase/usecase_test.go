package usecase

import (
	domain "park_2020/2020_2_tmp_name/api/likes"
	"park_2020/2020_2_tmp_name/api/likes/mock"
	"park_2020/2020_2_tmp_name/models"

	"github.com/golang/mock/gomock"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewLikeUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var l domain.LikeRepository
	lu := NewLikeUsecase(l)
	require.Empty(t, lu)
}

func TestLikeUsecase_LikeSuccess(t *testing.T) {
	like := models.Like{
		ID:   0,
		Uid1: 1,
		Uid2: 2,
	}

	user := models.User{
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
		Uid1:    user.ID,
		Uid2:    like.Uid2,
		LastMsg: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockLikeRepository(ctrl)
	mock.EXPECT().InsertLike(user.ID, like.Uid2).Return(nil)
	mock.EXPECT().Match(user.ID, like.Uid2).Return(true)
	mock.EXPECT().CheckChat(chat).Return(false)
	mock.EXPECT().InsertChat(chat).Return(nil)

	ls := likeUsecase{
		likeRepo: mock,
	}

	err := ls.Like(user, like)

	require.NoError(t, err)
	require.Equal(t, nil, err)

}

func TestLikeUsecase_LikeFail(t *testing.T) {
	like := models.Like{
		ID:   0,
		Uid1: 1,
		Uid2: 2,
	}

	user := models.User{
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
		Uid1:    user.ID,
		Uid2:    like.Uid2,
		LastMsg: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockLikeRepository(ctrl)
	mock.EXPECT().InsertLike(user.ID, like.Uid2).Return(nil)
	mock.EXPECT().Match(user.ID, like.Uid2).Return(true)
	mock.EXPECT().CheckChat(chat).Return(false)
	mock.EXPECT().InsertChat(chat).Return(models.ErrInternalServerError)

	ls := likeUsecase{
		likeRepo: mock,
	}

	err := ls.Like(user, like)
	require.Equal(t, models.ErrInternalServerError, err)
}

func TestLikeUsecase_DislikeSuccess(t *testing.T) {
	user := models.User{
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
		Uid1: user.ID,
		Uid2: 2,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockLikeRepository(ctrl)
	mock.EXPECT().InsertDislike(user.ID, dislike.Uid2).Return(nil)

	ls := likeUsecase{
		likeRepo: mock,
	}

	err := ls.Dislike(user, dislike)

	require.NoError(t, err)
	require.Equal(t, nil, err)
}

func TestLikeUsecase_DislikeFail(t *testing.T) {
	user := models.User{
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
		Uid1: user.ID,
		Uid2: 2,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockLikeRepository(ctrl)
	mock.EXPECT().InsertDislike(user.ID, dislike.Uid2).Return(models.ErrInternalServerError)

	ls := likeUsecase{
		likeRepo: mock,
	}

	err := ls.Dislike(user, dislike)
	require.Equal(t, models.ErrInternalServerError, err)

}
