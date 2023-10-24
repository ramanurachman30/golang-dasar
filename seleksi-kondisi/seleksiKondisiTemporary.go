package main

import "fmt"

func main() {
	var point = 112233.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s Perfect !\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f% not bad\n", percent, "%")
	}
}
