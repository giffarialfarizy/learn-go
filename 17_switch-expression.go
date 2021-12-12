package main

import "fmt"

func main() {
	name := "Giffari"
	name = "Evan"

	switch name {
	case "Giffari":
		fmt.Println("Hello Giffari!")
	case "Evan":
		fmt.Println("Hai Evan!")
	default:
		fmt.Println("Maaf kami tidak mengenal nama Anda.")
	}

	// Switch dengan Short Statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama terlalu pendek")
	}

	// Switch Tanpa Kondisi
	length:= len(name)
	switch {
	case length > 5:
		fmt.Println("Nama terlalu panjang")
	case length > 3:
		fmt.Println("Nama agak panjang")
	default:
		fmt.Println("Nama terlalu pendek")
	}
}

// go build 17_switch-expression.go
// go run 17_switch-expression.go

// Pengecekan yg lebih sederhana dibanding if else