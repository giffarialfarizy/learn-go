package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

// func runApplication(value int) {
// 	fmt.Println("Run Application")
// 	result := 10/value
// 	fmt.Println("Result:", result)
// 	// ini cara manggil fungsi di dalam fungsi
// 	// logging tidak akan dieksekusi jika atasnya error
// 	logging()
// }

func runApplication(value int) {
	defer logging()
	fmt.Println("Run Application")
	result := 10/value
	fmt.Println("Result:", result)
}

// func runApplication() {
// 	defer logging()
// 	fmt.Println("Run Application")
// }

// PANIC
func endApp() {
	fmt.Println("Aplikasi selesai")
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi ERROR")
	}
	// Contoh yang salah
	// message := recover()
	// fmt.Println("Error dengan message:", message)
	fmt.Println("Aplikasi berjalan")
}

func main() {
	// runApplication(0)
	runApp(true)
	fmt.Println("Baris ini dieksekusi")
}

// Defer function adalah function yang bisa kita jadwalkan
// ...untuk dieksekusi setelah sebuah function lain selesai
// ...dieksekusi

// Defer function akan selalu dieksekusi walaupun terjadi
// ...ERROR di function yang dieksekusi

// Panic function adalah function yang bisa digunakan untuk
// ...menghentikan program

// Panic function biasanya dipanggil ketika terjadi error 
// ...pada saat program kita berjalan

// Saat panic function dipanggil, program akan terhanti,
// ...namun defer function tetap akan dieksekusi

// Recover adalah function yang bisa kita gunakan untuk
// ...menangkap data panic

// Dengan recover proses panic akan terhenti,
// ...sehingga program akan tetap berjalan

// Simpan recover di dalam defer function