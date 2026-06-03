package main

import "fmt"

type mahasiswa struct {
	name string
	nim  string
	ipk  float64
}

type arrMhs [100]mahasiswa

func InsertionSortStruct(T *arrMhs, n int) {
	var temp mahasiswa
	var j, i int

	for i = 1; i < n; i++ {
		temp = T[i]
		j = i - 1

		for j >= 0 && T[j].name > temp.name {
			T[j+1] = T[j]
			j--
		}

		T[j+1] = temp
	}
}

func main() {
	var data arrMhs
	var n, i int

	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Println("Masukkan data mahasiswa ke-", i+1)

		fmt.Print("Nama: ")
		fmt.Scan(&data[i].name)

		fmt.Print("NIM: ")
		fmt.Scan(&data[i].nim)

		fmt.Print("IPK: ")
		fmt.Scan(&data[i].ipk)
	}

	fmt.Println("\nData mahasiswa sebelum diurutkan:")
	for i = 0; i < n; i++ {
		fmt.Println(data[i].name, data[i].nim, data[i].ipk)
	}

	InsertionSortStruct(&data, n)

	fmt.Println("\nData mahasiswa setelah diurutkan:")
	for i = 0; i < n; i++ {
		fmt.Println(data[i].name, data[i].nim, data[i].ipk)
	}
}
