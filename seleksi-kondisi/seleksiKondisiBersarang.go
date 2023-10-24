package main

import "fmt"

func main() {
	var point = 5

	if point > 7 {
		switch point {
		case 10:
			fmt.Println("good job kids")
		default:
			fmt.Println("nice")
		}
	} else {
		if point == 5 {
			fmt.Println("Nice TRY Broh")
		} else if point == 3 {
			fmt.Println("Nice TRY for more kids")
		} else {
			fmt.Println("Ah nt Terus")
			if point == 0 {
				fmt.Println("Try yang sesungguhnyaaa!")
			}
		}
	}
}
