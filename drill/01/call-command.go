package main

import (
	"log"
	"os/exec"
)

func main() {
	log.Print("This is program for executing command!\n")
//	out, error := exec.Command("false").Output()
	out, error := exec.Command("ls", "-la").Output()

	if error != nil {
		log.Printf("error: %s", error)
        return
	}
	log. Printf("Now: %s", out)
}
