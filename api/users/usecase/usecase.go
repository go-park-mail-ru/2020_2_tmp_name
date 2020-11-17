package usecase

import (
	domain "park_2020/2020_2_tmp_name/api/users"
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
		return "", models.ErrInternalServerError
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
		return "", models.ErrBadParamInput
	}

	return SID.String(), nil
}

func (u *userUsecase) Logout(session string) error {
	return u.userRepo.DeleteSession(session)
}

func (u *userUsecase) Signup(user models.User) error {
	var check bool
	if check = u.userRepo.CheckUser(user.Telephone); check {
		return models.ErrUnauthorized
	}

	err := u.userRepo.InsertUser(user)
	if err != nil {
		return models.ErrInternalServerError
	}
	return nil
}

func (u *userUsecase) Settings(cookie string, userData models.User) error {
	telephone := u.userRepo.CheckUserBySession(cookie)
	user, err := u.userRepo.SelectUserFeed(telephone)
	if err != nil {
		return models.ErrInternalServerError
	}

	err = u.userRepo.UpdateUser(userData, user.ID)
	if err != nil {
		return models.ErrInternalServerError
	}
	return nil
}

func (u *userUsecase) Me(cookie string) (models.UserFeed, error) {
	telephone := u.userRepo.CheckUserBySession(cookie)
	user, err := u.userRepo.SelectUserFeed(telephone)
	if err != nil {
		return user, models.ErrInternalServerError
	}
	return user, nil
}

func (u *userUsecase) Feed(cookie string) ([]models.UserFeed, error) {
	var data []models.UserFeed
	telephone := u.userRepo.CheckUserBySession(cookie)
	user, err := u.userRepo.SelectUser(telephone)
	if err != nil {
		return data, models.ErrInternalServerError
	}

	data, err = u.userRepo.SelectUsers(user)
	if err != nil {
		return data, models.ErrInternalServerError
	}
	return data, nil
}

func (u *userUsecase) UserID(uid int) (models.UserFeed, error) {
	user, err := u.userRepo.SelectUserFeedByID(uid)
	if err != nil {
		return user, models.ErrInternalServerError
	}
	return user, nil
}
