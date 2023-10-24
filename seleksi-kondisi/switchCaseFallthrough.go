package main

import "fmt"

func main() {
	var point = 8

	switch {
	case point == 8:
		fmt.Println("Perfect")
	case point > 5 && point < 8:
		fmt.Println("Awesome")
		fallthrough
	case point < 5:
		fmt.Println("BELAJAR LAGI DEKDEK")
	default:
		{
			fmt.Println("Good Game")
			fmt.Println("Nice TRY")
		}
	}
}
