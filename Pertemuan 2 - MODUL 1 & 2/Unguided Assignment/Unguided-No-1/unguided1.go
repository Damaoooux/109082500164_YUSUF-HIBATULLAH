package main

import "fmt"

func main() {
	var tahun int
	var kabisat bool
	fmt.Print("Masukan tahun: ")
	fmt.Scan(&tahun)

	if (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0) {
		kabisat = true
	} else {
		kabisat = false
	}

	fmt.Print("kabisat: ", kabisat)

}
