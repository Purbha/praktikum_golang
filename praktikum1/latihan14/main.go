package main

import (
	"fmt"

	"latihan-module/praktikum1/latihan10/fungsi"
)

func main() {
	fungsi.Garis()
	nilaiA, nilaiB := 20, 30
	fmt.Println("nilaiA adalah", nilaiA)
	fmt.Println("nilaiB adalah", nilaiB)
	//Jika nilaiA lebih besar dari nilaiB
	if nilaiA > nilaiB {
		fmt.Println("nilaiA lebih besar dari nilaiB.")
	} else {
		fmt.Println("nilaiB lebih besar dari nilaiA.")
	}
	fungsi.Garis()
	jam := 14
	fmt.Println("Sekarang adalah pukul", jam, ":00")
	if jam < 10 {
		fmt.Println("Selamat Pagi.")
	} else if jam < 15 {
		fmt.Println("Selamat Siang.")
	} else {
		fmt.Println("Selamat Malam.")
	}
	fungsi.Garis()
	if jam <= 8 { //Contoh Penggunaan Nested If
		fmt.Println("Waktunya sarapan pagi.")
	} else if jam < 15 {
		fmt.Println("Waktunya makan siang.")
		if jam == 14 {
			fmt.Println("Setelah itu bobo siang ya.")
		}
	} else {
		fmt.Println("Waktunya makan malam.")
	}
}
