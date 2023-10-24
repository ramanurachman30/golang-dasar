package main

import "fmt"

func main() {
	var chickens = []map[string]string{
		map[string]string{"Name": "Muhammad Nurachman Musyaffa", "Gender": "Male"},
		map[string]string{"Name": "Dida", "Gender": "Female"},
		{"Name": "Zaki", "Gender": "Male"},
		{"Name": "Adiba", "Gender": "Female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["Name"], chicken["Gender"])
	}
}
