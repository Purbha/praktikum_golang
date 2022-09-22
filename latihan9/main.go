package main

import (
	"fmt"
)

func garis() {
	fmt.Println("=================================================")
}

func main() {
	/*
		Slice - Mirip seperti array tetapi lebih fleksible
		Panjang dari Slice bisa bertambah dan berkurang
		Contoh Deklarasi: slice_name := []datatype{values}
	*/
	garis()

	mataKuliah1 := []string{}
	fmt.Println("Panjang slice mataKuliah1 adalah:", len(mataKuliah1))
	fmt.Println("Isi dari mataKuliah1 adalah:", mataKuliah1)

	garis()

	mataKuliah2 := []string{"Pemrograman Mobile,", "Web Technology,", "Kapita Selekta,", "Skripsi"}
	fmt.Println("Panjang slice mataKuliah2 adalah:", len(mataKuliah2))
	fmt.Println("Isi dari mataKuliah2 adalah:", mataKuliah2)

	garis()

	//Mmebuat slice dari array
	dataNilai := [6]int{10, 11, 12, 13, 14, 15}
	sliceNilai := dataNilai[2:4]

	//%v - Prints the value in the default format
	fmt.Printf("Nilai dari dataNilai adalah: %v\n", dataNilai)
	fmt.Printf("Nilai dari sliceNilai[2:4] adalah: %v\n", sliceNilai)
	fmt.Printf("Panjang dari sliceNilai adalah: %d\n", len(sliceNilai))
	fmt.Printf("Kapasitas dari sliceNilai adalah: %d\n", cap(sliceNilai))

	garis()
}
