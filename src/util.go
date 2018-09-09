package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"math/rand"
)

// Generate uuid useing linux program
func GenerateUUID() string {
	response, err := http.Get("https://api.timo-reymann.de/uuid")

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	responseData, _ := ioutil.ReadAll(response.Body)

	var uuidResponse ApiUUIDResponse

	err = json.Unmarshal([]byte(string(responseData)), &uuidResponse)

	// fallback to local random string instead of using api
	if err != nil {
		return RandStringRunes(10)
	}

	return uuidResponse.Data.Uuid
}

// Log message for request
func Log(uuid string, msg string) {
	fmt.Println("[" + uuid + "] " + msg);
}

// Build file name for identifier
func BuildFileName(identifier string) string {
	return "./" + identifier + ".tar.gz"
}

// Letters for random string
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Create random string
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
