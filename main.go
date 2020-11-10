package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"park_2020/2020_2_tmp_name/middleware"
	"park_2020/2020_2_tmp_name/models"
	"park_2020/2020_2_tmp_name/server"
	"park_2020/2020_2_tmp_name/storage"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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

type Service struct {
	DB  *sql.DB
	Hub *server.Hub
}

func (app *application) initServer() {
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Content-Disposition"})
	originsOk := handlers.AllowedOrigins([]string{"http://95.163.213.222:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	s := server.NewServer()

	server.MyHub = server.NewHub()
	go s.Run()

	var err error
	s.DB = storage.DBConnection(&conf)

	middleware.MyCORSMethodMiddleware(app.serv)

	serv := &http.Server{
		Addr:         ":8080",
		Handler:      handlers.CORS(originsOk, headersOk, methodsOk, handlers.AllowCredentials())(app.serv),
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
