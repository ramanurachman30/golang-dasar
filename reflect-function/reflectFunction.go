package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = "nurachman"
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe variable nya adalah :", reflectValue.Type())

	if reflectValue.Kind() == reflect.String { //cek tipe data menggunakan fungsi Kind()
		fmt.Println("number nya adalah :", reflectValue.String())
	}

	//mengakses nilai dalam bentuk interface{}

	var numbers = "nurachman"
	var reflectValue2 = reflect.ValueOf(numbers)
	// var wowcuy = reflectValue2.Interface().(int)

	fmt.Println("tipe variable :", reflectValue2.Type())
	fmt.Println("tipe variable :", reflectValue2.Interface())

	// mengakses informasi property variable objek

	var s1 = &student{Name: "Nurachman", Grade: 19}
	s1.getPropertyInfo()
}

// mengakses informasi property variable objek

type student struct {
	Name  string
	Grade int
}

func (s *student) getPropertyInfo() {
	var reflectValue3 = reflect.ValueOf(s)

	if reflectValue3.Kind() == reflect.Ptr {
		reflectValue3 = reflectValue3.Elem()
	}

	var reflectType = reflectValue3.Type()

	for i := 0; i < reflectType.NumField(); i++ {
		fmt.Println("Name : ", reflectType.Field(i).Name)
		fmt.Println("Tipe data : ", reflectType.Field(i).Type)
		fmt.Println("Nilai : ", reflectValue3.Field(i).Interface())
		fmt.Println(" ")
	}
}
