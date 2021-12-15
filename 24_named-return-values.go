package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Mushlih"
	middleName = "Eshan"
	lastName = "Haazim"

	// Tidak perlu lagi menyebut ulang yang direturn
	// ... karena sudah dideklarasikan
	return
}

// func getCompleteName() (firstName string, middleName string, lastName string) {
// 	firstName = "Mushlih"
// 	middleName = "Eshan"
// 	lastName = "Haazim"

// 	return firstName, middleName, lastName
// }

func main() {
	namaDepan, namaTengah, namaBelakang := getCompleteName()
	fmt.Println(namaDepan, namaTengah, namaBelakang)
}