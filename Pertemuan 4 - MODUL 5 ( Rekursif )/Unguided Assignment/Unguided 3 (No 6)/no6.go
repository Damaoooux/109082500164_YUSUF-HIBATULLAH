package main

import "fmt"

func pangkat(x, n int) int {
	if n == 0 {
		return 1
	}
	return x * pangkat(x, n-1)
}

func main() {
	var x, n int
	fmt.Scanln(&x, &n)
	hasil := pangkat(x, n)
	fmt.Println(hasil)
}
