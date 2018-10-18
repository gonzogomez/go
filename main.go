package main

import (
		"encoding/base64"
		"crypto/sha512"
		"net/http"
		"encoding/json"
		"time"
		"log"
		)

func EncodePassword(password string) string {
	hash := sha512.New()
	hash.Write([]byte(password))
	sha512_hash :=  base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return sha512_hash
}

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
    		log.Printf("%s is encoding...\n", password)
    		// Keep socket open for 5 seconds before sendng response
    		time.Sleep(5 * time.Second)
    		json.NewEncoder(w).Encode(password)
    	default:
    		json.NewEncoder(w).Encode("Only POST method is supported for hash.")
    }
}

func catch_all(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	json.NewEncoder(w).Encode("Url is not supported.")
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/hash", hash)
	router.HandleFunc("/", catch_all)
	http.ListenAndServe(":8080", router)	
}