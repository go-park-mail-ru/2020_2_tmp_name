package http_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"park_2020/2020_2_tmp_name/api/users/mock"
	"park_2020/2020_2_tmp_name/models"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	userHttp "park_2020/2020_2_tmp_name/microservices/authorization/delivery/http"
)

func TestUserHandler_LoginHandlerSuccess(t *testing.T) {
	var byteData = []byte(`{
			"telephone" : "909-277-47-21",
			"password" : "qwerty"
		}`)

	login := models.LoginData{
		Telephone: "909-277-47-21",
		Password:  "qwerty",
	}
	body := bytes.NewReader(byteData)

	req, err := http.NewRequest("POST", "/login", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Login(login).Return("some uuid", nil)

	userHandler := userHttp.UserHandlerType{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.LoginHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.NoError(t, err)
	require.Equal(t, 200, status)

}

func TestUserHandler_LoginHandlerFail(t *testing.T) {
	var byteData = []byte(`{
			"telephone" : "909-277-47-21",
			"password" : "qwerty"
		}`)

	body := bytes.NewReader(byteData)

	login := models.LoginData{
		Telephone: "909-277-47-21",
		Password:  "qwerty",
	}

	req, err := http.NewRequest("POST", "/login", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	gomock.InOrder(
		mock.EXPECT().Login(login).Return("some uuid", models.ErrInternalServerError),
	)

	userHandler := userHttp.UserHandlerType{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.LoginHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code
	require.Equal(t, 500, status)
}

func TestUserHandler_LoginHandlerFailDecode(t *testing.T) {
	var byteData = []byte(``)

	body := bytes.NewReader(byteData)

	req, err := http.NewRequest("POST", "/login", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	userHandler := userHttp.UserHandlerType{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.LoginHandler)

	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 400, status)
}

func TestUserHandler_LogoutHandlerSuccess(t *testing.T) {
	var byteData = []byte(`{}`)
	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/logout", body)
	if err != nil {
		t.Fatal(err)
	}
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   "some uuid",
		Expires: time.Now().Add(10 * time.Hour),
	}
	cookie.HttpOnly = false
	cookie.Secure = false

	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Logout(cookie.Value).Return(nil)

	userHandler := userHttp.UserHandlerType{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.LogoutHandler)

	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.NoError(t, err)
	require.Equal(t, 200, status)
}

func TestUserHandler_LogoutHandlerFail(t *testing.T) {
	var byteData = []byte(`{}`)
	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/logout", body)
	if err != nil {
		t.Fatal(err)
	}
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   "some uuid",
		Expires: time.Now().Add(10 * time.Hour),
	}
	cookie.HttpOnly = false
	cookie.Secure = false

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Logout(cookie.Value).Return(errors.New("error"))

	userHandler := userHttp.UserHandlerType{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.LogoutHandler)
	for i := 0; i < 2; i++ {
		handler.ServeHTTP(rr, req)
		status := rr.Code
		req.AddCookie(cookie)

		require.NoError(t, err)
		require.Equal(t, 400, status)
	}
}
