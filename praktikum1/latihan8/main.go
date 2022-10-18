package main

import (
	"fmt"
)

func main() {
	//Arrays Initialize Only Specific Elements
	//Pemberian nilai hanya di elemen ke 1 dan ke 3 saja
	dataNilai := [5]int{1: 98, 3: 70}
	fmt.Println("Data nilai adalah:", dataNilai)

	//Perubahan nilai salah satu elemen array
	dataNilai[0] = 77
	fmt.Println("Data nilai sekarang adalah:", dataNilai)

	//Salah satu cara deklarasi multiple variable
	var (
		nilaiA int = 66
		nilaiB int = 90
	)

	//Memasukan nilai sebuah variable kedalam elemen array
	dataNilai[2] = nilaiA
	dataNilai[4] = nilaiB
	fmt.Println("Data nilai lengkap adalah:", dataNilai)
}
