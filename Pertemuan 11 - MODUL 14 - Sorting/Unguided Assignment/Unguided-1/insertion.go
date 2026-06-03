package main

import "fmt"

func main() {
	var arr []int
	var x int

	// Input data sampai bilangan negatif
	for {
		fmt.Scan(&x)

		if x < 0 {
			break
		}

		arr = append(arr, x)
	}

	// Insertion Sort
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}

	// Output array yang sudah diurutkan
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}

	fmt.Println()

	// Cek jarak antar data
	if len(arr) < 2 {
		fmt.Println("Data berjarak x")
		return
	}

	jarak := arr[1] - arr[0]
	tetap := true

	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != jarak {
			tetap = false
			break
		}
	}

	if tetap {
		fmt.Printf("Data berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
