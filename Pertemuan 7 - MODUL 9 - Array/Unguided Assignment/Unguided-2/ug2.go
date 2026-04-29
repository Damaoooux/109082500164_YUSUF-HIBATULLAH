package main

import (
	"fmt"
	"math"
)

const NMAX int = 100

type TabInt struct {
	tab [NMAX]int
	n   int
}

func tampilkanArray(arr TabInt) {
	fmt.Print("Isi array: ")
	for i := 0; i < arr.n; i++ {
		fmt.Print(arr.tab[i], " ")
	}
	fmt.Println()
}
func tampilkanIndeksGanjil(arr TabInt) {
	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 1; i < arr.n; i += 2 {
		fmt.Print(arr.tab[i], " ")
	}
	fmt.Println()
}

func tampilkanIndeksGenap(arr TabInt) {
	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < arr.n; i += 2 {
		fmt.Print(arr.tab[i], " ")
	}
	fmt.Println()
}

func tampilkanKelipatanX(arr TabInt, x int) {
	fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
	for i := 0; i < arr.n; i++ {
		if i%x == 0 {
			fmt.Print(arr.tab[i], " ")
		}
	}
	fmt.Println()
}

func hapusElemen(arr *TabInt, indeks int) {
	if indeks >= 0 && indeks < arr.n {
		for i := indeks; i < arr.n-1; i++ {
			arr.tab[i] = arr.tab[i+1]
		}
		arr.n--
		fmt.Println("Elemen berhasil dihapus")
	} else {
		fmt.Println("Indeks tidak valid")
	}
}

func rataRata(arr TabInt) float64 {
	if arr.n == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < arr.n; i++ {
		sum += arr.tab[i]
	}
	return float64(sum) / float64(arr.n)
}

func standarDeviasi(arr TabInt) float64 {
	if arr.n == 0 {
		return 0
	}
	rata := rataRata(arr)
	var variance float64 = 0
	for i := 0; i < arr.n; i++ {
		variance += math.Pow(float64(arr.tab[i])-rata, 2)
	}
	variance /= float64(arr.n)
	return math.Sqrt(variance)
}

func frekuensi(arr TabInt, bilangan int) int {
	count := 0
	for i := 0; i < arr.n; i++ {
		if arr.tab[i] == bilangan {
			count++
		}
	}
	return count
}

func main() {
	var arr TabInt
	var pilihan int

	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&arr.n)

	if arr.n > NMAX {
		fmt.Println("Jumlah elemen melebihi kapasitas maksimum")
		return
	}

	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < arr.n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&arr.tab[i])
	}

	for {
		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Tampilkan keseluruhan isi array")
		fmt.Println("2. Tampilkan elemen dengan indeks ganjil")
		fmt.Println("3. Tampilkan elemen dengan indeks genap")
		fmt.Println("4. Tampilkan elemen dengan indeks kelipatan x")
		fmt.Println("5. Hapus elemen pada indeks tertentu")
		fmt.Println("6. Tampilkan rata-rata")
		fmt.Println("7. Tampilkan standar deviasi")
		fmt.Println("8. Tampilkan frekuensi bilangan tertentu")
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tampilkanArray(arr)
		case 2:
			tampilkanIndeksGanjil(arr)
		case 3:
			tampilkanIndeksGenap(arr)
		case 4:
			var x int
			fmt.Print("Masukkan nilai x: ")
			fmt.Scan(&x)
			if x != 0 {
				tampilkanKelipatanX(arr, x)
			} else {
				fmt.Println("x tidak boleh 0")
			}
		case 5:
			var indeks int
			fmt.Print("Masukkan indeks yang akan dihapus: ")
			fmt.Scan(&indeks)
			hapusElemen(&arr, indeks)
		case 6:
			fmt.Printf("Rata-rata: %.2f\n", rataRata(arr))
		case 7:
			fmt.Printf("Standar deviasi: %.2f\n", standarDeviasi(arr))
		case 8:
			var bilangan int
			fmt.Print("Masukkan bilangan yang dicari: ")
			fmt.Scan(&bilangan)
			fmt.Printf("Frekuensi bilangan %d: %d\n", bilangan, frekuensi(arr, bilangan))
		case 0:
			fmt.Println("Program selesai")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
