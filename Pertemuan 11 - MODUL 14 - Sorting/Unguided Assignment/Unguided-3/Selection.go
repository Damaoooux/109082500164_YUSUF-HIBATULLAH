package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		arr := make([]int, m)

		// Input data
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}

		// Selection Sort
		for j := 0; j < m-1; j++ {
			min := j

			for k := j + 1; k < m; k++ {
				if arr[k] < arr[min] {
					min = k
				}
			}

			arr[j], arr[min] = arr[min], arr[j]
		}

		// Output hasil sorting
		for j := 0; j < m; j++ {
			fmt.Print(arr[j])

			if j != m-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
