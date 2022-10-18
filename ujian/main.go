package main

import (
	"fmt"
)

func cetakGaris(panjang int) {
	for i := 0; i < panjang; i++ {
		fmt.Print("=")
	}
	fmt.Print("\n")
}

type Pelanggan struct {
	no    int
	nama  string
	kelas int
	kamar string
	lama  int
}

func cetakKetentuan() {
	cetakGaris(50)
	fmt.Println("Kode \t| Nama Kelas\t| Biaya Permalam")
	cetakGaris(50)
	fmt.Println("1\t| VIP\t\t| 1.000.000")
	fmt.Println("2\t| Reguler\t| 800.000")
	fmt.Println("3\t| Ekonomi\t| 600.000")
	cetakGaris(50)
	fmt.Print("\n")
	cetakGaris(50)
	fmt.Println("Kode\t| Nama Kamar\t")
	cetakGaris(50)
	fmt.Println("A\t| Asther")
	fmt.Println("S\t| Serunir")
	fmt.Println("F\t| Flamboyan")
	cetakGaris(50)
	fmt.Print("\n")
	cetakGaris(50)
	fmt.Println("Diskon 10% jika lama sewa lebih dari 10 hari.")
	fmt.Println("Diskon 5% jika lama sewa lebih dari 5 hari.")
	cetakGaris(50)
	fmt.Print("\n")
}

var pel1, pel2, pel3 Pelanggan

func IsiData() {
	pel1.no = 1
	pel1.nama = "Andi"
	pel1.kelas = 2
	pel1.kamar = "F"
	pel1.lama = 12
	pel2.no = 2
	pel2.nama = "Joni"
	pel2.kelas = 1
	pel2.kamar = "A"
	pel2.lama = 8
	pel3.no = 3
	pel3.nama = "Budi"
	pel3.kelas = 3
	pel3.kamar = "S"
	pel3.lama = 4
}

func cetakPel(data Pelanggan) {
	fmt.Println("Data ke -", data.no)
	fmt.Println("Nama Pelanggan\t:", data.nama)
	fmt.Println("Kode Kelas\t:", data.kelas)
	fmt.Println("Kode Kamar\t:", data.kamar)
	fmt.Println("Lama Sewa\t:", data.lama)
	fmt.Print("\n")
}

func dataPel(data Pelanggan) {
	var kelas string = cetakKelas(data.kelas)
	var kamar string = cetakKamar(data.kamar)
	var harga float32 = cetakHarga(data.kelas)
	var subtotal float32 = harga * float32(data.lama)
	var diskon float32 = hitungDiskon(data.lama, subtotal)
	var total float32 = subtotal - diskon
	fmt.Print("|")
	fmt.Printf(" %v  |", data.no)
	fmt.Printf(" %v |", data.nama)
	fmt.Printf(" %v |", kelas)
	fmt.Printf(" %v |", kamar)
	fmt.Printf(" %8.0f |", harga)
	fmt.Printf(" %4.0f |", float32(data.lama))
	fmt.Printf(" %8.0f |", subtotal)
	fmt.Printf(" %6.0f |", diskon)
	fmt.Printf("    %8.0f |", total)
	fmt.Print("\n")
}

func hitungDiskon(lama int, subtot float32) (hasil float32) {
	if lama > 10 {
		hasil = 0.1 * float32(subtot)
	} else if lama > 5 {
		hasil = 0.05 * float32(subtot)
	} else {
		hasil = 0.0
	}
	return
}

func cetakKelas(kode int) (namaKelas string) {
	if kode == 1 {
		namaKelas = "VIP    "
	} else if kode == 2 {
		namaKelas = "Reguler"
	} else {
		namaKelas = "Ekonomi"
	}
	return
}

func cetakHarga(kode int) (hargaKelas float32) {
	if kode == 1 {
		hargaKelas = 1000000
	} else if kode == 2 {
		hargaKelas = 800000
	} else {
		hargaKelas = 600000
	}
	return
}

func cetakKamar(kode string) (namaKamar string) {
	if kode == "A" {
		namaKamar = "Asther   "
	} else if kode == "S" {
		namaKamar = "Serunir  "
	} else {
		namaKamar = "Flamboyan"
	}
	return
}

func cetakTotal() {
	cetakGaris(90)
	fmt.Println("| No | Nama | Kelas   | Kamar     |    Harga | Lama | Subtotal | Diskon | Total Bayar |")
	cetakGaris(90)
	dataPel(pel1)
	dataPel(pel2)
	dataPel(pel3)
	cetakGaris(90)
	fmt.Println("Terima kasih sudah menginap di Hotel Vinducia")
	cetakGaris(90)
}

func cetakJudul() {
	cetakGaris(50)
	fmt.Println("             Program Ujian Basic Go")
	cetakGaris(50)
	fmt.Println("                HOTEL VINDUCIA")
	fmt.Println("          Jl. Daan Mogot Jakarta Barat")
	cetakGaris(50)
	fmt.Print("\n")
}

func main() {
	cetakJudul()
	cetakKetentuan()
	IsiData()
	cetakPel(pel1)
	cetakPel(pel2)
	cetakPel(pel3)
	cetakTotal()
}
