package main

import "fmt"

type Person struct {
	Name string
}

// Ini inti course 35
type HasName interface {
	GetName() string
}

func (person Person) GetName() string {
	return person.Name
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func main() {
	// Tidak bisa begini
	// var giffari HasName
	// giffari.GetName()

	var giffari Person
	giffari.Name = "Giffari"

	SayHello(giffari)

}

// interface digunakan sebagai kontrak
// interface berisikan definisi-definisi method

// tida perlu mengimplementasikan interface secara manual