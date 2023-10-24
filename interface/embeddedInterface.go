package main

import (
	"fmt"
	"math"
)

type hitung2d interface {
	keliling() float64
	luas() float64
}

type hitung3d interface {
	volume() float64
}

type hitungs interface {
	hitung2d
	hitung3d
}

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

func main() {
	var bangunRuang hitungs = &kubus{12.0}

	fmt.Println("Volume Kubus   :", bangunRuang.volume())
	fmt.Println("Luas Kubus     :", bangunRuang.luas())
	fmt.Println("Keliling kubus :", bangunRuang.keliling())
}
