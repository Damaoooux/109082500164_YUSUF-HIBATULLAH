package main

import "fmt"

func hitungtotal(harga, jumlah int) {
	var total int
	total = harga * jumlah
	fmt.Println("Total Harga: ", total)

	hitungDiskon(total)
}

func hitungDiskon(total int) {
	var diskon, hargaAkhir int
	diskon = total * 10 / 100
	hargaAkhir = total - diskon

	fmt.Println("Diskon 10%: ", diskon)
	fmt.Println("Harga Akhir: ", hargaAkhir)
}

func main() {
	var price, qty int
	fmt.Print("Masukan harga barang: ")
	fmt.Scan(&price)
	fmt.Print("Masukan jumlah barang: ")
	fmt.Scan(&qty)

	hitungtotal(price, qty)
}
