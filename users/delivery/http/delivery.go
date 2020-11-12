package http

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"park_2020/2020_2_tmp_name/domain"
	"park_2020/2020_2_tmp_name/models"
	"strconv"
	"strings"

	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	UUsecase domain.UserUsecase
	Hub      *models.Hub
}

func NewUserHandler(r *mux.Router, us domain.UserUsecase) {
	handler := &UserHandler{
		UUsecase: us,
	}

	path := "/static/avatars/"
	http.Handle("/", r)
	r.PathPrefix(path).Handler(http.StripPrefix(path, http.FileServer(http.Dir("."+path))))

	http.HandleFunc("/health", handler.HealthHandler)

	r.HandleFunc("/health", handler.HealthHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/login", handler.LoginHandler).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/api/v1/logout", handler.LogoutHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/signup", handler.SignupHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/settings", handler.SettingsHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/upload", handler.UploadAvatarHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/add_photo", handler.AddPhotoHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/me", handler.MeHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/feed", handler.FeedHandler).Methods(http.MethodGet)

	r.HandleFunc("/api/v1/like", handler.LikeHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/dislike", handler.DislikeHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/comment", handler.CommentHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/comments/{user_id}", handler.CommentsByIdHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/chat", handler.ChatHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/message", handler.MessageHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/chats", handler.ChatsHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/chats/{chat_id}", handler.ChatIDHandler).Methods(http.MethodGet)

	r.HandleFunc("/api/v1/gochat", handler.GochatHandler).Methods(http.MethodGet, http.MethodPost)
}

func (u *UserHandler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func JSONError(message string) []byte {
	jsonError, err := json.Marshal(models.Error{Message: message})
	if err != nil {
		return []byte("")
	}
	return jsonError
}

func (u *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	loginData := models.LoginData{}
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	sidString, err := u.UUsecase.Login(loginData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sidString,
		Expires: time.Now().Add(10 * time.Hour),
	}
	cookie.HttpOnly = false
	cookie.Secure = false

	body, err := json.Marshal(loginData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	http.SetCookie(w, cookie)
	w.Write(body)
}

func (u *UserHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	err = u.UUsecase.Logout(session.Value)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal("logout success")
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	err = u.UUsecase.Signup(user)
	if err != nil {
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandler) SettingsHandler(w http.ResponseWriter, r *http.Request) {
	userData := models.User{}
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	cookie := r.Cookies()[0]
	err = u.UUsecase.Settings(cookie.Value, userData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (u *UserHandler) MeHandler(w http.ResponseWriter, r *http.Request) {
	cookie := r.Cookies()[0]
	user, err := u.UUsecase.Me(cookie.Value)
	if err != nil {
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandler) FeedHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	cookie := r.Cookies()[0]
	var feed models.Feed
	feed.Data, err = u.UUsecase.Feed(cookie.Value)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(feed)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandler) AddPhotoHandler(w http.ResponseWriter, r *http.Request) {
	photo := models.Photo{}
	err := json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	err = u.UUsecase.AddPhoto(photo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(photo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandler) UploadAvatarHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 1024)
	file, _, err := r.FormFile("photo")
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}
	defer file.Close()
	r.FormValue("photo")

	str, err := os.Getwd()
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	os.Chdir("/home/ubuntu/go/src/2020_2_tmp_name/static/avatars")

	photoID, err := u.UUsecase.UploadAvatar()
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	f, err := os.OpenFile(photoID.String(), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}
	defer f.Close()

	os.Chdir(str)

	body, err := json.Marshal(photoID.String())
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	io.Copy(f, file)
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandler) LikeHandler(w http.ResponseWriter, r *http.Request) {
	like := models.Like{}
	err := json.NewDecoder(r.Body).Decode(&like)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	cookie := r.Cookies()[0]

	err = u.UUsecase.Like(cookie.Value, like)

	body, err := json.Marshal(like)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandler) DislikeHandler(w http.ResponseWriter, r *http.Request) {
	dislike := models.Dislike{}
	err := json.NewDecoder(r.Body).Decode(&dislike)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	cookie := r.Cookies()[0]

	err = u.UUsecase.Dislike(cookie.Value, dislike)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(dislike)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandler) CommentHandler(w http.ResponseWriter, r *http.Request) {
	comment := models.Comment{}
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	cookie := r.Cookies()[0]
	err = u.UUsecase.Comment(cookie.Value, comment)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	comment.TimeDelivery = time.Now().Format("15:04")
	body, err := json.Marshal(comment)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandler) CommentsByIdHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/v1/comments/"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	comments, err := u.UUsecase.CommentsByID(userID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(comments)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandler) ChatHandler(w http.ResponseWriter, r *http.Request) {
	chat := models.Chat{}
	err := json.NewDecoder(r.Body).Decode(&chat)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	err = u.UUsecase.Chat(chat)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(chat)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandler) MessageHandler(w http.ResponseWriter, r *http.Request) {
	message := models.Message{}
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	cookie := r.Cookies()[0]
	err = u.UUsecase.Message(cookie.Value, message)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandler) ChatsHandler(w http.ResponseWriter, r *http.Request) {
	cookie := r.Cookies()[0]

	chatModel, err := u.UUsecase.Chats(cookie.Value)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(chatModel)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandler) ChatIDHandler(w http.ResponseWriter, r *http.Request) {
	chid, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/v1/chats/"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	cookie := r.Cookies()[0]
	chat, err := u.UUsecase.ChatID(cookie.Value, chid)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(chat)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandler) GochatHandler(w http.ResponseWriter, r *http.Request) {
	cookie := r.Cookies()[0]

	user, err := u.UUsecase.Gochat(cookie.Value)
	if err != nil {
		log.Println(err)
		w.WriteHeader(getStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	u.UUsecase.ServeWs(u.Hub, w, r, user.ID)
	w.WriteHeader(http.StatusOK)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	case domain.ErrUnauthorized:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
