package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64 = 15
	var area, circumReference = calculate(diameter)

	fmt.Printf("luas Lingkaran\t\t : %.2f \n", area)
	fmt.Printf("Keliling Lingkaran\t : %.2f \n", circumReference)
}

func calculate(n float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(n/2, 2)

	// hitung keliling
	var circumReference = math.Pi * n

	// kembalikan nilai
	return area, circumReference
}

func calculate2(m float64) (area2 float64, circumReference float64) { //fungsi lebih sederhana
	area2 = math.Pi * math.Pow(m/2, 2)
	circumReference = math.Pi * m
	return
}
