package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"google.golang.org/grpc"

	"park_2020/2020_2_tmp_name/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"

	metrics "park_2020/2020_2_tmp_name/prometheus"

	_chatDelivery "park_2020/2020_2_tmp_name/api/chats/delivery/http"
	_chatRepo "park_2020/2020_2_tmp_name/api/chats/repository/postgres"
	_chatUcase "park_2020/2020_2_tmp_name/api/chats/usecase"

	_commentClient "park_2020/2020_2_tmp_name/microservices/comments/delivery/grpc/client"
	_commentDelivery "park_2020/2020_2_tmp_name/microservices/comments/delivery/http"
	_commentRepo "park_2020/2020_2_tmp_name/microservices/comments/repository/postgres"
	_commentUcase "park_2020/2020_2_tmp_name/microservices/comments/usecase"

	_photoDelivery "park_2020/2020_2_tmp_name/api/photos/delivery/http"
	_photoRepo "park_2020/2020_2_tmp_name/api/photos/repository/postgres"
	_photoUcase "park_2020/2020_2_tmp_name/api/photos/usecase"

	_userDelivery "park_2020/2020_2_tmp_name/api/users/delivery/http"
	_userRepo "park_2020/2020_2_tmp_name/api/users/repository/postgres"
	_userUcase "park_2020/2020_2_tmp_name/api/users/usecase"

	_authClient "park_2020/2020_2_tmp_name/microservices/authorization/delivery/grpc/client"
	_authDelivery "park_2020/2020_2_tmp_name/microservices/authorization/delivery/http"
	_authRepo "park_2020/2020_2_tmp_name/microservices/authorization/repository/postgres"
	_authUcase "park_2020/2020_2_tmp_name/microservices/authorization/usecase"

	_faceClient "park_2020/2020_2_tmp_name/microservices/face_features/delivery/grpc/client"
)

type application struct {
	servicePort int
	serv        *mux.Router
}

func init() {
	err := godotenv.Load("envs/postgres.env")
	if err != nil {
		log.Fatal(err)
	}
}

func DBConnection() *sql.DB {
	connString := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable",
		os.Getenv("PostgresHost"),
		os.Getenv("PostgresUser"),
		os.Getenv("PostgresPassword"),
		os.Getenv("PostgresDBName"),
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(10)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func (app *application) initServer() {
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Content-Disposition"})
	originsOk := handlers.AllowedOrigins([]string{"https://mi-ami.ru"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	dbConn := DBConnection()

	router := mux.NewRouter()

	metricsProm := metrics.RegisterMetrics(router)
	middleware.NewLoggingMiddleware(metricsProm)

	logrus.SetFormatter(&logrus.TextFormatter{DisableColors: true})
	logrus.WithFields(logrus.Fields{
		"logger": "LOGRUS",
		"host":   "95.163.213.222",
		"port":   ":8080",
	}).Info("Starting server")

	AccessLogOut := new(middleware.AccessLogger)

	contextLogger := logrus.WithFields(logrus.Fields{
		"mode":   "[access_log]",
		"logger": "LOGRUS",
	})
	logrus.SetFormatter(&logrus.JSONFormatter{})
	AccessLogOut.LogrusLogger = contextLogger

	router.Use(AccessLogOut.AccessLogMiddleware(router))

	ar := _authRepo.NewPostgresAuthRepository(dbConn)
	grpcConnAuth, err := grpc.Dial("0.0.0.0:8081", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return
	}

	grpcAuthClient := _authClient.NewAuthClient(grpcConnAuth)
	au := _authUcase.NewAuthUsecase(ar)
	_authDelivery.NewAuthHandler(router, au, grpcAuthClient)

	chr := _chatRepo.NewPostgresChatRepository(dbConn)
	chu := _chatUcase.NewChatUsecase(chr)
	_chatDelivery.NewChatHandler(router, chu, grpcAuthClient)

	cr := _commentRepo.NewPostgresCommentRepository(dbConn)
	grpcConnComments, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
		return
	}

	grpcCommentClient := _commentClient.NewCommentsClientGRPC(grpcConnComments)
	cu := _commentUcase.NewCommentUsecase(cr)
	_commentDelivery.NewCommentHandler(router, cu, grpcCommentClient, grpcAuthClient)

	grpcConnFace, err := grpc.Dial("localhost:8083", grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
		return
	}
	grpcFaceClient := _faceClient.NewFaceClient(grpcConnFace)

	pr := _photoRepo.NewPostgresPhotoRepository(dbConn)
	pu := _photoUcase.NewPhotoUsecase(pr)
	_photoDelivery.NewPhotoHandler(router, pu, grpcAuthClient, grpcFaceClient)

	middleware.MyCORSMethodMiddleware(router)

	serv := &http.Server{
		Addr:         ":8080",
		Handler:      handlers.CORS(originsOk, headersOk, methodsOk, handlers.AllowCredentials())(router),
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	ur := _userRepo.NewPostgresUserRepository(dbConn)
	uu := _userUcase.NewUserUsecase(ur)
	_userDelivery.NewUserHandler(router, uu, grpcAuthClient, grpcFaceClient)

	fmt.Println("Starting server at: 8080")
	err = serv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func newApplication() *application {
	return &application{
		servicePort: 8080,
		serv:        mux.NewRouter().StrictSlash(true),
	}
}

func main() {
	app := newApplication()
	app.initServer()
}
