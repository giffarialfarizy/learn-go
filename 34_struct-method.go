package main

import "fmt"

type Customer struct {
	Name string
	Address string
	Age int
}

func (customer Customer) sayHi() {
	fmt.Println("Hi, my name is", customer.Name)
}

func main() {
	giffari := Customer {
		Name: "Giffari",
		Address: "Tasikmalaya",
		Age: 28,
	}
	giffari.sayHi()
}

// Digunakan spesifik untuk customer tertentu