package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	port    = "8080"
	portEnv = "PORT"
)

func main() {
	log.Print("Starting the service...")
	// Getting correct port
	if env := os.Getenv(portEnv); env != "" {
		port = env
	}
	router := getHandler()
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

func getHandler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	// Liveness check endpoint
	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})

	// Readiness check endpoint
	r.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})

	return r
}
