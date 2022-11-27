package model_barang

import (
	"encoding/json"
	"net/http"

	config "github.com/purbha/latihan/praktikum2/latihan17/config"
)

func TambahBarang(w http.ResponseWriter, r *http.Request) {
	var idbar, sqlStatement string
	var barang Barang
	var err error
	err = json.NewDecoder(r.Body).Decode(&barang)
	if err != nil {
		panic(err.Error())
	}
	sqlStatement = "INSERT INTO barang (bar_nama, bar_harga, bar_stok, bar_create, bar_update) " +
		"VALUES ($1, $2, $3, current_timestamp, current_timestamp) RETURNING idbar"
	config.Buka_koneksi()
	err = config.Db.QueryRow(sqlStatement, barang.Bar_Nama, barang.Bar_Harga, barang.Bar_Stok).Scan(&idbar)
	if err != nil {
		panic(err)
	}
	config.Tutup_koneksi()
	sqlStatement = "SELECT * FROM barang WHERE idbar = " + idbar
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
