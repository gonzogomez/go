package main

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
    router := http.NewServeMux()
	router.HandleFunc("/hash", hash)
	router.HandleFunc("/shutdown", shutdown)
	router.HandleFunc("/stats", stats)
	router.HandleFunc("/", catchAll)
    return router
}