package main

import (
	f "fmt"
	s "strings"
)

func main() {
	var secrets interface{}
	// teknik type assertions
	secrets = 2
	var number = secrets.(int) * 10
	f.Println(secrets, "multiplied by 10 is : ", number)

	secrets = []string{"apple", "durian", "manggo"}
	var gruits = s.Join(secrets.([]string), " , ")
	f.Println(gruits, " is my favorite fruits")

	// objek pointer di interface kosong

	type person struct {
		name string
		age  int
	}

	var secret2 interface{} = &person{name: "Nurachman Ganteng", age: 17}
	var nama = secret2.(*person).name
	f.Println(nama)

	// kombinasi slice, map dan interface{}

	var person2 = []map[string]interface{}{
		{"name": "Nurachman", "age": 12},
		{"name": "Nurachman", "age": 12},
		{"name": "Nurachman", "age": 12},
	}

	for _, each := range person2 {
		f.Println(each["name"], "Age Is", each["age"])
	}

	var fruits2 = []interface{}{
		map[string]interface{}{"nama": "durian", "total": 10},
		[]string{"durian", "buah"},
		"mangga",
	}

	for _, buah := range fruits2 {
		f.Println(buah)
	}
}
