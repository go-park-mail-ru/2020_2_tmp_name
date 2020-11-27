package http_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"park_2020/2020_2_tmp_name/api/photos/mock"
	"park_2020/2020_2_tmp_name/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	photoHttp "park_2020/2020_2_tmp_name/api/photos/delivery/http"
)

func TestNewUserHandler(t *testing.T) {
	router := mux.NewRouter()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockPhotoUsecase(ctrl)
	photoHttp.NewPhotoHandler(router, mock)
}

func TestPhotoHandler_AddPhotoHandlerSuccess(t *testing.T) {
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

	mock := mock.NewMockPhotoUsecase(ctrl)
	mock.EXPECT().AddPhoto(photo).Return(nil)

	photoHandler := photoHttp.PhotoHandlerType{
		PUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(photoHandler.AddPhotoHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)
}

func TestPhotoHandler_AddPhotoHandlerFail(t *testing.T) {
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

	mock := mock.NewMockPhotoUsecase(ctrl)
	mock.EXPECT().AddPhoto(photo).Return(models.ErrInternalServerError)

	photoHandler := photoHttp.PhotoHandlerType{
		PUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(photoHandler.AddPhotoHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestPhotoHandler_AddPhotoHandlerFailDecode(t *testing.T) {
	var byteData = []byte(``)
	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/add_photo", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockPhotoUsecase(ctrl)

	photoHandler := photoHttp.PhotoHandlerType{
		PUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(photoHandler.AddPhotoHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 400, status)
}
