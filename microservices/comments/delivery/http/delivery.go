package http

import (
	"context"
	"encoding/json"
	"net/http"
	_authClientGRPC "park_2020/2020_2_tmp_name/microservices/authorization/delivery/grpc/client"
	domain "park_2020/2020_2_tmp_name/microservices/comments"
	_commentClientGRPC "park_2020/2020_2_tmp_name/microservices/comments/delivery/grpc/client"
	"park_2020/2020_2_tmp_name/middleware"
	"park_2020/2020_2_tmp_name/models"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type CommentHandlerType struct {
	CUsecase      domain.CommentUsecase
	AuthClient    _authClientGRPC.AuthClientInterface
	CommentClient _commentClientGRPC.CommentClientInterface
}

func NewCommentHandler(r *mux.Router, cs domain.CommentUsecase, cc _commentClientGRPC.CommentClientInterface, ac _authClientGRPC.AuthClientInterface) {
	handler := &CommentHandlerType{
		CUsecase:      cs,
		AuthClient:    ac,
		CommentClient: cc,
	}

	r.HandleFunc("/api/v1/comment", middleware.CheckCSRF(handler.CommentHandler)).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/comments/{user_id}", middleware.SetCSRF(handler.CommentsByIdHandler)).Methods(http.MethodGet)
}

func JSONError(message string) []byte {
	jsonError, err := json.Marshal(models.Error{Message: message})
	if err != nil {
		return []byte("")
	}
	return jsonError
}

func (c *CommentHandlerType) CommentHandler(w http.ResponseWriter, r *http.Request) {
	comment := models.Comment{}
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	user, err := c.AuthClient.CheckSession(context.Background(), r.Cookies())
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	err = c.CommentClient.Comment(context.Background(), user, comment)
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	comment.TimeDelivery = time.Now().Format("15:04")
	body, err := json.Marshal(comment)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (c *CommentHandlerType) CommentsByIdHandler(w http.ResponseWriter, r *http.Request) {
	_, err := c.AuthClient.CheckSession(context.Background(), r.Cookies())
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	userID, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/v1/comments/"))
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	comments, err := c.CommentClient.CommentsByID(context.Background(), userID)
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(comments)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
