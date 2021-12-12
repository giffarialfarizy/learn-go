package main

import "fmt"

func main() {
	name := "Giffari"
	// name = "Evanti"

	if name == "Giffari" {
		fmt.Println("Hello, Giffari")
	} else if name =="Evanti" {
		fmt.Println("Hello, Evanti")
	} else {
		fmt.Println("Anda tidak terautentikasi?")
	}

	// If Short Statement
	if length := len(name); length > 5 {
		fmt.Println("Nama Anda lebih dari 5 karakter")
		} else {
		fmt.Println("Nama Anda KURANG dari 5 karakter")
			
	}
}

// go build 16_if-expression.go
// go run 16_if-expression.go

// IF SHORT STATEMENT
// if length := len(name); length > 5 {...}
// length := len(name) adalah if short statement
// length > 5 adalah kondisinya