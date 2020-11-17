package http

import (
	"encoding/json"
	"log"
	"net/http"
	domain "park_2020/2020_2_tmp_name/api/likes"
	"park_2020/2020_2_tmp_name/models"

	"github.com/gorilla/mux"
)

type LikeHandlerType struct {
	LUsecase domain.LikeUsecase
}

func NewLikeHandler(r *mux.Router, us domain.LikeUsecase) {
	handler := &LikeHandlerType{
		LUsecase: us,
	}

	r.HandleFunc("/api/v1/like", handler.LikeHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/dislike", handler.DislikeHandler).Methods(http.MethodPost)
}

func JSONError(message string) []byte {
	jsonError, err := json.Marshal(models.Error{Message: message})
	if err != nil {
		return []byte("")
	}
	return jsonError
}

func (l *LikeHandlerType) LikeHandler(w http.ResponseWriter, r *http.Request) {
	like := models.Like{}
	err := json.NewDecoder(r.Body).Decode(&like)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	cookie := r.Cookies()[0]

	err = l.LUsecase.Like(cookie.Value, like)
	if err != nil {
		log.Println(err)
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(like)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (l *LikeHandlerType) DislikeHandler(w http.ResponseWriter, r *http.Request) {
	dislike := models.Dislike{}
	err := json.NewDecoder(r.Body).Decode(&dislike)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	cookie := r.Cookies()[0]

	err = l.LUsecase.Dislike(cookie.Value, dislike)
	if err != nil {
		log.Println(err)
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(dislike)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
