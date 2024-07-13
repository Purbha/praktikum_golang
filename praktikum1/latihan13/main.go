package main

import (
	"fmt"

	"latihan-module/praktikum1/latihan10/fungsi"
)

/*
Efisiensi Memori
Ketika menggunakan slice, GO menuangkan seluruh elemen slice kedalam memori.
Jika isi data slice terlalu besar dan kita hanya perlu beberapa elemen saja maka
lebih baik kita meng-copy elemen tersebut ke slice baru menggunakan copy().
Ini akan mengurangi penggunaan memori pada program.
*/
func main() {
	fungsi.Garis()
	//Sebuah Slice Besar
	dataAngka := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	fmt.Printf("Isi dataAngka \t: %v\n", dataAngka)
	fmt.Printf("Panjang dataAngka \t: %d\n", len(dataAngka))
	fmt.Printf("Kapasitas dataAngka \t: %d\n", cap(dataAngka))
	fungsi.Garis()

	//Copy dataAngka dari elemen ke 2 sampai panjang elemen dataAngka dikurang 10
	dataCut := dataAngka[2 : len(dataAngka)-10]
	fmt.Printf("Isi dataCut \t: %v\n", dataCut)

	//Cek panjang dan kapasitas dataCut
	fmt.Printf("Panjang dataCut \t: %d\n", len(dataCut))
	fmt.Printf("Kapasitas dataCut \t: %d\n", cap(dataCut))
	fungsi.Garis()

	//Buat slice dengan panjang berupa panjang dari dataCut
	dataCopy := make([]int, len(dataCut))

	//Copy isi dataCut ke dataCopy
	copy(dataCopy, dataCut)

	fmt.Printf("Isi dataCopy \t: %v\n", dataCopy)
	fmt.Printf("Panjang dataCopy \t: %d\n", len(dataCopy))
	fmt.Printf("Kapasitas dataCopy \t: %d\n", cap(dataCopy))
	fungsi.Garis()
}
