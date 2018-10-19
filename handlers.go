package main

import (
	"net/http"
	"encoding/json"
	"time"
	"log"
	"syscall"
)

// Endpoint method that expects the password field to be set.
// If password field is set then EncodePassword is call to ecode and is set sent back.
// Otherwise error message is sent if password field is not set or wrong request type is found.
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

// Endpoint method for gracefull shutdown of the server.
// If request is correct method sends interupt signal to shutdown server.
// Otherwise error message is sent if wrong request type is found.
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

// Endpoint method that returns basic information about password hashes seen so far.
// Error message is sent if wrong request type is found.
func stats(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
    	case "GET":
    		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
    		if err := r.ParseForm(); err != nil {
    			log.Fatal(w, "ParseForm() err: %v", err)
    			return
    		}

    		log.Printf("Getting stats...\n")
    		encode_stats := GetStats()
    		sendJson(w, http.StatusOK, encode_stats, 5)
    	default:
    		sendError(w, http.StatusBadRequest, "Only GET method is supported for this shutdown")		
    }
}

// Endpoint method that handles any paths not supported.
// Error message is sent with an approriate message.
func catchAll(w http.ResponseWriter, r *http.Request) {
	sendError(w, http.StatusBadRequest, "URL is not supported")
}


// Helper function to send errors
func sendError(w http.ResponseWriter, code int, message string) {
	sendJson(w, code, map[string]string{"error": message}, 5)
}

// Helper function to send json encoded message
func sendJson(w http.ResponseWriter, code int, payload interface{}, sleeptime int) {
	// Keep socket open for 5 seconds before sendng response
	time.Sleep(time.Duration(sleeptime) * time.Second)
	response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}