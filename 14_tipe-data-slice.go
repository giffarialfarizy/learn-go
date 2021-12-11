package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Ini Array
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// Jika mengubah isi Array, akan mengubah isi Slice juga
	months[5] = "|Harusnya Juni|"
	fmt.Println(slice1)

	// Jika mengubah isi Slice juga akan mengubah isi Array
	slice1[0] = "|Harusnya Mei|"
	fmt.Println(months)

	// Coba append
	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "|Bulan baru|")
	fmt.Println((slice3))
	slice3[1] = "|Harusnya Desember|"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	// Perhatikan kapasitasnya

	// Coba eksplor lainnya
	fmt.Println("=====================")
	slice2 = months[2:4]
	var slice4 = append(slice2, "|Bulan baru|")
	fmt.Println((slice4))
	slice4[1] = "|Harusnya Desember|"
	fmt.Println(slice4)

	fmt.Println(slice2)
	fmt.Println(months)

	// NEW SLICE
	fmt.Println("===================== NEW SLICE =====================")
	// terbuat array baru dengan tipe data string, length nya 2, capacity nya 5
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Giffari"
	newSlice[1] = "Alfarizy"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copysLice := make([]string, len(newSlice), cap(newSlice))

	copy(copysLice, newSlice)
	fmt.Println(copysLice)

	// Hati2 saat membuat array.
	// Jika salah yang kita buat bukanlah Array, melainkan Slice

	// Array kapasitasnya ditentukan di awal (boleh kasih "..." atau angka "5" dst)
	iniArray := [...]int{1, 2, 3, 4, 5}

	// Slice tidak
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	// Cari tipe data
	fmt.Println(reflect.TypeOf(iniArray))
	fmt.Println(reflect.TypeOf(iniSlice))

	// Coba test untuk number
	angka := 1
	fmt.Println(reflect.TypeOf(angka))
}

// go build 14_tipe-data-slice.go
// go run 14_tipe-data-slice.go

// Slice adalah potongan dari Array
// ukuran Slice bisa berubah
// Slice dan Array selalu terkoneksi, di mana Slice adalah data
// ... yang mengakses sebagian atau seluruh data di Array

// Slice memiliki 3 data yaitu pointer, length, dan capacity
// Ponter adalah penunjuk data pertama di array para slice
// Length adalah panjang dari slice
// Capacity adalah kapasitas dari slice. Length <= Capacity

// Membuat Slice dari Array
// array[low:high]			: membuat slice dari array dimulai dari index low sampai index sebelum high
// array[low:]				: membuat slice dari array dimulai dari index low sampai index terakhir di array
// array[:high]				: membuat slice daru array dimulai dari 0 hingga index sebelum high
// array[:]					: membuat slice dari array dimulai dari index 0 sampai index akhir di array

// FUNCTION SLICE
// len(slice)							: untuk mendapatkan panjang
// cap(slice)							: untuk mendapatkan kapasitas
// append(slice, data)					: membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh, maka akan membuat array baru
// make([]TypeData, length, capacity)	: Membuat slice baru
// copy(destination, source)			: Menyalin slice dari source ke destination

