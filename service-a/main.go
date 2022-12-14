package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

var (
	port    = "8080"
	portEnv = "PORT"
)

func main() {
	log.Print("Starting service A...")
	// Getting correct port
	if env := os.Getenv(portEnv); env != "" {
		port = env
	}
	router := getHandler()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	log.Printf("The service is ready to listen and serve on port %s", port)

	go func() {
		log.Fatal(printBitcoinInLoop())
	}()
	log.Printf("Printing Bitcoin")

	killSignal := <-interrupt
	switch killSignal {
	case os.Interrupt:
		log.Print("Got SIGINT...")
	case syscall.SIGTERM:
		log.Print("Got SIGTERM...")
	}

	log.Print("The service is shutting down...")
	srv.Shutdown(context.Background())
	log.Print("Done")
}

func getHandler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World from Service A!")
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
