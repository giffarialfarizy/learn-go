package main

import "fmt"

func main() {
	counter := 0

	// Hati2 dengan function di dalam function
	increment := func() {
		// variabel di dalam function perhatikan dengan yang di luar function
		
		// Usahakan bikin variable yang namanya beda
		// Atau buat lagi (deklarasi variabel) counter := 0
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
}

// Closure adalah kemampuan sebuah function berinteraksi dengan
// ...data-data di sekitarnya dalam scope yang sama

// Gunakan fitur closure dengan bijak, karena kalau sampai salah
// ...implement kita bisa tidak sengaja mengubah data