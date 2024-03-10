package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Hellooowww")
	os.Exit(1)
	fmt.Println("Hai jugaaa")
}
