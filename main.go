package main

import (
	"net/http"
	"os"
	"os/signal"
	"log"
	"context"
)

func main() {
	// create router class and lunch server
	router := NewRouter()
	server := http.Server{Addr: ":8080", Handler: router}

	// Setting up to listen to interupt signal to shutdown gracefully
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := server.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Printf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}