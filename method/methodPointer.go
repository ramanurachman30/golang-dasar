package main

import "fmt"

type murid struct {
	nama  string
	kelas int
}

func (m murid) changeMurid1(name string) {
	fmt.Println("Change name ", name)
	m.nama = name
}

func (m *murid) changeMurid2(name string) {
	fmt.Println("Change Name 2 ", name)
	m.nama = name
}

func main() {
	var s1 = murid{"Nurachman Musyaffa", 18}
	s1.changeMurid1(s1.nama)
	fmt.Println("Nama mu adalah ", s1.nama)

	s1.changeMurid1("Nurachman ganteng")
	fmt.Println("Nama Anda yang asli adalah", s1.nama)

	s1.changeMurid2("Nurachman Keren cuk")
	fmt.Println("Nama Anda juga ini ", s1.nama)
}
