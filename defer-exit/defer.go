package main

import "fmt"

func orderSomeFood(menu string) {
	defer fmt.Println("Terima kasih telah memesan")
	if menu == "Pizza" {
		fmt.Print("PIZZA ANDA SEDANG KAMI BUAT ", " ")
		fmt.Print("PIZZA ANDA SUDAH JADI SILAHKAN AMBIL ", "\n")
		return
	}

	fmt.Println("Terima pesanan ", menu)
}

func main() {
	orderSomeFood("Pizza")
	orderSomeFood("Burger")

	// defer X IIFE
	number := 3

	if number == 3 {
		fmt.Println("Number 1")
		func() {
			fmt.Println("number 3")
		}()
	}

	fmt.Println("number 2")
}
