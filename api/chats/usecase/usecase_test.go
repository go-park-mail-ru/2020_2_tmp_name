package usecase

import (
	domain "park_2020/2020_2_tmp_name/api/chats"
	"park_2020/2020_2_tmp_name/api/chats/mock"
	"park_2020/2020_2_tmp_name/models"

	"github.com/golang/mock/gomock"

	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewChatUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var ch domain.ChatRepository
	chu := NewChatUsecase(ch)
	require.Empty(t, chu)
}

func TestChatUsecase_MessageSuccess(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	message := models.Message{
		ID:           0,
		Text:         "Save me from tests",
		TimeDelivery: time.Time{},
		ChatID:       1,
		UserID:       2,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().InsertMessage(message.Text, message.ChatID, user.ID).Return(nil)

	chs := chatUsecase{
		chatRepo: mock,
	}

	err := chs.Message(user, message)

	require.NoError(t, err)
	require.Equal(t, nil, err)
}

func TestChatUsecase_MessageFail(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	message := models.Message{
		ID:           0,
		Text:         "Save me from tests",
		TimeDelivery: time.Time{},
		ChatID:       1,
		UserID:       2,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().InsertMessage(message.Text, message.ChatID, user.ID).Return(models.ErrInternalServerError)

	chs := chatUsecase{
		chatRepo: mock,
	}

	err := chs.Message(user, message)
	require.Equal(t, models.ErrInternalServerError, err)

}

func TestChatUsecase_MsgSuccess(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	message := models.Msg{
		Message:      "Save me from tests",
		TimeDelivery: "time",
		ChatID:       1,
		UserID:       2,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().InsertMessage(message.Message, message.ChatID, user.ID).Return(nil)

	chs := chatUsecase{
		chatRepo: mock,
	}

	err := chs.Msg(user, message)

	require.NoError(t, err)
	require.Equal(t, nil, err)
}

func TestChatUsecase_MsgFail(t *testing.T) {
	user := models.User{}

	message := models.Msg{}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().InsertMessage(message.Message, message.ChatID, user.ID).Return(models.ErrInternalServerError)

	chs := chatUsecase{
		chatRepo: mock,
	}

	err := chs.Msg(user, message)
	require.Equal(t, models.ErrInternalServerError, err)
}

func TestChatUsecase_ChatSuccess(t *testing.T) {
	chat := models.Chat{
		ID:      0,
		Uid1:    4,
		Uid2:    3,
		LastMsg: "Save me from tests",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().InsertChat(chat).Times(1).Return(nil)

	chs := chatUsecase{
		chatRepo: mock,
	}

	err := chs.Chat(chat)

	require.NoError(t, err)
	require.Equal(t, nil, err)
}

func TestChatUsecase_ChatFail(t *testing.T) {
	chat := models.Chat{
		ID:      0,
		Uid1:    4,
		Uid2:    3,
		LastMsg: "Save me from tests",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().InsertChat(chat).Times(1).Return(models.ErrInternalServerError)

	chs := chatUsecase{
		chatRepo: mock,
	}

	err := chs.Chat(chat)

	require.NotEqual(t, nil, err)
}

func TestChatUsecase_СhatsSuccess(t *testing.T) {
	var chatModel models.ChatModel
	user := models.User{
		ID:         1,
		Name:       "Andrey",
		DateBirth:  20,
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	msg1 := models.Msg{
		UserID:       1,
		ChatID:       3,
		Message:      "Hi",
		TimeDelivery: "18:40",
	}

	msg2 := models.Msg{
		UserID:       6,
		ChatID:       3,
		Message:      "Hi",
		TimeDelivery: "18:41",
	}

	var chats []models.ChatData

	chat1 := models.ChatData{
		ID: 1,
		Partner: models.UserFeed{
			ID:         6,
			Name:       "Natasha",
			DateBirth:  20,
			LinkImages: nil,
			Job:        "",
			Education:  "BMSTU",
			AboutMe:    "",
		},
		Messages: []models.Msg{msg1, msg2},
	}

	chat2 := models.ChatData{
		ID: 2,
		Partner: models.UserFeed{
			ID:         4,
			Name:       "Dasha",
			DateBirth:  20,
			LinkImages: nil,
			Job:        "",
			Education:  "BMSTU",
			AboutMe:    "",
		},
		Messages: []models.Msg{msg1, msg2},
	}

	chats = append(chats, chat1, chat2)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().SelectChatsByID(user.ID).Times(1).Return(chats, nil)
	chatModel.Data = chats

	chs := chatUsecase{
		chatRepo: mock,
	}

	result, err := chs.Chats(user)

	require.NoError(t, err)
	require.Equal(t, result, chatModel)
}

func TestChatUsecase_СhatsFail(t *testing.T) {
	user := models.User{
		ID:         1,
		Name:       "Andrey",
		DateBirth:  20,
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	msg1 := models.Msg{
		UserID:       1,
		ChatID:       3,
		Message:      "Hi",
		TimeDelivery: "18:40",
	}

	msg2 := models.Msg{
		UserID:       6,
		ChatID:       3,
		Message:      "Hi",
		TimeDelivery: "18:41",
	}

	var chats []models.ChatData

	chat1 := models.ChatData{
		ID: 1,
		Partner: models.UserFeed{
			ID:         6,
			Name:       "Natasha",
			DateBirth:  20,
			LinkImages: nil,
			Job:        "",
			Education:  "BMSTU",
			AboutMe:    "",
		},
		Messages: []models.Msg{msg1, msg2},
	}

	chat2 := models.ChatData{
		ID: 2,
		Partner: models.UserFeed{
			ID:         4,
			Name:       "Dasha",
			DateBirth:  20,
			LinkImages: nil,
			Job:        "",
			Education:  "BMSTU",
			AboutMe:    "",
		},
		Messages: []models.Msg{msg1, msg2},
	}

	chats = append(chats, chat1, chat2)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().SelectChatsByID(user.ID).Times(1).Return(chats, models.ErrNotFound)

	chs := chatUsecase{
		chatRepo: mock,
	}

	_, err := chs.Chats(user)

	require.NotEqual(t, err, nil)
}

func TestChatUsecase_СhatID(t *testing.T) {
	var chid = 1
	user := models.User{
		ID:         1,
		Name:       "Andrey",
		DateBirth:  20,
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	msg1 := models.Msg{
		UserID:       1,
		ChatID:       3,
		Message:      "Hi",
		TimeDelivery: "18:40",
	}

	msg2 := models.Msg{
		UserID:       6,
		ChatID:       3,
		Message:      "Hi",
		TimeDelivery: "18:41",
	}

	chat := models.ChatData{
		ID: 1,
		Partner: models.UserFeed{
			ID:         6,
			Name:       "Natasha",
			DateBirth:  20,
			LinkImages: nil,
			Job:        "",
			Education:  "BMSTU",
			AboutMe:    "",
		},
		Messages: []models.Msg{msg1, msg2},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().SelectChatByID(user.ID, chid).Times(1).Return(chat, nil)

	chs := chatUsecase{
		chatRepo: mock,
	}

	result, err := chs.ChatID(user, chid)

	require.NoError(t, err)
	require.Equal(t, result, chat)
}

func TestChatUsecase_СhatIDFail(t *testing.T) {
	var chid = 1
	user := models.User{}
	chat := models.ChatData{}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().SelectChatByID(user.ID, chid).Times(1).Return(chat, models.ErrInternalServerError)

	chs := chatUsecase{
		chatRepo: mock,
	}

	_, err := chs.ChatID(user, chid)

	require.NotEqual(t, err, nil)
}

func TestChatUsecase_СhatSelectIDFail(t *testing.T) {
	var chid = 1
	user := models.User{}

	msg1 := models.Msg{}

	msg2 := models.Msg{}

	chat := models.ChatData{
		ID:       1,
		Partner:  models.UserFeed{},
		Messages: []models.Msg{msg1, msg2},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().SelectChatByID(user.ID, chid).Times(1).Return(chat, models.ErrInternalServerError)

	chs := chatUsecase{
		chatRepo: mock,
	}

	_, err := chs.ChatID(user, chid)

	require.NotEqual(t, err, nil)
}

func TestChatUsecase_UserSuccess(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	telephone := "(944) 546 98 24"
	sid := "something-like-this"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Return(telephone)
	mock.EXPECT().SelectUser(telephone).Return(user, nil)

	chs := chatUsecase{
		chatRepo: mock,
	}

	result, err := chs.User(sid)

	require.NoError(t, err)
	require.Equal(t, result, user)
}

func TestChatUsecase_UserFail(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	telephone := "(944) 546 98 24"
	sid := "something-like-this"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Return(telephone)
	mock.EXPECT().SelectUser(telephone).Return(user, models.ErrNotFound)

	chs := chatUsecase{
		chatRepo: mock,
	}

	_, err := chs.User(sid)

	require.Equal(t, err, models.ErrNotFound)
}
