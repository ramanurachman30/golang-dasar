package main

import "fmt"

func main() {
	var chicken = map[string]int{"Januari": 40, "Februari": 30}
	var value, isExist = chicken["Januari"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Value is not found")
	}
}
