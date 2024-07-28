package main

import (
	"fmt"
	"net/http"

	"latihan-module/praktikum2/latihan4/about" //Virtual Path
)

func index(w http.ResponseWriter, r *http.Request) {
	var htmlData string = "<h1>Halo Apa Kabar Dunia</h1>"
	htmlData += "<p>Nama saya Irsyad dan tempat tinggal saya di Jatiasih.</p>"
	htmlData += "<p>Silakan main ke rumah jika ada waktu.</p>"
	fmt.Fprint(w, htmlData)
}

func page_main(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, about.Page_about())
}

func main() {
	var portNumber string = ":3000"
	http.HandleFunc("/", index)
	http.HandleFunc("/about", page_main)
	fmt.Println("Server Running di Port", portNumber)
	http.ListenAndServe(portNumber, nil)
}
