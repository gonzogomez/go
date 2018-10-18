package main

import (
	"net/http"
)

func main() {
	// create router class and lunch server
	router := NewRouter()
	http.ListenAndServe(":8080", router)	
}