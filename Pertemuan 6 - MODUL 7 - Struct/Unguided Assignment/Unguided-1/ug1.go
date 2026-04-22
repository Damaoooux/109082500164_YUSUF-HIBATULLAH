package main

import "fmt"

type suhu float64

func CelsiusToReamur(celcius suhu) suhu {
	return celcius * 4 / 5
}

func CelsiusToFahrenheit(celcius suhu) suhu {
	return (celcius * 9 / 5) + 32
}

func CelsiusToKelvin(celcius suhu) suhu {
	return celcius + 273.15
}

func main() {
	var celcius suhu

	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukkan suhu (celcius): ")
	fmt.Scan(&celcius)

	r := CelsiusToReamur(celcius)
	f := CelsiusToFahrenheit(celcius)
	k := CelsiusToKelvin(celcius)

	fmt.Printf("\n%.0f celcius = %.1f reamur\n", celcius, r)
	fmt.Printf("%.0f celcius = %.1f fahrenheit\n", celcius, f)
	fmt.Printf("%.0f celcius = %.2f kelvin\n", celcius, k)
}
