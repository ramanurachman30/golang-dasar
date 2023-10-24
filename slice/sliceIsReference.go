package main

import "fmt"

func main() {
	var fruits = []string{"melon", "apel", "durian", "semangka", "manggis"}

	var aFruits = fruits[0:4]
	var bFruits = fruits[1:3]

	var aaFruits = aFruits[1:3]
	var bbFruits = bFruits[0:1]

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(bbFruits)

	aaFruits[1] = "Mangga"

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(bbFruits)
}
