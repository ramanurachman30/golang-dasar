package main

import "fmt"

func main() {
	var point = 1

	switch point {
	case 8:
		fmt.Println("Perfect")
	case 5, 6, 7:
		fmt.Println("Awesome")
	default:
		fmt.Println("bim bim bam bam")
	}
}
