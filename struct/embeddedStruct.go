package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	grade int
	person
}

func main() {
	var s1 = student{}
	s1.name = "Nurachman"
	s1.age = 20
	s1.grade = 9

	fmt.Println("Student s1 Name \t: ", s1.name)
	fmt.Println("Student s1 age \t\t: ", s1.age)
	fmt.Println("Student s1 age \t\t: ", s1.person.age)
	fmt.Println("Student s1 Grade \t: ", s1.grade)
}
