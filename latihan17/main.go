package main

import (
	"fmt"

	"github.com/purbha/latihan/latihan10/fungsi"
)

func main() {
	fungsi.Garis()
	//===========Nested Loops=============//
	//Buat Array
	sifat := [2]string{"Besar", "Manis"}
	buah := [3]string{"Apel", "Jeruk", "Anggur"}

	//Ulang selama panjang array sifat
	for i := 0; i < len(sifat); i++ {

		//Ulang selama panjang array buah
		for j := 0; j < len(buah); j++ {
			fmt.Println(buah[j], sifat[i])
		}

	}
	fungsi.Garis()
	//===========The Range Keyword=============//
	namaBuah := [5]string{"Apel", "Anggur", "Jeruk", "Pisang", "Jambu"}
	for indeks, nilai := range namaBuah {
		fmt.Printf("Indek ke-%v adalah %v\n", indeks, nilai)
	}
	fungsi.Garis()
	for _, nilai := range namaBuah {
		fmt.Printf("Nilainya adalah %v\n", nilai)
	}

}
