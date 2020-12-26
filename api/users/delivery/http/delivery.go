package http

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	domain "park_2020/2020_2_tmp_name/api/users"
	authClient "park_2020/2020_2_tmp_name/microservices/authorization/delivery/grpc/client"
	faceClient "park_2020/2020_2_tmp_name/microservices/face_features/delivery/grpc/client"
	"park_2020/2020_2_tmp_name/middleware"
	"park_2020/2020_2_tmp_name/models"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type UserHandlerType struct {
	UUsecase   domain.UserUsecase
	AuthClient authClient.AuthClientInterface
	FaceClient *faceClient.FaceClient
}

func NewUserHandler(r *mux.Router, us domain.UserUsecase, ac authClient.AuthClientInterface, fc *faceClient.FaceClient) {
	handler := &UserHandlerType{
		UUsecase:   us,
		AuthClient: ac,
		FaceClient: fc,
	}

	r.HandleFunc("/health", handler.HealthHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/signup", middleware.CheckCSRF(handler.SignupHandler)).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/settings", middleware.CheckCSRF(handler.SettingsHandler)).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/me", middleware.SetCSRF(handler.MeHandler)).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/feed", middleware.SetCSRF(handler.FeedHandler)).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/user/{user_id}", middleware.SetCSRF(handler.UserIDHandler)).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/is_premium", middleware.SetCSRF(handler.IsPremiumHandler)).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/telephone", middleware.CheckCSRF(handler.TelephoneHandler)).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/upload", middleware.CheckCSRF(handler.UploadAvatarHandler)).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/get_premium", handler.GetPremiumHandler).Methods(http.MethodPost)
}

func (u *UserHandlerType) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func JSONError(message string) []byte {
	jsonError, err := json.Marshal(models.Error{Message: message})
	if err != nil {
		return []byte("")
	}
	return jsonError
}

func (u *UserHandlerType) UploadAvatarHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 10*1024*1024)
	err := r.ParseMultipartForm(10 * 1024 * 1024)
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

	photoPath := "/app/static/avatars"
	os.Chdir(photoPath)

	photoID, err := uuid.NewRandom()
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	f, err := os.OpenFile(photoID.String()+".jpg", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}
	defer f.Close()

	os.Chdir(str)

	body, err := json.Marshal("https://mi-ami.ru/static/avatars/" + photoID.String() + ".jpg")
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	photoModel := &models.Photo{
		Path: photoPath + "/" + photoID.String() + ".jpg",
		Mask: "",
	}

	haveFace, err := u.FaceClient.HaveFace(context.Background(), photoModel)
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

	err = u.UUsecase.ResizePhoto(photoPath + "/" + photoID.String() + ".jpg")
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandlerType) SignupHandler(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	err = u.UUsecase.Signup(user)
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(user)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandlerType) SettingsHandler(w http.ResponseWriter, r *http.Request) {
	userData := models.User{}
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	user, err := u.AuthClient.CheckSession(context.Background(), r.Cookies())
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	err = u.UUsecase.Settings(user.ID, userData)
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(userData)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandlerType) IsPremiumHandler(w http.ResponseWriter, r *http.Request) {
	user, err := u.AuthClient.CheckSession(context.Background(), r.Cookies())
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	var premium models.Premium
	premium.IsPremium = u.UUsecase.IsPremium(user.ID)

	body, err := json.Marshal(premium)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandlerType) MeHandler(w http.ResponseWriter, r *http.Request) {
	user, err := u.AuthClient.CheckSession(context.Background(), r.Cookies())
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(user)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandlerType) FeedHandler(w http.ResponseWriter, r *http.Request) {
	user, err := u.AuthClient.CheckSession(context.Background(), r.Cookies())
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	var feed models.Feed
	feed.Data, err = u.UUsecase.Feed(user)
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(feed)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandlerType) UserIDHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/v1/user/"))
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	user, err := u.UUsecase.UserID(userID)
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(user)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandlerType) TelephoneHandler(w http.ResponseWriter, r *http.Request) {
	phoneData := models.Phone{}
	err := json.NewDecoder(r.Body).Decode(&phoneData)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	hasUser := u.UUsecase.Telephone(phoneData.Telephone)

	result := models.HasTelephone{
		Telephone: hasUser,
	}

	body, err := json.Marshal(result)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	logrus.Println(body)

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandlerType) GetPremiumHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	label := r.PostFormValue("label")
	logrus.Println(label)
	userId, err := strconv.Atoi(label)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))

		return
	}

	err = u.UUsecase.GetPremium(userId)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))

		return
	}

	w.WriteHeader(http.StatusOK)
}
