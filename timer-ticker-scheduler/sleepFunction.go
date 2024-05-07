package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting...")
	time.Sleep(time.Second * 4)
	fmt.Println("Done")

	// scheduler

	for true {
		fmt.Println("Starting...")
		time.Sleep(1 * time.Second)
	}
}
