package main

import "fmt"

func main() {
	const firstName string = "Giffari"
	const lastName = "Alfarizy"
	const value = 1000

	fmt.Println(firstName)

	// Akan error
	// firstName = "Evanti"
	// lastName = "Arosyani"

	// Deklarasi multiple const
	const (
		namaDepan string	= "Evanti"
		namaBelakang		= "Arosyani"
		umur				= 27
	)

	fmt.Println(namaBelakang)
}

// go build 07_constant.go
// go run 07_constant.go

// const harus dideklarasikan bersama valuenya
// Jika tidak digunakan pun tidak akan "diprotes" sama Go nya
// Tidak bisa pakai := karena akan jadi variable