package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	var htmlData string = "<h1>Halo Apa Kabar Dunia</h1>"
	htmlData += "<p>Nama saya Irsyad dan tempat tinggal saya di Jatiasih.</p>"
	htmlData += "<p>Silakan main ke rumah jika ada waktu.</p>"
	fmt.Fprint(w, htmlData)
}

func about(w http.ResponseWriter, r *http.Request) {
	var htmlData string = "<h1>Halaman About</h1>"
	htmlData += "<p>Ini adalah halaman about.</p>"
	htmlData += "<p>Hobi saya adalah bermain musik.</p>"
	fmt.Fprint(w, htmlData)
}

func main() {
	var portNumber string = ":3000"
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server Running di Port", portNumber)
	http.ListenAndServe(portNumber, nil)
}
