package main

import "fmt"

func main() {
	var fruits = []string{
		"apel",
		"semangka",
		"durian",
		"melon",
	}

	var aFruits = append(fruits, "mangga")

	fmt.Println(fruits)
	fmt.Println(aFruits)
}
