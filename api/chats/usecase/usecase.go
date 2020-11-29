package usecase

import (
	domain "park_2020/2020_2_tmp_name/api/chats"
	"park_2020/2020_2_tmp_name/models"
	"time"
)

type chatUsecase struct {
	chatRepo       domain.ChatRepository
	contextTimeout time.Duration
}

func NewChatUsecase(ch domain.ChatRepository) domain.ChatUsecase {
	return &chatUsecase{
		chatRepo: ch,
	}
}

func (ch *chatUsecase) Chat(chat models.Chat) error {
	err := ch.chatRepo.InsertChat(chat)
	if err != nil {
		return models.ErrInternalServerError
	}
	return nil
}

func (ch *chatUsecase) Message(user models.User, message models.Message) error {
	err := ch.chatRepo.InsertMessage(message.Text, message.ChatID, user.ID)
	if err != nil {
		return models.ErrInternalServerError
	}
	return nil
}

func (ch *chatUsecase) Msg(user models.User, message models.Msg) error {
	err := ch.chatRepo.InsertMessage(message.Message, message.ChatID, user.ID)
	if err != nil {
		return models.ErrInternalServerError
	}
	return nil
}

func (ch *chatUsecase) Sessions(uid int) ([]string, error) {
	sessions, err := ch.chatRepo.SelectSessions(uid)
	if err != nil {
		return sessions, models.ErrNotFound
	}
	return sessions, nil
}

func (ch *chatUsecase) Chats(user models.User) (models.ChatModel, error) {
	var chatModel models.ChatModel
	chats, err := ch.chatRepo.SelectChatsByID(user.ID)
	if err != nil {
		return chatModel, models.ErrNotFound
	}

	chatModel.Data = chats
	return chatModel, nil
}

func (ch *chatUsecase) ChatID(user models.User, chid int) (models.ChatData, error) {
	var chat models.ChatData
	chat, err := ch.chatRepo.SelectChatByID(user.ID, chid)
	if err != nil {
		return chat, models.ErrNotFound
	}

	return chat, nil
}

func (ch *chatUsecase) Partner(user models.User, chid int) (models.UserFeed, error) {
	partner, err := ch.chatRepo.SelectUserByChat(user.ID, chid)
	if err != nil {
		return partner, models.ErrNotFound
	}
	return partner, nil
}

func (ch *chatUsecase) User(cookie string) (models.User, error) {
	telephone := ch.chatRepo.CheckUserBySession(cookie)
	user, err := ch.chatRepo.SelectUser(telephone)
	if err != nil {
		return user, models.ErrNotFound
	}
	return user, nil
}

func (ch *chatUsecase) UserFeed(cookie string) (models.UserFeed, error) {
	telephone := ch.chatRepo.CheckUserBySession(cookie)
	user, err := ch.chatRepo.SelectUserFeed(telephone)
	if err != nil {
		return user, models.ErrNotFound
	}
	return user, nil
}

func (ch *chatUsecase) Like(user models.User, like models.Like) error {
	err := ch.chatRepo.InsertLike(user.ID, like.Uid2)
	if err != nil {
		return models.ErrInternalServerError
	}
	return nil
}

func (ch *chatUsecase) MatchUser(user models.User, like models.Like) (models.Chat, bool, error) {
	var chat models.Chat
	if ch.chatRepo.Match(user.ID, like.Uid2) {
		chat.Uid1 = user.ID
		chat.Uid2 = like.Uid2
		if !ch.chatRepo.CheckChat(chat) {
			err := ch.chatRepo.InsertChat(chat)
			if err != nil {
				return chat, false, models.ErrInternalServerError
			}
			chat.ID, err = ch.chatRepo.SelectChatID(user.ID, like.Uid2)
			if err != nil {
				return chat, false, models.ErrNotFound
			}
			return chat, true, nil
		}
		return chat, false, nil
	}
	return chat, false, nil
}

func (ch *chatUsecase) Dislike(user models.User, dislike models.Dislike) error {
	err := ch.chatRepo.InsertDislike(user.ID, dislike.Uid2)
	if err != nil {
		return models.ErrInternalServerError
	}
	return nil
}

func (ch *chatUsecase) SuperLike(user models.User, superLike models.SuperLike) error {
	err := ch.chatRepo.InsertSuperLike(user.ID, superLike.Uid2)
	if err != nil {
		return models.ErrInternalServerError
	}

	return nil
}