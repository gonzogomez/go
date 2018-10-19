package main

import ()

type PasswordHash struct {
    Length  int `json:"length"`
    Time   int    `json:"time"`
}

type Stats struct {
	Total int `json:"Total"`
	Average float64 `json:"Average"`
}