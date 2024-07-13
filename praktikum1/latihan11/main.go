package main

import (
	"fmt"

	"latihan-module/praktikum1/latihan10/fungsi"
)

func main() {
	fungsi.Garis()

	//Mengubah elemen dari Slice
	harga := []int{10000, 20000, 30000}
	fmt.Println("Isi slice harga adalah:", harga)
	harga[2] = 50000
	fmt.Println("Isi slice harga sekarang adalah:", harga)
	fungsi.Garis()

	//Menambahkan elemen dari Slice
	//Syntax --> slice_name = append(slice_name, element1, element2, ...)
	dataAngka := make([]int, 5, 10)
	dataAngka[0] = 7
	dataAngka[1] = 5
	dataAngka[2] = 9
	fmt.Printf("Isi slice dataAngka adalah  %v\n", dataAngka)
	fmt.Printf("Panjang slice dataAngka adalah %d\n", len(dataAngka))
	fmt.Printf("Kapasitas slice dataAngka adalah %d\n", cap(dataAngka))
	fungsi.Garis()

	dataAngka = append(dataAngka, 2, 8)
	fmt.Printf("Isi slice dataAngka sekarang adalah  %v\n", dataAngka)
	fmt.Printf("Panjang slice dataAngka sekarang adalah %d\n", len(dataAngka))
	fungsi.Garis()
}
