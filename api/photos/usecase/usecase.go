package usecase

import (
	domain "park_2020/2020_2_tmp_name/api/photos"
	"park_2020/2020_2_tmp_name/models"

	"github.com/google/uuid"
)

type photoUsecase struct {
	photoRepo domain.PhotoRepository
}

func NewPhotoUsecase(p domain.PhotoRepository) domain.PhotoUsecase {
	return &photoUsecase{
		photoRepo: p,
	}
}

func (p *photoUsecase) AddPhoto(photo models.Photo) error {
	user, err := p.photoRepo.SelectUserFeed(photo.Telephone)
	if err != nil {
		return models.ErrInternalServerError
	}

	user.LinkImages = append(user.LinkImages, photo.Path)

	err = p.photoRepo.InsertPhoto(photo.Path, user.ID)
	if err != nil {
		return models.ErrInternalServerError
	}
	return nil
}

func (p *photoUsecase) UploadAvatar() (uuid.UUID, error) {
	photoID, err := uuid.NewRandom()
	if err != nil {
		return photoID, models.ErrInternalServerError
	}

	return photoID, nil
}
