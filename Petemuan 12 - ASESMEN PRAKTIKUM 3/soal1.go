package main

import "fmt"

const NMAX = 1000000

type arrInt [NMAX]int

func sorting(T *arrInt, n int) {
	var temp int
	var pass int
	var idx_min int

	for pass = 1; pass <= n-1; pass++ {
		idx_min = pass - 1
		for i := pass; i <= n-1; i++ {
			if T[idx_min] > T[i] {
				idx_min = i
			}
		}

		temp = T[idx_min]
		T[idx_min] = T[pass-1]
		T[pass-1] = temp
	}
}

func median(T *arrInt, n int) float64 {
	var mid int
	mid = n / 2

	if n%2 == 0 {
		return float64(T[mid-1]+T[mid]) / 2.0
	} else {
		return float64(T[mid])
	}
}

func main() {
	var x int
	var n int = 0
	var A arrInt

	fmt.Scan(&x)
	for x != -5313541 && n < NMAX {
		if x == 0 {
			sorting(&A, n)
			fmt.Println("Median: ", median(&A, n))
		} else {
			A[n] = x
			n++
		}
		fmt.Scan(&x)
	}
}
