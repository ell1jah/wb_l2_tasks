package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
)

const (
	TimeOnly       = "15:04:05"
	TimeExactMilli = "15:04:05.000"
)

func main() {
	ntptime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatal(err) // пакет log по умолчанию пишет в stderr, fatal завершает программу вызвав os.Exit(1)
	}
	fmt.Println("current time:", ntptime.Format(TimeOnly))
	fmt.Println("exact time:", ntptime.Format(TimeExactMilli))
}
