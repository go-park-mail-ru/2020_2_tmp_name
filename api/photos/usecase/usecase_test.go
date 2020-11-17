package usecase

import (
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

func TestAddPhotoSuccess(t *testing.T) {
	photo := models.Photo{
		Path:      "path",
		Telephone: "944-739-32-28",
	}

	user := models.UserFeed{
		ID:         1,
		Name:       "Andrey",
		DateBirth:  20,
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().SelectUserFeed(photo.Telephone).Times(1).Return(user, nil)
	mock.EXPECT().InsertPhoto(photo.Path, user.ID).Times(1).Return(nil)

	us := userUsecase{
		userRepo: mock,
	}

	err := us.AddPhoto(photo)

	require.NoError(t, err)
}

func TestAddPhotoFail(t *testing.T) {
	photo := models.Photo{
		Path:      "path",
		Telephone: "944-739-32-28",
	}

	user := models.UserFeed{
		ID:         1,
		Name:       "Andrey",
		DateBirth:  20,
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().SelectUserFeed(photo.Telephone).Times(1).Return(user, models.ErrInternalServerError)

	us := userUsecase{
		userRepo: mock,
	}

	err := us.AddPhoto(photo)

	require.NotEqual(t, err, nil)
}

func TestAddPhotoFailSelect(t *testing.T) {
	photo := models.Photo{
		Path:      "path",
		Telephone: "944-739-32-28",
	}

	user := models.UserFeed{
		ID:         1,
		Name:       "Andrey",
		DateBirth:  20,
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().SelectUserFeed(photo.Telephone).Times(1).Return(user, nil)
	mock.EXPECT().InsertPhoto(photo.Path, user.ID).Times(1).Return(models.ErrInternalServerError)

	us := userUsecase{
		userRepo: mock,
	}

	err := us.AddPhoto(photo)

	require.NotEqual(t, err, nil)
}

func TestUploadAvatarSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)

	us := userUsecase{
		userRepo: mock,
	}

	uid, err := us.UploadAvatar()

	require.NoError(t, err)
	require.NotEqual(t, uid.String(), "")
}
