package main

import "fmt"

func main() {
	var data map[string]int

	data["one"] = 1
	// akan error

	data = map[string]int{}
	data["one"] = 1
	// tidak error

	// cara horizontal
	var chicken = map[string]int{"Januari": 50, "februari": 40}
	fmt.Println(chicken)

	// cara Vertikal
	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}
	fmt.Println(chicken2)

	var chicken3 = map[string]int{}
	fmt.Println(chicken3)
	var chicken4 = make(map[string]int)
	fmt.Println(chicken4)
	var chicken5 = *new(map[string]int)
	fmt.Println(chicken5)
}
