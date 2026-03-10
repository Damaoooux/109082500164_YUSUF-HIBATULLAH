package main

import "fmt"

func main() {
	var berat int
	var biayaKG, biayaGR int
	fmt.Print("Masukan Berat Parcel (gram): ")
	fmt.Scan(&berat)

	kg := berat / 1000
	sisa := berat % 1000
	biayaKG = kg * 10000

	if sisa >= 500 {
		biayaGR = sisa * 5
	} else {
		biayaGR = sisa * 15
	}

	if kg > 10 {
		biayaGR = 0
	}

	fmt.Println("---Rincian---")
	fmt.Print("Detail Berat: ", kg, "kg + ", sisa, "gram\n")
	fmt.Print("Detail Biaya: ", biayaKG, " + ", biayaGR, "\n")
	fmt.Print("Biaya pengiriman: ", biayaKG+biayaGR)
}
