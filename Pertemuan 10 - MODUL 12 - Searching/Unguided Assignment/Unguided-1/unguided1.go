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

	fmt.Println("Suara masuk:", totalData)
	fmt.Println("Suara sah:", suaraSah)
	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}
