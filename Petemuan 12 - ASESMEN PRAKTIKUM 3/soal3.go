package main

import "fmt"

const NMAX = 1000000

type partai struct {
	nama  int
	suara int
}

type tabPartai [NMAX]partai

func posisi(t tabPartai, n int, nama int) int {
	var ketemu int
	var i int
	ketemu = -1
	for i = 0; i < n && ketemu == -1; i++ {
		if t[i].nama == nama {
			ketemu = i
		}
	}
	return ketemu
}

func main() {
	var p tabPartai
	var n int = 0
	var idx int
	var x int

	fmt.Print("Masukkan Proses Input Suara: ")
	fmt.Scan(&x)
	for x != -1 {
		idx = posisi(p, n, x)
		if idx == -1 {
			p[n].nama = x
			p[n].suara = 1
			n++
		} else {
			p[idx].suara++
		}
		fmt.Scan(&x)
	}

	var k int
	var pass int
	var temp partai

	for pass = 1; pass <= n-1; pass++ {
		k = pass
		temp = p[k]
		for k > 0 && temp.suara > p[k-1].suara {
			p[k] = p[k-1]
			k--
		}
		p[k] = temp
	}

	fmt.Print("\n", "Hasil Perhitungan Suara: ")
	fmt.Println()

	for k = 0; k < n; k++ {
		fmt.Printf("%v(%v) ", p[k].nama, p[k].suara)
	}
}
