package main

import (
	"fmt"
	"strings"
)

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func main() {
	var data = []string{"wick", "jason", "ethan"}
	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	var dataContains5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("Ambil data : \t\t", data)

	fmt.Println("Ambil data yang berfilter \"o\"\t O: ", dataContainsO)

	fmt.Println("Ambil data yang jumlah huruf 5 \"5\"\t :", dataContains5)
}

type FilterCallback func(string) bool

func filter(data []string, callback FilterCallback) []string{
	//
}
