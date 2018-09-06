package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type Api struct {
	Server *http.Server
	Done   chan bool
	Mux    *http.ServeMux
}

func New(done chan bool, port string) Api {
	httpServer := &http.Server{
		Addr: ":" + port,
	}
	serv := Api{
		Done:   done,
		Server: httpServer,
	}
	return serv
}

func (serv Api) RegisterRoutes(mux *http.ServeMux) {
	serv.Server.Handler = mux
}

// Start starts the api server
func (serv Api) Start() {
	log.Println("listening on :8080")
	if err := serv.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error listening on :8080 %v\n", err)
	}
}

func (serv Api) ShutDown() {
	fmt.Println("shutting the server down")
	serv.Server.SetKeepAlivesEnabled(false)
	if err := serv.Server.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
	serv.Done <- true
}
