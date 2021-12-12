package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{...}
	person := map[string]string{
		"name":    "Giffari",
		"address": "Bandung",
	}

	fmt.Println(person)

	person["title"] = "Programmer"

	fmt.Println(person["title"])

	// FUNCTION MAP
	book := make(map[string]string)
	book["title"] = "Go-Lang"
	book["author"] = "John Doe"
	book["wrong"] = "Oops"

	delete(book, "wrong")

	fmt.Println(book)
	fmt.Println(len(book))
}

// go build 15_tipe-data-map.go
// go run 15_tipe-data-map.go

// map adalah tipe data kumpulan key-value, kata kuncinya unik
// jumlah data yang dimasukkan ke dalam Map boleh sebanyak-banyaknya

// FUNCTION MAP
// len(map)						: untuk mendapatkan jumlah data di map
// map[key]						: Mengambil data di map dengan key
// map[key] = value				: Mengubah data di map dengan key
// make(map[TypeKey]TypeValue)	: Membuat map baru
// delelte(map, key)			: Menghapus data di map dengan key