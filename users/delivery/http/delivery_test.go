package http_test

import (
	"bytes"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"park_2020/2020_2_tmp_name/domain/mock"
	"park_2020/2020_2_tmp_name/models"
	"testing"
	"time"

	userHttp "park_2020/2020_2_tmp_name/users/delivery/http"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	userHandler := &userHttp.UserHandler{}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.HealthHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

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

	userHandler := userHttp.UserHandler{
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
	var byteData = []byte(``)
	var byteData2 = []byte(`{
			"telephone" : "909-277-47-21",
			"password" : "qwerty"
		}`)

	body := bytes.NewReader(byteData)
	body_sec := bytes.NewReader(byteData2)


	login := models.LoginData{
		Telephone: "909-277-47-21",
		Password:  "qwerty",
	}


	req, err := http.NewRequest("POST", "/login", body)
	if err != nil {
		t.Fatal(err)
	}
	req2, err := http.NewRequest("POST", "/login", body_sec)
	if err != nil {
		t.Fatal(err)
	}
	requests := make([]*http.Request, 0, 1)
	requests = append(requests, req)
	requests = append(requests, req2)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()


	mock := mock.NewMockUserUsecase(ctrl)
	gomock.InOrder(
		mock.EXPECT().Login(login).Return("some uuid", errors.New("error uuid")),
		)


	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}


	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.LoginHandler)
	for i := 0; i < 2; i++ {
		handler.ServeHTTP(rr, requests[i])
		status := rr.Code

		require.NoError(t, err)
		require.Equal(t, 500, status)
	}

}
func TestNewUserHandler(t *testing.T) {
	router := mux.NewRouter()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()


	mock := mock.NewMockUserUsecase(ctrl)
	userHttp.NewUserHandler(router, mock)
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

	userHandler := userHttp.UserHandler{
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

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.LogoutHandler)
	for i := 0; i < 2; i++ {
		handler.ServeHTTP(rr, req)
		status := rr.Code
		req.AddCookie(cookie)

		require.NoError(t, err)
		require.Equal(t, 500, status)
	}
}

func TestUserHandler_SignupHandlerSuccess(t *testing.T) {
	user := models.User{
		Name:       "Misha",
		Telephone:  "909-277-47-21",
		Password:   "1234",
		Sex:        "male",
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}
	var byteData = []byte(`{
		"id":         "0",
		"name":       "Misha",
		"telephone":  "909-277-47-21",
		"password":   "1234",
		"sex":        "male",
		"linkImages": null,
		"job":        "Fullstack",
		"education":  "BMSTU",
		"aboutMe":    ""
	}`)
	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/signup", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()


	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Signup(user).Return(nil)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.SignupHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)

}

func TestUserHandler_SignupHandlerFail(t *testing.T) {
	user := models.User {
		Name:       "Misha",
		Telephone:  "909-277-47-21",
		Password:   "1234",
		Sex:        "male",
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}
	var byteData = []byte(`{
		"id":         "0",
		"name":       "Misha",
		"telephone":  "909-277-47-21",
		"password":   "1234",
		"sex":        "male",
		"linkImages": null,
		"job":        "Fullstack",
		"education":  "BMSTU",
		"aboutMe":    ""
	}`)
	var ByteData = []byte(``)

	body2 := bytes.NewReader(ByteData)
	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/signup", body2)
	if err != nil {
		t.Fatal(err)
	}

	req2, err := http.NewRequest("POST", "/signup", body)
	if err != nil {
		t.Fatal(err)
	}

	requests := make([]*http.Request, 0, 1)
	requests = append(requests, req)
	requests = append(requests, req2)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Signup(user).Return(errors.New("error"))

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.SignupHandler)
	for i := 0; i < 2; i++ {
		handler.ServeHTTP(rr, requests[i])
		status := rr.Code

		require.NoError(t, err)
		require.Equal(t, 500, status)
	}

}