package barang_struct

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	config_db "latihan-module/praktikum2/latihan15/config"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Barang struct {
	IDBar      string
	Bar_Nama   string
	Bar_Harga  float64
	Bar_Stok   int16
	Bar_Create time.Time
	Bar_Update time.Time
}

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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(barangs)
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
