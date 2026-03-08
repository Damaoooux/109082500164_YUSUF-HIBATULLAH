package main

import "fmt"

func main() {
	var angka int
	fmt.Println("---MENU---")
	fmt.Println("1. Cek angka == 10")
	fmt.Println("2. Cek genap atau ganjil")
	fmt.Print("Masukan pilihan: ")
	var pilihan int
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		fmt.Print("Masukan angka: ")
		fmt.Scan(&angka)
		if angka == 10 {
			fmt.Println("Angka adalah 10")
		} else {
			fmt.Println("Angka bukan 10")
		}

	case 2:
		fmt.Print("Masukan angka: ")
		fmt.Scan(&angka)
		if angka%2 == 0 {
			fmt.Println("Angka adalah genap")
		} else {
			fmt.Println("Angka adalah ganjil")
		}

	default:
		fmt.Println("Pilihan tidak valid")
	}
}
