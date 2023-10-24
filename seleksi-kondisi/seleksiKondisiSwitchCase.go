package main

import "fmt"

func main() {
	var point = 7

	switch point {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("good")
	default:
		fmt.Println("not bad")
	}
}
