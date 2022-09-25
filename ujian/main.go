package main

import (
	"fmt"

	"github.com/purbha/latihan/latihan10/fungsi"
	"github.com/purbha/latihan/latihan22/tampil"
)

func garis2() {
	fmt.Println("=====================================================")
}

func garis3(panjang int) {
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

func cetakHarga() {
	garis2()
	fmt.Println("Kode Kelas\t| Nama Kelas\t| Biaya Permalam")
	garis2()
	fmt.Println("1\t\t| VIP\t\t| 1.000.000")
	fmt.Println("2\t\t| Reguler\t| 800.000")
	fmt.Println("3\t\t| Ekonomi\t| 600.000")
	garis2()
	fmt.Print("\n")
	garis2()
	fmt.Println("Kode Kamar\t| Nama Kamar\t")
	garis2()
	fmt.Println("A\t\t| Asther")
	fmt.Println("S\t\t| Serunir")
	fmt.Println("F\t\t| Flamboyan")
	garis2()
	fmt.Print("\n")
	garis2()
	fmt.Println("Diskon 10% jika lama sewa lebih dari 10 hari.")
	fmt.Println("Diskon 5% jika lama sewa lebih dari 5 hari.")
	garis2()
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
	pel3.lama = 7
}

func cetakPel(data Pelanggan) {
	fmt.Println("Data ke -", data.no)
	fmt.Println("Nama Pelanggan\t:", data.nama)
	fmt.Println("Kode Kelas\t:", data.kelas)
	fmt.Println("Kode Kamar\t:", data.kamar)
	fmt.Println("Lama Sewa\t:", data.lama)
	fmt.Print("\n")
}

func cetakPel2(data Pelanggan) {
	var kelas string = cetakKelas(data.kelas)
	var kamar string = cetakKamar(data.kamar)
	var harga int = nilaiHarga(data.kelas)
	var subtotal int = harga * data.lama
	var dis float32 = diskon(data.lama, subtotal)
	var total int = subtotal - int(dis)
	fmt.Print(data.no, "\t", data.nama, "\t", kelas, "\t\t", kamar, "\t", harga, "\t\t")
	fmt.Println(data.lama, "\t", subtotal, "\t", dis, "\t", total)
}

func diskon(lama int, subtot int) (hasil float32) {
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

func nilaiHarga(kode int) (hargaKelas int) {
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
	garis3(120)
	fmt.Println("No\tNama\tKelas\t\tKamar\t\tHarga\t\tLama\tSubtotal\tDiskon\t\tTotal Bayar")
	garis3(120)
	cetakPel2(pel1)
	cetakPel2(pel2)
	cetakPel2(pel3)
	garis3(120)
}

func main() {
	tampil.Cetakjudul("Program Ujian Basic Go")
	fmt.Println("\tHOTEL VINDUCIA")
	fmt.Println("Jl. Daan Mogot Jakarta Barat")
	fungsi.Garis()
	fmt.Print("\n")
	cetakHarga()
	IsiData()
	cetakPel(pel1)
	cetakPel(pel2)
	cetakPel(pel3)
	cetakTotal()
}
