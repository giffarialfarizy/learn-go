package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	
	// Gak akan muat, akan mulai dari batas bawahnya dia (-32768)
	fmt.Println(nilai16)
	fmt.Println(nilai8)

	// Konversi Karakter
	// Awalnya character
	var name = "giffari"
	// Pas dipanggil indeksnya keluarnya ASCII
	var g byte = name[0]
	// Mengubah dari ASCII ke string
	var gString string = string(g)

	fmt.Println(name)
	fmt.Println(g)
	fmt.Println(gString)
}

// go build 08_konversi-tipe-data.go
// go run 08_konversi-tipe-data.go