package usecase

import (
	domain "park_2020/2020_2_tmp_name/microservices/authorization"
	"park_2020/2020_2_tmp_name/models"

	"github.com/google/uuid"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(u domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: u,
	}
}

func (u *userUsecase) Login(data models.LoginData) (string, error) {
	var check bool
	if check = u.userRepo.CheckUser(data.Telephone); !check {
		return "", models.ErrUnauthorized
	}

	user, err := u.userRepo.SelectUser(data.Telephone)
	if err != nil {
		return "", models.ErrNotFound
	}

	if !models.CheckPasswordHash(data.Password, user.Password) {
		return "", models.ErrUnauthorized
	}

	SID, err := uuid.NewRandom()
	if err != nil {
		return "", models.ErrInternalServerError
	}

	err = u.userRepo.InsertSession(SID.String(), data.Telephone)
	if err != nil {
		return "", models.ErrInternalServerError
	}

	return SID.String(), nil
}

func (u *userUsecase) Logout(session string) error {
	return u.userRepo.DeleteSession(session)
}
