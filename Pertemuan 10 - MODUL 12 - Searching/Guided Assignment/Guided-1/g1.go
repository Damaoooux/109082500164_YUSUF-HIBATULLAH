package main

import "fmt"

func SequentialSearch(arrBuah [5]string, dataDicari string) int {
	var idx_found = -1
	for i := 0; i < len(arrBuah); i++ {
		if arrBuah[i] == dataDicari {
			idx_found = i
			break
		}
	}

	return idx_found
}

func main() {
	var arrBuah [5]string
	for i := 0; i < len(arrBuah); i++ {
		fmt.Printf("Masukan data buah indeks ke-%d : ", i)
		fmt.Scan(&arrBuah[i])
	}

	var dataDicari string
	fmt.Print("Masukan data buah yang ingin dicari : ")
	fmt.Scan(&dataDicari)

	var index_data int
	index_data = SequentialSearch(arrBuah, dataDicari)

	if index_data > -1 {
		fmt.Printf("Data %s ditemukan pada indeks ke-%d\n", dataDicari, index_data)
	} else {
		fmt.Printf("Data %s tidak ditemukan\n", dataDicari)
	}
}
