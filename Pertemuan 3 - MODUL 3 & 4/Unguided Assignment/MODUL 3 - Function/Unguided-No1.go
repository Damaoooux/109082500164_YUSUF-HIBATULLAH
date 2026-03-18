package main

import "fmt"

func factorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutasi(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func kombinasi(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Print("Masukan nilai A: ")
	fmt.Scan(&a)
	fmt.Print("Masukan nilai B: ")
	fmt.Scan(&b)
	fmt.Print("Masukan nilai C: ")
	fmt.Scan(&c)
	fmt.Print("Masukan nilai D: ")
	fmt.Scan(&d)

	fmt.Println("Hasil Permutasi A terhadap C:", permutasi(a, c))
	fmt.Println("Hasil Kombinasi A terhadap C:", kombinasi(a, c))
	fmt.Println("Hasil Permutasi B terhadap D:", permutasi(b, d))
	fmt.Println("Hasil Kombinasi B terhadap D:", kombinasi(b, d))
}
