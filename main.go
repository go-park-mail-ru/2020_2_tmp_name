package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"park_2020/2020_2_tmp_name/middleware"
	"park_2020/2020_2_tmp_name/models"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"

	_userHttpDelivery "park_2020/2020_2_tmp_name/users/delivery/http"
	_userRepo "park_2020/2020_2_tmp_name/users/repository/postgres"
	_userUcase "park_2020/2020_2_tmp_name/users/usecase"
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
	originsOk := handlers.AllowedOrigins([]string{"http://95.163.213.222:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	var err error
	dbConn := DBConnection(&conf)

	router := mux.NewRouter()

	u := _userRepo.NewPostgresUserRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	uu := _userUcase.NewUserUsecase(u, timeoutContext)

	_userHttpDelivery.NewUserHandler(router, uu)

	middleware.MyCORSMethodMiddleware(router)

	serv := &http.Server{
		Addr:         ":8080",
		Handler:      handlers.CORS(originsOk, headersOk, methodsOk, handlers.AllowCredentials())(router),
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

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
