package main

import (
	"fmt"

	"latihan-module/praktikum1/latihan4/aritmatik"
)

func main() {
	/*
		"Var"
		- Bisa digunakan diluar ataupun didalam sebuah fungsi
		- Deklarasi variable dan pemberian nilai bisa dilakukan secara terpisah

		":="
		- Hanya bisa digunakan didalam sebuah fungsi
		- Deklarasi variable dan pemberian nilai tidak bisa dilakukan secara terpisah
		  (harus dilakukan dalam baris yang sama)
	*/

	//Go Multiple Variable Declaration
	NilaiA, NilaiB := 2, 3
	var testJumlah int = aritmatik.Penjumlahan(NilaiA, NilaiB)

	fmt.Println("Nilai A adalah:", NilaiA)
	fmt.Println("Nilai B adalah:", NilaiB)

	fmt.Println("Nilai A ditambah Nilai B adalah:", testJumlah)

}
