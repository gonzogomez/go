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
    			json.NewEncoder(w).Encode("Password field was not set.")
    			time.Sleep(5 * time.Second)
    			return
    		}

    		// Encode passwword
    		password = EncodePassword(password)
    		log.Printf("Encoding Password...\n")
    		// Keep socket open for 5 seconds before sendng response
    		time.Sleep(5 * time.Second)
    		json.NewEncoder(w).Encode(password)
    	default:
    		json.NewEncoder(w).Encode("Only POST method is supported for hash.")
    }
}

func catchAll(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	json.NewEncoder(w).Encode("Url is not supported.")
}