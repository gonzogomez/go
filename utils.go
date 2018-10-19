package main

import (
	"encoding/base64"
	"encoding/json"
	"crypto/sha512"
	"time"
	"log"
	"os"
)

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

func WriteToFile(password_length int, encode_time int) {
	var jsonArray []PasswordHash
	if _, existserr := os.Stat("./data.json"); !os.IsNotExist(existserr) {
		jsonArray = LoadJSONFile("./data.json")// path/to/whatever exists
	}

	passhash := PasswordHash{Length: password_length, Time: encode_time}
	jsonArray = append(jsonArray, passhash)	

	jdata1, err := json.MarshalIndent(jsonArray, "", "  ")
	if err != nil {
		log.Printf("error:", err)
	}

	jsonFile, err := os.Create("./data.json")
	jsonFile.Write(jdata1)
}

func LoadJSONFile(filename string) []PasswordHash {
	// passhash :=PasswordHash{}
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