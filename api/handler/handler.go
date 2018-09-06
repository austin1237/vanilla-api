package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/user/api/hasher"
	"github.com/user/api/server"
	"github.com/user/api/stats"
	"github.com/user/api/validator"
)

type key int

var (
	startTimeKey key = 1
)

func Stats(sStats *stats.ServerStats) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientJSON, err := json.Marshal(sStats)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(clientJSON)
	})
	// Output:
	// [{"average": 5000, "total": 1}]
}

func Hash(sStats *stats.ServerStats) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		err := validator.ValidateFormPassword(r.Form)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		userPassword := r.Form["password"][0]
		hashStr := hasher.GenerateHash(userPassword)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(hashStr))
		startTime, ok := r.Context().Value("startTime").(time.Time)
		if !ok {
			log.Println("Error: start time was not found in context, skipping metrics")
		} else {
			sStats.SuccessfulRequest(startTime)
		}
	})
}

func ShutDown(serv server.Api) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Shutting Down"))
		go func() {
			serv.ShutDown()
		}()
	})
}
