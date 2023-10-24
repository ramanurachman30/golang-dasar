package main

import "fmt"

func main() {
	var fruits = []string{"apel", "durian", "semangka"}            // slice
	var fruitsB = [4]string{"apel", "durian", "semangka", "melon"} //array
	var fruitsC = [...]string{"apel", "durian"}                    //array

	var newFruits = fruits[0:2]

	fmt.Println("Fruits : ", fruits[0])
	fmt.Println("Fruits : ", fruitsB)
	fmt.Println("Fruits : ", fruitsC)
	fmt.Println("Fruits : ", newFruits)
}
