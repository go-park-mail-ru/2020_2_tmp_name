package usecase_test

import (
	"park_2020/2020_2_tmp_name/domain/mock"
	"park_2020/2020_2_tmp_name/models"
	userHttp "park_2020/2020_2_tmp_name/users/delivery/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestLoginHandler(t *testing.T) {
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
