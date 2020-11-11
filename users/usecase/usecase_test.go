package usecase

import (
	"park_2020/2020_2_tmp_name/domain/mock"
	"park_2020/2020_2_tmp_name/models"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestLoginFail(t *testing.T) {
	login := models.LoginData{
		Telephone: "944-739-32-28",
		Password:  "password",
	}

	user := models.User{
		ID:         1,
		Name:       "Andrey",
		Telephone:  "944-739-32-28",
		Password:   "password",
		DateBirth:  time.Time{},
		Sex:        "male",
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUser(login.Telephone).Times(1).Return(true)
	mock.EXPECT().SelectUser(login.Telephone).Times(1).Return(user, nil)

	us := userUsecase{
		userRepo: mock,
	}

	_, err := us.Login(login)

	require.NotEqual(t, err, nil)
}

func TestLogout(t *testing.T) {
	sid := "something-like-this"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().DeleteSession(sid).Times(1).Return(nil)

	us := userUsecase{
		userRepo: mock,
	}

	err := us.Logout(sid)
	require.NoError(t, err)
}

func TestAddPhotoSuccess(t *testing.T) {
	photo := models.Photo{
		Path:      "path",
		Telephone: "944-739-32-28",
	}

	user := models.UserFeed{
		ID:         1,
		Name:       "Andrey",
		DateBirth:  20,
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().SelectUserFeed(photo.Telephone).Times(1).Return(user, nil)
	mock.EXPECT().InsertPhoto(photo.Path, user.ID).Times(1).Return(nil)

	us := userUsecase{
		userRepo: mock,
	}

	err := us.AddPhoto(photo)

	require.NoError(t, err)
}

func TestMe(t *testing.T) {
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

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Times(1).Return(telephone)
	mock.EXPECT().SelectUserFeed(telephone).Times(1).Return(user, nil)

	us := userUsecase{
		userRepo: mock,
	}

	me, err := us.Me(sid)

	require.NoError(t, err)
	require.Equal(t, me, user)
}

func TestFeed(t *testing.T) {
	sid := "something-like-this"

	user := models.User{
		ID:         1,
		Name:       "Andrey",
		Telephone:  "944-739-32-28",
		Password:   "password",
		DateBirth:  time.Time{},
		Sex:        "male",
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	var users []models.UserFeed
	user1 := models.UserFeed{
		ID:         3,
		Name:       "Masha",
		DateBirth:  20,
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	user2 := models.UserFeed{
		ID:         4,
		Name:       "Dasha",
		DateBirth:  20,
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	users = append(users, user1, user2)

	telephone := "944-739-32-28"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Times(1).Return(telephone)
	mock.EXPECT().SelectUser(telephone).Times(1).Return(user, nil)
	mock.EXPECT().SelectUsers(user).Times(1).Return(users, nil)

	us := userUsecase{
		userRepo: mock,
	}

	feed, err := us.Feed(sid)

	require.NoError(t, err)
	require.Equal(t, feed, users)
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

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Times(1).Return(telephone)
	mock.EXPECT().SelectUserFeed(telephone).Times(1).Return(user, nil)
	mock.EXPECT().SelectChatsByID(user.ID).Times(1).Return(chats, nil)
	chatModel.Data = chats

	us := userUsecase{
		userRepo: mock,
	}

	result, err := us.Chats(sid)

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

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Times(1).Return(telephone)
	mock.EXPECT().SelectUserFeed(telephone).Times(1).Return(user, nil)
	mock.EXPECT().SelectChatByID(user.ID, chid).Times(1).Return(chat, nil)

	us := userUsecase{
		userRepo: mock,
	}

	result, err := us.ChatID(sid, chid)

	require.NoError(t, err)
	require.Equal(t, result, chat)
}
