package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("starting after functions...")

	<-time.After(4 * time.Second)

	fmt.Println("Ended after functions...")
}
