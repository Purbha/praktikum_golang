package main

import (
	"fmt"

	"github.com/purbha/latihan/praktikum1/latihan10/fungsi"
	"github.com/purbha/latihan/praktikum1/latihan22/tampil"
)

func main() {
	tampil.Cetakjudul("Program Maps")
	var motor1 = map[string]string{"merk": "Honda", "model": "Vario", "tahun": "2018"}
	//Maps References
	//Jika sebuah map mereferensikan ke map yang lain maka mengubah nilai map
	//tersebut akan mengubah nilai map asalnya.
	motor2 := motor1

	fmt.Println("Data motor1 adalah ", motor1)
	fmt.Println("Data motor2 adalah ", motor2)

	motor2["tahun"] = "2020"
	fungsi.Garis()

	fmt.Println("Data motor1 sekarang adalah ", motor1)
	fmt.Println("Data motor2 sekarang adalah ", motor2)
	fungsi.Garis()

	//Looping di dalam map
	for kunci, nilai := range motor1 {
		fmt.Printf("%v:%v, ", kunci, nilai)
	}
	fmt.Print("\n")
}
