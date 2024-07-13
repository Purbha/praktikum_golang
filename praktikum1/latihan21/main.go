package main

import (
	"fmt"

	"latihan-module/praktikum1/latihan10/fungsi"
)

/*
Go Structures
Sebuah struct (kependekan dari struktur) digunakan untuk membuat kumpulan dari elemen-
elemen yang memiliki tipe data yang berbeda-beda kedalam sebuah single variable.
Struct sangat berguna untuk meng-grouping data-data menjadi sebuah record.
*/

type Orang struct {
	nama         string
	umur         int
	pekerjaan    string
	ukuranSepatu float32
}

func CetakJudul() {
	fungsi.Garis()
	fmt.Println("\tProgram Struct")
	fungsi.Garis()
}

func main() {
	CetakJudul()
	//Deklarasi Variable
	var orang1 Orang
	var orang2 Orang

	//Memberi nilai pada Struct
	orang1.nama = "Budi"
	orang1.umur = 22
	orang1.pekerjaan = "Nelayan"
	orang1.ukuranSepatu = 43.5

	//Memberi nilai pada Struct
	orang2.nama = "Eko"
	orang2.umur = 24
	orang2.pekerjaan = "Petani"
	orang2.ukuranSepatu = 41.5

	//Mengakses nilai pada Struct
	fmt.Print("Nama orang pertama adalah ", orang1.nama, ". Saat ini dia berumur ", orang1.umur, " tahun")
	fmt.Print(". Sehari-hari pekerjaannya adalah seorang ", orang1.pekerjaan)
	fmt.Print(" dan ukuran sepatunya adalah ", orang1.ukuranSepatu)
	fmt.Print("\n")
	fmt.Print("Nama orang kedua adalah ", orang2.nama, ". Saat ini dia berumur ", orang2.umur, " tahun")
	fmt.Print(". Sehari-hari pekerjaannya adalah seorang ", orang2.pekerjaan)
	fmt.Print(" dan ukuran sepatunya adalah ", orang2.ukuranSepatu)
	fmt.Print("\n")

}
