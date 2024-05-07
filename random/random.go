package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randomizer := rand.New(rand.NewSource(20))

	fmt.Println("Angka Rando ke 1 ", randomizer.Int())
	fmt.Println("Angka Rando ke 2 ", randomizer.Int())
	fmt.Println("Angka Rando ke 3 ", randomizer.Int())

	// unix number

	randomizers := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	fmt.Println("Angka Random ke 1 ", randomizers.Int())
	fmt.Println("Angka Random ke 2 ", randomizers.Float32())
	fmt.Println("Angka Random ke 3 ", randomizers.Uint32())
	fmt.Println("Angka Random ke 4 ", randomizers.Intn(100))

	fmt.Println("String Random ", randomString(5))
}

func randomString(lenght int) string {
	// Random Tipe data String
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var randomizerss = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	b := make([]rune, lenght)
	for i := range b {
		b[i] = letters[randomizerss.Intn(len(letters))]
	}
	return string(b)
}
