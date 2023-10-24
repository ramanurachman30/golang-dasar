package main

import "fmt"

func main() {
	var numbers1 = [3][3]int{[3]int{2, 3, 4}, [3]int{5, 6, 7}, [3]int{8, 9, 10}}
	var numbers2 = [3][3]int{{2, 3, 4}, {5, 6, 7}, {8, 9, 10}}

	fmt.Println(numbers1)
	fmt.Println(numbers2)
}
