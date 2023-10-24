package main

import "fmt"

func main() {
	var point = 3

	if point == 10 {
		fmt.Print("Lulus dengan point sempurna")
	} else if point > 5 {
		fmt.Print("Lulus")
	} else if point == 4 {
		fmt.Print("Tidak lulus")
	} else {
		fmt.Printf("Tidak lulus. nilai anda %d\n ", point)
	}
}
