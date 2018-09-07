package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

// ArtificalWait will wait for 5 seconds
func ArtificalWait(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		next.ServeHTTP(w, r)
	})
}

// PostOnly will return 404s on any non Post Request
func PostOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// GetOnly will return 404s on any non get request
func GetOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// StartTime will attach the current time to the request's contexxt
func StartTime(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ctx := context.WithValue(r.Context(), "startTime", start)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ReqInfo will log basic info about the a request sent to the server
func ReqInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			requestID, ok := r.Context().Value("requestIDKey").(string)
			if !ok {
				requestID = "unknown"
			}
			log.Println("ReqID:"+requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
		}()
		next.ServeHTTP(w, r)
	})
}

// ReqID attaches either a requestId string to the requests context
// ReqID with either use the clients X-Request-Id or generate a new one
func ReqID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-Id")
		if requestID == "" {
			requestID = fmt.Sprintf("%d", time.Now().UnixNano())
		}
		ctx := context.WithValue(r.Context(), "requestIDKey", requestID)
		w.Header().Set("X-Request-Id", requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
