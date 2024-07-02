package main

import "fmt"

func main() {

	//var variableName type = value
	var nama string = "Irsyad"
	var umur int = 37
	var ukuranSepatu float32 = 42.5

	//const constName type = value
	const isPria bool = true

	fmt.Println("Nama saya adalah:", nama)
	fmt.Println("Umur saya adalah:", umur)
	// \t artinya mencetak karakter TAB
	fmt.Println("Ukuran sepatu saya adalah \t\t\t:", ukuranSepatu)
	fmt.Println("Jenis kelamin saya adalah (true = pria)) \t:", isPria)

}
