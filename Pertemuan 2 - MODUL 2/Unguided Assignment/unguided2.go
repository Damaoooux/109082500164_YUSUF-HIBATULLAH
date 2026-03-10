package main

import "fmt"

func main() {
	var urutan [4]string = [4]string{"merah", "kuning", "hijau", "ungu"}
	var status bool = false
	var percobaan1 [4]string
	var percobaan2 [4]string
	var percobaan3 [4]string
	var percobaan4 [4]string
	var percobaan5 [4]string

	fmt.Print("Percobaan 1: ")
	for i := 0; i < len(urutan); i++ {
		fmt.Scan(&percobaan1[i])
	}
	fmt.Print("Percobaan 2: ")
	for i := 0; i < len(urutan); i++ {
		fmt.Scan(&percobaan2[i])
	}
	fmt.Print("Percobaan 3: ")
	for i := 0; i < len(urutan); i++ {
		fmt.Scan(&percobaan3[i])
	}
	fmt.Print("Percobaan 4: ")
	for i := 0; i < len(urutan); i++ {
		fmt.Scan(&percobaan4[i])
	}
	fmt.Print("Percobaan 5: ")
	for i := 0; i < len(urutan); i++ {
		fmt.Scan(&percobaan5[i])
	}

	if percobaan1 == urutan && percobaan2 == urutan && percobaan3 == urutan && percobaan4 == urutan && percobaan5 == urutan {
		status = true
	} else {
		status = false
	}

	fmt.Print("Berhasil?: ", status)
}
