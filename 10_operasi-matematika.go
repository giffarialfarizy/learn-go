package main

import "fmt"

func main() {
	var a int16 = 5
	var b int16 = 8
	var c int16

	c = a + b
	fmt.Println(c)

	// Augmented Assignments
	a = a+10
	fmt.Println(a)
	a += 10
	fmt.Println(a)

	// Unary Operator
	var i = 0
	i++
	fmt.Println(i)
}

// go build 10_operasi-matematika.go
// go run 10_operasi-matematika.go

// AUGMENTED ASSIGNMENTS
// +=
// -=
// *=
// /=
// %=

// UNARY OPERATOR
// ++
// --
// - (negative)
// + (positive)
// ! (Boolean kebalikan)
