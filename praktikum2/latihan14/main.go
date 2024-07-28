package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	config_db "latihan-module/praktikum2/latihan14/config"
	barang_struct "latihan-module/praktikum2/latihan14/model"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func ambilBarangs(w http.ResponseWriter, r *http.Request) {
	var barangs []barang_struct.Barang
	result, err := db.Query("SELECT * FROM barang")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var barang barang_struct.Barang
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

func ambilBarang(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	//id ini harus sama dengan id yang ada di r.HandleFunc("/barang/{id}"
	idbar := params["id"]
	var barang barang_struct.Barang
	// buat sql query dengan parameter idbar
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

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config_db.DBHost, config_db.DBPort, config_db.DBUser, config_db.DBPass, config_db.DBName)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Koneksi Sukses")
	r := mux.NewRouter()
	r.HandleFunc("/", ambilBarangs).Methods("GET")
	r.HandleFunc("/barang/{id}", ambilBarang).Methods("GET")
	var portNumber string = ":3000"
	fmt.Println("Server Running di Port", portNumber)
	log.Fatal(http.ListenAndServe(portNumber, r))
}
