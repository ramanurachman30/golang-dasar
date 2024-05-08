package main

import (
	"fmt"
	"time"
)

// Pada chapter ini kita akan belajar tentang tipe data untuk pengolahan durasi waktu yaitu time.Duration.

// Tipe time.Duration ini merepresentasikan durasi, contohnya seperti 1 menit, 2 jam 5 detik, dst. Data dengan tipe ini bisa dihasilkan dari operasi pencarian delta atau selisih dari dua buah objek time.Time, atau bisa juga kita buat sendiri.

// Tipe ini sangat berguna untuk banyak hal, salah satunya untuk benchmarking ataupun operasi-operasi lainnya yang membutuhkan informasi durasi waktu.

func main() {
	start := time.Now()

	time.Sleep(5 * time.Second)

	duration := time.Since(start)

	fmt.Println("Waktu per Detik ", duration.Seconds())
	fmt.Println("Waktu per Menit ", duration.Minutes())
	fmt.Println("Waktu per Jam ", duration.Hours())

	// Kalkulasi Durasi Antara 2 Objek Waktu

	t1 := time.Now()
	time.Sleep(5 * time.Second)
	t2 := time.Now()

	duration2 := t2.Sub(t1)

	fmt.Println("Waktu per detiks ", duration2.Seconds())
	fmt.Println("Waktu per menits ", duration2.Minutes())
	fmt.Println("Waktu per Jams ", duration2.Hours())

	// Konversi Angka ke time.Duration

	// 12 * time.Minute             // 12 menit
	// 65 * time.Hour                 // 65 jam
	// 150000 * time.Milisecond     // 150k milidetik atau 150 detik
	// 45 * time.Microsecond         // 45 microdetik
	// 233 * time.Nanosecond         // 233 nano detik

	// Tipe time.Duration diciptakan menggunakan tipe ìnt64. Jadi jika ingin mengalikan time.Duration dengan suatu angka, maka pastikan tipe-nya juga sama yaitu time.Duration. Jika angka tersebut tidak ditampung dalam variabel terlebih dahulu (contohnya seperti di atas) maka bisa langsung kalikan saja. Jika ditampung ke variabel terlebih dahulu, maka pastikan tipe variabelnya adalah time.Duration. Contoh:

	// var n time.Duration = 5
	// duration := n * time.Second

	// Atau bisa manfaatkan casting untuk mengkonversi data numerik ke tipe time.Duration. Contoh:

	// n := 5
	// duration := time.Duration(n) * time.Second
}
