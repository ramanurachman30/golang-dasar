package main

import "fmt"

type murid struct {
	nama  string
	kelas int
}

func main() {
	var s1 = murid{}
	s1.nama = "Nurachman"
	s1.kelas = 20

	var s2 = murid{"nurachman", 1}

	var s3 = murid{nama: "zaki"}
	var s4 = murid{nama: "ucok", kelas: 4}
	var s5 = murid{kelas: 4, nama: "aep"}

	fmt.Println("Nama : ", s1.nama)
	fmt.Println("Nama : ", s2.nama)
	fmt.Println("Nama : ", s3.nama)
}
