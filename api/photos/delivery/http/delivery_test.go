package http_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"park_2020/2020_2_tmp_name/domain"
	"park_2020/2020_2_tmp_name/domain/mock"
	"park_2020/2020_2_tmp_name/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	userHttp "park_2020/2020_2_tmp_name/api/photos/delivery/http"
)

func TestNewUserHandler(t *testing.T) {
	router := mux.NewRouter()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	userHttp.NewUserHandler(router, mock)
}

func TestUserHandler_AddPhotoHandlerSuccess(t *testing.T) {
	photo := models.Photo{
		Path:      "./static/avatars/4.jpg",
		Telephone: "909-277-47-21",
	}

	var byteData = []byte(`{
		"linkImages": "./static/avatars/4.jpg",
		"telephone":  "909-277-47-21"
	}`)

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/add_photo", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().AddPhoto(photo).Return(nil)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.AddPhotoHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)
}

func TestUserHandler_AddPhotoHandlerFail(t *testing.T) {
	photo := models.Photo{
		Path:      "./static/avatars/4.jpg",
		Telephone: "909-277-47-21",
	}

	var byteData = []byte(`{
		"linkImages": "./static/avatars/4.jpg",
		"telephone":  "909-277-47-21"
	}`)

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/add_photo", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().AddPhoto(photo).Return(domain.ErrInternalServerError)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.AddPhotoHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestUserHandler_AddPhotoHandlerFailDecode(t *testing.T) {
	var byteData = []byte(``)
	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/add_photo", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.AddPhotoHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestUserHandler_UploadAvatarHandlerFail(t *testing.T) {
	req, err := http.NewRequest("POST", "/upload", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.UploadAvatarHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}
