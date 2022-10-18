package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Halo Apa Kabar Dunia</h1>")
	fmt.Fprintf(w, "<p>Nama saya Irsyad dan tempat tinggal saya di Jatiasih.</p>")
}

func main() {
	var portNumber string = ":3000"
	http.HandleFunc("/", index)
	fmt.Println("Server Running di Port", portNumber)
	http.ListenAndServe(portNumber, nil)
}
