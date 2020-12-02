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
	fmt.Println("я в ините")
	models.LoadConfig(&conf)
}

func DBConnection(conf *models.Config) *sql.DB {
	connString := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable",
		conf.SQLDataBase.Server,
		conf.SQLDataBase.UserID,
		conf.SQLDataBase.Password,
		conf.SQLDataBase.Database,
	)

	fmt.Println(connString)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	
	db.SetMaxOpenConns(10)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("я в мейне")
	return db
}

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

	grpcServer.StartCommentsGRPCServer(cu, ":8082")
}
