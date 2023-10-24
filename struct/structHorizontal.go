package main

type buah struct {
	nama   string
	jenis  string
	jumlah int
}

func main() {
	var b1 = struct {
		nama  string
		jenis string
	}{"Durian", "Musangking"} // horizontal
	var b2 = struct {
		nama  string
		jenis string
	}{nama: "Rambutan", jenis: "markonah"} // horizontal
}
