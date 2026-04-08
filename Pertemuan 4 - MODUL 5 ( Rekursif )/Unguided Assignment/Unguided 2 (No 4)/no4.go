package main

import "fmt"

func turun(n int) {
	if n == 1 {
		return
	}
	fmt.Print(n, " ")
	turun(n - 1)
}

func naik(n, current int) {
	if current > n {
		return
	}
	fmt.Print(current, " ")
	naik(n, current+1)
}

func barisan(n int) {
	if n == 0 {
		return
	}

	turun(n)
	naik(n, 1)
	fmt.Println()
}

func main() {
	var n int
	fmt.Scan(&n)
	barisan(n)
}
