package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"

	"park_2020/2020_2_tmp_name/middleware"
	"park_2020/2020_2_tmp_name/models"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"

	metrics "park_2020/2020_2_tmp_name/prometheus"

	_chatDelivery "park_2020/2020_2_tmp_name/api/chats/delivery/http"
	_chatRepo "park_2020/2020_2_tmp_name/api/chats/repository/postgres"
	_chatUcase "park_2020/2020_2_tmp_name/api/chats/usecase"

	_commentClientGRPC "park_2020/2020_2_tmp_name/microservices/comments/delivery/grpc/client"
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
)

type application struct {
	servicePort int
	serv        *mux.Router
}

var conf models.Config

func init() {
	models.LoadConfig(&conf)
}

func DBConnection(conf *models.Config) *sql.DB {
	connString := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable",
		conf.SQLDataBase.Server,
		conf.SQLDataBase.UserID,
		conf.SQLDataBase.Password,
		conf.SQLDataBase.Database,
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

	dbConn := DBConnection(&conf)

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

	// router.Use(AccessLogOut.AccessLogMiddleware(router))
	ar := _authRepo.NewPostgresUserRepository(dbConn)
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
	cu := _commentUcase.NewCommentUsecase(cr)
	grpcConnComments, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
		return
	}
	ccGRPC := _commentClientGRPC.NewCommentsClientGRPC(grpcConnComments)
	_commentDelivery.NewCommentHandler(router, cu, ccGRPC, grpcAuthClient)

	pr := _photoRepo.NewPostgresPhotoRepository(dbConn)
	pu := _photoUcase.NewPhotoUsecase(pr)
	_photoDelivery.NewPhotoHandler(router, pu, grpcAuthClient)

	middleware.MyCORSMethodMiddleware(router)

	serv := &http.Server{
		Addr:         ":8080",
		Handler:      handlers.CORS(originsOk, headersOk, methodsOk, handlers.AllowCredentials())(router),
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	ur := _userRepo.NewPostgresUserRepository(dbConn)
	uu := _userUcase.NewUserUsecase(ur)
	_userDelivery.NewUserHandler(router, uu, grpcAuthClient)

	fmt.Println("Starting server at: 8080")
	err = serv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func newApplication(conf models.Config) *application {
	return &application{
		servicePort: 8080,
		serv:        mux.NewRouter().StrictSlash(true),
	}
}

func main() {
	app := newApplication(conf)
	app.initServer()

}
