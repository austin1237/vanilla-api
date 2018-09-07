package server

import (
	"context"
	"log"
	"net/http"
)

type Api struct {
	Server *http.Server
	Done   chan bool
	Mux    *http.ServeMux
}

func New(port string) Api {
	done := make(chan bool, 1)
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

// Start the api server beings listening for requests
// Blocks until Shutdown has finished
func (serv Api) Start() {
	log.Println("listening on " + serv.Server.Addr)
	if err := serv.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error listening on %v %v\n", serv.Server.Addr, err)
	}
	<-serv.Done
}

func (serv Api) ShutDown() {
	log.Println("shutting the server down")
	serv.Server.SetKeepAlivesEnabled(false)
	if err := serv.Server.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
	log.Println("server down has shutdown")
	serv.Done <- true
}
