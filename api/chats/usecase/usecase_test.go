package usecase

import (
	"errors"
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

	var u domain.ChatRepository
	uu := NewChatUsecase(u)
	require.NotEmpty(t, uu)
}

func TestChatUsecase_MessageSuccess(t *testing.T) {
	cookie := "Something-like-uuid"
	telephone := "909-277-47-21"

	userFeed := models.UserFeed{
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
	mock.EXPECT().CheckUserBySession(cookie).Times(1).Return(telephone)
	mock.EXPECT().SelectUserFeed(telephone).Return(userFeed, nil)
	mock.EXPECT().InsertMessage(message.Text, message.ChatID, userFeed.ID).Return(nil)

	chs := chatUsecase{
		chatRepo: mock,
	}

	err := chs.Message(cookie, message)

	require.NoError(t, err)
	require.Equal(t, nil, err)

}

func TestChatUsecase_MessageFail(t *testing.T) {
	cookie := "Something-like-uuid"
	telephone := "909-277-47-21"

	userFeed := models.UserFeed{
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
	mock.EXPECT().CheckUserBySession(cookie).Times(2).Return(telephone)
	gomock.InOrder(
		mock.EXPECT().SelectUserFeed(telephone).Return(userFeed, errors.New("error")),
		mock.EXPECT().SelectUserFeed(telephone).Return(userFeed, nil),
		mock.EXPECT().InsertMessage(message.Text, message.ChatID, userFeed.ID).Return(errors.New("error")),
	)

	chs := chatUsecase{
		chatRepo: mock,
	}

	for i := 0; i < 2; i++ {
		err := chs.Message(cookie, message)
		require.Equal(t, models.ErrInternalServerError, err)
	}
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

func TestСhats(t *testing.T) {
	sid := "something-like-this"
	var chatModel models.ChatModel
	user := models.UserFeed{
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

	telephone := "944-739-32-28"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Times(1).Return(telephone)
	mock.EXPECT().SelectUserFeed(telephone).Times(1).Return(user, nil)
	mock.EXPECT().SelectChatsByID(user.ID).Times(1).Return(chats, nil)
	chatModel.Data = chats

	chs := chatUsecase{
		chatRepo: mock,
	}

	result, err := chs.Chats(sid)

	require.NoError(t, err)
	require.Equal(t, result, chatModel)
}

func TestСhatID(t *testing.T) {
	sid := "something-like-this"
	var chid = 1
	user := models.UserFeed{
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

	telephone := "944-739-32-28"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Times(1).Return(telephone)
	mock.EXPECT().SelectUserFeed(telephone).Times(1).Return(user, nil)
	mock.EXPECT().SelectChatByID(user.ID, chid).Times(1).Return(chat, nil)

	chs := chatUsecase{
		chatRepo: mock,
	}

	result, err := chs.ChatID(sid, chid)

	require.NoError(t, err)
	require.Equal(t, result, chat)
}

func TestСhatIDFail(t *testing.T) {
	sid := "something-like-this"
	var chid = 1
	user := models.UserFeed{}

	telephone := "944-739-32-28"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Times(1).Return(telephone)
	mock.EXPECT().SelectUserFeed(telephone).Times(1).Return(user, models.ErrInternalServerError)

	chs := chatUsecase{
		chatRepo: mock,
	}

	_, err := chs.ChatID(sid, chid)

	require.NotEqual(t, err, nil)
}

func TestСhatSelectIDFail(t *testing.T) {
	sid := "something-like-this"
	var chid = 1
	user := models.UserFeed{}

	msg1 := models.Msg{}

	msg2 := models.Msg{}

	chat := models.ChatData{
		ID:       1,
		Partner:  models.UserFeed{},
		Messages: []models.Msg{msg1, msg2},
	}

	telephone := "944-739-32-28"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Times(1).Return(telephone)
	mock.EXPECT().SelectUserFeed(telephone).Times(1).Return(user, nil)
	mock.EXPECT().SelectChatByID(user.ID, chid).Times(1).Return(chat, models.ErrInternalServerError)

	chs := chatUsecase{
		chatRepo: mock,
	}

	_, err := chs.ChatID(sid, chid)

	require.NotEqual(t, err, nil)
}

func TestGochat(t *testing.T) {
	sid := "something-like-this"
	user := models.UserFeed{
		ID:         1,
		Name:       "Andrey",
		DateBirth:  20,
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	telephone := "944-739-32-28"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Times(1).Return(telephone)
	mock.EXPECT().SelectUserFeed(telephone).Times(1).Return(user, nil)

	chs := chatUsecase{
		chatRepo: mock,
	}

	result, err := chs.Gochat(sid)

	require.NoError(t, err)
	require.Equal(t, result, user)
}

func TestGochatFail(t *testing.T) {
	sid := "something-like-this"
	user := models.UserFeed{}

	telephone := "944-739-32-28"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Times(1).Return(telephone)
	mock.EXPECT().SelectUserFeed(telephone).Times(1).Return(user, models.ErrInternalServerError)

	chs := chatUsecase{
		chatRepo: mock,
	}

	_, err := chs.Gochat(sid)

	require.NotEqual(t, err, nil)
}
