package main

import "fmt"

// return type nya string
func getHello(name string) string {
	return "Hello " + name
}

func main() {
	result := getHello("Giffari")
	fmt.Println(result)
	result = getHello("")
	fmt.Println(result)
	// getHello("Giffari") // error kalau langsung begini
}