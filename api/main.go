package main

import (
	"fmt"
	"log"
	"os"

	"github.com/user/api/router"
	"github.com/user/api/server"
	"github.com/user/api/stats"
)

var (
	//ENV VARIABLES
	apiPort string
)

func init() {
	apiPort = os.Getenv("API_PORT")
	if apiPort == "" {
		fmt.Println("API_PORT ENV var was not set.")
		os.Exit(1)
	}
}

func main() {
	done := make(chan bool, 1)
	sStats := stats.New()
	server := server.New(done, apiPort)
	router := router.CreateRouter(sStats, server)
	server.RegisterRoutes(router)
	server.Start()
	<-done
	log.Println("api shutdown exiting")
	os.Exit(0)
}
