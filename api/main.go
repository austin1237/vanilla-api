package main

import (
	"log"
	"os"

	"github.com/user/api/server"
)

func main() {
	done := make(chan bool, 1)
	server.Start(done)
	<-done
	log.Println("api shutdown exiting")
	os.Exit(0)
}
