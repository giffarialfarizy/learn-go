package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func main() {
	var giffari Customer
	giffari.Name = "Giffari Alfarizy"
	giffari.Address = "Tasikmalaya"
	giffari.Age = 28

	fmt.Println(giffari)

	// Struct Literals
	// Recommended
	evanti := Customer {
		Name: "Evanti Arosyani",
		Address: "Gununghalu",
		Age: 27,
	}
	fmt.Println(evanti)

	// Not recommended, harus sesuai urutan.
	// Akan error jika jumlah field bertambah
	eshan := Customer{"Mushlih Eshan", "Bandung", 1}
	fmt.Println(eshan)
}

// template data yang digunakan untuk menggabungkan nol
// ...atau lebih tipe data dalam satu kesatuan

// Struct biasanya representasi data dalam program
// ...aplikasi yang kita buat

// Data di struct disimpan dalam field

// Sederhananya struct adalah kumpulan dari field

// Struct tidak bisa digunakan langsung karena dia adalah template data