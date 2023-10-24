package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var s1 = struct { //anonymous struct tanpa pengisian property
		person
		grade int
	}{}

	var s2 = struct { //anonymous struct dengan pengisian property
		person
		major string
	}{
		person: person{"Nurachman", 2},
		major:  "Rekayasa Perangkat Lunak",
	}

	s1.person = person{name: "Nurachman", age: 20}
	s1.grade = 9

	fmt.Println("Student S1 Name : ", s1.person.name)
	fmt.Println("Student S1 Age \t: ", s1.person.age)
	fmt.Println("Student S1 grade : ", s1.grade)
}
