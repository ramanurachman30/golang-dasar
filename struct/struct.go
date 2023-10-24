package main

import "fmt"

type student struct {
	name  string
	grade int
}

func main() {
	var s1 student
	s1.name = "Nurachman"
	s1.grade = 19

	fmt.Println("Name : ", s1.name)
	fmt.Println("Grade : ", s1.grade)
}
