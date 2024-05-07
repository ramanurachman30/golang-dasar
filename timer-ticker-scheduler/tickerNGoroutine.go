package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var timeout = 5
	var ch = make(chan bool)

	go timer(timeout, ch)
	go waither(timeout, ch)

	var input string
	fmt.Println("Ada Berapa Juz pada Al-Qur'an ? ")
	fmt.Scan(&input)

	if input == "30" {
		fmt.Println("Wow Jawaban Kamu benar ")
	} else {
		fmt.Println("Maaf Jawaban Kamu salah Geningan")
	}
}

func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func waither(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("Anda Telalu lama menjawab dan melebih waktu yang ditentukan yaitu: ", timeout, " Detik")
	os.Exit(0)
}
