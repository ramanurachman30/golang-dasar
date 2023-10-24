package main

type person3 struct {
	nama  string
	kelas string
}

type student1 struct {
	person3 struct {
		jurusan string
	}

	nama2 string
	umur  []int
}
