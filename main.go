package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("notify-send", "test")
	err := cmd.Run()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
}
