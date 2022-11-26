package model_barang

import (
	"encoding/json"
	"fmt"
	"net/http"

	config_db "github.com/purbha/latihan/praktikum2/latihan19/config"
	decrypt_teks "github.com/purbha/latihan/praktikum2/latihan19/encrypt"
)

func Cekbarang(w http.ResponseWriter, r *http.Request) {
	var brgenc BrgEnc
	_ = json.NewDecoder(r.Body).Decode(&brgenc)
	decText, err := decrypt_teks.Decrypt(brgenc.Data, config_db.MySecret)
	if err != nil {
		fmt.Println("error decrypting your encrypted text: ", err)
	}
	//decText = replaceJsonString(decText)
	brgenc.Data = decText
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(brgenc)
}

/*
func replaceJsonString(x string) string {
	result := strings.Replace(x, "\f", "", -1)
	return result
}
*/
