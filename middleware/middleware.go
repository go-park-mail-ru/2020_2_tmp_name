package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type AccessLogger struct {
	LogrusLogger *logrus.Entry
}

// const SessionCookieName = "authCookie"

// type SessionMiddleware struct {
// 	sessionRepo domain.UserRepository
// }

// func NewSessionMiddleware(u domain.UserRepository) *SessionMiddleware {
// 	return &SessionMiddleware{
// 		sessionRepo: u,
// 	}
// }

func MyCORSMethodMiddleware(_ *mux.Router) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			fmt.Printf("URL: %s, METHOD: %s", req.RequestURI, req.Method)
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Allow-Origin", "https://mi-ami.ru")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			if req.Method == "OPTIONS" {
				return
			}
			next.ServeHTTP(w, req)
		})
	}
}

func (ac *AccessLogger) AccessLogMiddleware(_ *mux.Router) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)

			ac.LogrusLogger.WithFields(logrus.Fields{
				"method":      r.Method,
				"remote_addr": r.RemoteAddr,
				"work_time":   time.Since(start),
			}).Info(r.URL.Path)
		})
	}
}

// func (s *SessionMiddleware) SessionMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if len(r.Cookies()) != 0 {
// 			telephone, err := s.sessionRepo.SelectUserBySession(r.Cookies()[0].Value)
// 			if err != nil {
// 				return
// 			}

// 			if found := s.sessionRepo.CheckUserBySession(r.Cookies()[0].Value); !found {
// 				err = s.sessionRepo.InsertSession(r.Cookies()[0].Value, telephone)
// 				if err != nil {
// 					return
// 				}
// 			}

// 			r = r.WithContext(context.WithValue(r.Context(), "session", r.Cookies()[0].Value))
// 		}

// 		next.ServeHTTP(w, r)

// 	})
// }
