package main

import "fmt"

func main() {
	for i := 0; i < 6; i++ {
		for j := i; j < 6; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
