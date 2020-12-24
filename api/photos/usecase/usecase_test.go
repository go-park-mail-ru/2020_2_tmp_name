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

func TestPhotoUsecase_TestAddPhotoSuccess(t *testing.T) {
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

func TestPhotoUsecase_TestAddPhotoFail(t *testing.T) {
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

func TestPhotoUsecase_TestAddPhotoFailSelect(t *testing.T) {
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

func TestPhotoUsecase_TestRemovePhotoSuccess(t *testing.T) {
	path := "path"
	uid := 1

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockPhotoRepository(ctrl)
	mock.EXPECT().DeletePhoto(path, uid).Times(1).Return(nil)

	ps := photoUsecase{
		photoRepo: mock,
	}

	err := ps.RemovePhoto(path, uid)

	require.NoError(t, err)
}

func TestPhotoUsecase_TestRemovePhotoFail(t *testing.T) {
	path := "path"
	uid := 1

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockPhotoRepository(ctrl)
	mock.EXPECT().DeletePhoto(path, uid).Times(1).Return(models.ErrInternalServerError)

	ps := photoUsecase{
		photoRepo: mock,
	}

	err := ps.RemovePhoto(path, uid)
	require.NotEqual(t, err, nil)
}

func TestPhotoUsecase_TestFindPhotoWithMaskSuccess(t *testing.T) {
	path := "path"
	links := []string{"link1", "link2"}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockPhotoRepository(ctrl)
	mock.EXPECT().SelectPhotoWithMask(path).Times(1).Return(links, nil)

	ps := photoUsecase{
		photoRepo: mock,
	}

	result, err := ps.FindPhotoWithMask(path)

	require.NoError(t, err)
	require.Equal(t, links, result)
}

func TestPhotoUsecase_TestFindPhotoWitouthMaskSuccess(t *testing.T) {
	path := "path"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockPhotoRepository(ctrl)

	ps := photoUsecase{
		photoRepo: mock,
	}

	_, err := ps.FindPhotoWithoutMask(path)

	require.NotEqual(t, err, nil)
}
