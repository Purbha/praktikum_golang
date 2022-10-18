package main

import (
	"fmt"
)

func main() {
	// Arrays
	/*
		Cara mendeklarasikan array pertama

		var array_name = [length]datatype{values} --> Panjang array di definisikan
		var array_name = [...]datatype{values} --> Panjang array bebas

		Cara mendeklarasikan array kedua

		array_name := [length]datatype{values} --> Panjang array di definisikan
		array_name := [...]datatype{values} --> Panjang array bebas
	*/
	var buah = [2]string{"Anggur", "Nanas"}
	fmt.Printf("Buah yang pertama adalah %s dan buah yang kedua adalah %s \n", buah[0], buah[1])

	/*
		Integer Formatting Verbs	-	"%d"	-	Base 10
		String Formatting Verbs		-	"%s"	-	Prints the value as plain string
		Boolean Formatting Verbs	-	"%t"	-	Value of the boolean operator in true or false format
		Float Formatting Verbs		-	"%.2f"	-	Default width, precision 2
	*/
}
