package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Barang struct {
	ID    string
	Nama  string
	Harga int
}

// Ambil data seluruh barang
func ambilBarangs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(barangs)
}

// Ambil data 1 barang berdasarkan ID
func ambilBarang(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//Ambil Parameter
	params := mux.Vars(r)
	//Cari didalam slice barang
	for _, item := range barangs {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Barang{})
}

// Tambah 1 barang
func tambahBarang(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var barang Barang
	_ = json.NewDecoder(r.Body).Decode(&barang)
	barang.ID = strconv.Itoa(rand.Intn(100000000))
	barangs = append(barangs, barang)
	json.NewEncoder(w).Encode(barang)
}

var barangs []Barang

func main() {
	barangs = append(barangs, Barang{ID: "1", Nama: "Flashdisk", Harga: 75000})
	barangs = append(barangs, Barang{ID: "2", Nama: "Mouse", Harga: 120000})
	r := mux.NewRouter()
	r.HandleFunc("/", ambilBarangs).Methods("GET")
	r.HandleFunc("/barang/{id}", ambilBarang).Methods("GET")
	r.HandleFunc("/tambahbarang", tambahBarang).Methods("POST")
	var portNumber string = ":3000"
	fmt.Println("Server Running di Port", portNumber)
	log.Fatal(http.ListenAndServe(portNumber, r))
}