package main

type person2 struct {
	name string
	age  int
}

func main() {
	var student struct { //hanya deklarasi
		person2
		major string
	}

	var student2 = struct { //deklarasi sekaligus inisiasi
		grade int
	}{
		10,
	}
}
