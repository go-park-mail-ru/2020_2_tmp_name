package usecase

import (
	domain "park_2020/2020_2_tmp_name/api/users"
	"park_2020/2020_2_tmp_name/models"
	"time"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(u domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: u,
	}
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

func (u *userUsecase) Settings(uid int, userData models.User) error {
	err := u.userRepo.UpdateUser(userData, uid)
	if err != nil {
		return models.ErrInternalServerError
	}
	return nil
}

func (u *userUsecase) IsPremium(uid int) bool {
	return u.userRepo.CheckPremium(uid)
}

func (u *userUsecase) Feed(user models.User) ([]models.UserFeed, error) {
	data, err := u.userRepo.SelectUsers(user)
	if err != nil {
		return data, models.ErrNotFound
	}

	for idxUserFeed, userFeed := range data {
		data[idxUserFeed].IsSuperlike = u.userRepo.CheckSuperLikeMe(user.ID, userFeed.ID)
	}

	return data, nil
}

func (u *userUsecase) UserID(uid int) (models.UserFeed, error) {
	user, err := u.userRepo.SelectUserFeedByID(uid)
	if err != nil {
		return user, models.ErrNotFound
	}
	return user, nil
}

func (u *userUsecase) Telephone(telephone string) bool {
	return u.userRepo.CheckUser(telephone)
}

func (u *userUsecase) GetPremium(uid int) error {
	return u.userRepo.InsertPremium(uid, time.Now(), time.Now())
}
