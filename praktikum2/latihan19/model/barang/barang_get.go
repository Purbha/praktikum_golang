package model_barang

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	config_db "github.com/purbha/latihan/praktikum2/latihan19/config"
	encrypt_teks "github.com/purbha/latihan/praktikum2/latihan19/encrypt"
)

func getpsqlInfo() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config_db.DBHost, config_db.DBPort, config_db.DBUser, config_db.DBPass, config_db.DBName)
}

func AmbilBarangs(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", getpsqlInfo())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Koneksi Sukses")
	var barangs []Barang
	//var brgdata BrgData
	var brgenc BrgEnc
	result, err := db.Query("SELECT * FROM barang")
	if err != nil {
		panic(err.Error())
	}
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
	e, err := json.Marshal(barangs)
	if err != nil {
		panic(err.Error())
	}
	encText, err := encrypt_teks.Encrypt(string(e), config_db.MySecret)
	if err != nil {
		panic(err.Error())
	}
	brgenc.Status = "OK"
	brgenc.Data = encText
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(brgenc)
}

func AmbilBarang(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", getpsqlInfo())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Koneksi Sukses")
	params := mux.Vars(r)
	idbar := params["id"]
	var barang Barang
	sqlStatement := "SELECT * FROM barang WHERE idbar = " + idbar
	result, err := db.Query(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
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
