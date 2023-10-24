package main

import "fmt"

func main() {
	var fruits = []string{
		"melon",
		"apel",
		"durian",
		"Semangka",
		"Mangga",
	}

	var aFruits = fruits[0:4]
	fmt.Println("Fruits : ", len(aFruits))
	fmt.Println("Fruits : ", cap(aFruits))

	var bFruits = fruits[1:5]
	fmt.Println("Fruits : ", len(bFruits))
	fmt.Println("Fruits : ", cap(bFruits))
}
