package main

import (
	"log"
    "fmt"
	"os/exec"
)

func main() {
	log.Print("This is program for executing command!\n")
    command := exec.Command("sleep", "2")
    error := command.Start()

	if error != nil {
        panic(fmt.Sprintf("error1: %v", error)) // when the command does not exist
	}
	log. Printf("Start command")

    error = command.Wait()

    if error != nil {
        panic(fmt.Sprintf("error2: %v", error)) // when the command finish with error status   
    }

    log.Print("Finish command")
}
