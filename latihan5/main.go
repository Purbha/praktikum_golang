package main

import (
	"fmt"
)

/*
	Sebuah fungsi adalah sebuah blok dari pernyataan-pernyataan yang bisa digunakan
	secara berulang-ulang dalam sebuah program.
	Sebuah fungsi tidak akan di eksekusi secara otomatis ketika program dijalankan.
	Sebuah fungsi akan di eksekusi ketika fungsi tersebut dipanggil.
*/

func tampilSalam(nama string) string {
	return "Assalamualaikum " + nama
}

func jumlahin(num1, num2 int) int {
	return num1 + num2
}

func main() {

	const nama string = "Irsyad"
	const a, b int = 2, 3

	fmt.Println(tampilSalam(nama) + " apa kabar?")
	fmt.Println("Hasil dari ", a, " ditambah ", b, " adalah ", jumlahin(a, b))

}
