package main

import (
	"fmt"
	"net/http"
)

/*
Fprintf formats according to a format specifier and writes to w.
It returns the number of bytes written and any write error encountered.
*/
func index(w http.ResponseWriter, r *http.Request) {
	//https://pkg.go.dev/fmt#Fprintf
	fmt.Fprintf(w, "<h1>Halo Apa Kabar</h1>")
}

func main() {
	//Router
	http.HandleFunc("/", index)
	fmt.Println("Server Running di Port 3000")
	http.ListenAndServe(":3000", nil)
}
