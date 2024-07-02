/*
Nama : Ananda Rakhma Aulia
Nim : 3420220024
Prodi : Teknik informatika
*/

package main

import (
	"fmt"
)

type Orang struct {
	nama       string
	kodekelas  int
	namakelas  string
	kodekamar  string
	namakamar  string
	harga      int
	lamasewa   int
	subtotal   int
	diskon     int
	totalbayar int
}

func cetakOrang(org Orang) {

	fmt.Println("Nama Pelanggan : ", org.nama)
	fmt.Println("Kode Kelas 	: ", org.kodekelas)
	fmt.Println("Kode Kamar 	: ", org.kodekamar)
	fmt.Println("Lama Sewa 		: ", org.lamasewa)

}

func cetakbon(org Orang) {

	fmt.Println(org.nama, "\t", org.namakelas, "\t", org.namakamar, "\t", org.harga, "\t", org.lamasewa, "\t", org.subtotal, "\t", org.diskon, "\t", org.totalbayar)

}

func garis() {
	fmt.Println("===========================================================")

}

func awal() {
	garis()
	fmt.Println("Kode  | Nama Kelas | Biaya Per-malam")
	garis()
	fmt.Println("1     | VIP        | Rp. 1.250.000")
	fmt.Println("1     | Reguler    | Rp. 900.000")
	fmt.Println("1     | Ekonomi    | Rp. 750.000")
	garis()

	garis()
	fmt.Println("Kode  | Nama Kamar ")
	garis()
	fmt.Println("1     | Asther ")
	fmt.Println("1     | Seruni")
	fmt.Println("1     | Flamboyan")
	garis()

	garis()
	fmt.Println("Diskon 10% jika lama menginap lebih dari 10 hari ")
	fmt.Println("Diskon 5% jika lama menginap lebih dari 5 hari")
	garis()

}

func main() {
	awal()

	var orang1 Orang
	var orang2 Orang
	var orang3 Orang

	orang1.nama = "Son"
	orang1.kodekelas = 2
	switch orang1.kodekelas {
	case 1:
		orang1.namakelas = "  VIP  "
		orang1.harga = 1250000
		break
	case 2:
		orang1.namakelas = "  Reguler"
		orang1.harga = 900000
		break
	case 3:
		orang1.namakelas = "  Ekonomi	"
		orang1.harga = 750000
		break

	}

	orang1.kodekamar = "F"
	if orang1.kodekamar == "A" {
		orang1.namakamar = "Asther"
	} else if orang1.kodekamar == "S" {
		orang1.namakamar = "Seruni"
	} else if orang1.kodekamar == "F" {
		orang1.namakamar = "  Flamboyan"
	}

	orang1.lamasewa = 12
	orang1.subtotal = orang1.harga * orang1.lamasewa

	if orang1.lamasewa > 9 {
		orang1.diskon = orang1.subtotal * 10 / 100
	} else if orang1.lamasewa > 4 {
		orang1.diskon = orang1.subtotal * 5 / 100
	}
	orang1.totalbayar = orang1.subtotal - orang1.diskon

	orang2.nama = "Cho"
	orang2.kodekelas = 1
	switch orang2.kodekelas {
	case 1:
		orang2.namakelas = "  VIP        "
		orang2.harga = 1250000
		break
	case 2:
		orang2.namakelas = "  Reguler"
		orang2.harga = 900000
		break
	case 3:
		orang2.namakelas = "  Ekonomi"
		orang2.harga = 750000
		break

	}

	orang2.kodekamar = "A"
	if orang2.kodekamar == "A" {
		orang2.namakamar = "  Asther"
	} else if orang2.kodekamar == "S" {
		orang2.namakamar = "Seruni"
	} else if orang2.kodekamar == "F" {
		orang2.namakamar = "Flamboyan"
	}

	orang2.lamasewa = 8
	orang2.subtotal = orang2.harga * orang2.lamasewa

	if orang2.lamasewa > 9 {
		orang2.diskon = orang2.subtotal * 10 / 100
	} else if orang2.lamasewa > 4 {
		orang2.diskon = orang2.subtotal * 5 / 100
	}
	orang2.totalbayar = orang2.subtotal - orang2.diskon

	orang3.nama = "Song"
	orang3.kodekelas = 3
	switch orang3.kodekelas {
	case 1:
		orang3.namakelas = "  VIP        "
		orang3.harga = 1250000
		break
	case 2:
		orang3.namakelas = "  Reguler"
		orang3.harga = 900000
		break
	case 3:
		orang3.namakelas = "  Ekonomi"
		orang3.harga = 7500000
		break

	}

	orang3.kodekamar = "S"
	if orang3.kodekamar == "A" {
		orang3.namakamar = "Asther"
	} else if orang3.kodekamar == "S" {
		orang3.namakamar = "  Seruni"
	} else if orang3.kodekamar == "F" {
		orang3.namakamar = "Flamboyan"
	}

	orang3.lamasewa = 4
	orang3.subtotal = orang3.harga * orang3.lamasewa

	if orang3.lamasewa > 9 {
		orang3.diskon = orang3.subtotal * 10 / 100
	} else if orang3.lamasewa > 4 {
		orang3.diskon = orang3.subtotal * 5 / 100
	}
	orang3.totalbayar = orang3.subtotal - orang3.diskon

	fmt.Println("Data ke - 1")
	cetakOrang(orang1)
	fmt.Println("Data ke - 2")
	cetakOrang(orang2)
	fmt.Println("Data ke - 3")
	cetakOrang(orang3)

	garis()
	fmt.Print("Nama\t|| Kelas\t || Kamar\t|| Harga\t|| Lama || Subtotal ||	Diskon\t || Total Bayar || \n")
	garis()
	cetakbon(orang1)
	cetakbon(orang2)
	cetakbon(orang3)
	garis()

}