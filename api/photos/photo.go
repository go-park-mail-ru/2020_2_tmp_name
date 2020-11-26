package domain

import (
	"park_2020/2020_2_tmp_name/models"
)

//go:generate mockgen -destination=./mock/mock_usecase.go -package=mock park_2020/2020_2_tmp_name/api/photos PhotoUsecase
//go:generate mockgen -destination=./mock/mock_repo.go -package=mock park_2020/2020_2_tmp_name/api/photos PhotoRepository

type PhotoUsecase interface {
	AddPhoto(photo models.Photo) error
}

type PhotoRepository interface {
	SelectUserFeed(telephone string) (models.UserFeed, error) // Tested
	SelectImages(uid int) ([]string, error)                   // Tested
	InsertPhoto(path string, uid int) error                   // Tested
}
