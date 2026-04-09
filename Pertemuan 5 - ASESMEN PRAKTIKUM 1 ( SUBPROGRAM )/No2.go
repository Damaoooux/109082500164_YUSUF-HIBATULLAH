package main

import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	if tujuan == "domestik" && jumlahHari > 3 {
		jumlahHari = 3
	} else if tujuan == "mancanegara" && jumlahHari > 8 {
		jumlahHari = 8
	}
	return jumlahHari
}

func biayaPerHari(jumlahMhs int) int {
	return jumlahMhs * (70000 + 250000 + 300000)
}

func perhitunganBiaya(jumlahMhs, lamaPerjalanan int, tujuan string, totalBiaya *float64) {
	var lama int = tanggunganHari(lamaPerjalanan, tujuan)
	*totalBiaya = float64(biayaPerHari(jumlahMhs))
	*totalBiaya = *totalBiaya * float64(lama)

	if tujuan == "mancanegara" {
		*totalBiaya = 1.5 * *totalBiaya
	}
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64

	fmt.Print("Masukan Jumlah Mahasiswa: ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukan Lama hari study tour: ")
	fmt.Scan(&lama)
	fmt.Print("Masukan Tujuan (domestik/mancanegara): ")
	fmt.Scan(&tujuan)

	if tujuan != "domestik" && tujuan != "mancanegara" {
		fmt.Println("Tujuan tidak valid. Harap masukkan 'domestik' atau 'mancanegara'.")
		return
	}

	fmt.Print("\n")

	perhitunganBiaya(jumlah, lama, tujuan, &biaya)
	fmt.Print("Biaya perjalanan yang harus dikeluarkan Tel-U : ")
	fmt.Printf("Rp %.0f", biaya)
}
