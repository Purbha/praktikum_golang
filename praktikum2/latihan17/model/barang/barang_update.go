package model_barang

import (
	"encoding/json"
	"net/http"

	config "latihan-module/praktikum2/latihan17/config"

	"github.com/gorilla/mux"
)

func UpdateBarang(w http.ResponseWriter, r *http.Request) {
	var sqlStatement string
	var barang Barang
	params := mux.Vars(r)
	barang.IDBar = params["id"]
	err := json.NewDecoder(r.Body).Decode(&barang)
	if err != nil {
		panic(err.Error())
	}
	sqlStatement = "UPDATE barang SET bar_nama = $1, bar_harga = $2, bar_stok = $3, bar_update=current_timestamp " +
		"WHERE idbar = $4"
	config.Buka_koneksi()
	_, err = config.Db.Exec(sqlStatement, barang.Bar_Nama, barang.Bar_Harga, barang.Bar_Stok, barang.IDBar)
	if err != nil {
		panic(err)
	}
	config.Tutup_koneksi()
	sqlStatement = "SELECT * FROM barang WHERE idbar = " + barang.IDBar
	config.Buka_koneksi()
	result, err := config.Db.Query(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	config.Tutup_koneksi()
	defer result.Close()
	for result.Next() {
		err := result.Scan(&barang.IDBar, &barang.Bar_Nama, &barang.Bar_Harga, &barang.Bar_Stok,
			&barang.Bar_Create, &barang.Bar_Update)
		if err != nil {
			panic(err.Error())
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(barang)
}
