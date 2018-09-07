package main

import (
	"log"
	"os"

	"github.com/user/api/router"
	"github.com/user/api/server"
	"github.com/user/api/stats"
)

func init() {

}

func main() {
	done := make(chan bool, 1)
	sStats := stats.New()
	server := server.New(done, "8080")
	router := router.CreateRouter(sStats, server)
	server.RegisterRoutes(router)
	server.Start()
	<-done
	log.Println("api shutdown exiting")
	os.Exit(0)
}
