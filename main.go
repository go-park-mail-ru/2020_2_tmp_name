package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"park_2020/2020_2_tmp_name/middleware"
	"park_2020/2020_2_tmp_name/models"
	"park_2020/2020_2_tmp_name/server"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type application struct {
	servicePort int
	s           *mux.Router
}

var conf models.Config

func init() {
	models.LoadConfig(&conf)
}

func (app *application) initServer() {

	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Content-Disposition"})
	originsOk := handlers.AllowedOrigins([]string{"http://95.163.213.222:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	app.s = mux.NewRouter().StrictSlash(true)

	s := server.NewServer()

	middleware.MyCORSMethodMiddleware(app.s)

	http.Handle("/", app.s)
	path := "/static/avatars/"
	app.s.PathPrefix(path).Handler(http.StripPrefix(path, http.FileServer(http.Dir("."+path))))

	app.s.HandleFunc("/health", s.HealthHandler).Methods(http.MethodGet)
	app.s.HandleFunc("/api/v1/login", s.Login).Methods(http.MethodGet, http.MethodPost)
	app.s.HandleFunc("/api/v1/logout", s.Logout).Methods(http.MethodPost, http.MethodOptions)
	app.s.HandleFunc("/api/v1/signup", s.Signup).Methods(http.MethodPost, http.MethodOptions)
	app.s.HandleFunc("/api/v1/settings", s.Settings).Methods(http.MethodPost)
	app.s.HandleFunc("/api/v1/upload", s.UploadAvatar).Methods(http.MethodPost)
	app.s.HandleFunc("/api/v1/add_photo", s.AddPhoto).Methods(http.MethodPost)
	app.s.HandleFunc("/api/v1/me", s.MeHandler).Methods(http.MethodGet, http.MethodOptions)
	app.s.HandleFunc("/api/v1/feed", s.FeedHandler).Methods(http.MethodGet, http.MethodOptions)

	serv := &http.Server{
		Addr:         ":8080",
		Handler:      handlers.CORS(originsOk, headersOk, methodsOk, handlers.AllowCredentials())(app.s),
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	fmt.Println("Starting server at: 8080")
	err := serv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func newApplication(conf models.Config) *application {

	return &application{
		servicePort: 8080,
	}
}

func main() {

	app := newApplication(conf)
	app.initServer()
}
