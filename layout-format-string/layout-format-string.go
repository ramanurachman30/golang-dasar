package main

import "fmt"

func main() {
	type student struct {
		name       string
		height     float64
		age        int32
		isGraduate bool
		hobbies    []string
	}

	var data = student{
		name:       "Muhammad Nurachman Musyaffa",
		height:     170,
		age:        20,
		isGraduate: true,
		hobbies:    []string{"Futsal", "Football", "Badminton", "Coding"},
	}

	// Layout Format %b
	fmt.Printf("%b\n", data.age)
	// 10100

	// Layout Format %c
	fmt.Printf("%c\n", 1400)
	// n
	fmt.Printf("%c\n", 1235)
	// ӓ

	// Layout Format %d
	fmt.Printf("%d\n", data.age)
	// 20

	// Layout Format %e dan %E
	fmt.Printf("%e\n", data.height)
	// 1.700000e+02
	fmt.Printf("%E\n", data.height)
	// 1.700000E+02

	// Layout Format %f dan %F
	fmt.Printf("%f\n", data.height)
	// 170.000000
	fmt.Printf("%.9F\n", data.height)
	// 170.000000000
	fmt.Printf("%.2f\n", data.height)
	// 170.00
	fmt.Printf("%.f\n", data.height)
	// 170
}
