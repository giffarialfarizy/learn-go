package main

import "fmt"

func main() {
	var ujian = 90
	var presensi = 80

	var lulusUjian = ujian >= 80
	var lulusPresensi = presensi >= 80

	var lulus = lulusUjian && lulusPresensi
	fmt.Println(lulus)

}

// go build 12_operasi-boolean.go
// go run 12_operasi-boolean.go

// OPERASI BOOLEAN
// && (Dan, 2 boolean)
// || (Atau, 2 boolean)
// ! (Kebalikan, 1 boolean)