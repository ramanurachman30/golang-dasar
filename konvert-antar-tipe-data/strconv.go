package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str = "123"
	var num, err = strconv.Atoi(str) //method atoi from string to int

	var num2 = 124
	var str2 = strconv.Itoa(num2) // method Itoa from int to string

	var str3 = "125"
	var num3, err2 = strconv.ParseInt(str3, 10, 64) // method Parse in akan mengconversi string ke numerik dengan baris nya 10 dan tipe data int64(ada di parameter ketiga)

	var str4 = "1010"
	var num4, err3 = strconv.ParseInt(str4, 2, 8)

	var num5 = int64(24)
	var str5 = strconv.FormatInt(num5, 10) // Berguna untuk konversi data numerik int64 ke string dengan basis numerik bisa ditentukan sendiri.

	var str6 = "24.12"
	var num6, err4 = strconv.ParseFloat(str6, 32)

	// Pada contoh di atas, string "24.12" dikonversi ke float dengan lebar tipe data float32. Hasil konversi strconv.ParseFloat adalah sesuai dengan standar IEEE Standard for Floating-Point Arithmetic.

	var num7 = float64(24.12)
	var str7 = strconv.FormatFloat(num7, 'f', 7, 64) //Berguna untuk konversi data bertipe float64 ke string dengan format eksponen, lebar digit desimal, dan lebar tipe data bisa ditentukan.

	var str8 = "true"
	var bul, err5 = strconv.ParseBool(str8) // konversi string ke boolean

	var bul2 = true
	var str9 = strconv.FormatBool(bul2) // konversi boolean ke string

	fmt.Println(str9)

	if err5 == nil {
		fmt.Println(bul)
	}

	fmt.Println(str7)

	if err4 == nil {
		fmt.Println(num6)
	}

	fmt.Println(str5)

	if err2 == nil {
		fmt.Println(num3) //125
	}

	if err3 == nil {
		fmt.Println(num4)
	}

	fmt.Println(str2) // "123"

	if err == nil {
		fmt.Println(num) //123
	}
}
