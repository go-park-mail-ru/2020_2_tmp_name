package usecase

import (
	domain "park_2020/2020_2_tmp_name/api/likes"
	"park_2020/2020_2_tmp_name/models"
)

type likeUsecase struct {
	likeRepo domain.LikeRepository
}

func NewLikeUsecase(u domain.LikeRepository) domain.LikeUsecase {
	return &likeUsecase{
		likeRepo: u,
	}
}

func (l *likeUsecase) Like(cookie string, like models.Like) error {
	telephone := l.likeRepo.CheckUserBySession(cookie)
	user, err := l.likeRepo.SelectUserFeed(telephone)
	if err != nil {
		return models.ErrInternalServerError
	}

	err = l.likeRepo.InsertLike(user.ID, like.Uid2)
	if err != nil {
		return models.ErrInternalServerError
	}

	if res := l.likeRepo.Match(user.ID, like.Uid2); res {
		var chat models.Chat
		chat.Uid1 = user.ID
		chat.Uid2 = like.Uid2
		if !l.likeRepo.CheckChat(chat) {
			err := l.likeRepo.InsertChat(chat)
			if err != nil {
				return models.ErrInternalServerError
			}
		}
	}
	return nil
}

func (l *likeUsecase) Dislike(cookie string, dislike models.Dislike) error {
	telephone := l.likeRepo.CheckUserBySession(cookie)
	user, err := l.likeRepo.SelectUserFeed(telephone)
	if err != nil {
		return models.ErrInternalServerError
	}

	err = l.likeRepo.InsertDislike(user.ID, dislike.Uid2)
	if err != nil {
		return models.ErrInternalServerError
	}
	return nil
}