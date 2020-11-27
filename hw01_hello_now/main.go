package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	var pool string = "0.beevik-ntp.pool.ntp.org"
	ntpTime, err := ntp.Time(pool)

	if err != nil {
		log.Fatalf("FAILED connection to %s", pool)
	}

	fmt.Println("current time:", time.Now().Round(0))
	fmt.Println("exact time:", ntpTime.Round(0))
}
