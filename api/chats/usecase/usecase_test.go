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

func TestChatUsecase_Sessions(t *testing.T) {
	uid := 1

	sessions := []string{"aaa", "bbb"}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().SelectSessions(uid).Return(sessions, nil)

	chs := chatUsecase{
		chatRepo: mock,
	}

	result, err := chs.Sessions(uid)

	require.NoError(t, err)
	require.Equal(t, result, sessions)
}

func TestChatUsecase_SessionsF(t *testing.T) {
	uid := 1

	sessions := []string{"aaa", "bbb"}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().SelectSessions(uid).Return(sessions, models.ErrNotFound)

	chs := chatUsecase{
		chatRepo: mock,
	}

	_, err := chs.Sessions(uid)

	require.Equal(t, models.ErrNotFound, err)
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

func TestChatUsecase_PartnerSuccess(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	userfeed := models.UserFeed{
		ID:         0,
		Name:       "Masha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	chid := 1

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().SelectUserByChat(user.ID, chid).Return(userfeed, nil)

	chs := chatUsecase{
		chatRepo: mock,
	}

	result, err := chs.Partner(user, chid)

	require.NoError(t, err)
	require.Equal(t, result, userfeed)
}

func TestChatUsecase_PartnerFail(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	userfeed := models.UserFeed{
		ID:         0,
		Name:       "Masha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	chid := 1

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().SelectUserByChat(user.ID, chid).Return(userfeed, models.ErrNotFound)

	chs := chatUsecase{
		chatRepo: mock,
	}

	_, err := chs.Partner(user, chid)

	require.Error(t, models.ErrNotFound, err)
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

func TestChatUsecase_UserFeedSuccess(t *testing.T) {
	user := models.UserFeed{
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
	mock.EXPECT().SelectUserFeed(telephone).Return(user, nil)

	chs := chatUsecase{
		chatRepo: mock,
	}

	result, err := chs.UserFeed(sid)

	require.NoError(t, err)
	require.Equal(t, result, user)
}

func TestChatUsecase_UserFeedFail(t *testing.T) {
	user := models.UserFeed{
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
	mock.EXPECT().SelectUserFeed(telephone).Return(user, models.ErrNotFound)

	chs := chatUsecase{
		chatRepo: mock,
	}

	_, err := chs.UserFeed(sid)

	require.Equal(t, err, models.ErrNotFound)
}

func TestChatUsecase_LikeSuccess(t *testing.T) {
	like := models.Like{
		ID:   0,
		Uid1: 1,
		Uid2: 2,
	}

	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().CheckLike(user.ID, like.Uid2).Return(false)
	mock.EXPECT().CheckDislike(user.ID, like.Uid2).Return(true)
	mock.EXPECT().DeleteDislike(user.ID, like.Uid2).Return(nil)
	mock.EXPECT().InsertLike(user.ID, like.Uid2).Return(nil)

	chs := chatUsecase{
		chatRepo: mock,
	}

	err := chs.Like(user, like)

	require.NoError(t, err)
	require.Equal(t, nil, err)

}

func TestChatUsecase_LikeFailInsert(t *testing.T) {
	like := models.Like{
		ID:   0,
		Uid1: 1,
		Uid2: 2,
	}

	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().CheckLike(user.ID, like.Uid2).Return(false)
	mock.EXPECT().CheckDislike(user.ID, like.Uid2).Return(true)
	mock.EXPECT().DeleteDislike(user.ID, like.Uid2).Return(nil)
	mock.EXPECT().InsertLike(user.ID, like.Uid2).Return(models.ErrInternalServerError)

	chs := chatUsecase{
		chatRepo: mock,
	}

	err := chs.Like(user, like)
	require.Equal(t, models.ErrInternalServerError, err)
}

func TestChatUsecase_LikeFailDD(t *testing.T) {
	like := models.Like{
		ID:   0,
		Uid1: 1,
		Uid2: 2,
	}

	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().CheckLike(user.ID, like.Uid2).Return(false)
	mock.EXPECT().CheckDislike(user.ID, like.Uid2).Return(true)
	mock.EXPECT().DeleteDislike(user.ID, like.Uid2).Return(models.ErrInternalServerError)

	chs := chatUsecase{
		chatRepo: mock,
	}

	err := chs.Like(user, like)
	require.Equal(t, models.ErrInternalServerError, err)
}

func TestChatUsecase_LikeFailCL(t *testing.T) {
	like := models.Like{
		ID:   0,
		Uid1: 1,
		Uid2: 2,
	}

	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().CheckLike(user.ID, like.Uid2).Return(true)

	chs := chatUsecase{
		chatRepo: mock,
	}

	err := chs.Like(user, like)
	require.NoError(t, err)
}

func TestChatUsecase_MatchSuccess(t *testing.T) {
	like := models.Like{
		ID:   0,
		Uid1: 1,
		Uid2: 0,
	}

	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	chat := models.Chat{
		ID:   0,
		Uid1: 0,
		Uid2: 0,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	chid := 1

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().Match(user.ID, like.Uid2).Return(true)
	mock.EXPECT().CheckChat(chat).Return(false)
	mock.EXPECT().InsertChat(chat).Return(nil)
	mock.EXPECT().SelectChatID(user.ID, like.Uid2).Return(chid, nil)

	chs := chatUsecase{
		chatRepo: mock,
	}

	_, _, err := chs.MatchUser(user, like)

	require.NoError(t, err)
}

func TestChatUsecase_MatchFailSelect(t *testing.T) {
	like := models.Like{
		ID:   0,
		Uid1: 1,
		Uid2: 0,
	}

	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	chat := models.Chat{
		ID:   0,
		Uid1: 0,
		Uid2: 0,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	chid := 1

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().Match(user.ID, like.Uid2).Return(true)
	mock.EXPECT().CheckChat(chat).Return(false)
	mock.EXPECT().InsertChat(chat).Return(nil)
	mock.EXPECT().SelectChatID(user.ID, like.Uid2).Return(chid, models.ErrNotFound)

	chs := chatUsecase{
		chatRepo: mock,
	}

	_, _, err := chs.MatchUser(user, like)

	require.Equal(t, models.ErrNotFound, err)
}

func TestChatUsecase_MatchFailInsert(t *testing.T) {
	like := models.Like{
		ID:   0,
		Uid1: 1,
		Uid2: 0,
	}

	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	chat := models.Chat{
		ID:   0,
		Uid1: 0,
		Uid2: 0,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().Match(user.ID, like.Uid2).Return(true)
	mock.EXPECT().CheckChat(chat).Return(false)
	mock.EXPECT().InsertChat(chat).Return(models.ErrInternalServerError)

	chs := chatUsecase{
		chatRepo: mock,
	}

	_, _, err := chs.MatchUser(user, like)

	require.Equal(t, models.ErrInternalServerError, err)
}

func TestChatUsecase_MatchSuccessCheck(t *testing.T) {
	like := models.Like{
		ID:   0,
		Uid1: 1,
		Uid2: 0,
	}

	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	chat := models.Chat{
		ID:   0,
		Uid1: 0,
		Uid2: 0,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().Match(user.ID, like.Uid2).Return(true)
	mock.EXPECT().CheckChat(chat).Return(true)

	chs := chatUsecase{
		chatRepo: mock,
	}

	_, _, err := chs.MatchUser(user, like)

	require.Equal(t, nil, err)
}

func TestChatUsecase_MatchSuccessMatch(t *testing.T) {
	like := models.Like{
		ID:   0,
		Uid1: 1,
		Uid2: 0,
	}

	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().Match(user.ID, like.Uid2).Return(false)

	chs := chatUsecase{
		chatRepo: mock,
	}

	_, _, err := chs.MatchUser(user, like)

	require.Equal(t, nil, err)
}

func TestChatUsecase_DislikeSuccess(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	dislike := models.Dislike{
		ID:   0,
		Uid1: user.ID,
		Uid2: 2,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().CheckDislike(user.ID, dislike.Uid2).Return(false)
	mock.EXPECT().CheckLike(user.ID, dislike.Uid2).Return(true)
	mock.EXPECT().DeleteLike(user.ID, dislike.Uid2).Return(nil)
	mock.EXPECT().InsertDislike(user.ID, dislike.Uid2).Return(nil)

	chs := chatUsecase{
		chatRepo: mock,
	}

	err := chs.Dislike(user, dislike)

	require.NoError(t, err)
	require.Equal(t, nil, err)
}

func TestChatUsecase_DislikeFail(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	dislike := models.Dislike{
		ID:   0,
		Uid1: user.ID,
		Uid2: 2,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().CheckDislike(user.ID, dislike.Uid2).Return(false)
	mock.EXPECT().CheckLike(user.ID, dislike.Uid2).Return(true)
	mock.EXPECT().DeleteLike(user.ID, dislike.Uid2).Return(nil)
	mock.EXPECT().InsertDislike(user.ID, dislike.Uid2).Return(models.ErrInternalServerError)

	chs := chatUsecase{
		chatRepo: mock,
	}

	err := chs.Dislike(user, dislike)
	require.Equal(t, models.ErrInternalServerError, err)

}

func TestChatUsecase_DislikeFailDelete(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	dislike := models.Dislike{
		ID:   0,
		Uid1: user.ID,
		Uid2: 2,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().CheckDislike(user.ID, dislike.Uid2).Return(false)
	mock.EXPECT().CheckLike(user.ID, dislike.Uid2).Return(true)
	mock.EXPECT().DeleteLike(user.ID, dislike.Uid2).Return(models.ErrInternalServerError)

	chs := chatUsecase{
		chatRepo: mock,
	}

	err := chs.Dislike(user, dislike)
	require.Equal(t, models.ErrInternalServerError, err)
}

func TestChatUsecase_DislikeFailCD(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	dislike := models.Dislike{
		ID:   0,
		Uid1: user.ID,
		Uid2: 2,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().CheckDislike(user.ID, dislike.Uid2).Return(true)

	chs := chatUsecase{
		chatRepo: mock,
	}

	err := chs.Dislike(user, dislike)
	require.NoError(t, err)

}

func TestChatUsecase_SuperlikeSuccess(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	superlike := models.Superlike{
		ID:   0,
		Uid1: user.ID,
		Uid2: 2,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().CheckDislike(user.ID, superlike.Uid2).Return(true)
	mock.EXPECT().DeleteDislike(user.ID, superlike.Uid2).Return(nil)
	mock.EXPECT().InsertSuperlike(user.ID, superlike.Uid2).Return(nil)
	mock.EXPECT().CheckLike(user.ID, superlike.Uid2).Return(false)
	mock.EXPECT().InsertLike(user.ID, superlike.Uid2).Return(nil)

	chs := chatUsecase{
		chatRepo: mock,
	}

	err := chs.Superlike(user, superlike)

	require.NoError(t, err)
	require.Equal(t, nil, err)
}

func TestChatUsecase_SuperlikeSuccessCL(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	superlike := models.Superlike{
		ID:   0,
		Uid1: user.ID,
		Uid2: 2,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().CheckDislike(user.ID, superlike.Uid2).Return(true)
	mock.EXPECT().DeleteDislike(user.ID, superlike.Uid2).Return(nil)
	mock.EXPECT().InsertSuperlike(user.ID, superlike.Uid2).Return(nil)
	mock.EXPECT().CheckLike(user.ID, superlike.Uid2).Return(true)

	chs := chatUsecase{
		chatRepo: mock,
	}

	err := chs.Superlike(user, superlike)

	require.NoError(t, err)
	require.Equal(t, nil, err)
}

func TestChatUsecase_SuperlikeFailInsert(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	superlike := models.Superlike{
		ID:   0,
		Uid1: user.ID,
		Uid2: 2,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().CheckDislike(user.ID, superlike.Uid2).Return(true)
	mock.EXPECT().DeleteDislike(user.ID, superlike.Uid2).Return(nil)
	mock.EXPECT().InsertSuperlike(user.ID, superlike.Uid2).Return(nil)
	mock.EXPECT().CheckLike(user.ID, superlike.Uid2).Return(false)
	mock.EXPECT().InsertLike(user.ID, superlike.Uid2).Return(models.ErrInternalServerError)

	chs := chatUsecase{
		chatRepo: mock,
	}

	err := chs.Superlike(user, superlike)

	require.Equal(t, models.ErrInternalServerError, err)
}

func TestChatUsecase_SuperlikeFailInsertSuper(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	superlike := models.Superlike{
		ID:   0,
		Uid1: user.ID,
		Uid2: 2,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().CheckDislike(user.ID, superlike.Uid2).Return(true)
	mock.EXPECT().DeleteDislike(user.ID, superlike.Uid2).Return(nil)
	mock.EXPECT().InsertSuperlike(user.ID, superlike.Uid2).Return(models.ErrInternalServerError)

	chs := chatUsecase{
		chatRepo: mock,
	}

	err := chs.Superlike(user, superlike)

	require.Equal(t, models.ErrInternalServerError, err)
}

func TestChatUsecase_SuperlikeFailDD(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	superlike := models.Superlike{
		ID:   0,
		Uid1: user.ID,
		Uid2: 2,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatRepository(ctrl)
	mock.EXPECT().CheckDislike(user.ID, superlike.Uid2).Return(true)
	mock.EXPECT().DeleteDislike(user.ID, superlike.Uid2).Return(models.ErrInternalServerError)

	chs := chatUsecase{
		chatRepo: mock,
	}

	err := chs.Superlike(user, superlike)

	require.Equal(t, models.ErrInternalServerError, err)
}
