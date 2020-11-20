package usecase

import (
	domain "park_2020/2020_2_tmp_name/api/photos"
	"park_2020/2020_2_tmp_name/api/photos/mock"
	"park_2020/2020_2_tmp_name/models"

	"github.com/golang/mock/gomock"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPhotoUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var p domain.PhotoRepository
	pu := NewPhotoUsecase(p)
	require.Empty(t, pu)
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

	mock := mock.NewMockPhotoRepository(ctrl)
	mock.EXPECT().SelectUserFeed(photo.Telephone).Times(1).Return(user, nil)
	mock.EXPECT().InsertPhoto(photo.Path, user.ID).Times(1).Return(nil)

	ps := photoUsecase{
		photoRepo: mock,
	}

	err := ps.AddPhoto(photo)

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

	mock := mock.NewMockPhotoRepository(ctrl)
	mock.EXPECT().SelectUserFeed(photo.Telephone).Times(1).Return(user, models.ErrInternalServerError)

	ps := photoUsecase{
		photoRepo: mock,
	}

	err := ps.AddPhoto(photo)

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

	mock := mock.NewMockPhotoRepository(ctrl)
	mock.EXPECT().SelectUserFeed(photo.Telephone).Times(1).Return(user, nil)
	mock.EXPECT().InsertPhoto(photo.Path, user.ID).Times(1).Return(models.ErrInternalServerError)

	ps := photoUsecase{
		photoRepo: mock,
	}

	err := ps.AddPhoto(photo)

	require.NotEqual(t, err, nil)
}

func TestUploadAvatarSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockPhotoRepository(ctrl)

	ps := photoUsecase{
		photoRepo: mock,
	}

	uid, err := ps.UploadAvatar()

	require.NoError(t, err)
	require.NotEqual(t, uid.String(), "")
}
