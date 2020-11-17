package domain

import (
	"park_2020/2020_2_tmp_name/models"

	"github.com/google/uuid"
)

//go:generate mockgen -destination=./mock/mock_usecase.go -package=mock park_2020/2020_2_tmp_name/domain UserUsecase
//go:generate mockgen -destination=./mock/mock_repo.go -package=mock park_2020/2020_2_tmp_name/domain UserRepository

type PhotoUsecase interface {
	AddPhoto(photo models.Photo) error
	UploadAvatar() (uuid.UUID, error)
}

type PhotoRepository interface {
	SelectUserFeed(telephone string) (models.UserFeed, error)
	SelectImages(uid int) ([]string, error)
	InsertPhoto(path string, uid int) error
}
