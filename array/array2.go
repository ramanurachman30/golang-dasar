package main

import "fmt"

func main() {
	//cara horizontal
	var fruits = [4]string{"apel", "banana", "nangka", "durian"}

	//cara vertikal
	fruits = [4]string{
		"apel",
		"banana",
		"nangka",
		"durian",
	}

	fmt.Println("Jumlah buah adalah \t\t", len(fruits))
	fmt.Println("buah nya yaitu \t", fruits)
}
