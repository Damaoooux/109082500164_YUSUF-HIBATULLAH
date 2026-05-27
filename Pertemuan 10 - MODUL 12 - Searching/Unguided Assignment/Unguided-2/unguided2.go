package main

import "fmt"

const MAX = 21

func main() {
	var suara [MAX]int
	var x int
	var totalData int
	var suaraSah int

	for {
		fmt.Scan(&x)

		if x == 0 {
			break
		}

		totalData++
		if x >= 1 && x <= 20 {
			suara[x]++
			suaraSah++
		}
	}

	var ketua int = 1
	var wakil int = 1

	for i := 2; i <= 20; i++ {
		if suara[i] > suara[ketua] {
			wakil = ketua
			ketua = i

		} else if suara[i] == suara[ketua] && i < ketua {
			wakil = ketua
			ketua = i

		} else if i != ketua {
			if suara[i] > suara[wakil] || wakil == ketua {
				wakil = i
			} else if suara[i] == suara[wakil] && i < wakil {
				wakil = i
			}
		}
	}

	fmt.Println("Suara masuk:", totalData)
	fmt.Println("Suara sah:", suaraSah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}
