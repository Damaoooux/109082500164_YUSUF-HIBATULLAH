package main

import "fmt"

const nProv int = 10

type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	var i int
	for i = 0; i < nProv; i++ {
		fmt.Print("Masukan data ke-", i+1, " : ")
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	var idx int = 0
	var i int
	for i = 1; i < nProv; i++ {
		if tumbuh[idx] < tumbuh[i] {
			idx = i
		}
	}
	return idx
}
func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	var i int
	var result float64
	for i = 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			result = (tumbuh[i] + 1) * float64(pop[i])
			fmt.Println(prov[i], result)
		}
	}
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	var found int = -1
	var i int = 0
	for i < nProv && found == -1 {
		if prov[i] == nama {
			found = i
		}
		i++
	}
	return found
}

func main() {
	var TProvinsi NamaProv
	var TPopulasi PopProv
	var TPertumbuhan TumbuhProv
	var cari string
	var idxTercepat, idxProvinsi int
	fmt.Println("========= Masukkan nama provinsi, populasi, dan pertumbuhan =========")
	InputData(&TProvinsi, &TPopulasi, &TPertumbuhan)

	fmt.Print("Masukan nama provinsi yang ingin dicari: ")
	fmt.Scan(&cari)

	idxTercepat = ProvinsiTercepat(TPertumbuhan)
	fmt.Print("Provinsi dengan pertumbuhan tercepat: ")
	fmt.Println(TProvinsi[idxTercepat])

	idxProvinsi = IndeksProvinsi(TProvinsi, cari)
	fmt.Println("Data provinsi", cari, ":")
	fmt.Println(TProvinsi[idxProvinsi])

	fmt.Println("========= Prediksi populasi untuk provinsi dengan pertumbuhan diatas 2% =========")
	Prediksi(TProvinsi, TPopulasi, TPertumbuhan)
}
