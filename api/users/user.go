package domain

import (
	"park_2020/2020_2_tmp_name/models"
	"time"
)

//go:generate mockgen -destination=./mock/mock_usecase.go -package=mock park_2020/2020_2_tmp_name/api/users UserUsecase
//go:generate mockgen -destination=./mock/mock_repo.go -package=mock park_2020/2020_2_tmp_name/api/users UserRepository

type UserUsecase interface {
	Signup(user models.User) error
	Settings(uid int, user models.User) error
	IsPremium(uid int) bool
	Feed(user models.User) ([]models.UserFeed, error)
	UserID(uid int) (models.UserFeed, error)
	Telephone(telephone string) bool
	GetPremium(uid int) error
}

type UserRepository interface {
	InsertUser(user models.User) error                   // Tested
	InsertSession(sid, telephone string) error           // Tested
	UpdateUser(user models.User, uid int) error          // Tested
	SelectUserByID(uid int) (models.User, error)         // Tested
	SelectUserFeedByID(uid int) (models.UserFeed, error) // Tested
	SelectImages(uid int) ([]string, error)              // Tested
	DeleteSession(sid string) error                      // Tested
	CheckUser(telephone string) bool
	CheckPremium(uid int) bool
	SelectUsers(user models.User) ([]models.UserFeed, error)           // Tested
	InsertPremium(uid int, dateFrom time.Time, dateTo time.Time) error // Tested
	CheckSuperLikeMe(me, userId int) bool
}
