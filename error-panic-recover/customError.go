package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("GA BOLEH KOSONG!")
	}
	return true, nil
}

// recover
func catch() {
	if r := recover(); r != nil {
		fmt.Println("Terjadi Error! ", r)
	} else {
		fmt.Println("Aplikasi berjalan lancar")
	}
}

func main() {
	// defer catch()
	// pemanfaatan recover pada IIFE
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi Error nih : ", r)
		} else {
			fmt.Println("Ga ada Error ah")
		}
	}()
	var input string
	fmt.Print("Input sesuatu : ")
	fmt.Scanln(&input)

	if valid, err := validate(input); valid {
		fmt.Println("Nih ada ", input)
	} else {
		// PANIC
		panic(err.Error())
	}

	// handle error tanpa menghentikan looping
	data := []string{"superman", "aquaman", "wonder woman"}
	for _, each := range data {
		func() {
			// recover untuk IIFE dalam perulangan
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Terjadi Error nih : ", each, " | Message : ", r)
				} else {
					fmt.Println("Ga ada Error ah")
				}
			}()
		}()
	}
	panic("Ada error cenah bros")
}
