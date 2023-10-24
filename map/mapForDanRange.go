package main

import "fmt"

func main() {
	var chicken = map[string]int{
		"februari": 50,
		"Januari":  40,
		"mei":      30,
		"juli":     20,
	}

	for key, val := range chicken {
		fmt.Println(key, " \t:", val)
	}
}
