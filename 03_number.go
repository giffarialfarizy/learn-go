package main

import "fmt"

func main()  {
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga Koma Lima = ", 3.5)
}

// go build 03_number.go
// ./03_number
 
// Tipe Data Number
	// 1. Integer
	// 2. Floating Point

// Integer
	// int8 (-128 s/d 127)
	// int16 (-32,768 s/d 32,767)
	// int32 (-2,147,483,648 s/d 2,147,483,647)
	// int64 (-9,223,372,036,854,775,808 s/d -9,223,372,036,854,775,807)

// Unsigned Integer
	// uint8 (0 s/d 255)
	// uint16 (0 s/d 65535)
	// uint32 (0 s/d 4,294,967,295)
	// uint64 (0 s/d 18,446,744,073,709,551,615)

// Floating Point
	// float32 (1.18e-38 s/d 3.4e38)
	// float64 (2.23e-308 s/d 1.8e308)
	// complex64 (float32 with real and imaginary)
	// complex128 (float64 with real and imaginary)

// Alias
	// byte = uint8
	// rune = int32
	// int = minimal int32 (mengukuti bit OS)
	// uint = minimal uint32 (mengikuti bit OS)