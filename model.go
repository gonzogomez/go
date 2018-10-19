package main

import ()

// JSON object for password hashes to be written to a file
type PasswordHash struct {
    Length  int `json:"length"`
    Time   int    `json:"time"`
}

// JSON object for stat info on password hashes
type Stats struct {
	Total int `json:"Total"`
	Average float64 `json:"Average"`
}