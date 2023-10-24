package main

import "fmt"

type murid struct {
	nama  string
	kelas int
}

func main() {
	var s1 = murid{nama: "nurachman", kelas: 5}

	var s2 *murid = &s1
	fmt.Println("student 1, Name : ", s1.nama)
	fmt.Println("Student 4, Name : ", s2.nama)

	s2.nama = "ucok"

	fmt.Println("Student 1, nama : ", s1.nama)
	fmt.Println("Student 4, nama : ", s2.nama)
}
