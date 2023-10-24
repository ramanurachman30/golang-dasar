package main

import "fmt"

func main() {
	var chicken = map[string]int{"Januari": 40, "Februari": 30}

	fmt.Println(len(chicken))
	fmt.Println(chicken)

	delete(chicken, "Januari")

	fmt.Println(len(chicken))
	fmt.Println(chicken)
}
