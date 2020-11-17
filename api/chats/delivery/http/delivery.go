package http

import (
	"encoding/json"
	"log"
	"net/http"
	domain "park_2020/2020_2_tmp_name/api/chats"
	"park_2020/2020_2_tmp_name/models"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type ChatHandlerType struct {
	ChUsecase domain.ChatUsecase
	Hub       *models.Hub
}

func NewChatHandler(r *mux.Router, chs domain.ChatUsecase) {
	handler := &ChatHandlerType{
		ChUsecase: chs,
	}

	r.HandleFunc("/api/v1/chat", handler.ChatHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/message", handler.MessageHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/chats", handler.ChatsHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/chats/{chat_id}", handler.ChatIDHandler).Methods(http.MethodGet)

	r.HandleFunc("/api/v1/gochat", handler.GochatHandler).Methods(http.MethodGet, http.MethodPost)
}

func JSONError(message string) []byte {
	jsonError, err := json.Marshal(models.Error{Message: message})
	if err != nil {
		return []byte("")
	}
	return jsonError
}

func (ch *ChatHandlerType) ChatHandler(w http.ResponseWriter, r *http.Request) {
	chat := models.Chat{}
	err := json.NewDecoder(r.Body).Decode(&chat)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	err = ch.ChUsecase.Chat(chat)
	if err != nil {
		log.Println(err)
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(chat)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (ch *ChatHandlerType) MessageHandler(w http.ResponseWriter, r *http.Request) {
	message := models.Message{}
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	cookie := r.Cookies()[0]
	err = ch.ChUsecase.Message(cookie.Value, message)
	if err != nil {
		log.Println(err)
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (ch *ChatHandlerType) ChatsHandler(w http.ResponseWriter, r *http.Request) {
	cookie := r.Cookies()[0]

	chatModel, err := ch.ChUsecase.Chats(cookie.Value)
	if err != nil {
		log.Println(err)
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(chatModel)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (ch *ChatHandlerType) ChatIDHandler(w http.ResponseWriter, r *http.Request) {
	chid, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/v1/chats/"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	cookie := r.Cookies()[0]
	chat, err := ch.ChUsecase.ChatID(cookie.Value, chid)
	if err != nil {
		log.Println(err)
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(chat)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (ch *ChatHandlerType) GochatHandler(w http.ResponseWriter, r *http.Request) {
	cookie := r.Cookies()[0]

	user, err := ch.ChUsecase.Gochat(cookie.Value)
	if err != nil {
		log.Println(err)
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	ch.ChUsecase.ServeWs(ch.Hub, w, r, user.ID)
	w.WriteHeader(http.StatusOK)
}
