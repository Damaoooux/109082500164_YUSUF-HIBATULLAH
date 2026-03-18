package main

import "fmt"

func hitungBiayaMotor(jamMasuk, jamKeluar int) int {
	totalBiaya := 0

	for jam := jamMasuk; jam < jamKeluar; jam++ {
		if jam >= 17 && jam < 24 {
			totalBiaya += 5000
		} else {
			totalBiaya += 4000
		}
	}

	return totalBiaya
}

func hitungBiayaMobil(jamMasuk, jamKeluar int) int {
	totalBiaya := 0

	for jam := jamMasuk; jam < jamKeluar; jam++ {
		if jam >= 17 && jam < 24 {
			totalBiaya += 7000
		} else {
			totalBiaya += 6000
		}
	}

	return totalBiaya
}

func hitungBiaya(jenisKendaraan string, jamMasuk, jamKeluar int) int {
	if jenisKendaraan == "motor" {
		return hitungBiayaMotor(jamMasuk, jamKeluar)
	} else if jenisKendaraan == "mobil" {
		return hitungBiayaMobil(jamMasuk, jamKeluar)
	}
	return 0
}

func main() {
	totalPendapatan := 0
	nomorKendaraan := 1

	fmt.Println("==== Rekap Tarif Parkir Cafe Per Hari ====")

	for {
		fmt.Printf("\n*Kendaraan %d\n", nomorKendaraan)

		var jenisKendaraan string
		fmt.Print("Kendaraan apa? (motor / mobil / - untuk selesai): ")
		fmt.Scan(&jenisKendaraan)

		if jenisKendaraan == "-" {
			break
		}

		var jamMasuk, jamKeluar int
		fmt.Print("Masukkan Jam Masuk Kendaraan (0-24): ")
		fmt.Scan(&jamMasuk)
		fmt.Print("Masukkan Jam Keluar Kendaraan (0-24): ")
		fmt.Scan(&jamKeluar)

		biaya := hitungBiaya(jenisKendaraan, jamMasuk, jamKeluar)
		totalPendapatan += biaya

		fmt.Printf("Biaya parkir %s %d: %d\n", jenisKendaraan, nomorKendaraan, biaya)
		fmt.Println("========================================")
		nomorKendaraan++
	}

	fmt.Printf("\n*** Total Pendapatan Parkir Hari Ini Adalah %d ***\n", totalPendapatan)
}
