package main

import (
	"database/sql"
	"fmt"
	"log"

	"park_2020/2020_2_tmp_name/middleware"
	"park_2020/2020_2_tmp_name/models"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"

	metrics "park_2020/2020_2_tmp_name/prometheus"

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

	dbConn := DBConnection(&conf)

	router := mux.NewRouter()

	metricsProm := metrics.RegisterMetrics(router)
	middleware.NewLoggingMiddleware(metricsProm)

	ar := _authRepo.NewPostgresAuthRepository(dbConn)
	au := _authUcase.NewAuthUsecase(ar)

	grpcServer.StartAuthGRPCServer(au, "localhost:8081")
}

func newApplication(conf models.Config) *application {
	return &application{
		servicePort: 8080,
		serv:        mux.NewRouter().StrictSlash(true),
	}
}

func main() {
	dbConn := DBConnection(&conf)

	router := mux.NewRouter()

	metricsProm := metrics.RegisterMetrics(router)
	middleware.NewLoggingMiddleware(metricsProm)

	ar := _authRepo.NewPostgresAuthRepository(dbConn)
	au := _authUcase.NewAuthUsecase(ar)

	grpcServer.StartAuthGRPCServer(au, "localhost:8081")
}
