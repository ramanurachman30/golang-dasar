package main

import "fmt"

func main() {
	var fruits = [3]string{"apel", "mangga", "durian"}

	for _, fruit := range fruits {
		fmt.Printf("Fruits : %s\n", fruit)
	}
}
