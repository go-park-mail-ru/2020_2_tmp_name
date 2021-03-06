package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"park_2020/2020_2_tmp_name/middleware"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"

	metrics "park_2020/2020_2_tmp_name/prometheus"

	_authRepo "park_2020/2020_2_tmp_name/microservices/authorization/repository/postgres"
	_authUcase "park_2020/2020_2_tmp_name/microservices/authorization/usecase"

	grpcServer "park_2020/2020_2_tmp_name/microservices/authorization/delivery/grpc/server"
)

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

func main() {
	dbConn := DBConnection()

	router := mux.NewRouter()

	metricsProm := metrics.RegisterMetrics(router)
	middleware.NewLoggingMiddleware(metricsProm)

	ar := _authRepo.NewPostgresAuthRepository(dbConn)
	au := _authUcase.NewAuthUsecase(ar)

	fmt.Println("Starting server at: 8081")
	grpcServer.StartAuthGRPCServer(au, "auth:8081")
}
