package main

import "fmt"

func main() {
	// Type assertions merupakan teknik untuk mengambil tipe data konkret dari data yang terbungkus dalam interface{} atau any. Lebih jelasnya silakan cek contoh berikut.

	// Variabel data disiapkan bertipe map[string]interface{}, map tersebut berisikan beberapa item dengan tipe data value-nya berbeda satu sama lain, sementara tipe data untuk key-nya sama yaitu string.

	var data = map[string]interface{}{
		"name":    "Nurachman",
		"age":     20,
		"propesi": "tidur",
		"tinggi":  170.5,
		"berat":   60.5,
		"isMale":  true,
		"hobbies": []string{"Turu", "ngetik"},
	}

	fmt.Println(data["name"].(string))
	fmt.Println(data["age"].(int))
	fmt.Println(data["propesi"].(string))
	fmt.Println(data["tinggi"].(float64))
	fmt.Println(data["berat"].(float64))
	fmt.Println(data["isMale"].(bool))
	fmt.Println(data["hobbies"].([]string))

	// pengecekkan jika kita tidak tau tipe data asli
	for _, val := range data {
		switch val.(type) {
		case string:
			fmt.Println(val.(string))
		case int:
			fmt.Println(val.(int))
		case float64:
			fmt.Println(val.(float64))
		case bool:
			fmt.Println(val.(bool))
		case []string:
			fmt.Println(val.([]string))
		default:
			fmt.Println(val.(int))
		}
	}
}
