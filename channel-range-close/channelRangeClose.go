package main

import (
	"fmt"
	"runtime"
)

func sendMessage(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
	close(ch)
}

func printMessage(ch <-chan string) {
	for messages := range ch {
		fmt.Println(messages)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)
	go sendMessage(messages)
	printMessage(messages)
}
