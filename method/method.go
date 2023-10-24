package main

import (
	"fmt"
	"strings"
)

type students struct {
	name  string
	grade int
}

func (s students) sayHello() {
	fmt.Println("Your name : ", s.name)
}

func (s students) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

func main() {
	var s1 = students{"Muhammad Nurachman Musyaffa", 12}
	s1.sayHello()

	var name = s1.getNameAt(3)
	fmt.Println("Name nya ada ", name)
}
