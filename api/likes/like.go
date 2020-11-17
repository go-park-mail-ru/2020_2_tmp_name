package domain

import (
	"park_2020/2020_2_tmp_name/models"
)

//go:generate mockgen -destination=./mock/mock_usecase.go -package=mock park_2020/2020_2_tmp_name/api/likes LikeUsecase
//go:generate mockgen -destination=./mock/mock_repo.go -package=mock park_2020/2020_2_tmp_name/api/likes LikeRepository

type LikeUsecase interface {
	Like(cookie string, like models.Like) error
	Dislike(cookie string, dislike models.Dislike) error
}

type LikeRepository interface {
	SelectUserFeed(telephone string) (models.UserFeed, error)
	CheckUserBySession(sid string) string
	CheckChat(chat models.Chat) bool
	InsertChat(chat models.Chat) error
	Match(uid1, uid2 int) bool
	InsertLike(uid1, uid2 int) error
	InsertDislike(uid1, uid2 int) error
	SelectImages(uid int) ([]string, error)
}
