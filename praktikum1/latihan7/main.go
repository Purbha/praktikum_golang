package main

import (
	"fmt"
)

func main() {
	// Arrays
	var namaBand = [3]string{}

	//Pemberian nilai menggunakan index array
	namaBand[0] = "Muse"
	namaBand[1] = "Nirvana"
	namaBand[2] = "The Cranberries"
	fmt.Printf("Band-band favorit %s \n", namaBand)

	macamBuah := [...]string{"Apel", "Jeruk", "Anggur", "Nanas"}

	//Menentukan panjang sebuah array
	fmt.Printf("Jumlah elemen dalam array macamBuah adalah: %d\n", len(macamBuah))
	fmt.Println("Pengambilan nilai dengan Slice: ", macamBuah[1:3])

	//Inisialisasi Array Langsung
	angka1 := [5]int{}              //Belum diinisialisasi
	angka2 := [5]int{1, 2}          //Diinisialisasi sebagian
	angka3 := [5]int{1, 2, 3, 4, 5} //Diinisialisasi lengkap

	//Cetak Array
	fmt.Println(angka1)
	fmt.Println(angka2)
	fmt.Println(angka3)

}
