package tampil

import (
	"fmt"

	"latihan-module/praktikum1/latihan10/fungsi"
)

func Cetakjudul(judul string) {
	fungsi.Garis()
	fmt.Println("\t", judul)
	fungsi.Garis()
}
