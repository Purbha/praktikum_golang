package main

import (
	"fmt"
	"latihan-module/praktikum1/latihan4/aritmatik"

	"github.com/enescakir/emoji"
)

/*
	Sebuah fungsi adalah sebuah blok dari pernyataan-pernyataan yang bisa digunakan
	secara berulang-ulang dalam sebuah program.
	Sebuah fungsi tidak akan di eksekusi secara otomatis ketika program dijalankan.
	Sebuah fungsi akan di eksekusi ketika fungsi tersebut dipanggil.
*/

func tampilSalam(nama string) string {
	return "Assalamualaikum " + nama + emoji.ThumbsUp.String()
}

func main() {

	const nama string = "Irsyad"
	const a, b int = 2, 3
	c := aritmatik.Penjumlahan(a, b)

	fmt.Println(tampilSalam(nama) + " apa kabar?")
	fmt.Println("Hasil dari", a, "ditambah", b, "adalah", c)

}
