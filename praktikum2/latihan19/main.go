package main

import (
	"fmt"
	"log"
	"net/http"

	model_barang "latihan-module/praktikum2/latihan19/model/barang"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", model_barang.AmbilBarangs).Methods("GET")
	r.HandleFunc("/barang/{id}", model_barang.AmbilBarang).Methods("GET")
	r.HandleFunc("/cekbarang", model_barang.Cekbarang).Methods("POST")
	var portNumber string = ":3000"
	fmt.Println("Server Running di Port", portNumber)
	log.Fatal(http.ListenAndServe(portNumber, r))
}
