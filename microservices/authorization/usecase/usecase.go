package usecase

import (
	domain "park_2020/2020_2_tmp_name/microservices/authorization"
	authClient "park_2020/2020_2_tmp_name/microservices/authorization/delivery/grpc/client"
	"park_2020/2020_2_tmp_name/models"

	"github.com/google/uuid"
)

type userUsecase struct {
	userRepo   domain.UserRepository
	authClient *authClient.AuthClient
}

func NewAuthUsecase(u domain.UserRepository, ac *authClient.AuthClient) *userUsecase {
	return &userUsecase{
		userRepo:   u,
		authClient: ac,
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

func (u *userUsecase) CheckSession(cookie string) (models.User, error) {
	telephone := u.userRepo.CheckUserBySession(cookie)
	user, err := u.userRepo.SelectUser(telephone)
	if err != nil {
		return user, models.ErrNotFound
	}
	return user, nil
}
