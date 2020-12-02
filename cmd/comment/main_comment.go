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

	_commentRepo "park_2020/2020_2_tmp_name/microservices/comments/repository/postgres"
	_commentUcase "park_2020/2020_2_tmp_name/microservices/comments/usecase"

	grpcServer "park_2020/2020_2_tmp_name/microservices/comments/delivery/grpc/server"
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
	//headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Content-Disposition"})
	//originsOk := handlers.AllowedOrigins([]string{"https://mi-ami.ru"})
	//methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	dbConn := DBConnection(&conf)

	//router := mux.NewRouter()

	logrus.SetFormatter(&logrus.TextFormatter{DisableColors: true})
	logrus.WithFields(logrus.Fields{
		"logger": "LOGRUS",
		"host":   "95.163.213.222",
		"port":   ":8082",
	}).Info("Starting server")

	AccessLogOut := new(middleware.AccessLogger)

	contextLogger := logrus.WithFields(logrus.Fields{
		"mode":   "[access_log]",
		"logger": "LOGRUS",
	})
	logrus.SetFormatter(&logrus.JSONFormatter{})
	AccessLogOut.LogrusLogger = contextLogger

	// router.Use(AccessLogOut.AccessLogMiddleware(router))

	cr := _commentRepo.NewPostgresCommentRepository(dbConn)
	cu := _commentUcase.NewCommentUsecase(cr)

	go grpcServer.StartCommentsGRPCServer(cu, "localhost:8082")
	// _commentDelivery.NewCommentHandler(router, cu)

	//middleware.MyCORSMethodMiddleware(router)
	//
	//serv := &http.Server{
	//	Addr:         "localhost:8083",
	//	Handler:      handlers.CORS(originsOk, headersOk, methodsOk, handlers.AllowCredentials())(router),
	//	WriteTimeout: 60 * time.Second,
	//	ReadTimeout:  60 * time.Second,
	//}

	//fmt.Println("Starting server at: 8083")
	//err := serv.ListenAndServe()
	//if err != nil {
	//	log.Fatal(err)
	//}
}
//
//func newApplication(conf models.Config) *application {
//	return &application{
//		servicePort: 8083,
//		serv:        mux.NewRouter().StrictSlash(true),
//	}
//}

func main() {
	dbConn := DBConnection(&conf)

	logrus.SetFormatter(&logrus.TextFormatter{DisableColors: true})
	logrus.WithFields(logrus.Fields{
		"logger": "LOGRUS",
		"host":   "95.163.213.222",
		"port":   ":8082",
	}).Info("Starting server")

	AccessLogOut := new(middleware.AccessLogger)

	contextLogger := logrus.WithFields(logrus.Fields{
		"mode":   "[access_log]",
		"logger": "LOGRUS",
	})
	logrus.SetFormatter(&logrus.JSONFormatter{})
	AccessLogOut.LogrusLogger = contextLogger

	cr := _commentRepo.NewPostgresCommentRepository(dbConn)
	cu := _commentUcase.NewCommentUsecase(cr)

	grpcServer.StartCommentsGRPCServer(cu, "localhost:8082")
}
