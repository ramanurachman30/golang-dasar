package main

import "fmt"

func main() {
	type buah2 struct {
		nama  string
		jenis string
	}

	type buahh = buah2

	var b1 = buah2{"Durian", "Musangking"}
	fmt.Println(b1)

	var b2 = buahh{"Durian", "duri hitam"}
	fmt.Println(b2)

	buah3 := buah2{"Durian", "belanda"}
	fmt.Println(buah2(buah3))

	buah4 := buah2{"Durian", "merah"}
	fmt.Println(buah2(buah4))

	type person6 struct {
		nama string
		umur int
	}

	type person7 = struct {
		nama string
		umur string
	}

	type Number = int
	var num Number = 12
	fmt.Println(num)
}
