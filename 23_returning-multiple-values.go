package main

import "fmt"

func getFullName() (string, string) {
	return "Giffari", "Alfarizy"
}

func tigaNama() (string, string, string) {
	return "Mushlih", "Eshan", "Haazim"
}

func main() {
	// firstName, lastName := getFullName()
	// fmt.Println(firstName, lastName)
	// firstName, middleName, lastName := tigaNama()
	// fmt.Println(firstName, middleName, lastName)
	firstName, _, lastName := tigaNama()
	fmt.Println(firstName, lastName)
	 
}