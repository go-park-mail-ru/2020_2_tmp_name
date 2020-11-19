package usecase

import (
	domain "park_2020/2020_2_tmp_name/api/chats"
	"park_2020/2020_2_tmp_name/models"
	"time"
)

type chatUsecase struct {
	chatRepo       domain.ChatRepository
	Hub            *models.Hub
	client         *models.Client
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

func (ch *chatUsecase) User(cookie string) (models.User, error) {
	telephone := ch.chatRepo.CheckUserBySession(cookie)
	user, err := ch.chatRepo.SelectUser(telephone)
	if err != nil {
		return user, models.ErrNotFound
	}
	return user, nil
}
