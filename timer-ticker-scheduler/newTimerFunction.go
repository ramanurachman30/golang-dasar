package main

import (
	"fmt"
	"time"
)

func main() {
	var timer = time.NewTimer(time.Second * 4)
	fmt.Println("Starting Timer...")
	<-timer.C
	fmt.Println("Ended Timer...")
}
