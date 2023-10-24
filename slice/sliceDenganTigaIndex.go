package main

import "fmt"

func main() {
	var fruits = []string{
		"semangka",
		"melon",
		"durian",
		"apel",
	}

	var aFruits = fruits[0:3]
	var bFruits = fruits[0:3:3]

	fmt.Println(fruits)
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))

	fmt.Println(aFruits)
	fmt.Println(len(aFruits))
	fmt.Println(cap(aFruits))

	fmt.Println(bFruits)
	fmt.Println(len(bFruits))
	fmt.Println(cap(bFruits))
}
