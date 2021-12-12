package main

import "fmt"

func main() {
	counter :=  1

	// Kok mirip while ya?
	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	// For dengan Statement
	for i:=1 ; i<=10; i++ {
		fmt.Println("i =", i)
	}

	// Bentuk lain
	slice := []string{"Giffari", "Evanti", "Eshan"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// For Range
	// Mirip python
	// index dan name boleh diganti apa saja
	names := []string{"A", "B", "C", "D", "E", "F", "G"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	// jika tidak butuh index pakai _ saja
	for _, value := range names {
		fmt.Println(value)
	}

	// Pakai map
	person := make(map[string]string)
	person["name"] = "Giffari"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

	
}

// go build 18_for-loops.go
// go run 18_for-loops.go

// di GO hanya ada perulangan for. Tidak ada while maupun do-while

// FOR STATEMENT
// 1. Init Statement, dieksekusi 1x sebelum perulangan
// 2. Post Statement, selalu dieksekusi di akhir tiap perulangan