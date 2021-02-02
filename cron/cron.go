package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	c.AddFunc("*/1 * * * * *", func() { fmt.Println("1 second!") })
	c.AddFunc("0 */1 * * * *", func() { fmt.Println("1 minues!") })
	c.Start()

	for {
		time.Sleep(10000000000000)
		fmt.Println("sleep")
	}
}
