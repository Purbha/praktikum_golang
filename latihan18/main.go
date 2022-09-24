package main

import (
	"fmt"

	"github.com/purbha/latihan/latihan10/fungsi"
)

// Go Function Named Return Values
func luasSegitiga(alas int, tinggi int) (luas float32) {
	//Didalam GO tidak bisa melakukan operasi matematika yang berbeda tipe valuenya
	//luas = (alas * tinggi) * 0.5
	luas = (float32(alas) * float32(tinggi)) * 0.5
	return
}

func CetakJudul() {
	fungsi.Garis()
	fmt.Println("Program Menghitung Luas Segitiga")
	fungsi.Garis()
}

func main() {
	CetakJudul()
	panjangAlas, panjangTinggi := 10, 3
	fmt.Println("Panjang alas adalah:", panjangAlas)
	fmt.Println("Panjang tinggi adalah:", panjangTinggi)

	//Hitung luas segitiga dan dimasukan ke variable luas
	var luas float32 = luasSegitiga(panjangAlas, panjangTinggi)
	fmt.Println("Luas segitiga adalah:", luas)

	fungsi.Garis()
}
