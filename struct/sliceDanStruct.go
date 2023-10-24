package main

import "fmt"

type person1 struct {
	nama  string
	age   int
	major string
}

func main() {
	var s1 = []person1{
		{"Nurachman", 20, "Rekayasa Perangkat Lunak"},
		{"Usep", 24, "Rekayasa Perangkat Lunak"},
		{"Aan", 10, "Rekayasa Perangkat Lunak"},
	}

	var allStudent = []struct {
		person1
		grade int
	}{
		{person1: person1{"Nurachman", 12, "TKJ"}, grade: 5},
		{person1: person1{"Nurachman Musyaffa", 15, "TKJ"}, grade: 5},
		{person1: person1{"Muhammad Nurachman Musyaffa", 17, "TKJ"}, grade: 5},
	}

	for _, students := range allStudent {
		fmt.Println(students)
	}

	for _, student := range s1 {
		fmt.Println(student.nama, student.age, student.major)
	}
}
