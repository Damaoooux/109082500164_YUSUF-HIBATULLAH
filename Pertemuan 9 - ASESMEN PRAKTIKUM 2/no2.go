package main

import "fmt"

const nMax int = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func inputMahasiswa(T *arrayMahasiswa, N *int) {
	var i int = 0
	fmt.Print("Masukkan jumlah Data: ")
	fmt.Scan(N)

	for i < *N {
		fmt.Print("Masukkan Data ke-", i+1, " : ")
		fmt.Scan(&T[i].NIM, &T[i].nama, &T[i].nilai)
		i++
	}
}

func first(T arrayMahasiswa, N int, NIM string) int {
	var i int = 0
	var ketemu bool = false

	for i < N && !ketemu {
		ketemu = T[i].NIM == NIM
		if !ketemu {
			i++
		}
	}

	if ketemu {
		return i
	}
	return -1
}

func MAX(T arrayMahasiswa, N int, NIM string) int {
	var ketemu = first(T, N, NIM)
	var i int
	var IDXMAX int

	if ketemu != -1 {
		IDXMAX = ketemu
		for i = ketemu + 1; i < N; i++ {
			if T[i].NIM == NIM && T[i].nilai > T[IDXMAX].nilai {
				IDXMAX = i
			}
		}
		return IDXMAX
	} else {
		return ketemu
	}

}

func main() {
	var A arrayMahasiswa
	var M, IDX1, IDX2 int
	var NIM string

	inputMahasiswa(&A, &M)

	fmt.Print("Masukkan NIM yang ingin dicari: ")
	fmt.Scan(&NIM)

	IDX1 = first(A, M, NIM)
	IDX2 = MAX(A, M, NIM)

	fmt.Println("Nilai pertama dari NIM", NIM, "adalah", A[IDX1].nilai)
	fmt.Println("Nilai terbesar dari NIM", NIM, "adalah", A[IDX2].nilai)
}
