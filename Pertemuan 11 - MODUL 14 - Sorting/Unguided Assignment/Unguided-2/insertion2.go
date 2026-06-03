package main

import "fmt"

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(
			&pustaka[i].id,
			&pustaka[i].judul,
			&pustaka[i].penulis,
			&pustaka[i].penerbit,
			&pustaka[i].eksemplar,
			&pustaka[i].tahun,
			&pustaka[i].rating,
		)
	}
}

func CetakFavorit(pustaka DaftarBuku, n int) {
	max := pustaka[0]

	for i := 1; i < n; i++ {
		if pustaka[i].rating > max.rating {
			max = pustaka[i]
		}
	}

	fmt.Println("Buku Favorit:")
	fmt.Println(max.judul, max.penulis, max.penerbit, max.tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var temp Buku
	var j int

	for i := 1; i < n; i++ {
		temp = pustaka[i]
		j = i - 1

		for j >= 0 && pustaka[j].rating < temp.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}

		pustaka[j+1] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Println("5 Buku Rating Tertinggi:")

	batas := 5
	if n < 5 {
		batas = n
	}

	for i := 0; i < batas; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	kiri := 0
	kanan := n - 1
	ketemu := false

	for kiri <= kanan && !ketemu {
		tengah := (kiri + kanan) / 2

		if pustaka[tengah].rating == r {
			fmt.Println("Data Buku Ditemukan:")
			fmt.Println(
				pustaka[tengah].judul,
				pustaka[tengah].penulis,
				pustaka[tengah].penerbit,
				pustaka[tengah].tahun,
				pustaka[tengah].eksemplar,
				pustaka[tengah].rating,
			)
			ketemu = true
		} else if r > pustaka[tengah].rating {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if !ketemu {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var pustaka DaftarBuku
	var n, ratingCari int
	fmt.Scan(&n)

	DaftarkanBuku(&pustaka, n)
	fmt.Scan(&ratingCari)
	CetakFavorit(pustaka, n)
	UrutBuku(&pustaka, n)
	Cetak5Terbaru(pustaka, n)
	CariBuku(pustaka, n, ratingCari)
}
