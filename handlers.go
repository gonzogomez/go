package main

import (
	"net/http"
	"encoding/json"
	"time"
	"log"
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
    		sendJson(w, http.StatusOK, map[string]string{"password": password})
    	default:
    		sendError(w, http.StatusBadRequest, "Only POST method is supported for hash")		
    }
}

func catchAll(w http.ResponseWriter, r *http.Request) {
	sendError(w, http.StatusBadRequest, "URL is not supported")
}

func sendError(w http.ResponseWriter, code int, message string) {
	sendJson(w, code, map[string]string{"error": message})
}

func sendJson(w http.ResponseWriter, code int, payload interface{}) {
	// Keep socket open for 5 seconds before sendng response
	time.Sleep(5 * time.Second)
	response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}