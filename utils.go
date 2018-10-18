package main

import (
	"encoding/base64"
	"crypto/sha512"
)

func EncodePassword(password string) string {
	hash := sha512.New()
	hash.Write([]byte(password))
	sha512_hash :=  base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return sha512_hash
}