package main

import "fmt"

func main() {
	var point = 1

	switch {
	case point == 8:
		fmt.Println("perfect")
	case point > 5 && point < 8: //tanda kurung () pada kondisi if else di go optional boss
		fmt.Println("Awesome")
	default:
		{
			fmt.Println("Not Bad")
			fmt.Println("Nice TRY")
		}
	}
}
