package usecase

import (
	"context"
	domain "park_2020/2020_2_tmp_name/microservices/authorization"
	"park_2020/2020_2_tmp_name/microservices/authorization/mock"
	"park_2020/2020_2_tmp_name/models"

	"github.com/golang/mock/gomock"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewUserUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var u domain.AuthRepository
	uu := NewAuthUsecase(u)
	require.Empty(t, uu)
}

func TestUserUsecase_LoginFail(t *testing.T) {
	login := models.LoginData{
		Telephone: "944-739-32-28",
		Password:  "password",
	}

	user := models.User{
		ID:         1,
		Name:       "Andrey",
		Telephone:  "944-739-32-28",
		Password:   "password",
		DateBirth:  20,
		Sex:        "male",
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := mock.NewMockAuthRepository(ctrl)
	mock.EXPECT().CheckUser(login.Telephone).Times(1).Return(true)
	mock.EXPECT().SelectUser(login.Telephone).Times(1).Return(user, nil)

	us := authUsecase{
		userRepo: mock,
	}

	_, err := us.Login(context.Background(), login)

	require.NotEqual(t, err, nil)
}

func TestUserUsecase_LoginFailSelect(t *testing.T) {
	login := models.LoginData{
		Telephone: "944-739-32-28",
		Password:  "password",
	}

	user := models.User{
		ID:         1,
		Name:       "Andrey",
		Telephone:  "944-739-32-28",
		Password:   "password",
		DateBirth:  20,
		Sex:        "male",
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := mock.NewMockAuthRepository(ctrl)
	mock.EXPECT().CheckUser(login.Telephone).Times(1).Return(true)
	mock.EXPECT().SelectUser(login.Telephone).Times(1).Return(user, models.ErrNotFound)

	us := authUsecase{
		userRepo: mock,
	}

	_, err := us.Login(context.Background(), login)

	require.NotEqual(t, err, nil)
}

func TestUserUsecase_LoginFailCheck(t *testing.T) {
	login := models.LoginData{
		Telephone: "944-739-32-28",
		Password:  "password",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := mock.NewMockAuthRepository(ctrl)
	mock.EXPECT().CheckUser(login.Telephone).Times(1).Return(false)

	us := authUsecase{
		userRepo: mock,
	}

	_, err := us.Login(context.Background(), login)

	require.NotEqual(t, err, nil)
}

func TestUserUsecase_Logout(t *testing.T) {
	sid := "something-like-this"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockAuthRepository(ctrl)
	mock.EXPECT().DeleteSession(sid).Times(1).Return(nil)

	us := authUsecase{
		userRepo: mock,
	}

	err := us.Logout(context.Background(), sid)
	require.NoError(t, err)
}

func TestUserUsecase_CheckSessionSuccess(t *testing.T) {
	cookie := "cookie"
	telephone := "telephone"

	user := models.User{
		ID:         1,
		Name:       "Andrey",
		Telephone:  "944-739-32-28",
		Password:   "password",
		DateBirth:  20,
		Sex:        "male",
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := mock.NewMockAuthRepository(ctrl)
	mock.EXPECT().CheckUserBySession(cookie).Times(1).Return(telephone)
	mock.EXPECT().SelectUser(telephone).Times(1).Return(user, nil)

	us := authUsecase{
		userRepo: mock,
	}

	_, err := us.CheckSession(context.Background(), cookie)

	require.Equal(t, err, nil)
}

func TestUserUsecase_CheckSessionFail(t *testing.T) {
	cookie := "cookie"
	telephone := "telephone"

	user := models.User{
		ID:         1,
		Name:       "Andrey",
		Telephone:  "944-739-32-28",
		Password:   "password",
		DateBirth:  20,
		Sex:        "male",
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := mock.NewMockAuthRepository(ctrl)
	mock.EXPECT().CheckUserBySession(cookie).Times(1).Return(telephone)
	mock.EXPECT().SelectUser(telephone).Times(1).Return(user, models.ErrNotFound)

	us := authUsecase{
		userRepo: mock,
	}

	_, err := us.CheckSession(context.Background(), cookie)

	require.NotEqual(t, err, nil)
}
