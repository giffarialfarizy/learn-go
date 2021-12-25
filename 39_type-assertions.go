package main

import "fmt"

func random() interface {} {
	return 100
}

func main() {
	result := random();
	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int)	// panic
	// fmt.Println(resultInt)

	// Main aman
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}

/*	
	Type Assertions merupakan kemampuan merubah tipe data
	menjadi tipe data yang diinginkan
*/

/*
	Fitur ini sering sekali digunakan ketika kita bertemu
	dengan data interface kosong
*/

/*
	Agar lebih aman gunakan switch expression untuk type
	assertions
*/