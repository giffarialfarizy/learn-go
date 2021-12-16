package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	// index ignored
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println("Variadic Function")
	total := sumAll(10, 10, 10, 10, 10)
	fmt.Println("sumAll:", total)

	// Jika terlanjur memiliki slice
	slice := []int{10, 20, 30}
	total = sumAll(slice...)
	fmt.Println("Slice Sum:", total)
}

// ciri-cirinya adalah di function:
// ...1. di posisi akhir
// ...2. ada titik 3

// variadic function:
// ...1. bisa menerima lebih dari 1 input, anggap saja semacam Array
// ...2. Bukan tipe data array biasa, karena Array biasa harus membuat Array terlebih dahulu
// ...3. Memiliki kemampuan dijadikan sebuah varargs