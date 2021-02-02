package main

import (
	"os"
)

const fileName = "samplefile"

func createFile() error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
func writeFile(s string) error {

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	output := s
	file.Write(([]byte)(output))
	return nil
}
