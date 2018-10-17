package main

import (
		"fmt"
		"encoding/base64"
		"crypto/sha512"
		)

func EncodePassword(password string) string {
	hash := sha512.New()
	hash.Write([]byte(password))
	sha512_hash :=  base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return sha512_hash
}

func main() {
	var password string
	for password != "q" {
		fmt.Printf("Enter text or q to quit: ")
		fmt.Scanln(&password)
		fmt.Printf(EncodePassword(password))
	}
}