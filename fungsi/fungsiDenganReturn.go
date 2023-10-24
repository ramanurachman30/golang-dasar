package main

import (
	"fmt"
	"math/rand"
	"time"
)

var randomize = rand.New(rand.NewSource(time.Now().Unix()))

func main() {
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("Random 1 : ", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("Random 2 : ", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("Random 3 : ", randomValue)

}

func randomWithRange(min, max int) int {
	var value = randomize.Int()%(max-min+1) + min
	return value
}
