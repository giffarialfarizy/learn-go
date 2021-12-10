package main

import "fmt"

func main() {
	var family [3]string
	family[0] = "Giffari"
	family[1] = "Evanti"
	family[2] = "Eshan"
	// Error ketika melebihi daya tampung
	// family[3] = "Eshan"

	fmt.Println(family[0])
	fmt.Println(family[1])
	fmt.Println(family[2])

	// Membuat Array secara langsung
	var values = [3]int{
		90, 80, 95,
	}

	fmt.Println(values)

	// len itu panjang array, bukan jumlah data
	fmt.Println(len(values))
 
	var contoh [10]string
	fmt.Println(len(contoh))
}

// go build 13_tipe-data-array.go
// go run 13_tipe-data-array.go

// Array berisikan kumpulan data dengan tipe yang sama
// Saat membuat array, perlu ditentukan jumlah data yang bisa ditampung
// Data tampung array tidak bisa bertambah setelah array dibuat

// len(array)			: mendapatkan panjang array
// array[index]			: mendapat data di posisi index
// array[index] = value	: mengubah data di posisi index 