package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	model_barang "github.com/purbha/latihan/praktikum2/latihan19/model/barang"
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
