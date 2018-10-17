package main

import (
		"fmt"
		"encoding/base64"
		"crypto/sha512"
		"bufio"
		"os"
		)

func EncodePassword(password string) string {
	hash := sha512.New()
	hash.Write([]byte(password))
	sha512_hash :=  base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return sha512_hash
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	fmt.Printf(EncodePassword(password))
}