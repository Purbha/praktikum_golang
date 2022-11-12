package main

/*
Install the github.com/lib/pq package
go get -u github.com/lib/pq
*/
import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Buat konstanta untuk koneksi
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "p@ssword"
	dbname   = "dbjual"
)

func main() {
	/*
		Buat Connection String
		sslmode di set disable karena pada lib/pg default valuenya tidak disable
		Jika sslmode tidak di disable maka akan berpotensi error -> pq: SSL is not enabled on the server
	*/
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//Buka Koneksi ke Database
	db, err := sql.Open("postgres", psqlInfo)

	//Pengecekan jika ada import yang kurang misalnya tidak import github.com/lib/pq
	if err != nil {
		panic(err)
	}
	defer db.Close()

	/*
		Sangat penting untuk menggunakan Ping() karena metode sql.Open() tidak pernah
		membuat koneksi ke database tetapi hanya memvalidasi argumen yang di gunakan
	*/
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Koneksi Sukses")
}
