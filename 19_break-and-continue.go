package main

import "fmt"

func main() {
	// break
	for i:=0; i<10; i++ {
		fmt.Println("break", i)
		if i==5 {
			break
		}
	}

	// continue
	for i:=0; i<10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("continue", i)
	}
}

// go build 19_break-and-continue.go
// go run 19_break-and-continue.go

// break untuk menghentikan seluruh perulangan,
// ...tidak peduli sisa perulangan

// continue tidak akan mengerjakan perulangan tertentu saja
