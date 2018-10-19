package main

import (
	"encoding/base64"
	"encoding/json"
	"crypto/sha512"
	"time"
	"log"
	"os"
)


// EncodePassword takes in a string and provides the sha512 of it in encoded in base64.
// Time to encode and password length is noted and written to a json file.
func EncodePassword(password string) string {
	start := time.Now()
	hash := sha512.New()
	hash.Write([]byte(password))
	sha512_hash :=  base64.StdEncoding.EncodeToString(hash.Sum(nil))
	t := time.Now()
	elapsed := t.Sub(start)
	WriteToFile(len(password), int(elapsed/1000))
	return sha512_hash
}

// GetStatus calculate average encode time of hashes and return total and average.
// Reads from a json file to calculate average encode time from all hashes seen so far.
// If file is empty then {0,0} is returned.
func GetStats() Stats {
	// Check if file exists if not then there are no stats return {0,0}
	if _, existserr := os.Stat("./data.json"); !os.IsNotExist(existserr) {
		var jsonArray []PasswordHash
		jsonArray = LoadJSONFile("./data.json")// path/to/whatever exists
		sum := 0

		// Get sum
		for i:=0; i<len(jsonArray); i++ {
			sum += jsonArray[i].Time
		}

		// Calculate average
		average := float64(float64(sum)/float64(len(jsonArray)))
		stat := Stats{Total: len(jsonArray), Average: average}
		return stat
	} else {
		stat := Stats{Total: 0, Average:0}
		return stat
	}
}


// Helper function to write password length and encode time to a json file.
// WriteToFile checks if the json file exists, if so then appends to the json object.
func WriteToFile(password_length int, encode_time int) {
	var jsonArray []PasswordHash
	// Load jsonArray if file exists
	if _, existserr := os.Stat("./data.json"); !os.IsNotExist(existserr) {
		jsonArray = LoadJSONFile("./data.json")// path/to/whatever exists
	}

	// Add to json array
	passhash := PasswordHash{Length: password_length, Time: encode_time}
	jsonArray = append(jsonArray, passhash)	

	// Encode json array
	jdata1, err := json.MarshalIndent(jsonArray, "", "  ")
	if err != nil {
		log.Printf("error:", err)
	}

	// Write to file
	jsonFile, err := os.Create("./data.json")
	jsonFile.Write(jdata1)
}


// Helper function to read json file and return an array of json objects of type PasswordHash.
func LoadJSONFile(filename string) []PasswordHash {
	var jsonArray []PasswordHash
	jsonFile, err := os.Open(filename)
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	// err = decoder.Decode(&passhash)
	err = decoder.Decode(&jsonArray)

    if err != nil {
        log.Printf(err.Error())
    }
    return jsonArray
}