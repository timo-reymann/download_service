package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
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

	if err != nil {
		panic(err)
	}

	return uuidResponse.Data.Uuid
}

// Log message for request
func Log(uuid string, msg string) {
	fmt.Println("[" + uuid + "] " + msg);
}

func BuildFileName(identifier string) string {
	return "./" + identifier + ".tar.gz"
}
