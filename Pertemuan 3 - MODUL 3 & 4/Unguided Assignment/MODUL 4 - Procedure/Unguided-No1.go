package main

import "fmt"

func deret(n int) {
	for n > 1 {
		fmt.Print(n, " ")
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Print(n)
}

func main() {
	var n int
	fmt.Print("Masukan angka: ")
	fmt.Scan(&n)
	deret(n)
}
