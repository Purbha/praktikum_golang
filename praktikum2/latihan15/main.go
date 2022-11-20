package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	barang_struct "github.com/purbha/latihan/praktikum2/latihan15/model"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", barang_struct.AmbilBarangs).Methods("GET")
	r.HandleFunc("/barang/{id}", barang_struct.AmbilBarang).Methods("GET")
	var portNumber string = ":3000"
	fmt.Println("Server Running di Port", portNumber)
	log.Fatal(http.ListenAndServe(portNumber, r))
}
