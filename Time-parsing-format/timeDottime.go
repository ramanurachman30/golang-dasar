package main

import (
	"fmt"
	"time"
)

func main() {
	var time1 = time.Now()
	fmt.Printf("Time nya %v\n", time1)

	var time2 = time.Date(2024, 03, 12, 03, 16, 0, 0, time.UTC)
	fmt.Printf("Time 2 %v\n", time2)

	var now = time.Now()
	fmt.Println("Year : ", now.Year(), " Month : ", now.Month(), " Day : ", now.Day(), now.Weekday(), now.Local(), now.Format("DD-MM-YYYY HH:mm:ss"), now.Nanosecond(), now.Location())
}
