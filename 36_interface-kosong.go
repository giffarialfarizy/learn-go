package main

import "fmt"

func Ups(i int) interface {} {
	if i==1 {
		return 1
	} else if i==2 {
		return true
	} else {
		return "Ups"
	}
}

func main(){
	// Tidak bisa begini karena outputnya ada 3 kemungkinan tipe
	// var data int = Ups(1)
	// var data bool = Ups(2)
	// var data string = Ups(3)

	// Sehingga harus pakai interface kosong
	var data interface{} = Ups(1)
	fmt.Println(data)
}