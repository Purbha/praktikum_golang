package main

import (
	"fmt"

	"github.com/purbha/latihan/latihan10/fungsi"
)

func main() {
	//Fungsi garis yang di buat dalam package
	fungsi.Garis()

	fmt.Println("Contoh Membuat Slice Dengan Make")
	dataNilai1 := make([]int, 5, 10)
	fmt.Printf("Isi dari dataNilai1: %v\n", dataNilai1)
	fmt.Printf("Panjang dari dataNilai1: %d\n", len(dataNilai1))
	fmt.Printf("Kapasitas dari dataNilai1: %d\n", cap(dataNilai1))

	fungsi.Garis()

	//Jika parameter kapasitas tidak di definisikan maka kapasitas
	//akan sama dengan panjang dari slice
	dataNilai2 := make([]int, 5)
	fmt.Printf("Isi dari dataNilai2: %v\n", dataNilai2)
	fmt.Printf("Panjang dari dataNilai2: %d\n", len(dataNilai2))
	fmt.Printf("Kapasitas dari dataNilai2: %d\n", cap(dataNilai2))

	fungsi.Garis()
}
