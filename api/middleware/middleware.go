package middleware

import (
	"net/http"
	"time"
)

type key int

const (
	requestIDKey key = 0
	startTimeKey key = 1
)

// func logging(logger *log.Logger) func(http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			defer func() {
// 				requestID, ok := r.Context().Value(requestIDKey).(string)
// 				if !ok {
// 					requestID = "unknown"
// 				}
// 				logger.Println(requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
// 			}()
// 			next.ServeHTTP(w, r)
// 		})
// 	}
// }

func ArtificalWait(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		next.ServeHTTP(w, r)
	})
}

func PostOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}
