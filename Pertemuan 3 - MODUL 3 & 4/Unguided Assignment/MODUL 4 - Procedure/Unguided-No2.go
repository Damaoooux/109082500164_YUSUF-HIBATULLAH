package main

import "fmt"

func hitungPersegi(sisi int) {
	var luas, keliling int
	luas = sisi * sisi
	keliling = 4 * sisi
	fmt.Println("Luas Persegi: ", luas)
	fmt.Println("Keliling Persegi: ", keliling)
}

func hitungPersegiPanjang(panjang int, lebar int) {
	var luas, keliling int
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	fmt.Println("Luas Persegi Panjang: ", luas)
	fmt.Println("Keliling Persegi Panjang: ", keliling)
}

func hitungLingkaran(jariJari float64) {
	var luas, keliling float64
	luas = 3.14 * jariJari * jariJari
	keliling = 2 * 3.14 * jariJari
	fmt.Println("Luas Lingkaran: ", luas)
	fmt.Println("Keliling Lingkaran: ", keliling)
}

func main() {
	var pilihan int

	fmt.Println("##### Program Bangun Datar #####")
	fmt.Println("1. Luas dan Keliling Persegi")
	fmt.Println("2. Luas dan Keliling Persegi Panjang")
	fmt.Println("3. Luas dan Keliling Lingkaran")
	fmt.Print("Pilihan : ")
	fmt.Scanln(&pilihan)
	fmt.Print("\n")

	switch pilihan {
	case 1:
		var sisi int
		fmt.Print("Masukan sisi: ")
		fmt.Scanln(&sisi)
		hitungPersegi(sisi)
	case 2:
		var panjang, lebar int
		fmt.Print("Masukan panjang: ")
		fmt.Scanln(&panjang)
		fmt.Print("Masukan lebar: ")
		fmt.Scanln(&lebar)
		hitungPersegiPanjang(panjang, lebar)
	case 3:
		var jariJari float64
		fmt.Print("Masukan jari-jari: ")
		fmt.Scanln(&jariJari)
		hitungLingkaran(jariJari)
	default:
		fmt.Println("Pilihan tidak valid!")
	}
}
