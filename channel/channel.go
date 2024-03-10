package main

import (
	"fmt"
	"runtime"
)

func printMessage(what chan string) {
	fmt.Println(<-what)
}

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("Hallo %s", who)
		messages <- data
	}

	for _, each := range []string{"rama", "ucok", "jaki"} {
		go func(who string) {
			var data = fmt.Sprintf("Hallo %s", who)
			messages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}

	go sayHelloTo("ucup kupling")
	go sayHelloTo("Saeput cahya")
	go sayHelloTo("Asep cahyo")

	var message1 = <-messages
	fmt.Println(message1)
	var message2 = <-messages
	fmt.Println(message2)
	var message3 = <-messages
	fmt.Println(message3)
}
