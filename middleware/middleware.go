package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	metrics "park_2020/2020_2_tmp_name/prometheus"

	uuid "github.com/nu7hatch/gouuid"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type AccessLogger struct {
	LogrusLogger *logrus.Entry
}

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

func NewLoggingMiddleware(metrics *metrics.PromMetrics) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			msg := fmt.Sprintf("URL: %s, METHOD: %s", r.RequestURI, r.Method)
			log.Println(msg)

			reqTime := time.Now()
			next.ServeHTTP(w, r)
			respTime := time.Since(reqTime)
			if r.URL.Path != "/metrics" {
				metrics.Hits.WithLabelValues(strconv.Itoa(http.StatusOK), r.URL.String(), r.Method).Inc()
				metrics.Timings.WithLabelValues(
					strconv.Itoa(http.StatusOK), r.URL.String(), r.Method).Observe(respTime.Seconds())
			}
		})
	}
}

func generateCsrfLogic(w http.ResponseWriter) {
	csrf, err := uuid.NewV4()
	if err != nil {

		return
	}
	timeDelta := time.Now().Add(time.Hour * 24 * 30)
	cookie := &http.Cookie{Name: "csrf", Value: csrf.String(), Path: "/", HttpOnly: true, Expires: timeDelta}

	http.SetCookie(w, cookie)
	w.Header().Set("csrf", csrf.String())

}

func SetCSRF(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			generateCsrfLogic(w)
			next.ServeHTTP(w, r)
		})
}

func CheckCSRF(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			csrf := r.Header.Get("X-Csrf-Token")
			csrfCookie, err := r.Cookie("csrf")

			if err != nil || csrf == "" || csrfCookie.Value == "" || csrfCookie.Value != csrf {
				return
			}
			generateCsrfLogic(w)
			next.ServeHTTP(w, r)
		})

}
