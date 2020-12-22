package http

import (
	"context"
	"encoding/json"
	"net/http"
	domain "park_2020/2020_2_tmp_name/microservices/comments"
	grpcClient "park_2020/2020_2_tmp_name/microservices/comments/delivery/grpc/client"
	"park_2020/2020_2_tmp_name/models"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type CommentHandlerType struct {
	CUsecase domain.CommentUsecase
	ClientGRPC *grpcClient.CommentClient
}

func NewCommentHandler(r *mux.Router, cs domain.CommentUsecase, clientGRPC *grpcClient.CommentClient) {
	handler := &CommentHandlerType{
		CUsecase: cs,
		ClientGRPC: clientGRPC,
	}

	r.HandleFunc("/api/v1/comment", handler.CommentHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/comments/{user_id}", handler.CommentsByIdHandler).Methods(http.MethodGet)
}

func JSONError(message string) []byte {
	jsonError, err := json.Marshal(models.Error{Message: message})
	if err != nil {
		return []byte("")
	}
	return jsonError
}

func (c *CommentHandlerType) CommentHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	comment := models.Comment{}
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	if len(r.Cookies()) == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(JSONError("User not authorized"))
		return
	}

	user, err := c.CUsecase.User(ctx, r.Cookies()[0].Value)
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	err = c.ClientGRPC.Comment(ctx, user, comment)
	//err = c.CUsecase.Comment(ctx, user, comment)
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
	ctx := context.Background()

	userID, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/v1/comments/"))
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	//comments, err := c.CUsecase.CommentsByID(ctx, userID)
	comments, err := c.ClientGRPC.CommentsById(ctx, userID)
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
