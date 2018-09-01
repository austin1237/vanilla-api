package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/user/api/router"
)

func main() {
	fmt.Println("we out here boy1")
	mux := http.NewServeMux()
	router.SetUpRoutes(mux)
	log.Printf("serving on port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
