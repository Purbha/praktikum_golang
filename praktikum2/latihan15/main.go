package main

import (
	"fmt"
	"log"
	"net/http"

	barang_struct "latihan-module/praktikum2/latihan15/model"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", barang_struct.AmbilBarangs).Methods("GET")
	r.HandleFunc("/barang/{id}", barang_struct.AmbilBarang).Methods("GET")
	var portNumber string = ":3000"
	fmt.Println("Server Running di Port", portNumber)
	log.Fatal(http.ListenAndServe(portNumber, r))
}
