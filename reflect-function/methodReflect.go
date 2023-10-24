package main

import (
	"fmt"
	"reflect"
)

type students struct {
	nama  string
	kelas int
}

func (s *students) SetName(name string) {
	s.nama = name
}

func main() {
	var s1 = &students{nama: "Nurachman", kelas: 12}
	fmt.Println("nama : ", s1.nama)

	var reflectValue = reflect.ValueOf(s1)
	var method = reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("asek"),
	})

	fmt.Println("nama : ", s1.nama)
}
