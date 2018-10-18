package main

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
    router := http.NewServeMux()
	router.HandleFunc("/hash", hash)
	router.HandleFunc("/", catchAll)
    return router
}