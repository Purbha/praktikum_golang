package main

import (
	"fmt"

	"github.com/purbha/latihan/latihan10/fungsi"
	"github.com/purbha/latihan/latihan22/tampil"
)

/*
Maps
Maps digunakan untuk menyimpan data nilai dalam format pasangan key:value
Setiap elemen di dalam sebuah map ada lah sebuah key:value
Sebuah map adalah sebuah koleksi yang tidak tersusun (unordered),
	tidak bisa diubah dan tidak mengizinkan adalah duplikasi
*/

func main() {
	tampil.Cetakjudul("Program Maps")
	//Deklarasi Maps
	//var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
	//b := map[KeyType]ValueType{key1:value1, key2:value2,...}
	var handphone = map[string]string{"merk": "samsung", "model": "Galaxy A51", "tahun": "2020"}
	kota := map[string]int{"Jakarta": 1, "Bandung": 2, "Medan": 3, "Surabaya": 4}
	fmt.Printf("Data handphone adalah %v\n", handphone)

	//Urutan dari map elemen yang terdapat di kode berbeda dengan bagaimana elemen-elemen itu
	//disimpan didalam memori. Data di simpan dalam memori dengan urutan yang paling efisien
	fmt.Printf("Data kota adalah %v\n", kota)
	fungsi.Garis()

	//Deklarasi Maps Menggunakan Make
	//var a = make(map[KeyType]ValueType)
	//b := make(map[KeyType]ValueType)
	var komputer = make(map[string]string) //Map dalam kondisi kosong
	komputer["prosesor"] = "Intel Core i7 Gen 12"
	komputer["ram"] = "16GB DDR5 5200Mhz"
	komputer["ssd"] = "1TB PCI Gen4x4"
	motor := make(map[string]int)
	motor["Vario"] = 1
	motor["Beat"] = 2
	motor["N-Max"] = 3
	motor["PCX"] = 4
	fmt.Printf("Data komputer adalah %v\n", komputer)
	fmt.Printf("Data motor adalah %v\n", motor)

}
