package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh  nol")
	} else {
		hasil := nilai / pembagi
		return hasil, nil
	}
}

func main() {
	fmt.Println("error")
	var contohError error = errors.New("Ups Error")
	fmt.Println(contohError.Error())

	hasil, err := Pembagian(100, 10)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}

// Go-Lang sudah menyediakan library untuk membuat helper
// ...dengan mudah, yang terdapat di package errors