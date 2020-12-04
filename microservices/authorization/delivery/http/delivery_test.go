package http_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"park_2020/2020_2_tmp_name/microservices/authorization/mock"
	"park_2020/2020_2_tmp_name/models"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	authClient "park_2020/2020_2_tmp_name/microservices/authorization/delivery/grpc/client"
	mockClient "park_2020/2020_2_tmp_name/microservices/authorization/delivery/grpc/client/mock"
	authHttp "park_2020/2020_2_tmp_name/microservices/authorization/delivery/http"
)

func TestNewAuthHandler(t *testing.T) {
	router := mux.NewRouter()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authClient := &authClient.AuthClient{}
	mock := mock.NewMockAuthUsecase(ctrl)
	authHttp.NewAuthHandler(router, mock, authClient)
}

func TestAuthHandler_LoginHandlerSuccess(t *testing.T) {
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

	sid := "some uuid"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockAuthUsecase(ctrl)
	clientMock := mockClient.NewMockAuthClientInterface(ctrl)
	clientMock.EXPECT().Login(context.Background(), &login).Return(sid, nil)

	authHandler := authHttp.AuthHandlerType{
		AUsecase:   mock,
		AuthClient: clientMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(authHandler.LoginHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.NoError(t, err)
	require.Equal(t, 200, status)

}

func TestAuthHandler_LoginHandlerFail(t *testing.T) {
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

	sid := "some uuid"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockAuthUsecase(ctrl)
	clientMock := mockClient.NewMockAuthClientInterface(ctrl)
	clientMock.EXPECT().Login(context.Background(), &login).Return(sid, models.ErrBadRequest)

	authHandler := authHttp.AuthHandlerType{
		AUsecase:   mock,
		AuthClient: clientMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(authHandler.LoginHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.NoError(t, err)
	require.Equal(t, 400, status)

}

func TestAuthHandler_LoginHandlerFailDecode(t *testing.T) {
	var byteData = []byte(``)

	body := bytes.NewReader(byteData)

	req, err := http.NewRequest("POST", "/login", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockAuthUsecase(ctrl)
	authHandler := authHttp.AuthHandlerType{
		AUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(authHandler.LoginHandler)

	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 400, status)
}

func TestAuthHandler_LogoutHandlerSuccess(t *testing.T) {
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

	mock := mock.NewMockAuthUsecase(ctrl)
	clientMock := mockClient.NewMockAuthClientInterface(ctrl)
	clientMock.EXPECT().Logout(context.Background(), cookie.Value).Return(nil)

	authHandler := authHttp.AuthHandlerType{
		AUsecase:   mock,
		AuthClient: clientMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(authHandler.LogoutHandler)

	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.NoError(t, err)
	require.Equal(t, 200, status)
}

func TestAuthHandler_LogoutHandlerFail(t *testing.T) {
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

	mock := mock.NewMockAuthUsecase(ctrl)
	clientMock := mockClient.NewMockAuthClientInterface(ctrl)
	clientMock.EXPECT().Logout(context.Background(), cookie.Value).Return(models.ErrBadRequest)

	authHandler := authHttp.AuthHandlerType{
		AUsecase:   mock,
		AuthClient: clientMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(authHandler.LogoutHandler)

	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.NoError(t, err)
	require.Equal(t, 400, status)
}

func TestAuthHandler_LogoutHandlerFailCookie(t *testing.T) {
	var byteData = []byte(`{}`)
	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/logout", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockAuthUsecase(ctrl)
	clientMock := mockClient.NewMockAuthClientInterface(ctrl)

	authHandler := authHttp.AuthHandlerType{
		AUsecase:   mock,
		AuthClient: clientMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(authHandler.LogoutHandler)

	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.NoError(t, err)
	require.Equal(t, 400, status)
}
