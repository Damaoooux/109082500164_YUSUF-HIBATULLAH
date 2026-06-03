package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		ganjil := []int{}
		genap := []int{}

		for j := 0; j < m; j++ {
			var x int
			fmt.Scan(&x)

			if x%2 == 1 {
				ganjil = append(ganjil, x)
			} else {
				genap = append(genap, x)
			}
		}

		for j := 0; j < len(ganjil)-1; j++ {
			min := j

			for k := j + 1; k < len(ganjil); k++ {
				if ganjil[k] < ganjil[min] {
					min = k
				}
			}

			ganjil[j], ganjil[min] = ganjil[min], ganjil[j]
		}

		for j := 0; j < len(genap)-1; j++ {
			max := j

			for k := j + 1; k < len(genap); k++ {
				if genap[k] > genap[max] {
					max = k
				}
			}

			genap[j], genap[max] = genap[max], genap[j]
		}

		for j := 0; j < len(ganjil); j++ {
			fmt.Print(ganjil[j], " ")
		}

		for j := 0; j < len(genap); j++ {
			fmt.Print(genap[j])

			if j != len(genap)-1 {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}
