package main

import (
	"fmt"

	"latihan-module/praktikum1/latihan10/fungsi"
	"latihan-module/praktikum1/latihan22/tampil"
)

func main() {
	tampil.Cetakjudul("Program Maps")
	var motor = make(map[string]string)
	motor["merk"] = "Honda"
	motor["model"] = "Vario"
	motor["tahun"] = "2018"
	fmt.Println("Data maps motor adalah", motor)
	fungsi.Garis()

	//Mengubah elemen tahun
	motor["tahun"] = "2020"
	//Menambah elemen warna
	motor["warna"] = "Hitam"
	fmt.Println("Data maps motor sekarang adalah", motor)

	//Menghapus elemen model
	delete(motor, "model")
	fmt.Println("Data maps motor sekarang adalah", motor)
	fungsi.Garis()

	motor["cc"] = ""
	nilai1, status1 := motor["merk"]  //Cek dari key yang ada dan nilainya
	nilai2, status2 := motor["model"] //Cek dari key yang tidak ada dan nilainya
	nilai3, status3 := motor["cc"]    //Cek dari key yang ada dan nilainya
	_, status4 := motor["tahun"]      // Hanya mengecek statusnya saja

	fmt.Println("nilai1 adalah", nilai1, "dan statusnya adalah", status1)
	fmt.Println("nilai2 adalah", nilai2, "dan statusnya adalah", status2)
	fmt.Println("nilai3 adalah", nilai3, "dan statusnya adalah", status3)
	fmt.Println("nilai4 statusnya adalah", status4)
}
