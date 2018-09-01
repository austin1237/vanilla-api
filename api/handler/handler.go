package handler

import (
	"net/http"
	"time"

	"github.com/user/api/hasher"
)

func GetHash() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		// Validate form here
		userPassword := r.Form["password"][0]
		hashStr := hasher.HashString(userPassword)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(hashStr))
	})
}

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
