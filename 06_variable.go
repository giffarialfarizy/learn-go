package main

import "fmt"

func main() {
	var name string

	name = "Giffari Alfarizy"
	fmt.Println(name)

	var age = 38
	fmt.Println(age)

	// jika pakai := justru tidak boleh pakai vat
	// := hanya untuk deklarasi awal
	country := "Indonesia"
	fmt.Println(country)

	// Deklarasi Multiple Variable
	var (
		firstName = "Giffari"
		lastName = "Alfarizy"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}

// go build 06_variable.go
// ./06_variable

// go run 06_variable.go

// VARIABLE
// Hanya bisa menyimpan tipe data yang sama
// dimulai dengan var diikuti dengan nama variable dan tipe datanya
// var tidak wajib, kita bisa inisialisasi variabel namun dengan :=
// Tidak bisa ubah TIPE data di tengah jalan untuk variable yang sama pada program yang sama
// Bisa ubah VALUE data

// Pada Go jika variable tidak dipakai maka akan error 