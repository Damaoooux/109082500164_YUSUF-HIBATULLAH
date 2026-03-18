package main

import "fmt"

func tambahvalue(x int) {
	x = x + 10
	fmt.Println("Nilai x di dalam Prosedur tambahvalue (pass by value): ", x)
}

func tambahreference(x *int) {
	*x = *x + 10
	fmt.Println("Nilai x di dalam Prosedur tambahreference (pass by reference): ", *x)
}

func main() {
	var y int = 5
	fmt.Println("Nilai awal: ", y)

	tambahvalue(y)
	fmt.Println("Nilai setelah pass by value: ", y)

	tambahreference(&y)
	fmt.Println("Nilai setelah pass by reference: ", y)
}
