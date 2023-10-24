package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	person
	age  int
	name string
}

func main() {
	var s1 = student{}
	s1.name = "Nurachman"
	s1.age = 20
	s1.person.age = 25
	var p1 = person{name: "nurachman", age: 18}
	var s2 = student{person: p1, age: 30} //struct

	fmt.Println(s2.person.name)
	fmt.Println(s1.name)
	fmt.Println(s1.age)
	fmt.Println(s1.person.age)
}
