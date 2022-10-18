package main

import (
	"fmt"

	"github.com/purbha/latihan/praktikum1/latihan10/fungsi"
)

func main() {
	fungsi.Garis()
	namaHari := 4 //Go switch Statement Single-Case
	fmt.Print("Hari ke ", namaHari, " adalah hari ")
	switch namaHari {
	case 1:
		fmt.Println("Senin")
	case 2:
		fmt.Println("Selasa")
	case 3:
		fmt.Println("Rabu")
	case 4:
		fmt.Println("Kamis")
	case 5:
		fmt.Println("Jumat")
	case 6:
		fmt.Println("Sabtu")
	default:
		fmt.Println("Minggu")
	}
	fungsi.Garis()
	jenisHari := 5 //Go Multi-case switch Statement
	fmt.Println("jenisHari adalah", jenisHari)
	fmt.Print(jenisHari, " adalah ")
	switch jenisHari {
	case 1, 3, 5:
		fmt.Println("hari kerja ganjil.")
	case 2, 4:
		fmt.Println("hari kerja genap.")
	case 6, 7:
		fmt.Println("hari libur.")
	default:
		fmt.Println("jenisHari tidak valid.")
	}
}
