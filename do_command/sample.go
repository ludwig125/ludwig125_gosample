package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("ls", "-la").Output()
	if err != nil {
		log.Println(err)
	}
	log.Println(string(out))
}
