package main

import (
	f "fmt" //alias
	. "golang-level-akses/library"
)

func main() {
	// var s1 = Student{"nurachman", 21}
	f.Println("Hello ", Student.Nama)
	f.Println("Umurku ", Student.Umur)
	SayHello("Nurachman")
	// library.introduce("nurachman")
	sayHelloPartial("nurachman")
}
