package http_test

import (
	"bytes"
	"context"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"park_2020/2020_2_tmp_name/api/photos/mock"
	"park_2020/2020_2_tmp_name/models"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	photoHttp "park_2020/2020_2_tmp_name/api/photos/delivery/http"
	authClient "park_2020/2020_2_tmp_name/microservices/authorization/delivery/grpc/client"
	faceClient "park_2020/2020_2_tmp_name/microservices/face_features/delivery/grpc/client"

	mockClient "park_2020/2020_2_tmp_name/microservices/authorization/delivery/grpc/client/mock"
)

func TestNewUserHandler(t *testing.T) {
	router := mux.NewRouter()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockPhotoUsecase(ctrl)
	authClient := &authClient.AuthClient{}
	faceClient := &faceClient.FaceClient{}
	photoHttp.NewPhotoHandler(router, mock, authClient, faceClient)
}

func TestPhotoHandler_AddPhotoHandlerSuccess(t *testing.T) {
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
		"linkImages": "./static/avatars/4.jpg",
		"telephone":  "909-277-47-21"
	}`)

	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/add_photo", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockPhotoUsecase(ctrl)
	clientMock := mockClient.NewMockAuthClientInterface(ctrl)
	clientMock.EXPECT().CheckSession(context.Background(), req.Cookies()).Return(user, nil)

	photoHandler := photoHttp.PhotoHandlerType{
		PUsecase:   mock,
		AuthClient: clientMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(photoHandler.AddPhotoHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 400, status)
}
func TestPhotoHandler_AddPhotoHandlerFailUser(t *testing.T) {
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
	clientMock := mockClient.NewMockAuthClientInterface(ctrl)
	clientMock.EXPECT().CheckSession(context.Background(), req.Cookies()).Return(user, models.ErrUnauthorized)

	photoHandler := photoHttp.PhotoHandlerType{
		PUsecase:   mock,
		AuthClient: clientMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(photoHandler.AddPhotoHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 401, status)
}

func TestPhotoHandler_RemovePhotoHandlerSuccess(t *testing.T) {
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

	image := models.Image{
		LinkImage: "./static/avatars/4.jpg",
	}

	var byteData = []byte(`{
		"link_image": "./static/avatars/4.jpg"
	}`)

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/remove_photo", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockPhotoUsecase(ctrl)
	clientMock := mockClient.NewMockAuthClientInterface(ctrl)
	clientMock.EXPECT().CheckSession(context.Background(), req.Cookies()).Return(user, nil)
	mock.EXPECT().RemovePhoto(image.LinkImage, user.ID).Return(nil)

	photoHandler := photoHttp.PhotoHandlerType{
		PUsecase:   mock,
		AuthClient: clientMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(photoHandler.RemovePhotoHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)
}

func TestPhotoHandler_RemovePhotoHandlerFail(t *testing.T) {
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

	image := models.Image{
		LinkImage: "./static/avatars/4.jpg",
	}

	var byteData = []byte(`{
		"link_image": "./static/avatars/4.jpg"
	}`)

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/remove_photo", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockPhotoUsecase(ctrl)
	clientMock := mockClient.NewMockAuthClientInterface(ctrl)
	clientMock.EXPECT().CheckSession(context.Background(), req.Cookies()).Return(user, nil)
	mock.EXPECT().RemovePhoto(image.LinkImage, user.ID).Return(models.ErrInternalServerError)

	photoHandler := photoHttp.PhotoHandlerType{
		PUsecase:   mock,
		AuthClient: clientMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(photoHandler.RemovePhotoHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestPhotoHandler_RemovePhotoHandlerFailUser(t *testing.T) {
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
		"link_image": "./static/avatars/4.jpg"
	}`)

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/remove_photo", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockPhotoUsecase(ctrl)
	clientMock := mockClient.NewMockAuthClientInterface(ctrl)
	clientMock.EXPECT().CheckSession(context.Background(), req.Cookies()).Return(user, models.ErrUnauthorized)

	photoHandler := photoHttp.PhotoHandlerType{
		PUsecase:   mock,
		AuthClient: clientMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(photoHandler.RemovePhotoHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 401, status)
}

func TestPhotoHandler_RemovePhotoHandlerFailDecode(t *testing.T) {
	var byteData = []byte(``)
	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/remove_photo", body)
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
	handler := http.HandlerFunc(photoHandler.RemovePhotoHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 400, status)
}

func TestPhotoHandler_MaskHandlerFail(t *testing.T) {
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
		"linkImages": "./static/avatars/4.jpg",
		"telephone":  "909-277-47-21",
		"mask": "mask"
	}`)

	sid := "something-like-this"
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(10 * time.Hour),
	}

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/mask", body)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockPhotoUsecase(ctrl)
	clientMock := mockClient.NewMockAuthClientInterface(ctrl)
	clientMock.EXPECT().CheckSession(context.Background(), req.Cookies()).Return(user, models.ErrUnauthorized)

	photoHandler := photoHttp.PhotoHandlerType{
		PUsecase:   mock,
		AuthClient: clientMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(photoHandler.MaskHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 401, status)
}

func TestPhotoHandler_MaskHandlerFailDecode(t *testing.T) {
	var byteData = []byte(``)

	sid := "something-like-this"
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(10 * time.Hour),
	}

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/mask", body)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockPhotoUsecase(ctrl)
	clientMock := mockClient.NewMockAuthClientInterface(ctrl)

	photoHandler := photoHttp.PhotoHandlerType{
		PUsecase:   mock,
		AuthClient: clientMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(photoHandler.MaskHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 400, status)
}
