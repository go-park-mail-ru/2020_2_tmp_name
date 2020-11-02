package main

import (
	"fmt"
	"log"
	"net/http"

	"park_2020/2020_2_tmp_name/chat"
	"park_2020/2020_2_tmp_name/middleware"
	"park_2020/2020_2_tmp_name/models"
	"park_2020/2020_2_tmp_name/server"
	"park_2020/2020_2_tmp_name/storage"

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

func (app *application) initServer() {
	// headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Content-Disposition"})
	// originsOk := handlers.AllowedOrigins([]string{"http://95.163.213.222:3000"})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	s := server.NewServer()

	server := chat.NewServer("/entry")
	go server.Listen()

	http.Handle("/", http.FileServer(http.Dir("webroot")))

	var err error
	s.DB, err = storage.DBConnection(&conf)
	if err != nil {
		log.Fatalln("database connection failed")
	}

	middleware.MyCORSMethodMiddleware(app.serv)

	path := "/static/avatars/"
	// http.Handle("/", app.serv)
	// app.serv.PathPrefix(path).Handler(http.StripPrefix(path, http.FileServer(http.Dir("."+path))))

	// app.serv.HandleFunc("/health", s.HealthHandler).Methods(http.MethodGet)
	// app.serv.HandleFunc("/api/v1/login", s.Login).Methods(http.MethodGet, http.MethodPost)
	// app.serv.HandleFunc("/api/v1/logout", s.Logout).Methods(http.MethodPost)
	// app.serv.HandleFunc("/api/v1/signup", s.Signup).Methods(http.MethodPost)
	// app.serv.HandleFunc("/api/v1/settings", s.Settings).Methods(http.MethodPost)
	// app.serv.HandleFunc("/api/v1/upload", s.UploadAvatar).Methods(http.MethodPost)
	// app.serv.HandleFunc("/api/v1/add_photo", s.AddPhoto).Methods(http.MethodPost)
	// app.serv.HandleFunc("/api/v1/me", s.MeHandler).Methods(http.MethodGet)
	// app.serv.HandleFunc("/api/v1/feed", s.FeedHandler).Methods(http.MethodGet)

	// app.serv.HandleFunc("/api/v1/like", s.Like).Methods(http.MethodPost)
	// app.serv.HandleFunc("/api/v1/dislike", s.Dislike).Methods(http.MethodPost)
	// app.serv.HandleFunc("/api/v1/comment", s.Comment).Methods(http.MethodPost)

	http.Handle(path, http.StripPrefix(path, http.FileServer(http.Dir("."+path))))

	http.HandleFunc("/health", s.HealthHandler)
	http.HandleFunc("/api/v1/login", s.Login)
	http.HandleFunc("/api/v1/logout", s.Logout)
	http.HandleFunc("/api/v1/signup", s.Signup)
	http.HandleFunc("/api/v1/settings", s.Settings)
	http.HandleFunc("/api/v1/upload", s.UploadAvatar)
	http.HandleFunc("/api/v1/add_photo", s.AddPhoto)
	http.HandleFunc("/api/v1/me", s.MeHandler)
	http.HandleFunc("/api/v1/feed", s.FeedHandler)

	http.HandleFunc("/api/v1/like", s.Like)
	http.HandleFunc("/api/v1/dislike", s.Dislike)
	http.HandleFunc("/api/v1/comment", s.Comment)

	// serv := &http.Server{
	// 	Addr:         ":8080",
	// 	Handler:      handlers.CORS(originsOk, headersOk, methodsOk, handlers.AllowCredentials())(app.serv),
	// 	WriteTimeout: 60 * time.Second,
	// 	ReadTimeout:  60 * time.Second,
	// }

	fmt.Println("Starting server at: 8080")
	err = http.ListenAndServe(":8080", nil)
	// err = serv.ListenAndServe()
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
