package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	model_barang "github.com/purbha/latihan/praktikum2/latihan17/model/barang"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", model_barang.AmbilBarangs).Methods("GET")
	r.HandleFunc("/barang/{id}", model_barang.AmbilBarang).Methods("GET")
	r.HandleFunc("/tambahbarang", model_barang.TambahBarang).Methods("POST")

	//r.HandleFunc("/hapusbarang/{id}", hapusBarang).Methods("DELETE")
	//r.HandleFunc("/ubahbarang/{id}", ubahBarang).Methods("PUT")

	var portNumber string = ":3000"
	fmt.Println("Server Running di Port", portNumber)
	log.Fatal(http.ListenAndServe(portNumber, r))
}
