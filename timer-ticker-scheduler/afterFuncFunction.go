package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan bool)

	time.AfterFunc(4*time.Second, func() {
		fmt.Println("Kadaluwarsa")

		ch <- true
	})

	fmt.Println("starting Function...")
	<-ch
	fmt.Println("Finish the function")
}
