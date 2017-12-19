package main

import(
	"os/exec"
	"strings"
	"fmt"
)

// Generate uuid useing linux program
func GenerateUUID() string  {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		panic(err)
	}
	return strings.TrimRight(string(out),"\r\n");
}

// Log message for request
func Log(uuid string, msg string) {
	fmt.Println("[" + uuid + "] " + msg);
}
