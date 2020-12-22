package usecase

import (
	"context"
	domain "park_2020/2020_2_tmp_name/microservices/authorization"
	"park_2020/2020_2_tmp_name/models"

	"github.com/google/uuid"
)

type authUsecase struct {
	userRepo domain.AuthRepository
}

func NewAuthUsecase(u domain.AuthRepository) *authUsecase {
	return &authUsecase{
		userRepo: u,
	}
}

func (a *authUsecase) Login(ctx context.Context, data models.LoginData) (string, error) {
	var check bool
	if check = a.userRepo.CheckUser(data.Telephone); !check {
		return "", models.ErrUnauthorized
	}

	user, err := a.userRepo.SelectUser(data.Telephone)
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

	err = a.userRepo.InsertSession(SID.String(), data.Telephone)
	if err != nil {
		return "", models.ErrInternalServerError
	}

	return SID.String(), nil
}

func (a *authUsecase) Logout(ctx context.Context, session string) error {
	return a.userRepo.DeleteSession(session)
}

func (a *authUsecase) CheckSession(ctx context.Context, cookie string) (models.User, error) {
	telephone := a.userRepo.CheckUserBySession(cookie)
	user, err := a.userRepo.SelectUser(telephone)
	if err != nil {
		return user, models.ErrNotFound
	}
	return user, nil
}
