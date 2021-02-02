package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() {
		fnA()
	}()
	if err := fnB(false); err != nil {
		log.Println(err)
		return
	}
}

func fnA() {
	fmt.Println("A")
}

func fnB(b bool) error {
	fmt.Println("b:", b)
	if b {
		return nil
	}
	return fmt.Errorf("error")
}
