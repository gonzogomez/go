package main

import (
	"net/http"
)

// Returns a http.ServeMux with endpoints and methods to those endpoints
func NewRouter() *http.ServeMux {
    router := http.NewServeMux()
	router.HandleFunc("/hash", hash)
	router.HandleFunc("/shutdown", shutdown)
	router.HandleFunc("/stats", stats)
	router.HandleFunc("/", catchAll)
    return router
}