package main

import (
	"database/sql"
	"fmt"
	"log"

	"park_2020/2020_2_tmp_name/middleware"
	"park_2020/2020_2_tmp_name/models"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"

	_authRepo "park_2020/2020_2_tmp_name/microservices/authorization/repository/postgres"
	_authUcase "park_2020/2020_2_tmp_name/microservices/authorization/usecase"

	grpcServer "park_2020/2020_2_tmp_name/microservices/authorization/delivery/grpc/server"
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
	// headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Content-Disposition"})
	// originsOk := handlers.AllowedOrigins([]string{"https://mi-ami.ru"})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	dbConn := DBConnection(&conf)

	// router := mux.NewRouter()

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
	au := _authUcase.NewAuthUsecase(ar)

	grpcServer.StartAuthGRPCServer(au, "localhost:8081")

	// grpcAuthConn, err := grpc.Dial("0.0.0.0:8081", grpc.WithInsecure())
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// grpcAuthClient := authClient.NewAuthClient(grpcAuthConn)

	// _authDelivery.NewUserHandler(router, au)

	// middleware.MyCORSMethodMiddleware(router)

	// sessMiddleware := middleware.NewSessionMiddleware(ar)
	// router.Use(sessMiddleware.SessionMiddleware)

	// serv := &http.Server{
	// 	Addr:         ":8082",
	// 	Handler:      handlers.CORS(originsOk, headersOk, methodsOk, handlers.AllowCredentials())(router),
	// 	WriteTimeout: 60 * time.Second,
	// 	ReadTimeout:  60 * time.Second,
	// }

	// fmt.Println("Starting server at: 8082")
	// err = serv.ListenAndServe()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func newApplication(conf models.Config) *application {
	return &application{
		servicePort: 8080,
		serv:        mux.NewRouter().StrictSlash(true),
	}
}

func main() {
	dbConn := DBConnection(&conf)

	logrus.SetFormatter(&logrus.TextFormatter{DisableColors: true})
	logrus.WithFields(logrus.Fields{
		"logger": "LOGRUS",
		"host":   "95.163.213.222",
		"port":   ":8081",
	}).Info("Starting server")

	AccessLogOut := new(middleware.AccessLogger)

	contextLogger := logrus.WithFields(logrus.Fields{
		"mode":   "[access_log]",
		"logger": "LOGRUS",
	})
	logrus.SetFormatter(&logrus.JSONFormatter{})
	AccessLogOut.LogrusLogger = contextLogger

	ar := _authRepo.NewPostgresUserRepository(dbConn)
	au := _authUcase.NewAuthUsecase(ar)

	grpcServer.StartAuthGRPCServer(au, "localhost:8081")

}