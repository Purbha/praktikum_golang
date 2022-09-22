package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		Pemberian tipe float64 karena fungsi Floor, Ceil dan Sqrt memiliki tipe parameter Float64
		Reff Floor : https://golangbyexample.com/floor-number-golang/
		Reff Ceil  : https://golangbyexample.com/ceil-number-golang/
		Reff Sqrt  : https://pkg.go.dev/math#Sqrt
	*/

	//Go Multiple Variable Declaration
	var NilaiA, NilaiB float64 = 2.7, 2.7

	//Pemberian nilai dengan tanda ":="
	//Pemberian nilai 16.0 agar value yang diberikan tidak dianggap int
	NilaiC := 16.0

	fmt.Println("Nilai A adalah:", NilaiA)
	fmt.Println("Nilai Floor A adalah:", math.Floor(NilaiA))

	fmt.Println("Nilai B adalah:", NilaiB)
	fmt.Println("Nilai Ceil B adalah:", math.Ceil(NilaiB))

	fmt.Println("Nilai C adalah:", NilaiC)
	fmt.Println("Nilai Sqrt C adalah:", math.Sqrt(NilaiC))

}
