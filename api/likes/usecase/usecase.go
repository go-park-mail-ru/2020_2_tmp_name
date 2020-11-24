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

func (l *likeUsecase) Like(user models.User, like models.Like) error {
	err := l.likeRepo.InsertLike(user.ID, like.Uid2)
	if err != nil {
		return models.ErrInternalServerError
	}
	return nil
}

func (l *likeUsecase) MatchUser(user models.User, like models.Like) (models.Chat, bool, error) {
	var chat models.Chat
	if l.likeRepo.Match(user.ID, like.Uid2) {
		chat.Uid1 = user.ID
		chat.Uid2 = like.Uid2
		if !l.likeRepo.CheckChat(chat) {
			err := l.likeRepo.InsertChat(chat)
			if err != nil {
				return chat, false, models.ErrInternalServerError
			}
			chat.ID, err = l.likeRepo.SelectChatID(user.ID, like.Uid2)
			if err != nil {
				return chat, false, models.ErrNotFound
			}
			return chat, true, nil
		}
		return chat, false, nil
	}
	return chat, false, nil
}

func (l *likeUsecase) Partner(user models.User, chid int) (models.UserFeed, error) {
	partner, err := l.likeRepo.SelectUserByChat(user.ID, chid)
	if err != nil {
		return partner, models.ErrNotFound
	}
	return partner, nil
}

func (l *likeUsecase) Dislike(user models.User, dislike models.Dislike) error {
	err := l.likeRepo.InsertDislike(user.ID, dislike.Uid2)
	if err != nil {
		return models.ErrInternalServerError
	}
	return nil
}

func (l *likeUsecase) User(cookie string) (models.User, error) {
	telephone := l.likeRepo.CheckUserBySession(cookie)
	user, err := l.likeRepo.SelectUser(telephone)
	if err != nil {
		return user, models.ErrNotFound
	}
	return user, nil
}
