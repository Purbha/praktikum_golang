package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	config_db "github.com/purbha/latihan/praktikum2/latihan13/config"
	barang_struct "github.com/purbha/latihan/praktikum2/latihan13/model"
)

/*
// const ini dipindahkan ke package config biar rapi
const (

	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "p@ssword"
	dbname   = "dbjual"

)
*/

/*
// struct ini dipindahkan ke package model biar rapi
type Barang struct {
	IDBar      string
	Bar_Nama   string
	Bar_Harga  float64
	Bar_Stok   int16
	Bar_Create time.Time
	Bar_Update time.Time
}
*/

func ambilBarangs(w http.ResponseWriter, r *http.Request) {
	//var barangs []Barang
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

var db *sql.DB
var err error

func main() {
	/*
		Portnya dirubah menjadi %s karena di config_db tipe port dirubah ke string.
		Kenapa dirubah ke string? Ga ada alasan cuma kepengen aja.
	*/
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
	var portNumber string = ":3000"
	fmt.Println("Server Running di Port", portNumber)
	log.Fatal(http.ListenAndServe(portNumber, r))
}
