package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

func main() {
	start := time.Now()

	res := fmt.Sprintf("cpu: %d\n", runtime.NumCPU())
	res += fmt.Sprintf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))

	token := mustGetenv("SLACK_TOKEN")
	channel := mustGetenv("SLACK_CHANNEL")
	if err := sendSlackMsg(token, channel, res, start); err != nil {
		log.Println(err)
	}
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("%s environment variable not set.", k)
	}
	log.Printf("%s environment variable set.", k)
	return v
}
