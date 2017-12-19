package main

import(
	"os/exec"
	"strings"
	"fmt"
)

func GenerateUUID() string  {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		panic(err)
	}
	return strings.TrimRight(string(out),"\r\n");
}

func Log(uuid string, msg string) {
	fmt.Println("[" + uuid + "] " + msg);
}
