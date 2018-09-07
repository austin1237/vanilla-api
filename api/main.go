package main

import (
	"fmt"
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
	sStats := stats.New()
	server := server.New(apiPort)
	router := router.CreateRouter(sStats, server)
	server.RegisterRoutes(router)
	server.Start()
	os.Exit(0)
}
