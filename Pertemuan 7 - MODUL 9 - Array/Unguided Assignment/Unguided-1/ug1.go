package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	cx, cy int
	r      int
}

func jarak(p, q Titik) float64 {
	return math.Sqrt(math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2))
}

func didalam(c Lingkaran, p Titik) bool {
	pusatLingkaran := Titik{c.cx, c.cy}
	jarakDariPusat := jarak(p, pusatLingkaran)
	return jarakDariPusat < float64(c.r)
}

func main() {
	var lingkaran1, lingkaran2 Lingkaran
	var titikSembarang Titik

	fmt.Scanln(&lingkaran1.cx, &lingkaran1.cy, &lingkaran1.r)

	fmt.Scanln(&lingkaran2.cx, &lingkaran2.cy, &lingkaran2.r)

	fmt.Scanln(&titikSembarang.x, &titikSembarang.y)

	diDalamL1 := didalam(lingkaran1, titikSembarang)
	diDalamL2 := didalam(lingkaran2, titikSembarang)

	if diDalamL1 && diDalamL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalamL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalamL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
