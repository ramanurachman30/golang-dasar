package main

import "fmt"

func main() {
	var fruits = [5]string{"apel", "mangga", "durian", "semangka", "melon"}

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Fruits %d : %s\n", i, fruits[i])
	}
}
