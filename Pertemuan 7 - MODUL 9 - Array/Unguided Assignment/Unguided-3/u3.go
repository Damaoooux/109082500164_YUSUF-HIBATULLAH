package main

import "fmt"

const NMAX int = 100

type Klub struct {
	nama string
	skor int
}

type TabKlub struct {
	data [NMAX]Klub
	n    int
}

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang TabKlub
	var pertandingan int

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	pemenang.n = 0
	pertandingan = 0

	// Input
	for {
		pertandingan++
		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			pertandingan--
			break
		}

		if skorA > skorB {
			pemenang.data[pemenang.n].nama = klubA
			pemenang.data[pemenang.n].skor = skorA
			pemenang.n++
		} else if skorB > skorA {
			pemenang.data[pemenang.n].nama = klubB
			pemenang.data[pemenang.n].skor = skorB
			pemenang.n++
		} else {
			// seri
			pemenang.data[pemenang.n].nama = "Draw"
			pemenang.data[pemenang.n].skor = skorA
			pemenang.n++
		}
	}

	fmt.Println()
	for i := 0; i < pemenang.n; i++ {
		fmt.Printf("Hasil %d : %s\n", i+1, pemenang.data[i].nama)
	}

	fmt.Println("Pertandingan selesai")
}
