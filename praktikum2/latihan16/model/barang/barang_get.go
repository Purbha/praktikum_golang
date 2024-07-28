package model_barang

import (
	"encoding/json"
	"net/http"

	config "latihan-module/praktikum2/latihan16/config"

	"github.com/gorilla/mux"
)

func AmbilBarangs(w http.ResponseWriter, r *http.Request) {
	config.Buka_koneksi()
	var barangs []Barang
	result, err := config.Db.Query("SELECT * FROM barang")
	if err != nil {
		panic(err.Error())
	}
	config.Tutup_koneksi()
	defer result.Close()
	for result.Next() {
		var barang Barang
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

func AmbilBarang(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idbar := params["id"]
	var barang Barang
	sqlStatement := "SELECT * FROM barang WHERE idbar = " + idbar
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
