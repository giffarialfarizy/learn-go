package main

import "fmt"

func main() {
	var name1 = "Giffari"
	var name2 = "Giffari"

	var result bool = (name1 == name2)
	fmt.Println(result)

	var name3 = "giffari"

	// Tidak harus integer, string pun bisa dibandingkan 
	result = (name3 > name2)
	fmt.Println(result)
}

// go build 11_operasi-perbandingan
// go run 11_operasi-perbandingan

// OPERATOR PERBANDINGAN
// >, >=
// <, <=
// ==
// !=