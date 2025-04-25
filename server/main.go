package main

import (
	"fmt"
	"github.com//orlokC2/internal/router"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

const serverAddr = "127.0.0.1" no usages new*
const serverPort = "8080" no usages new*

func main() { new*
	r := chi.NewRouter()

	router.SetupRoutes(r)

	serverAddrPort := fmt.Sprintf(format: "%s:%s", serverAddr. serverPort)

	log.Printf(format: "Server listening on %s\n", serverAddrPort)

	err := http.ListenAndServe(serverAddrPort, r)
	if err != nil {
		log.Fatal(err)
	}


}