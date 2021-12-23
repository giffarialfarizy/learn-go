package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == ""{
		return nil
	} else {
		return map[string]string {
			"name": name,
		}
	}
}

func main() {
	var person map[string] string = nil
	fmt.Println(person)

	animal := NewMap("Cat")
	fmt.Println(animal)

	// Pengecekan menggunakan nil
	data := NewMap("")
	if data == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(data)
	}

	// Pengecekan tanpa nil
	var data2 map[string]string

	if data2["name"] == "" {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(data2)
	}
}

// Nil hanya bisa digunakan di
// ...1. interface,
// ...2. function,
// ...3. map,
// ...4. slice,
// ...5. pointer,
// ...6. dan channel

// Nil buat apa? Pengecekan