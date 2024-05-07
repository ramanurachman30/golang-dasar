package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	ticker := time.NewTicker(time.Second)

	go func() {
		time.Sleep(10 * time.Second)

		done <- true
	}()

	for {
		select {
		case <-done:
			ticker.Stop()
			return
		case c := <-ticker.C:
			fmt.Println("HEllooooo!!!! ", c)
		}
	}
}
