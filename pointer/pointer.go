package main

import "fmt"

func main() {
	var number1 int = 4
	var number2 *int = &number1

	fmt.Println("number 1 (value) : ", number1)
	fmt.Println("Number 1 (address) : ", &number1)

	fmt.Println("number 2 (value) :", number2)
	fmt.Println("number 2 (address) :", *number2)
}
