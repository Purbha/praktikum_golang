package main

import (
	"fmt"

	"latihan-module/praktikum1/latihan10/fungsi"
)

func main() {
	fungsi.Garis()

	//Menjumlahkan 2 buah slice
	//Syntax --> slice3 = append(slice1, slice2...)
	dataNama1 := []string{"Budi", "Edo", "Umar"}
	fmt.Printf("datanama1 adalah %v\n", dataNama1)
	dataNama2 := []string{"Bunga", "Intan", "Lisa"}
	fmt.Println("dataNama2 adalah", dataNama2)
	dataNama3 := append(dataNama1, dataNama2...)
	fmt.Print("dataNama3 adalah ", dataNama3, "\n")
	fmt.Printf("Panjang slice dataNama3 adalah %d\n", len(dataNama3))
	fmt.Printf("Kapasitas slice dataNama3 adalah %d\n", cap(dataNama3))
	fungsi.Garis()

	//Merubah panjang dari slice
	arrayAngka := [8]int{9, 10, 11, 12, 13, 14, 15, 16} //Sebuah array angka
	dataAngka1 := arrayAngka[5:6]                       //Membuat slice dari array
	fmt.Printf("Isi dataAngka1 adalah %v\n", dataAngka1)
	fmt.Printf("Panjang dataAngka1 adalah %d\n", len(dataAngka1))
	fmt.Printf("Kapasitas dataAngka1 adalah %d\n", cap(dataAngka1))
	fungsi.Garis()

	dataAngka1 = arrayAngka[1:6] //Merubah panjang slice dengan cara di slice ulang
	fmt.Printf("Isi dataAngka1 adalah %v\n", dataAngka1)
	fmt.Printf("Panjang dataAngka1 adalah %d\n", len(dataAngka1))   //Panjang slice berkurang
	fmt.Printf("Kapasitas dataAngka1 adalah %d\n", cap(dataAngka1)) //Kapasitas slice tidak berubah
	fungsi.Garis()

	//Kapasitas slice bertambah kelipatan 5 jika di bawah 10.
	//Jika diatas 10 kapasitasnya akan bertambah 2
	dataAngka1 = append(dataAngka1, 20, 21, 22) //Merubah panjang slice dengan cara menambahkan item
	fmt.Printf("Isi dataAngka1 adalah %v\n", dataAngka1)
	fmt.Printf("Panjang dataAngka1 adalah %d\n", len(dataAngka1)) //Panjang slice bertambah
	fmt.Printf("Kapasitas dataAngka1 adalah %d\n", cap(dataAngka1))
}
