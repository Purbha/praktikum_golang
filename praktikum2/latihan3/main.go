package main

import (
	"fmt"
	"net/http"
)

/*
https://stackoverflow.com/questions/71052339/what-does-printf-style-function-with-dynamic-format-string-and-no-further-argum
fmt.Fprintf() expects a format string that may contain verbs that will be substituted with arguments.
If you do not pass arguments, that hints you're likely not having / using a format string,
so you shouldn't use fmt.Fprintf() in the first place.
*/
func index(w http.ResponseWriter, r *http.Request) {
	var htmlData string = "<h1>Halo Apa Kabar Dunia</h1>"
	htmlData += "<p>Nama saya Irsyad dan tempat tinggal saya di Jatiasih.</p>"
	htmlData += "<p>Silakan main ke rumah jika ada waktu.</p>"
	fmt.Fprint(w, htmlData)
}

func main() {
	var portNumber string = ":3000"
	http.HandleFunc("/", index)
	fmt.Println("Server Running di Port", portNumber)
	http.ListenAndServe(portNumber, nil)
}
