package library

import "fmt"

func SayHello(name string) {
	fmt.Println("Hello World")
	introduce(name)
}

var Student = struct {
	Nama string
	Umur int
}{}

func introduce(name string) {
	fmt.Println("ucok", name)
}

func init() {
	Student.Nama = "Nurachman"
	Student.Umur = 21

	fmt.Println("imported")
}
