package main

import (
	"fmt"

	"latihan-module/praktikum1/latihan10/fungsi"
)

func tampilDeret(dataAwal int) int {
	if dataAwal == 11 {
		return 0
	}
	fmt.Println("Nilai variable dataAwal adalah: ", dataAwal)
	//Sebuah fungsi menjadi bersifat rekursif jika fungsi tersebut
	//memanggil dirinya sendiri
	return tampilDeret(dataAwal + 1)
}

func faktorial(x int) (y int) {
	if x > 0 {
		y = x * faktorial(x-1)
	} else {
		y = 1
	}
	return
}

func CetakJudul() {
	fungsi.Garis()
	fmt.Println("Program Rekursi")
	fungsi.Garis()
}

var nilaiA int = 4

func main() {
	CetakJudul()
	tampilDeret(1)
	fungsi.Garis()
	fmt.Println("Nilai dari faktorial ", nilaiA, "adalah", faktorial(nilaiA))
}
