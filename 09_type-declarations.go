package main

import "fmt"

func main() {
	type NoKTP string
	type isMarried bool

	var ktpGiffari NoKTP = "3273072806930004"
	fmt.Println(ktpGiffari)

	var marritalStatus isMarried = true
	fmt.Println(marritalStatus)
}

// go build 09_type-declarations.go
// go run 09_type-declarations.go

// TYPE DECLARATIONS
// - adalah kemampuan membuat ulang tipe data baru dari tipe
// ..data yang sudah ada
// - biasanya digunakan untuk membuat alias agar lebih mudah
// ..dimengerti