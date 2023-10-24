package main

import "fmt"

func main() {
	devineNumber(10, 2)
	devineNumber(4, 0)
	devineNumber(8, -4)
}

func devineNumber(m, n int) {
	if n == 0 {
		fmt.Printf("Invalid Divider. %d cannot divide by %d\n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}
