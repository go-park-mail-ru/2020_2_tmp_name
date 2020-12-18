package http

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
	domain "park_2020/2020_2_tmp_name/api/photos"
	_authClientGRPC "park_2020/2020_2_tmp_name/microservices/authorization/delivery/grpc/client"
	faceClient "park_2020/2020_2_tmp_name/microservices/face_features/delivery/grpc/client"
	"park_2020/2020_2_tmp_name/middleware"
	"park_2020/2020_2_tmp_name/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type PhotoHandlerType struct {
	PUsecase   domain.PhotoUsecase
	AuthClient _authClientGRPC.AuthClientInterface
	FaceClient *faceClient.FaceClient
}

func NewPhotoHandler(r *mux.Router, ps domain.PhotoUsecase, ac _authClientGRPC.AuthClientInterface, fc *faceClient.FaceClient) {
	handler := &PhotoHandlerType{
		PUsecase:   ps,
		AuthClient: ac,
		FaceClient: fc,
	}

	path := "/static/avatars/"
	http.Handle("/", r)
	r.PathPrefix(path).Handler(http.StripPrefix(path, http.FileServer(http.Dir("."+path))))

	r.HandleFunc("/api/v1/add_photo", middleware.CheckCSRF(handler.AddPhotoHandler)).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/remove_photo", middleware.CheckCSRF(handler.RemovePhotoHandler)).Methods(http.MethodPost)
}

func JSONError(message string) []byte {
	jsonError, err := json.Marshal(models.Error{Message: message})
	if err != nil {
		return []byte("")
	}
	return jsonError
}

func (p *PhotoHandlerType) AddPhotoHandler(w http.ResponseWriter, r *http.Request) {
	user, err := p.AuthClient.CheckSession(context.Background(), r.Cookies())
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 10*1024*1024)
	err = r.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	file, _, err := r.FormFile("photo")
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}
	defer file.Close()
	r.FormValue("photo")

	str, err := os.Getwd()
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	photoPath := "/home/ubuntu/go/src/park_2020/2020_2_tmp_name/static/avatars"
	os.Chdir(photoPath)

	photoID, err := uuid.NewRandom()
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	f, err := os.OpenFile(photoID.String(), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}
	defer f.Close()

	os.Chdir(str)

	_, err = io.Copy(f, file)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	photoModel := &models.Photo{
		Path: photoPath + "/" + photoID.String(),
		Mask: "",
	}

	haveFace, err := p.FaceClient.HaveFace(context.Background(), photoModel)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	if !haveFace {
		err = errors.New("no face on photo")
		logrus.Error(err.Error())
		w.WriteHeader(http.StatusForbidden)
		w.Write(JSONError(err.Error()))
		return
	}

	var photo models.Photo
	photo.Telephone = user.Telephone
	photo.Path = "https://mi-ami.ru/static/avatars/" + photoID.String()

	err = p.PUsecase.AddPhoto(photo)
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(photo)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (p *PhotoHandlerType) RemovePhotoHandler(w http.ResponseWriter, r *http.Request) {
	linkImage := models.Image{}
	err := json.NewDecoder(r.Body).Decode(&linkImage)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	user, err := p.AuthClient.CheckSession(context.Background(), r.Cookies())
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	err = p.PUsecase.RemovePhoto(linkImage.LinkImage, user.ID)
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(linkImage)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
