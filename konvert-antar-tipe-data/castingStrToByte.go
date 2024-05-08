package main

import "fmt"

func main() {
	// konversi menggunakan casting dari string ke byte

	var text = "Halloo"
	var b = []byte(text)

	fmt.Printf("%d %d %d %d \n", b[0], b[1], b[2], b[3])

	// konversi menggunakan castring dari byte ke string

	var b1 = []byte{72, 97, 108, 108}
	var text1 = string(b1)

	fmt.Println(text1)
}
