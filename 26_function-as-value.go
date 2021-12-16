package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	goodBye := getGoodBye
	fmt.Println(goodBye("Eko"))
}

// Function adalah first class citizen di go
// Function juga merupakan tipe data
// Function bisa disimpan di dalam variable