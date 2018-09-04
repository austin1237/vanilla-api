package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/user/api/router"
)

type key int

const (
	requestIDKey key = 0
	serverKey    key = 2
)

type api struct {
	Server *http.Server
	Done   chan bool
}

var serv api

func Start(done chan bool) {
	mux := http.NewServeMux()
	mux.Handle("/shutdown", shutDownHandler())
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Server is ready to handle requests at :8080")
	router.SetUpRoutes(mux)
	serv.Done = done
	serv.Server = &http.Server{
		Addr:     ":8080",
		ErrorLog: logger,
		Handler:  mux,
	}
	if err := serv.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Error listening on :8080 %v\n", err)
	}
}

func shutDownHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ShuttingDown"))
		go func() {
			ShutDown()
		}()
	})
}

func ShutDown() {
	fmt.Println("shutting the server down")
	serv.Server.SetKeepAlivesEnabled(false)
	if err := serv.Server.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
	serv.Done <- true
}
