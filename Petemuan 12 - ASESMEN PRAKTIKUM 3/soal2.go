package main

import "fmt"

type player struct {
	name   string
	gol    int
	assist int
}

type arrPlayer [1000]player

func main() {
	var T arrPlayer
	var n int
	var nama_depan, nama_belakang string

	fmt.Print("Masukan Jumlah Data: ")
	fmt.Scan(&n)
	for i := 0; i < n && i < 1000; i++ {
		fmt.Scan(&nama_depan, &nama_belakang, &T[i].gol, &T[i].assist)
		T[i].name = nama_depan + " " + nama_belakang
	}

	for pass := 1; pass <= n-1; pass++ {
		idx_max := pass - 1
		for i := pass; i < n; i++ {
			if T[idx_max].gol < T[i].gol || (T[idx_max].gol == T[i].gol && T[idx_max].assist <= T[i].assist) {
				idx_max = i
			}
		}

		temp := T[idx_max]
		T[idx_max] = T[pass-1]
		T[pass-1] = temp
	}

	fmt.Println("\n", "Hasil sorting: ")
	for i := 0; i < n; i++ {
		fmt.Println(T[i].name, T[i].gol, T[i].assist)
	}
}
