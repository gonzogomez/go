package main

import (
	"net/http"
	"encoding/json"
	"time"
	"log"
	"syscall"
)

func hash(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    	case "POST":
    		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
    		if err := r.ParseForm(); err != nil {
    			log.Fatal(w, "ParseForm() err: %v", err)
    			return
    		}

    		// Parse for password field
    		password := r.FormValue("password")
    		if len(password) == 0 {
    			sendError(w, http.StatusBadRequest, "Password field was not set")
    			return
    		}

    		// Encode passwword
    		password = EncodePassword(password)
    		log.Printf("Encoding Password...\n")
    		sendJson(w, http.StatusOK, map[string]string{"password": password}, 5)
    	default:
    		sendError(w, http.StatusBadRequest, "Only POST method is supported for hash")		
    }
}

func shutdown(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
    	case "POST":
    		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
    		if err := r.ParseForm(); err != nil {
    			log.Fatal(w, "ParseForm() err: %v", err)
    			return
    		}

    		log.Printf("Shutting Down server...\n")
    		sendJson(w, http.StatusOK, map[string]string{"message": "shutting down server...."}, 0)
    		// Send interupt signal
    		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
    	default:
    		sendError(w, http.StatusBadRequest, "Only POST method is supported for this shutdown")		
    }
}

func catchAll(w http.ResponseWriter, r *http.Request) {
	sendError(w, http.StatusBadRequest, "URL is not supported")
}

func sendError(w http.ResponseWriter, code int, message string) {
	sendJson(w, code, map[string]string{"error": message}, 5)
}

func sendJson(w http.ResponseWriter, code int, payload interface{}, sleeptime int) {
	// Keep socket open for 5 seconds before sendng response
	time.Sleep(time.Duration(sleeptime) * time.Second)
	response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}