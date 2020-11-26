package http

import (
	"encoding/json"
	"net/http"
	domain "park_2020/2020_2_tmp_name/api/photos"
	"park_2020/2020_2_tmp_name/models"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type PhotoHandlerType struct {
	PUsecase domain.PhotoUsecase
}

func NewPhotoHandler(r *mux.Router, ps domain.PhotoUsecase) {
	handler := &PhotoHandlerType{
		PUsecase: ps,
	}

	path := "/static/avatars/"
	http.Handle("/", r)
	r.PathPrefix(path).Handler(http.StripPrefix(path, http.FileServer(http.Dir("."+path))))

	r.HandleFunc("/api/v1/add_photo", handler.AddPhotoHandler).Methods(http.MethodPost)
}

func JSONError(message string) []byte {
	jsonError, err := json.Marshal(models.Error{Message: message})
	if err != nil {
		return []byte("")
	}
	return jsonError
}

func (p *PhotoHandlerType) AddPhotoHandler(w http.ResponseWriter, r *http.Request) {
	photo := models.Photo{}
	err := json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

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
