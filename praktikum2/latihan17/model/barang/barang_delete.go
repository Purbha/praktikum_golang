package model_barang

import (
	"encoding/json"
	"net/http"

	config "latihan-module/praktikum2/latihan17/config"

	"github.com/gorilla/mux"
)

func DeleteBarang(w http.ResponseWriter, r *http.Request) {
	var sqlStatement string
	var barang Barang
	params := mux.Vars(r)
	barang.IDBar = params["id"]
	sqlStatement = "DELETE FROM barang " +
		"WHERE idbar = $1"
	config.Buka_koneksi()
	_, err := config.Db.Exec(sqlStatement, barang.IDBar)
	if err != nil {
		panic(err)
	}
	config.Tutup_koneksi()
	sqlStatement = "SELECT * FROM barang"
	config.Buka_koneksi()
	result, err := config.Db.Query(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	config.Tutup_koneksi()
	defer result.Close()
	var barangs []Barang

	for result.Next() {
		err := result.Scan(&barang.IDBar, &barang.Bar_Nama, &barang.Bar_Harga, &barang.Bar_Stok,
			&barang.Bar_Create, &barang.Bar_Update)
		if err != nil {
			panic(err.Error())
		}
		barangs = append(barangs, barang)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(barangs)
}
