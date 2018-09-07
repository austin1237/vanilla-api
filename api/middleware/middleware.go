package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

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

func StartTime(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ctx := context.WithValue(r.Context(), "startTime", start)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

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
