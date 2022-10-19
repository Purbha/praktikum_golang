package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Barang struct {
	ID    string
	Nama  string
	Harga int
}

func ambilBarangs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(barangs)
}

var barangs []Barang

func main() {
	//https://go.dev/tour/moretypes/15
	/*
		Appending to a slice
		It is common to append new elements to a slice, and so Go provides a built-in append function
	*/
	barangs = append(barangs, Barang{ID: "1", Nama: "Flashdisk", Harga: 75000})
	barangs = append(barangs, Barang{ID: "2", Nama: "Mouse", Harga: 120000})
	//https://github.com/gorilla/mux
	/*
		Package gorilla/mux implements a request router and dispatcher for matching
		incoming requests to their respective handler.
	*/
	r := mux.NewRouter()
	r.HandleFunc("/", ambilBarangs).Methods("GET")
	var portNumber string = ":3000"
	fmt.Println("Server Running di Port", portNumber)
	log.Fatal(http.ListenAndServe(portNumber, r))
}
