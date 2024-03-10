package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Print("Input sebuah nomer : ")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)
	if err == nil {
		fmt.Println(number, "Is number")
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())
	}
}
