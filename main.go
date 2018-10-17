package main

import (
		"fmt"
		"encoding/base64"
		"crypto/sha512"
		"net/http"
		"encoding/json"
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
    			fmt.Fprintf(w, "ParseForm() err: %v", err)
    			return
    		}
    		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
    		password := r.FormValue("password")
    		password = EncodePassword(password)
    		json.NewEncoder(w).Encode(password)
    	default:
    		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
    }
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/hash", hash)
	http.ListenAndServe(":8080", router)	
}