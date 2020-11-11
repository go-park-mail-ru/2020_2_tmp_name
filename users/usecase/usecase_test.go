package usecase_test

import (
	"park_2020/2020_2_tmp_name/domain/mock"
	"park_2020/2020_2_tmp_name/models"
	userHttp "park_2020/2020_2_tmp_name/users/delivery/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	telephone := "944-739-32-28"
	loginData := &models.LoginData{
		Telephone: telephone,
		Password:  "password",
	}

	retModel := "string"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Login(*loginData).Times(1).Return(retModel, nil)

	uh := userHttp.UserHandler{
		UUsecase: mock,
	}

	models, err := uh.UUsecase.Login(*loginData)

	require.NoError(t, err)
	require.NotEqual(t, models, "")
}

func TestLogout(t *testing.T) {
	session := "hhxjxjjcxjj-hjxjxjx-xjjxjxjxj"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Logout(session).Times(1).Return(nil)

	uh := userHttp.UserHandler{
		UUsecase: mock,
	}

	err := uh.UUsecase.Logout(session)
	require.NoError(t, err)
}

func TestSettings(t *testing.T) {
	cookie := "hdisjsjs-sjsosksisi-jxjsjs"
	user := &models.User{
		Telephone: "958-475-21-69",
		Password:  "password",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Settings(cookie, *user).Times(1).Return(nil)

	uh := userHttp.UserHandler{
		UUsecase: mock,
	}

	err := uh.UUsecase.Settings(cookie, *user)

	require.NoError(t, err)
}
