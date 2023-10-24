package main

import "fmt"

func main() {
	var fruits = [5]string{"apel", "mangga", "semangka", "melon", "durian"}

	for i, fruit := range fruits {
		fmt.Printf("Fruits %d : %s\n", i, fruit)
	}
}
