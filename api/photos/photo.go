package domain

import (
	"park_2020/2020_2_tmp_name/models"
)

//go:generate mockgen -destination=./mock/mock_usecase.go -package=mock park_2020/2020_2_tmp_name/api/photos PhotoUsecase
//go:generate mockgen -destination=./mock/mock_repo.go -package=mock park_2020/2020_2_tmp_name/api/photos PhotoRepository

type PhotoUsecase interface {
	AddPhoto(photo models.Photo) error
	RemovePhoto(path string, uid int) error
	ClearPhotos(path string) error
	FindPhotoWithMask(path string) ([]string, error)
	FindPhotoWithoutMask(path string) (string, error)
}

type PhotoRepository interface {
	SelectUserFeed(telephone string) (models.UserFeed, error) // Tested
	SelectImages(uid int) ([]string, error)                   // Tested
	InsertPhoto(path string, uid int) error                   // Tested
	DeletePhoto(path string, uid int) error                   // Tested
	SelectPhotoWithMask(path string) ([]string, error)        // Tested
}
