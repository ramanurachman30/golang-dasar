package main

import "fmt"

func main() {
	var secret interface{}

	secret = "Nurachman Musyaffa"
	fmt.Println("Hello : ", secret)

	secret = []string{"Apel", "durian", "buah"}
	fmt.Println("buah : ", secret)

	secret = 12
	fmt.Println("Number : ", secret)

	var data map[string]interface{}

	data = map[string]interface{}{
		"name":    "Nurachman",
		"grade":   12,
		"man":     true,
		"hobbies": []string{"Football", "futsal", "badminton"},
	}

	fmt.Println("My datas is : ", data)

	var datas map[string]any //use any
	datas = map[string]any{
		"name":    "Nurachman",
		"grade":   12,
		"man":     true,
		"hobbies": []string{"Football", "futsal", "badminton"},
	}

	fmt.Println(datas)
}
