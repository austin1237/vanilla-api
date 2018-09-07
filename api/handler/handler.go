package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/user/api/hasher"
	"github.com/user/api/server"
	"github.com/user/api/stats"
	"github.com/user/api/validator"
)

// Stats will return the current servers stats in JSON format
func Stats(sStats *stats.ServerStats) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID, ok := r.Context().Value("requestIDKey").(string)
		if !ok {
			requestID = "unknown"
		}
		clientJSON, err := json.Marshal(sStats)
		if err != nil {
			logTxt := fmt.Sprintf("ReqID:%v Error:%v", requestID, err.Error())
			log.Println(logTxt)
			http.Error(w, "unable to marshal stats json", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(clientJSON)
	})
}

// Hash will return a hash of a password sent in and increment the server's stats
func Hash(sStats *stats.ServerStats) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID, ok := r.Context().Value("requestIDKey").(string)
		if !ok {
			requestID = "unknown"
		}
		r.ParseForm()
		err := validator.ValidateFormPassword(r.Form)
		if err != nil {
			logTxt := fmt.Sprintf("ReqID:%v Error:%v", requestID, err.Error())
			log.Println(logTxt)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		userPassword := r.Form["password"][0]
		hashStr := hasher.GenerateHash(userPassword)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(hashStr))
		startTime, ok := r.Context().Value("startTime").(time.Time)
		if !ok {
			logTxt := fmt.Sprintf("ReqID:%v Error:start time was not found in context, skipping metrics", requestID)
			log.Println(logTxt)
		} else {
			now := time.Now()
			sStats.SuccessfulRequest(startTime, now)
		}
	})
}

// ShutDown will gracefully shut the server down
func ShutDown(serv server.Api) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Shutting Down"))
		go func() {
			serv.ShutDown()
		}()
	})
}
