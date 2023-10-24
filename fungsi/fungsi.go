package main

import (
	"fmt"
	"strings"
)

func main() {
	var names = []string{"Nurachman", "Musyaffa"}
	sendMessage("Hello ", names)
}

func sendMessage(message string, arr []string) {
	var kirimMessage = strings.Join(arr, " ")
	fmt.Println(message, kirimMessage)
}
